package dockerClient

import (
	"github.com/gin-gonic/gin"
	"apiserver/struct/request"
	"apiserver/service/clusterServer/activeCluster"
	"apiserver/pkg/errno"
	"apiserver/service/initClient"
	"apiserver/service/dockerServer"
	. "apiserver/handler"
	"apiserver/enum/dockerEnum"
)

// post:
func DockerWithStop(c *gin.Context)  {
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

	dockerAction:=dockerServer.DockerAction{
		cli,
	}

	// 执行stop操作
	if err:=dockerAction.Stop(req.ContainerId);err!=nil{
		SendResponse(c,errno.ErrDockerAction,nil)
		return
	}

	SendResponse(c,errno.OK,nil)
}
