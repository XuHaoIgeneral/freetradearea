package dockerClient

import (
	"apiserver/enum/dockerEnum"
	"apiserver/pkg/errno"
	"apiserver/service/dockerNetwork"
	"apiserver/struct/request"
	"fmt"
	"github.com/gin-gonic/gin"
	. "apiserver/handler"
)

func NetworkCreate(c *gin.Context) {
	//集群network list
	var req request.NetworkCreate

	if err := c.BindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	tempDockerNetwork:=new(dockerNetwork.DockerNetwork)
	tempDockerNetwork.Ip="TCP://127.0.0.1:"+dockerEnum.DockerPort
	_,err:=tempDockerNetwork.Create(req.Networkname,req.Subnet,req.Gateway)
	if err!=nil{
		fmt.Println(err)
		SendResponse(c,errno.ErrDockerIp,nil)
		return
	}
	SendResponse(c,errno.OK,nil)
}