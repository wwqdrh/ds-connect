package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	docker "github.com/fsouza/go-dockerclient"

	"github.com/wwqdrh/ds-connect/pkg/common"
	"github.com/wwqdrh/ds-connect/pkg/ds/swarm"
	"github.com/wwqdrh/ds-connect/pkg/service/dns"
)

var client *docker.Client

func init() {
	var err error
	client, err = docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}
}

func getEnvs(dnsmode string, debug bool) []string {
	envs := []string{}

	localDomains := dns.GetLocalDomains()
	if localDomains != "" {
		envs = append(envs, fmt.Sprintf("%s=%s", common.EnvVarLocalDomains, localDomains))
	}

	if strings.HasPrefix(dnsmode, common.DnsModeLocalDns) {
		envs = append(envs, fmt.Sprintf("%s=%s", common.EnvVarDnsProtocol, "tcp"))
	} else {
		envs = append(envs, fmt.Sprintf("%s=%s", common.EnvVarDnsProtocol, "udp"))
	}
	if debug {
		envs = append(envs, fmt.Sprintf("%s=%s", common.EnvVarLogLevel, "debug"))
	} else {
		envs = append(envs, fmt.Sprintf("%s=%s", common.EnvVarLogLevel, "info"))
	}

	return envs
}

func updateNetwork(ctx context.Context) {
	t := time.NewTicker(10 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Println("do a update...")
			if err := swarm.ServiceUpdateNet(client, "ds-connect-shadow", swarm.GetServiceDiff(client, "ds-connect-shadow")); err != nil {
				fmt.Println(err.Error())
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	_, err := swarm.ServiceInspect(client, "ds-connect-shadow")
	if err != nil {
		envs := getEnvs("localDNS", false)
		id, err := swarm.ServiceCreate(client, "ds-connect-shadow", "wwqdrh/ds-connect-shadow:dev", envs)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("创建成功: " + id)
	} else {
		fmt.Println("已经存在")
	}

	ctx, cancel := context.WithCancel(context.Background())
	go updateNetwork(ctx)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	defer cancel()
}
