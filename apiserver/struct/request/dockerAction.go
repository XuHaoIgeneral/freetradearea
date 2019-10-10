package request

type DockerRun struct {
	ImageName     string   `json:"image_name"`     // 镜像名称
	NetworkName   string   `json:"network_name"`   // 选定网络名称
	ContainerName string   `json:"container_name"` // 容器命名
	ContainerIp   string   `json:"container_ip"`   // 容器自选ip
	HostIp        string   `json:"host_ip"`        // 指定宿主机ip
	ExternalPort  string   `json:"external_port"`  // 对外暴露端口
	InnerPort     string   `json:"inner_port"`     // 内部端口
	Cmd           []string `json:"cmd"`            // cmd 命令
	Auto          bool     `json:"auto"`
}

type DockerIp struct {
	Ip          string `json:"ip"`
	ContainerId string `json:"container_id"`
}

type NetworkCreate struct {
	Networkname string `json:"networkname"`
	Subnet      string `json:"subnet"`
	Gateway     string `json:"gateway"`
}

type NetworkDel struct {
	NetworkId string `json:"networkid"`
}
