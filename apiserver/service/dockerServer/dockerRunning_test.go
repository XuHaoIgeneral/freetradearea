package dockerServer

import (
	"testing"
	"apiserver/service/initClient"
)

func TestDockerRunning_DockerCreate(t *testing.T) {
	ts := new(DockerRunning)
	cli, err := initClient.InitDockerClient("tcp://192.168.56.201:5251")
	if err != nil {
		return
	}
	ts.Cli = cli
	imageName := "nginx"
	networkName := "ceshi1"
	containerName := "testNginx"
	exPort := "8081"
	inPort := "80"
	myIp := "10.0.1.15"
	respId, respWarnnings, err := ts.DockerCreate(imageName, networkName, containerName, exPort, inPort, []string{}, true, myIp)
	if err != nil {
		t.Error(err)
		t.Error(respWarnnings)
	} else {
		t.Log(respId)
	}
}

func TestDockerRunning_DockerPull(t *testing.T)  {
	ts := new(DockerRunning)
	cli, err := initClient.InitDockerClient("tcp://192.168.56.201:5251")
	if err != nil {
		return
	}
	ts.Cli = cli
	err=ts.DockerPull("alpine")
	if err!=nil{
		t.Error(err)
	}else {
		t.Log("ok")
	}
}
