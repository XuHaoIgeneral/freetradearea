package sendHttp

import (
	"apiserver/enum/sendToNode"
	"net/http"
	"github.com/lexkong/log"
	"io/ioutil"
	"github.com/tidwall/gjson"
	"apiserver/model/clusterModel"
	"github.com/mitchellh/mapstructure"
)

type SendToNodeModel struct {
	Ip   string                 //ip
	Port string                 //port
	Mode string                 //模式 get post
	Url  string                 //url    /sd/health
	Body map[string]interface{} //传递参数
}

func SendToFunc(agreement string, send *SendToNodeModel) (r clusterModel.HealthCheck, err error) {

	switch agreement {
	case sendToNode.HTTP:
		r, err = send.sendHttp()
	case sendToNode.HTTPS:
		r, err = send.sendHttps()
	case sendToNode.RPC:
		r, err = send.sendRpc()
	case sendToNode.TCP:
		r, err = send.sendTcp()
	}
	return
}

func (self *SendToNodeModel) sendHttp() (r clusterModel.HealthCheck, err error) {
	client := &http.Client{}
	url := "http://" + self.Ip + ":" + self.Port + self.Url
	reqest, err := http.NewRequest(self.Mode, url, nil)

	if err != nil {
		log.Infof(err.Error())
		return
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	//处理返回结果
	defer response.Body.Close()
	resp, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Info(err.Error())
		return
	}
	info := gjson.ParseBytes(resp)
	//断定正确
	if info.Get("code").Int() != 10000 {
		log.Info("code is fail")
		return
	}

	//对结构体进行绑定
	dataMp := info.Get("data").Value()
	if err := mapstructure.Decode(dataMp, &r); err != nil {
		log.Info("data is fail")
	}
	return

}

func (self *SendToNodeModel) sendHttps() (r clusterModel.HealthCheck, err error) {
	return
}

func (self *SendToNodeModel) sendRpc() (r clusterModel.HealthCheck, err error) {
	return
}

func (self *SendToNodeModel) sendTcp() (r clusterModel.HealthCheck, err error) {
	return
}
