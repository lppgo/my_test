package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// 使用计数器实现请求限流

// 限流的要求是在指定的时间间隔内，server 最多只能服务指定数量的请求。
// 实现的原理是我们启动一个计数器，每次服务请求会把计数器加一，同时到达指定的时间间隔后会把计数器清零
type NetWorkRequestLimitService struct {
	RequestInterval time.Duration // 指定间隔时间
	MaxCount        int           // 最大请求访问次数
	LockRequest     sync.Mutex    // 加锁
	ReqCount        int           // 当前请求访问的次数
	ReqIp           string        // 当前发起请求的Ip
}

var (
	netWorkRequestLimit          *NetWorkRequestLimitService
	doNetWorkRequestLimitService *NetWorkRequestLimitService
)

func init() {
	// 在服务请求的时候, 我们会对当前计数器和阈值进行比较，只有未超过阈值时才进行服务：
	doNetWorkRequestLimitService = DoNetWorkRequestLimitService(10*time.Second, 5)
}

func DoNetWorkRequestLimitService(interval time.Duration, maxCnt int) *NetWorkRequestLimitService {
	netWorkRequestLimit = &NetWorkRequestLimitService{
		RequestInterval: interval,
		MaxCount:        maxCnt,
		//ReqIp:reqIp,//当前发起请求的Ip是计算出来的
	}

	go func() {
		ticker := time.NewTicker(interval)

		for {
			// 每指定的时间间隔检查一次是否达到最大请求数
			<-ticker.C
			netWorkRequestLimit.LockRequest.Lock()
			fmt.Println("Reset Count...")

			netWorkRequestLimit.ReqCount = 0
			netWorkRequestLimit.LockRequest.Unlock()
		}
	}()

	return netWorkRequestLimit
}

// ReqIncreased在时间间隔内，每次请求计数器+1.
func (ls *NetWorkRequestLimitService) ReqIncreased() {
	ls.LockRequest.Lock()
	defer ls.LockRequest.Unlock()
	ls.ReqCount += 1
}

// ReqIsAvailabled判断是否达到最大请求数.
func (ls *NetWorkRequestLimitService) ReqIsAvailabled() bool {
	ls.LockRequest.Lock()
	defer ls.LockRequest.Unlock()

	return ls.ReqCount < ls.MaxCount
}

func ServeHTTPW(w http.ResponseWriter, r *http.Request) {
	// 判断是否达到访问设置的最大数量
	if doNetWorkRequestLimitService.ReqIsAvailabled() {
		fmt.Println("doNetWorkRequestLimitService:", doNetWorkRequestLimitService)
		// 访问一次让请求计数加1
		doNetWorkRequestLimitService.ReqIncreased()
		fmt.Println(doNetWorkRequestLimitService.ReqCount)
		fmt.Fprintln(w, fmt.Sprintf("当前请求：%d\n", doNetWorkRequestLimitService.ReqCount))
		return
	}

	fmt.Println("Reach request limiting!")
	fmt.Fprintln(w, "Reach request limit!")
}

func main() {
	http.HandleFunc("/robotReq", ServeHTTPW)
	http.ListenAndServe(":8080", nil)
}
