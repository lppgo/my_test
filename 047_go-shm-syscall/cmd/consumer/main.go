package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

const (
	// IpcCreate create if key is nonexistent
	IpcCreate = 00001000
)

// var mode = flag.Int("mode", 0, "0:write 1:read")

func main() {
	// flag.Parse()
	// fmt.Println("mode : ", *mode)

	shmid, _, err := syscall.Syscall(syscall.SYS_SHMGET, 2, 4, IpcCreate|0600)
	if err != 0 {
		fmt.Printf("syscall error, err: %v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("shmid: %v\n", shmid)

	shmaddr, _, err := syscall.Syscall(syscall.SYS_SHMAT, shmid, 0, 0)
	if err != 0 {
		fmt.Printf("syscall error, err: %v\n", err)
		os.Exit(-2)
	}
	fmt.Printf("shmaddr: %v\n", shmaddr)

	defer syscall.Syscall(syscall.SYS_SHMDT, shmaddr, 0, 0)

	fmt.Println("read mode")
	for {
		fmt.Println(*(*int)(unsafe.Pointer(uintptr(shmaddr))))
		time.Sleep(1 * time.Second)
	}
}
