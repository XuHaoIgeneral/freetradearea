package dockerServer

import (
	"github.com/docker/docker/client"
)

type DockerImage struct {
	cli *client.Client
}

//涉及去重
func (self DockerImage) List() {
	//cli := self.cli
	//ctx := context.Background()
	//options := types.ImageListOptions{}
	//respList, err := cli.ImageList(ctx, options)
	//if err != nil {
	//	log.Infof(err.Error())
	//	return
	//}
	//
	////image 集群镜像去重
	//var imageList = make(map[types.ImageSummary]bool, 0)
	//for _, v := range respList {
	//}
}
