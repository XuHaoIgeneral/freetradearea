package dockerServer

import (
	"apiserver/util"
	"apiserver/enum/dockerEnum"
	"apiserver/service/clusterServer/activeCluster"
	"apiserver/model/dockerModel"
	"github.com/docker/docker/client"
	"github.com/lexkong/log"
	"github.com/docker/docker/api/types"
	"context"
	"apiserver/struct/response"
)

var DockerPs dockerPsModel

type dockerPsModel struct {
}

func (dockerPsModel) ChoiceHealth() []dockerModel.DockerPsModel {
	nodeList := clusterHealth.Cluster.ClusterList()
	resp := make([]dockerModel.DockerPsModel, 0)
	for _, v := range nodeList {
		if v.Status == true {
			tempModel := dockerModel.DockerPsModel{
				HostName: v.Name,
				HostIp:   "tcp://" + util.IpSlice(v.Ip) + ":" + dockerEnum.DockerPort,
			}
			resp = append(resp, tempModel)
		}
	}
	return resp
}

//  并行go程处理
//  n-1 生产者消费者模型
//
func (dockerPsModel) SendToDocker(host dockerModel.DockerPsModel, options types.ContainerListOptions, ch chan<- []response.RespDockerPs) {

	cli, err := client.NewClient(host.HostIp, dockerEnum.Version, nil, map[string]string{"Content-type": "application/x-tar"})
	defer cli.Close()
	data := make([]response.RespDockerPs, 0)
	defer func() {
		ch <- data
	}()
	if err != nil {
		log.Infof(err.Error())
		return
	}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		log.Infof(err.Error())
		return
	}

	for _, container := range containers {
		tempContainer := response.RespDockerPs{
			Id:           container.ID[:10],
			Name:         container.Names,
			Image:        container.Image,
			Create:       container.Created,
			State:        container.State,
			NetworkMode:  container.HostConfig.NetworkMode,
			Ports:        container.Ports,
			HostNodeName: host.HostName,
		}
		data = append(data, tempContainer)
	}
}
