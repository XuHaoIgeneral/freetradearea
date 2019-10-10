package dockerClient

import (
	"apiserver/struct/request"
	"apiserver/service/clusterServer/activeCluster"
	"apiserver/pkg/errno"
	"apiserver/enum/dockerEnum"
	"apiserver/service/initClient"
	"apiserver/service/dockerServer"
	. "apiserver/handler"
	"github.com/gin-gonic/gin"
)

func DockerWithRestart(c *gin.Context)  {
	var req request.DockerIp
	if err:=c.BindJSON(&req);err!=nil{
		SendResponse(c, errno.ErrBind,nil)
		return
	}

	//验证ip输入的合法性
	if !clusterHealth.IpisFind(req.Ip){
		SendResponse(c,errno.ErrDockerIp,nil)
		return
	}

	//创建客户端
	hostIp := "tcp://" + req.Ip + ":"+dockerEnum.DockerPort
	cli,err:=initClient.InitDockerClient(hostIp)
	if err!=nil{
		SendResponse(c,errno.ErrDockerClient,nil)
		return
	}
	dockerAction:=dockerServer.DockerAction{
		cli,
	}
	// 执行start操作
	if err:=dockerAction.ReStart(req.ContainerId);err!=nil{
		SendResponse(c,errno.ErrDockerAction,nil)
		return
	}

	SendResponse(c,errno.OK,nil)
}
