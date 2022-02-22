package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func main() {
	DiskCheck()
	OSCheck()
	CPUCheck()
	RAMCheck()
}

//服务器硬盘使用量
func DiskCheck() {
	fmt.Println()
	fmt.Println("----------------服务器硬盘使用量:----------------")
	u, _ := disk.Usage("/")
	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)
	fmt.Printf("Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%\n", usedMB, usedGB, totalMB, totalGB, usedPercent)
}

//OS
func OSCheck() {
	fmt.Println()
	fmt.Println("----------------OS:----------------")
	fmt.Printf("goOs:%s,compiler:%s,numCpu:%d,version:%s,numGoroutine:%d\n", runtime.GOOS, runtime.Compiler, runtime.NumCPU(), runtime.Version(), runtime.NumGoroutine())
}

//CPU 使用量
func CPUCheck() {
	fmt.Println()
	fmt.Println("----------------CPUCheck:----------------")
	cores, _ := cpu.Counts(false)

	cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true)
	if err == nil {
		for i, c := range cpus {
			fmt.Printf("cpu%d : %f%%\n", i, c)
		}
	}
	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15
	fmt.Println(l1)
	fmt.Println(l5)
	fmt.Println(l15)
	fmt.Println(cores)
}

//内存使用量
func RAMCheck() {
	fmt.Println()
	fmt.Println("----------------RAMCheck:----------------")
	u, _ := mem.VirtualMemory()
	usedMB := int(u.Used) / MB
	totalMB := int(u.Total) / MB
	usedPercent := int(u.UsedPercent)
	fmt.Printf("used(MB):%d, total(MB):%d, usedPercent:%d\n", usedMB, totalMB, usedPercent)
}
