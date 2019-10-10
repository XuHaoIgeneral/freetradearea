package dockerServer

import (
	"fmt"
	"github.com/docker/docker/client"
	"github.com/lexkong/log"
	"github.com/docker/docker/api/types"
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/docker/docker/api/types/network"
	"os"
	"io"
)

type InDockerRunning interface {
	DockerPull(imageName string) error
	DockerCreate(imageName, networkName, containerName, externalPort, innerport string, cmd []string, tty bool, containerIp ...string) (respId string, respWarnings []string, err error)
	DockerStart(dockerId string) (err error)
}

type DockerRunning struct {
	Cli *client.Client
}

func (self DockerRunning) DockerPull(imageName string) (err error) {
	ctx := context.Background()
	imageName ="docker.io/library/"+imageName
	reader, err := self.Cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return
	}
	io.Copy(os.Stdout, reader)
	return
}

func (self DockerRunning) DockerCreate(imageName, networkName, containerName,
	externalPort, innerport string, cmd []string, tty bool,
	containerIp ...string) (respId string, respWarnings []string, err error) {


	fmt.Println("image==",imageName)
	fmt.Println("networkName",networkName)
	fmt.Println("containerName",containerName)
	fmt.Println("externalP",externalPort)
	fmt.Println("innerport",innerport)
	fmt.Println("cmd",cmd)
	fmt.Println("containnerIp",containerIp)

	ctx := context.Background()

	exposedPorts := nat.PortSet{
		//80
		nat.Port(innerport + "/tcp"): {},
	}

	portBindings := nat.PortMap{
		nat.Port(innerport + "/tcp"): []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: externalPort}},
	}

	networkSet := make(map[string]*network.EndpointSettings)

	var networkEp = network.EndpointSettings{}
	if len(containerIp) != 0 {
		networkEp = network.EndpointSettings{
			IPAddress: containerIp[0],
		}
	}

	networkSet[networkName] = &networkEp
	resp, err := self.Cli.ContainerCreate(ctx, &container.Config{
		Image:        imageName,
		Tty:          true,
		ExposedPorts: exposedPorts,
	}, &container.HostConfig{
		PortBindings: portBindings,
	},
		&network.NetworkingConfig{
			networkSet,
		},
		containerName)
	if err != nil {
		log.Info(err.Error())
		respWarnings = resp.Warnings
		return
	}

	respId = resp.ID
	respWarnings = resp.Warnings
	return
}

func (self DockerRunning) DockerStart(dockerId string) (err error) {
	ctx := context.Background()
	err = self.Cli.ContainerStart(ctx, dockerId, types.ContainerStartOptions{})
	return
}
