package swarm

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

func GetServiceDiff(engine *docker.Client, srvName string) []string {
	srvnets, err := ServiceNetwork(engine, srvName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	nets, err := ListNetworkByMode(engine, "overlay")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	srvset := map[string]struct{}{}
	for _, item := range srvnets {
		srvset[item] = struct{}{}
	}

	res := []string{}
	for _, item := range nets {
		if _, ok := srvset[item]; !ok {
			res = append(res, item)
		}
	}
	return res
}

// get service's networks, here is network id
func ServiceNetwork(client *docker.Client, name string) ([]string, error) {
	srv, err := client.InspectService(name)
	if err != nil {
		return nil, err
	}

	res := []string{}
	for _, item := range srv.Endpoint.VirtualIPs {
		res = append(res, item.NetworkID)
	}
	return res, nil
}

func ListNetworkByMode(client *docker.Client, mode string) ([]string, error) {
	networks, err := client.ListNetworks()
	if err != nil {
		return nil, err
	}

	res := []string{}
	for _, net := range networks {
		if net.Driver == mode && net.Name != "ingress" {
			res = append(res, net.ID)
		}
	}
	return res, nil
}
