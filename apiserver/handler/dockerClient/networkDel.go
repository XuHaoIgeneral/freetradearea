package dockerClient

import (
	"apiserver/enum/dockerEnum"
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/dockerNetwork"
	"apiserver/struct/request"
	"github.com/gin-gonic/gin"
)

func NetworkDel(c *gin.Context) {
	//集群network list
	var req request.NetworkDel
	if err := c.BindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	tempDockerNetwork := new(dockerNetwork.DockerNetwork)
	tempDockerNetwork.Ip = "TCP://127.0.0.1:" + dockerEnum.DockerPort
	err := tempDockerNetwork.Remove(req.NetworkId)
	if err != nil {
		SendResponse(c, errno.ErrDockerIp, nil)
		return
	}
	SendResponse(c, errno.OK, nil)
}
