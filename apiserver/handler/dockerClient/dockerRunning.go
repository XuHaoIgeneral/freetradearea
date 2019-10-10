package dockerClient

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"apiserver/struct/request"
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/schedule"
	"apiserver/service/initClient"
	"apiserver/service/dockerServer"
	"apiserver/service/clusterServer/activeCluster"
)

func DockerRunning(c *gin.Context) {
	var req request.DockerRun
	if err := c.BindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	//判定是否为自动分配
	if req.Auto {
		// 自动调度 选定节点
		req.HostIp = schedule.DefaultSchedule.Obtain()
	} else {
		// 断定指定宿主ip的合法性
		isLegal:=clusterHealth.IpisFind(req.HostIp)
		if !isLegal{
			SendResponse(c,errno.ErrDockerIp,nil)
			return
		}
	}
	fmt.Println(req.HostIp)
	//调度使用
	hostIp := "tcp://" + req.HostIp + ":5251"
	cli, err := initClient.InitDockerClient(hostIp)
	if err != nil {
		SendResponse(c, errno.ErrDockerClient, nil)
		return
	}

	//实现接口  当做参数传递
	var inDockerRunning dockerServer.InDockerRunning
	inDockerRunning = dockerServer.DockerRunning{
		cli,
	}

	//运行 running
	err=dockerServer.DefaultDockerAction.Run(inDockerRunning, req.ImageName, req.NetworkName,
		req.ContainerName, req.ExternalPort, req.InnerPort, req.Cmd, true, req.ContainerIp)
	if err!=nil{
		SendResponse(c, errno.ErrDockerRunning,err.Error())
		return
	}
	SendResponse(c, errno.OK, gin.H{
		"code": req,
	})
}
