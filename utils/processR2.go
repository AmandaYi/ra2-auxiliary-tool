package utils

import (
	"github.com/shirou/gopsutil/process"
	"syscall"
	"unsafe"
)

var MAXIMUM_ALLOWED = 0x2000000
var PROCESS_ALL_ACCESS = (0x000F0000 | 0x00100000 | 0xFFF)

//获取进程
func GetProcess(n string) int32 {
	pids, _ := process.Pids()
	for _, pid := range pids {
		p, _ := process.NewProcess(pid)
		name, _ := p.Name()
		if n == name {
			return pid
		}
	}
	return -1
}

func OpenProcess(bInheritHandle bool, dwProcesssId uint32) (uintptr, error) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	openProcess := kernel32.NewProc("OpenProcess")
	pHandle, _, err := openProcess.Call(uintptr(MAXIMUM_ALLOWED), uintptr(unsafe.Pointer(&bInheritHandle)), uintptr(dwProcesssId))
	return pHandle, err
}
