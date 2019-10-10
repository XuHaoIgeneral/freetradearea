package sd

import (
	"github.com/gin-gonic/gin"
	"nodeserver/server"
	"nodeserver/model"
	"nodeserver/handler"
	"nodeserver/pkg/errno"
)

// @Summary Shows OK as the ping-pong result
// @Description Shows OK as the ping-pong result
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {string} plain "OK"
// @Router /node/status [get]
func HealthCheck(c *gin.Context) {
	var inNode server.NodeInterface
	inNode = new(server.NodeServer)
	data := model.HealthCheck{
		Cpu:  inNode.Cpu(),
		Ram:  inNode.Ram(),
		Disk: inNode.Disk(),
	}
	handler.SendResponse(c, errno.OK, data)
}
