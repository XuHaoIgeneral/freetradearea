package dockerModel

type DockerPsModel struct {
	HostName string `json:"host_name"`
	HostIp   string `json:"host_ip"`
}

type DockerNetworlList struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Driver string `json:"driver"`
	Scope  string `json:"scope"`
}

type DockerImageList struct {
	Id string `json:"id"`
}
