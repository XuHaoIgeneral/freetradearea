package initClient

import (
	"github.com/docker/docker/client"
	"apiserver/enum/dockerEnum"
	"github.com/lexkong/log"
)

func InitDockerClient(hostIp string) (cli *client.Client, err error) {
	cli, err = client.NewClient(hostIp, dockerEnum.Version, nil, map[string]string{"Content-type": "application/x-tar"})
	if err != nil {
		log.Info(err.Error())
	}
	return
}
