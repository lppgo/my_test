/**
 * @Author: lucas
 * @Description: 给sync.WaitGroup封装一个超时，防止go阻塞而一直sync.Wait()
 * @File:  wgTimeOut.go
 * @Version: 1.0.0
 * @Date: 2020/7/6 10:00
 */
package wgTimeOut

import (
	"sync"
	"time"
)

type WaitGroupTimeOut struct{
	wg sync.WaitGroup
	done chan struct{}
	timeout time.Duration
}

// NewWaitGroupTimeOut
func NewWaitGroupTimeOut(timeout time.Duration) *WaitGroupTimeOut{
	return &WaitGroupTimeOut{
		done:    make(chan struct{}),
		timeout: timeout,
	}
}
func(wt *WaitGroupTimeOut) Add(delta int){
	wt.wg.Add(delta)
}

func (wt *WaitGroupTimeOut) Done(){
	wt.wg.Done()
}

func (wt *WaitGroupTimeOut) Wait(){
	go func(){
		wt.wg.Wait()
		close(wt.done)
	}()
	 select {
	case <-wt.done:
	case <-time.After(wt.timeout):
	}
}
