package main

import (
    "sync"
    "time"
)
// 滑动窗口限流
type RateLimiter struct {
    mu       sync.Mutex
    requests map[string][]time.Time
    limit    int
    window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
    return &RateLimiter{
        requests: make(map[string][]time.Time),
        limit:    limit,
        window:   window,
    }
}

func (rl *RateLimiter) Allow(userID string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()

    now := time.Now()
    windowStart := now.Add(-rl.window)

    // 清理过期的请求记录
    validRequests := make([]time.Time, 0)
    for _, t := range rl.requests[userID] {
        if t.After(windowStart) {
            validRequests = append(validRequests, t)
        }
    }

    // 检查是否超过限制
    iflen(validRequests) >= rl.limit {
        returnfalse
    }

    // 添加新请求
    validRequests = append(validRequests, now)
    rl.requests[userID] = validRequests

    // 定期清理（可选，防止内存泄漏）
    iflen(validRequests) == 1 {
        go rl.cleanup(userID, now)
    }

    returntrue
}

func (rl *RateLimiter) cleanup(userID string, baseTime time.Time) {
    time.Sleep(rl.window)
    rl.mu.Lock()
    defer rl.mu.Unlock()

    if requests, exists := rl.requests[userID]; exists {
        windowStart := baseTime.Add(-rl.window)
        validRequests := make([]time.Time, 0)
        for _, t := range requests {
            if t.After(windowStart) {
                validRequests = append(validRequests, t)
            }
        }
        iflen(validRequests) == 0 {
            delete(rl.requests, userID)
        } else {
            rl.requests[userID] = validRequests
        }
    }
}
