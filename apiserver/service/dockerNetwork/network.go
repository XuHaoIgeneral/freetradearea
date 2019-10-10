package dockerNetwork

import (
	"fmt"
	"github.com/docker/docker/client"
	"apiserver/enum/dockerEnum"
	"context"
	"github.com/docker/docker/api/types"
	"apiserver/model/dockerModel"
	"github.com/docker/docker/api/types/network"
)

//默认支持为ETCD下的overlay网络
var Default DockerNetwork

type DockerNetwork struct {
	Ip string
}

func (self DockerNetwork) List() (list []dockerModel.DockerNetworlList, err error) {

	cli, err := client.NewClient(self.Ip, dockerEnum.Version, nil, map[string]string{"Content-type": "application/x-tar"})
	defer cli.Close()
	if err != nil {
		return
	}
	ctx := context.Background()
	netLists, err := cli.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	list = make([]dockerModel.DockerNetworlList, 0)
	for _, v := range netLists {
		temp := dockerModel.DockerNetworlList{
			Id:     v.ID,
			Name:   v.Name,
			Driver: v.Driver,
			Scope:  v.Scope,
		}
		list = append(list, temp)
	}
	return
}

//subnet  10.0.0.0/24
//gateway 10.0.0.1
func (self DockerNetwork) Create(name, subnet, gateway string) (resp types.NetworkCreateResponse, err error) {
	cli, err := client.NewClient(self.Ip, dockerEnum.Version, nil, map[string]string{"Content-type": "application/x-tar"})
	defer cli.Close()
	if err != nil {
		return
	}
	ctx := context.Background()

	ipConfigList := make([]network.IPAMConfig, 0)
	ipConfig := new(network.IPAMConfig)
	ipConfig.Gateway = gateway
	ipConfig.Subnet = subnet
	ipConfigList = append(ipConfigList, *ipConfig)
	fmt.Println("AAA")
	options := types.NetworkCreate{
		CheckDuplicate: true,
		Attachable:     true,
		Driver:         "overlay",
		IPAM: &network.IPAM{
			Config: ipConfigList,
		},
	}
	resp, err = cli.NetworkCreate(ctx, name, options)
	return resp, err
}

func (self DockerNetwork) Remove(networkId string) (err error) {
	cli, err := client.NewClient(self.Ip, dockerEnum.Version, nil, map[string]string{"Content-type": "application/x-tar"})
	defer cli.Close()
	if err != nil {
		return
	}
	ctx := context.Background()
	err = cli.NetworkRemove(ctx, networkId)
	return
}

func (self DockerNetwork) ContainerAdd() {

}

func (self DockerNetwork) ContainerRemove() {

}
