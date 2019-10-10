package initClient

import (
	"github.com/spf13/viper"
	"time"
	"github.com/coreos/etcd/clientv3"
	"sync"
)

var (
	once       sync.Once
	EtcdClient *clientv3.Client
)

//创建etcd的 Client
//调用完毕后 close
func NewEtcdClient() *clientv3.Client {
	once.Do(func() {
		initEndpoints := viper.GetStringSlice("cluster")
		EtcdClient, _ = clientv3.New(clientv3.Config{
			Endpoints:   initEndpoints,
			DialTimeout: 5 * time.Second,
		})
	})
	return EtcdClient
}
