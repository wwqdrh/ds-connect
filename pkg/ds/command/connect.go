package command

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wwqdrh/nettool/common"
)

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

	return cmd
}

func Connect() error {
	fmt.Println("do a connect")
	return nil
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
	if !common.IsRunAsAdmin() {
		if common.IsWindows() {
			return fmt.Errorf("permission declined, please re-run connect command as Administrator")
		}
		return fmt.Errorf("permission declined, please re-run connect command with 'sudo'")
	}
	// if opt.Get().Connect.Mode == util.ConnectModeTun2Socks && opt.Get().Connect.DnsMode == util.DnsModePodDns {
	// 	return fmt.Errorf("dns mode '%s' is not available for connect mode '%s'", util.DnsModePodDns, util.ConnectModeTun2Socks)
	// }
	return nil
}
