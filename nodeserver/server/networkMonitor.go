package server

import (
	"github.com/shirou/gopsutil/net"
	"fmt"
)

func NetworkMonitor() {
	nv, _ := net.IOCounters(true)
	fmt.Println(nv)
}
