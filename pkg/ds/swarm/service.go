package swarm

import (
	"errors"
	"fmt"

	"github.com/docker/docker/api/types/swarm"
	docker "github.com/fsouza/go-dockerclient"
)

func ServiceInspect(engine *docker.Client, name string) (*swarm.Service, error) {
	srv, err := engine.InspectService(name)
	if err != nil {
		return nil, err
	}
	if srv.ID == "" {
		return nil, errors.New("not exist")
	}
	return srv, nil
}

func ServiceCreate(engine *docker.Client, name, image string, envs []string) (string, error) {
	srv, err := engine.CreateService(docker.CreateServiceOptions{
		ServiceSpec: swarm.ServiceSpec{
			Annotations: swarm.Annotations{
				Name: name,
			},
			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{
					Image: image,
					Env:   envs,
				},
			},
		},
	})
	if err != nil {
		return "", err
	}
	return srv.ID, nil
}

func ServiceVersion(engine *docker.Client, srvName string) (uint64, error) {
	srv, err := engine.InspectService(srvName)
	if err != nil {
		return 0, err
	}
	return srv.Version.Index, nil
}

func ServiceUpdateNet(engine *docker.Client, srvName string, networks []string) error {
	if len(networks) == 0 {
		fmt.Println("无需设置")
		return nil
	}

	srv, err := ServiceInspect(engine, srvName)
	if err != nil {
		return err
	}
	fmt.Println("cur version: ", srv.Version.Index)

	nets := []swarm.NetworkAttachmentConfig{}
	for _, item := range networks {
		nets = append(nets, swarm.NetworkAttachmentConfig{Target: item})
	}

	spec := srv.Spec
	spec.TaskTemplate.Networks = nets
	return engine.UpdateService(srvName, docker.UpdateServiceOptions{
		Version:     srv.Version.Index + 1,
		ServiceSpec: spec,
	})
}

func ServiceAddNet(engine *docker.Client, network string) {

}

func ServiceDelNet(engine *docker.Client, network string) {

}
