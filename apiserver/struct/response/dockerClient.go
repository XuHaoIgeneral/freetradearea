package response

type RespDockerPs struct {
	Id           string      `json:"id"`
	Name         interface{} `json:"name"`
	Image        string      `json:"image"`
	Create       int64       `json:"create"`
	State        string      `json:"state"`
	NetworkMode  string      `json:"network_mode"`
	Ports        interface{} `json:"ports"`
	HostNodeName string      `json:"host_node_name"`
}

type ports struct {
	IP          string `json:"ip"`
	PrivatePort int    `json:"private_port"`
	PublicPort  int    `json:"public_port"`
	Type        string `json:"type"`
}
