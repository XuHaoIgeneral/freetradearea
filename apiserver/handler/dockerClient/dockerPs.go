package dockerClient

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"apiserver/struct/response"
	"apiserver/service/dockerServer"
	"sync"
	"apiserver/model/dockerModel"
	"apiserver/enum/overtime"
	"github.com/lexkong/log"
	. "apiserver/handler"
	"apiserver/pkg/errno"
)

// @Summary  To list running containers (the equivalent of "docker ps")
// @Description select running dockers
// @Tags docker-client
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /test/ps [get]
func DockerPs(c *gin.Context) {

	//获取集群ip信息
	hostList := dockerServer.DockerPs.ChoiceHealth()

	wg := sync.WaitGroup{}
	options := types.ContainerListOptions{All: true}
	chData := make(chan []response.RespDockerPs, 1)
	chProducer := make(chan bool)
	chOverTime := make(chan struct{})

	for _, v := range hostList {
		temp := v
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, _ := context.WithTimeout(context.Background(), overtime.OvertimePs)
			go func(dockerModel.DockerPsModel) {
				dockerServer.DockerPs.SendToDocker(temp, options, chData)
				chOverTime <- struct{}{}
			}(temp)

			select {
			case <-ctx.Done():
				//timeout
				log.Infof("timeout==" + temp.HostName + temp.HostIp)
				chProducer <- false
			case <-chOverTime:
				chProducer <- true
			}
		}()
	}

	//消费者接受
	wg.Add(1)
	data := make([]response.RespDockerPs, 0)
	go func() {
		defer wg.Done()
		defer func() {

		}()
		for {
			dataList, isOk := <-chData
			if !isOk {
				break
			}
			data = append(data, dataList...)
		}
	}()
	
	//生产者管道关闭
	wg.Add(1)
	go func(count int) {
		defer wg.Done()
		for count > 0 {
			select {
			case <-chProducer:
				count--
			}
		}
		close(chProducer)
		close(chData)
		close(chOverTime)
	}(len(hostList))
	wg.Wait()

	SendResponse(c, errno.OK, data)
}
