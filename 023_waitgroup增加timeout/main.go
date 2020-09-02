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

//errorgroup

/*
当超时时间大于任务时间时，任务可以正常完成然后退出，
当超时时间小于任务时间时，任务没有执行完成就退出了。
*/
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
