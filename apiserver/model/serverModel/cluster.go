package serverModel

type ClusterNode struct {
	Id     uint64 `json:"id"`
	Ip     string `json:"ip"`
	Status bool   `json:"status"`
	Name   string `json:"name"`
}

type ClusterInfo struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name"`
	ClientUrls string `json:"client_urls"`
}

