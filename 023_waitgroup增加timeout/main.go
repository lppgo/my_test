/**
 * @Author: lucas
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2020/7/6 9:59
 */
package main

import (
	"023wt/wgTimeOut"
	"fmt"
	"time"
)

// 测试waitgroupTimeout
func main(){
	const TIMEOUT=10*time.Second
	const TASK_EXECUTE_TIME	=3*time.Second

	wt:=wgTimeOut.NewWaitGroupTimeOut(TIMEOUT)
	fmt.Println(wt)

	wt.Add(1)
	go func(){
		defer wt.Done()
		fmt.Println("start")
		time.Sleep(TASK_EXECUTE_TIME)
		fmt.Println("end")
	}()

	fmt.Println("wait exit")
	wt.Wait()
	fmt.Println("exit")
}
