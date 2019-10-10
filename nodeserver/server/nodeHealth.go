package server

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"nodeserver/model"
	. "nodeserver/enum"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
)

type NodeInterface interface {
	Cpu() model.Cpu
	Ram() model.Ram
	Disk() model.Disk
}

type NodeServer struct {
}

func (NodeServer) Cpu() model.Cpu {
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	l1 := a.Load1   //平均一分钟的负载
	l5 := a.Load5   //平均五分钟的负载
	l15 := a.Load15 //平均十五分钟的负载

	resp := model.Cpu{
		Status: Ok,
		L1:     l1,
		L5:     l5,
		L15:    l15,
		Cores:  cores,
	}

	// cpu 警戒值判定 平均单核cpu 0.7为临界点
	if l5/float64(cores) >= 0.7 {
		resp.Status = CRITICAL
	} else if l5/float64(cores) >= 1 {
		resp.Status = WARNING
	}
	return resp
}

func (NodeServer) Ram() model.Ram {
	u, _ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	resp := model.Ram{
		Status:      Ok,
		UserMb:      usedMB,
		UserGb:      usedGB,
		TotalMB:     totalMB,
		TotalGB:     totalGB,
		UsedPercent: usedPercent,
	}

	if usedPercent >= 95 {
		resp.Status = CRITICAL

	} else if usedPercent >= 90 {
		resp.Status = WARNING
	}
	return resp
}

func (NodeServer) Disk() model.Disk {
	u, _ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	resp := model.Disk{
		Status:      Ok,
		UserMb:      usedMB,
		UserGb:      usedGB,
		TotalMB:     totalMB,
		TotalGB:     totalGB,
		UsedPercent: usedPercent,
	}

	if usedPercent >= 95 {
		resp.Status = CRITICAL

	} else if usedPercent >= 90 {
		resp.Status = WARNING
	}
	return resp
}

//
//func (NodeServer) Network() {
//	nv, _ := net.IOCounters(true)
//}
