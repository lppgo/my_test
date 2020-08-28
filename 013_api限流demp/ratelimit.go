//这是ratelimit 限流方式
package ratelime
import (
    "go.uber.org/ratelimit"
    "sync"
)

var doorLimitMap sync.Map // (全局变量)
//返回键的现有值(如果存在)，否则存储并返回给定的值，如果是读取则返回true，如果是存储返回false。

func foo(){
r, _ := doorLimitMap.LoadOrStore(key, ratelimit.New(1))
r.(ratelimit.Limiter).Take()

// next  访问接口

}
