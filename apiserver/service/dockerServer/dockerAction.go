package dockerServer

import (
	"fmt"
	"github.com/docker/docker/client"
	"apiserver/enum/dockerEnum"
	"github.com/lexkong/log"
	"github.com/docker/docker/api/types"
	"context"
)


type InDockerAction interface {
	Start(string2 string) error
	Stop(string2 string) error
	Remove(string2 string) error
	ReStart(string2 string) error
}

var DefaultDockerAction DockerAction

type DockerAction struct {
	Cli *client.Client
}

func (self DockerAction) Stop(containerId string) (err error) {
	ctx := context.Background()
	d2 := dockerEnum.DockerStopTime // value is of type int
	err = self.Cli.ContainerStop(ctx, containerId, &d2)
	if err != nil {
		log.Info(err.Error())
	}
	return
}

func (self DockerAction) Remove(containerId string) (err error) {
	ctx := context.Background()
	options:=types.ContainerRemoveOptions{
		Force:true,
	}
	fmt.Println("containerID==",containerId)
	err = self.Cli.ContainerRemove(ctx, containerId,options )
	//defaults to SIGKILL
	if err != nil {
		log.Info(err.Error())
	}
	return
}

func (self DockerAction) Start(containerId string) (err error) {
	ctx := context.Background()
	if err := self.Cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{}); err != nil {
		log.Info(err.Error())
	}
	return
}

func (self DockerAction) ReStart(containerId string) (err error){
	ctx:=context.Background()
	d2 := dockerEnum.DockerStopTime // value is of type int
	if err=self.Cli.ContainerRestart(ctx,containerId,&d2);err!=nil{
		return err
	}
	return
}

func (DockerAction) Run(i InDockerRunning, imageName, networkName, containerName, externalPort,
	innerPort string, cmd []string, tty bool, containerIp ...string) (err error) {
	if err = i.DockerPull(imageName); err != nil {
		log.Error(err.Error(), err)
		return
	}
	
	dockerId, _, err := i.DockerCreate(imageName, networkName, containerName, externalPort, innerPort, cmd, tty, containerIp...)
	if err != nil {
		log.Error("", err)
		return
	}

	if err = i.DockerStart(dockerId); err != nil {
		log.Error("", err)
	}
	return
}