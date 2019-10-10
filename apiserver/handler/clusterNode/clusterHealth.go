package clusterNode

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"apiserver/pkg/errno"
	"apiserver/service/clusterServer/activeCluster"
	. "apiserver/handler"
)

func ClusterHealth(c *gin.Context) {
	fmt.Println("he")
	ll := clusterHealth.Cluster.ClusterList()
	SendResponse(c, errno.OK, ll)
}
