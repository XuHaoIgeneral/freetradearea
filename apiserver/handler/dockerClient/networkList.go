package dockerClient

import (
	"github.com/gin-gonic/gin"
	"apiserver/service/dockerNetwork"
	"apiserver/enum/dockerEnum"
	. "apiserver/handler"
	"apiserver/pkg/errno"
)

// 集群网络查看
// 使用了ETCD 所以只用查看集群中任意一台主机，即可了解集群下的网络状态
func NetworkList(c *gin.Context) {
	//集群network list
	tempDockerNetwork:=new(dockerNetwork.DockerNetwork)
	tempDockerNetwork.Ip="TCP://127.0.0.1:"+dockerEnum.DockerPort
	networkList,err:=tempDockerNetwork.List()
	if err!=nil{
		SendResponse(c,errno.ErrDockerIp,err)
		return
	}

	SendResponse(c,errno.OK,networkList)
}