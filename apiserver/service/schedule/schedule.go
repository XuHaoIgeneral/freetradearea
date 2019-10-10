package schedule

import (
	"apiserver/service/clusterServer/clusterStatus"
	. "apiserver/enum/sendToNode"
	. "apiserver/enum/getToNode"
	"apiserver/model/clusterModel"
	"fmt"
)

var DefaultSchedule Schedule

//选择调度算法
type Schedule struct {
}

type score struct {
	Ip  string
	Num int
}

//获取到所有可用节点信息
func (Schedule) Obtain() string{
	//默认http处理
	list := clusterStatus.NodeStatus(HTTP, PortHttp, GET, URL)
	fmt.Println(list)
	temp := score{
		Ip:  "192.168.56.200",
		Num: 0,
	}

	for _, v := range list {
		if v.List.Cpu.Status == WARNING || v.List.Cpu.Status == CRITICAL {
			continue
		} else if v.List.Disk.Status == WARNING || v.List.Disk.Status == CRITICAL {
			continue
		} else if v.List.Ram.Status == WARNING || v.List.Ram.Status == CRITICAL {
			continue
		}

		//计算得分
		tempV := v.List
		rangeTemp := achievement(tempV)
		if rangeTemp > temp.Num {
			temp.Ip = v.Ip
			temp.Num = rangeTemp
		}
	}

	return temp.Ip
}

//评分设定
// cpu  0.0~0.25==3  0.25~0.5==2 0.5~0.7=1 other=0
// disk
func achievement(temp clusterModel.HealthCheck) int {

	var allNums int
	cpuNums := temp.Cpu.L5 / float64(temp.Cpu.Cores)
	diskNums := float64(0.7 - float64(temp.Disk.UsedPercent)*0.01)
	ramsNums := float64(0.7 - float64(temp.Ram.UsedPercent)*0.01)

	//cpu 断定
	if cpuNums <= 0.25 {
		allNums = allNums + 3
	} else if cpuNums > 0.25 && cpuNums <= 0.5 {
		allNums = allNums + 2
	} else if cpuNums > 0.5 && cpuNums <= 0.7 {
		allNums = allNums + 1
	}

	allNums = int(ramsNums + diskNums)
	fmt.Println("得分",allNums)
	return allNums
}
