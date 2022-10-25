package command

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wwqdrh/clitool"
	"github.com/wwqdrh/ds-connect/pkg/ds/command/connect"
	"github.com/wwqdrh/logger"
	"github.com/wwqdrh/nettool/device/tun"
	"github.com/wwqdrh/nettool/server/dns"
	"github.com/wwqdrh/nettool/server/https"
	"github.com/wwqdrh/nettool/server/ssh"
	"github.com/wwqdrh/ostool"
)

var connectFlags = []clitool.OptionConfig{
	{
		Target:       "Shadow",
		DefaultValue: "127.0.0.1:18080",
		Description:  "docker swarm集群中的shadow服务地址",
	},
	{
		Target:       "SshUser",
		DefaultValue: "123456",
		Description:  "ssh用户名",
	},
	{
		Target:       "SshPass",
		DefaultValue: "",
		Description:  "ssh密码",
	},
	{
		Target:       "ProxyPort",
		DefaultValue: 2223,
		Description:  "(tun2socks mode only) Specify the local port which socks5 proxy should use",
	},
	{
		Target:       "DnsPort",
		DefaultValue: 53,
		Description:  "本地dns服务的端口",
	},
	{
		Target:       "DnsCacheTtl",
		DefaultValue: 60,
		Description:  "(local dns mode only) DNS cache refresh interval in seconds",
	},
}

// NewConnectCommand return new connect command
func NewConnectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connect",
		Short: "Create a network tunnel to docker swarm cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return fmt.Errorf("too many options specified (%s)", strings.Join(args, ","))
			}
			if err := preCheck(); err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return Connect()
		},
		Example: "ktctl connect [command options]",
	}

	clitool.SetOptions(cmd, cmd.Flags(), Opts().Connect, connectFlags)
	return cmd
}

// 1、start a socks5 local forward(本地转发模式)
// 2、tun2socks
// 3、为tun设置路由规则
// 4、启动dns server
func Connect() error {
	fmt.Println("do a connect")
	channel := ssh.NewChannel(ssh.ChannelWithAuth(
		Opts().Connect.SshUser, Opts().Connect.SshPass,
	))
	sockAddr := fmt.Sprintf("127.0.0.1:%d", Opts().Connect.ProxyPort)

	go func() {
		if err := channel.StartSocks5Proxy(sockAddr); err != nil {
			logger.DefaultLogger.Error(err.Error())
		}
	}()

	go func() {
		t := tun.Ins()
		// 设置路由规则
		go func() {
			if err := t.ToSocks(sockAddr); err != nil {
				logger.DefaultLogger.Error(err.Error())
			}
		}()
		connect.UpdateRoute()
	}()

	// 启动dns
	dnsserver := dns.NewDnsServer(func(domain string) [4]byte {
		res, err := https.DoReq(&https.HTTPOpt{
			Method: "GET",
			Url:    "/service/query",
			Query:  map[string][]string{"name": {domain}},
		})
		if err != nil {
			return [4]byte{}
		} else {
			return string2IP(res.Body)
		}
	}, func(ip string) string {
		return ip
	}, Opts().Connect.DnsPort)
	// TODO(暂时手动设置dns域名) 设置iptables规则
	dnsserver.Server(context.TODO())

	return nil
}

func string2IP(ip string) [4]byte {
	res := [4]byte{}
	for i, item := range strings.SplitN(ip, ".", 4) {
		ch, _ := strconv.Atoi(item)
		res[i] = byte(ch)
	}
	return res
}

func preCheck() error {
	if err := checkPermissionAndOptions(); err != nil {
		return err
	}
	// if pid := util.GetDaemonRunning(util.ComponentConnect); pid > 0 {
	// 	return fmt.Errorf("another connect process already running at %d, exiting", pid)
	// }
	return nil
}

func silenceCleanup() {

}

func checkPermissionAndOptions() error {
	if !ostool.IsRunAsAdmin() {
		if ostool.IsWindows() {
			return fmt.Errorf("permission declined, please re-run connect command as Administrator")
		}
		return fmt.Errorf("permission declined, please re-run connect command with 'sudo'")
	}
	// if opt.Get().Connect.Mode == util.ConnectModeTun2Socks && opt.Get().Connect.DnsMode == util.DnsModePodDns {
	// 	return fmt.Errorf("dns mode '%s' is not available for connect mode '%s'", util.DnsModePodDns, util.ConnectModeTun2Socks)
	// }
	return nil
}
