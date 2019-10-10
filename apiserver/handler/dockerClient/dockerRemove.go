package dockerClient

import (
	"github.com/gin-gonic/gin"
	"apiserver/struct/request"
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/clusterServer/activeCluster"
	"apiserver/service/initClient"
	"apiserver/service/dockerServer"
	"apiserver/enum/dockerEnum"
)

func DockerWithRemove(c *gin.Context)  {
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
	defer cli.Close()
	if err!=nil{
		SendResponse(c,errno.ErrDockerClient,nil)
		return
	}

	//执行remove操作
	dockerAction:=dockerServer.DockerAction{
		cli,
	}

	// 默认执行一次 stop操作
	dockerAction.Stop(req.ContainerId)

	// 执行remove操作
	if err:=dockerAction.Remove(req.ContainerId);err!=nil{
		SendResponse(c,errno.ErrDockerAction,nil)
		return
	}

	SendResponse(c,errno.OK,nil)
}