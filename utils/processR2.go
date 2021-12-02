package utils

import "github.com/shirou/gopsutil/process"

//获取进程
func GetProcess() int {
	pids, _ := process.Pids()
	return int(pids[0])
}
