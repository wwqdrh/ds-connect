package swarm

import (
	"fmt"
	"testing"

	docker "github.com/fsouza/go-dockerclient"
)

func TestServiceIPMap(t *testing.T) {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		t.Skip("docker环境无法使用")
	}

	ips, err := ServiceIPMap(client)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ips)
}
