package clusterModel

type HealthCheck struct {
	Cpu  Cpu  `json:"cpu"`
	Ram  Ram  `json:"ram"`
	Disk Disk `json:"disk"`
}

type Cpu struct {
	Status string  `json:"status"`
	L1     float64 `json:"l1"`
	L5     float64 `json:"l2"`
	L15    float64 `json:"l15"`
	Cores  int     `json:"cores"`
}

type Ram struct {
	Status      string `json:"status"`
	UserMb      int    `json:"userMb"`
	UserGb      int    `json:"userGb"`
	TotalMB     int    `json:"totalMb"`
	TotalGB     int    `json:"totalGb"`
	UsedPercent int    `json:"usedPercent"`
}

type Disk struct {
	Status      string `json:"status"`
	UserMb      int    `json:"userMb"`
	UserGb      int    `json:"userGb"`
	TotalMB     int    `json:"totalMb"`
	TotalGB     int    `json:"totalGb"`
	UsedPercent int    `json:"usedPercent"`
}
