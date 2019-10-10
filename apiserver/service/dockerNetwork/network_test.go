package dockerNetwork

import (
	"testing"
	"apiserver/enum/dockerEnum"
)

func TestDockerNetwork_List(t *testing.T) {
	test := new(DockerNetwork)
	test.Ip = "tcp://127.0.0.1:" + dockerEnum.DockerPort
	resp, err := test.List()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp)
	}
}

func TestDockerNetwork_Create(t *testing.T) {
	test := new(DockerNetwork)
	test.Ip = "tcp://127.0.0.1:" + dockerEnum.DockerPort
	resp, err := test.Create("ceshi1", "10.0.1.0/24", "10.0.1.1")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp)
	}
}

func TestDockerNetwork_Remove(t *testing.T) {
	test := new(DockerNetwork)
	test.Ip = "tcp://127.0.0.1:" + dockerEnum.DockerPort
	if err := test.Remove("ae24f74422f9"); err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}
