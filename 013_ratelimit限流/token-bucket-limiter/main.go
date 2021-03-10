package main

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

// 算法描述：用户配置的平均发送速率为r，则每隔1/r秒一个令牌被加入到桶中
// （每秒会有r个令牌放入桶中），桶中最多可以存放b个令牌。如果令牌到达时令牌桶已经满了，
// 那么这个令牌会被丢弃；

var limiter = rate.NewLimiter(10, 1000)

// 第一个参数r:代表每秒可以向token桶中产生多少token。
// 第二个参数b:代表 token 桶的容量大小

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if limiter.Allow() == false {
				http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
				return

			}
			next.ServeHTTP(w, r)
		})
}

func main() {
	// limiter.
	fmt.Println(limiter)
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)
	http.ListenAndServe(":4000", limit(mux))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handler ok"))
}

/* // Wait无可用token会阻塞住，直到获取一个token，或者超时或取消（基于context.Context)。
func wait() {
	// Limiter对象: 控制事件发生的频率。它实现了一个令牌桶。开始的时候为满的，大小为b。然后每秒补充r个令牌。
	// Limiter的默认初始化（Zero Value）是一个有效值，但是会拒绝所有的事件。需要使用NewLimiter来创建实际可用的限速器。
	limiter := rate.NewLimiter(3, 5)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	for i := 0; ; i++ {
		fmt.Printf("%03d %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return
		}
	}
}

// Reserve 无可用token，则返回一个或多个未来token的预订以及调用者在使用前必须等待的时长.
func reserve() {
	limiter := rate.NewLimiter(3, 5)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	for i := 0; ; i++ {
		fmt.Printf("%03d %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		reserve := limiter.Reserve()
		if !reserve.OK() {
			//返回是异常的，不能正常使用
			fmt.Println("Not allowed to act! Did you remember to set lim.burst to be > 0 ?")
			return
		}
		delayD := reserve.Delay()
		fmt.Println("sleep delay ", delayD)
		time.Sleep(delayD)
		select {
		case <-ctx.Done():
			fmt.Println("timeout, quit")
			return
		default:
		}
		//TODO 业务逻辑
	}
}

//Allow 无可用token则返回false
func allow() {
	limiter := rate.NewLimiter(3, 5)
	//n := 0
	m := 0
	for i := 0; i < 50; i++ {
		if limiter.Allow() {
			//n++
			//fmt.Printf("%03d %03d Ok  %s\n", n, i, time.Now().Format("2006-01-02 15:04:05.000"))
		} else {
			m++
			fmt.Printf("%03d %03d Err %s\n", m, i, time.Now().Format("2006-01-02 15:04:05.000"))
		}
		time.Sleep(100 * time.Millisecond)
	}
}
*/
