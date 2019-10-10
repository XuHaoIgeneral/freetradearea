package clusterStatus

import (
	"apiserver/service/clusterServer/activeCluster"
	"sync"
	"apiserver/pkg/sendHttp"
	"github.com/lexkong/log"
	"apiserver/enum/overtime"
	"context"
	"apiserver/struct/response"
)

var (
	lock      sync.Mutex
	nodeLists = make([]response.ClusterStatusResp, 0)
)

// 对所有节点的cpu，ram，disk情况进行检查
// agreement 发送方式 ex：http、https、tcp 枚举
// port 节点主机接收的端口号  枚举！
// mode 发送格式 ex：post、get  枚举！
// url  路由地址  枚举!
// body http or https body
func NodeStatus(agreement, port, mode, url string, body ...map[string]interface{}) []response.ClusterStatusResp {
	//清空
	nodeLists = nil
	defer func() {
		nodeLists = nil
	}()

	//获取正常节点ip
	normalLists := clusterHealth.Cluster.ClusterList()

	//开启go程对集群状态进行查询
	//设定 有超时判断
	wg := sync.WaitGroup{}

	for _, v := range normalLists {
		temp := v
		wg.Add(1)
		go func() {
			defer wg.Done()
			chSign := make(chan bool)

			ctx, _ := context.WithTimeout(context.Background(), overtime.OvertimePs)
			//单独对节点进行http请求，并接受返回值。加入超时判断机制
			go func() {
				tempSendModel := &sendHttp.SendToNodeModel{
					Ip:   temp.Ip,
					Port: port,
					Mode: mode,
					Url:  url,
				}

				health, err := sendHttp.SendToFunc(agreement, tempSendModel)
				if err != nil {
					log.Info(err.Error())
					chSign <- false
					return
				}

				temp := response.ClusterStatusResp{
					Ip:   temp.Ip,
					Name: temp.Name,
					List: health,
				}

				//加锁
				lock.Lock()

				nodeLists = append(nodeLists, temp)
				//解锁
				lock.Unlock()

				chSign <- true
			}()

			//超时判定
			select {
			case <-chSign:

			case <-ctx.Done():
				log.Info(v.Ip + "超时")
			}
		}()
	}
	wg.Wait()

	return nodeLists
}
