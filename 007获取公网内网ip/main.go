package main

import (
	"fmt"
	"os/exec"
	"strings"
)

//获取IP
func exeSysCommand(cmdStr string) string {
	cmd := exec.Command("sh", "-c", cmdStr)
	opBytes, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(opBytes)
}

func GetLocalIp() string {
	tmp := exeSysCommand("ifconfig eth1 | grep 'inet addr' | cut -d : -f 2 | cut -d ' ' -f 1")
	if len(tmp) == 0 {
		fmt.Println("GetLocalIp Failed")
		return ""
	}

	localip := strings.Trim(tmp, "\n") // 去除尾部换行符
	return localip
}

func main() {
	ip := GetLocalIp()
	fmt.Println("ip=", ip)
}

/*
获取公网ip  :
curl ipinfo.io
curl ip.sb
curl ifconfig.me
*/

/*
获取内网网ip  :
ifconfig
*/
