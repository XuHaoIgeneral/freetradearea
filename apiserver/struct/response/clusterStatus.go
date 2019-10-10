package response

import "apiserver/model/clusterModel"

type ClusterStatusResp struct {
	Ip   string                     `json:"ip"`
	Name string                     `json:"name"`
	List clusterModel.HealthCheck `json:"list"`
}
