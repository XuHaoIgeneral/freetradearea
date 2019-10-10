package clusterHealt

import (
	. "apiserver/model/serverModel"
	"apiserver/service/initClient"
	"apiserver/util"
	"context"
	"github.com/lexkong/log"
	"time"
)

var Cluster activeCluster

type activeCluster struct {
}

//需要优化
// 目标 节点并行
// docker ps 为节点全并行
func (activeCluster) ClusterList() (nodeList []ClusterNode) {
	cli := initClient.NewEtcdClient()

	endpointsList := etcdLists()

	nodeList = make([]ClusterNode, 0)
	for _, v := range endpointsList {
		ctx, _ := context.WithTimeout(context.Background(), 200*time.Millisecond)

		chSign := make(chan bool)
		str := util.IpSlice(v.Ip)
		go func() {
			//状态判断，失联的节点断开
			_, err := cli.Status(context.Background(), v.Ip)
			if err != nil {
				log.Infof(err.Error())
			}
			chSign <- true
		}()
		//超时器
		select {
		case <-chSign:
			tempModel := ClusterNode{
				Name:   v.Name,
				Id:     v.Id,
				Ip:     str,
				Status: true,
			}
			nodeList = append(nodeList, tempModel)
			continue
		case <-ctx.Done():
			tempModel := ClusterNode{
				Name:   v.Name,
				Id:     v.Id,
				Ip:     str,
				Status: false,
			}
			nodeList = append(nodeList, tempModel)
			continue
		}
	}
	return
}

func etcdLists() (listNode []ClusterNode) {
	listNode = make([]ClusterNode, 0)
	cli := initClient.NewEtcdClient()
	memberListPrt, err := cli.MemberList(context.Background())
	if err != nil {
		log.Info(err.Error())
		return
	}
	list := memberListPrt.Members
	for _, v := range list {
		tempModel := ClusterNode{
			Name: v.Name,
			Id:   v.ID,
			Ip:   util.HttpSlice(v.ClientURLs[1]),
		}
		listNode = append(listNode, tempModel)
	}
	return listNode
}

// 返回运行正常的节点ip
func IpList() (list []string) {
	nodeList := Cluster.ClusterList()
	for _, v := range nodeList {
		list = append(list, v.Ip)
	}
	return
}

//判定 ip值是否有效
func IpisFind(ip string) (isFind bool) {
	list := IpList()
	for _, v := range list {
		if ip == v {
			isFind = true
			break
		}
	}
	return
}
