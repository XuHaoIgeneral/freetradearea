package dockerEnum

import "time"

const (
	Version    = "1.39"
	DockerPort = "5251"
	DockerStopTime=time.Duration(200) * time.Millisecond
)
