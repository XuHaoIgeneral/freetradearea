package clusterNode

import (
	"github.com/gin-gonic/gin"
	"apiserver/service/clusterServer/clusterStatus"
	"apiserver/enum/sendToNode"
	"apiserver/pkg/errno"
	. "apiserver/handler"
)

func ClusterStatus(c *gin.Context) {
	ll := clusterStatus.NodeStatus(sendToNode.HTTP, sendToNode.PortHttp, sendToNode.GET, sendToNode.URL)
	SendResponse(c, errno.OK, ll)
}
