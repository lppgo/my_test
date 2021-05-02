- [1:B站微服务框架自适应限流模块分析](#1b站微服务框架自适应限流模块分析)
- [2: 限流接口](#2-限流接口)
- [3: aqm](#3-aqm)
- [4: Vegas](#4-vegas)

# 1:B站微服务框架自适应限流模块分析
微服务中限流模块是必不可少的，理想的情况是使系统维持在能承受的负载范围内，没有堆积请求，请求处理时间就成了一个很好地判断指标，如果请求处理的时间过长，说明发生了堆积。传统的做法可能是硬性限制qps或队列长度或者根据cpu使用率等指标限制，无法实现自适应限流，根据请求处理时间使用算法限流可以很好地自适应系统的变化。

# 2: 限流接口
```go
// Limiter limit interface.
type Limiter interface {
	Allow(ctx context.Context) (func(Op), error)
}

// Limiter use tcp vegas + codel for adaptive limit.
type Limiter struct {
	rate  *vegas.Vegas
	queue *aqm.Queue
}

// Allow immplemnet rate.Limiter.
// if error is returned,no need to call done()
func (l *Limiter) Allow(ctx context.Context) (func(rate.Op), error) {
	var (
		done func(time.Time, rate.Op)
		err  error
		ok   bool
	)
	if done, ok = l.rate.Acquire(); !ok {
		// NOTE exceed max inflight, use queue
		if err = l.queue.Push(ctx); err != nil {
			done(time.Time{}, rate.Ignore)
			return func(rate.Op) {}, err
		}
	}
	start := time.Now()
	return func(op rate.Op) {
		done(start, op)
		l.queue.Pop()
	}, nil
}
```
接口很简单，就一个Allow函数，它会返回一个done func，需要我们执行完请求后调用。可以看到具体的实现有两个组件，Vegas和aqm队列，首先使用Vegas，不成功放入aqm队列。

# 3: aqm
`主动队列管理（Active Queue Management，AQM）`，网络包传输的一种拥塞控制机制，使得在路由器缓存被耗尽前有计划的丢掉一部分分组，Codel是一种延时控制算法，通过检测排队延时来判断拥塞情况，可以自适应的调整丢包，这里有一段代码实现。

B站的go版本实现：
```go
type Config struct {
	Target   int64 // target queue delay (default 20 ms). 能容忍的最大排队延时
	Internal int64 // sliding minimum time window width (default 500 ms)  窗口大小
}

type packet struct {
	ch chan bool
	ts int64
}

// Queue queue is CoDel req buffer queue.
type Queue struct {
	pool    sync.Pool
	packets chan packet

	mux      sync.RWMutex
	conf     *Config
	count    int64
	dropping bool  // 	Equal to 1 if in drop state
	faTime   int64 //  丢弃状态触发启动时间
	dropNext int64 // 下一次开始丢弃的时间
}

// Push req into CoDel request buffer queue.
// if return error is nil,the caller must call q.Done() after finish request handling
func (q *Queue) Push(ctx context.Context) (err error) {
	r := packet{
		ch: q.pool.Get().(chan bool),
		ts: time.Now().UnixNano() / int64(time.Millisecond),
	}
	select {
	case q.packets <- r:
	default:
		err = ecode.LimitExceed //channel满直接丢弃
		q.pool.Put(r.ch)
	}
	if err == nil {
		select {
		case drop := <-r.ch: //阻塞等待pop判断是否丢弃
			if drop {
				err = ecode.LimitExceed
			}
			q.pool.Put(r.ch)
		case <-ctx.Done():
			err = ecode.Deadline
		}
	}
	return
}

// Pop req from CoDel request buffer queue.
func (q *Queue) Pop() {
	for {
		select {
		case p := <-q.packets:
			drop := q.judge(p) //如果判断为丢弃，循环检测
			select {
			case p.ch <- drop:
				if !drop {
					return
				}
			default:
				q.pool.Put(p.ch)
			}
		default:
			return
		}
	}
}

func (q *Queue) controlLaw(now int64) int64 {
	q.dropNext = now + int64(float64(q.conf.Internal)/math.Sqrt(float64(q.count))) //根据当前丢弃周期内的丢包数量计算下一次丢弃的时间，丢弃的数量越多，丢弃间隔越短。
	return q.dropNext
}

// judge decide if the packet should drop or not.
func (q *Queue) judge(p packet) (drop bool) {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	sojurn := now - p.ts
	q.mux.Lock()
	defer q.mux.Unlock()
	if sojurn < q.conf.Target { //延时小于容忍时间，停止丢弃
		q.faTime = 0
	} else if q.faTime == 0 { //延时较大后容忍一个窗口期
		q.faTime = now + q.conf.Internal
	} else if now >= q.faTime { //延时较大维持一个窗口期后开始进行是否进入丢弃状态的判读
		drop = true  //这里直接将返回值设为true，也就是说容忍一个窗口期后直接开始丢弃，而原始算法实现中丢弃只在状态下发生，如果最近没有发生丢弃应该继续容忍一个一个窗口期再开始丢弃的。
	}
	if q.dropping { //丢弃状态下延时较小后直接停止
		if !drop {
			// sojourn time below target - leave dropping state
			q.dropping = false
		} else if now > q.dropNext {
			q.count++
			q.dropNext = q.controlLaw(q.dropNext)
			drop = true
			return
		}   //延时较大维持一个窗口期后如果距上一次计算出的丢弃时间在一个窗口期内直接进入丢弃状态，如果一个窗口期内没有丢弃则再容忍一个窗口期
	} else if drop && (now-q.dropNext < q.conf.Internal || now-q.faTime >= q.conf.Internal) {
		q.dropping = true
		// If we're in a drop cycle, the drop rate that controlled the queue
		// on the last cycle is a good starting point to control it now.
		if now-q.dropNext < q.conf.Internal { //如果距上一次计算出的丢弃时间在一个窗口期内直接沿用之前的控制策略，否则从1开始重新调整
			if q.count > 2 {
				q.count = q.count - 2
			} else {
				q.count = 1
			}
		} else {
			q.count = 1
		}
		q.dropNext = q.controlLaw(now)
		drop = true
		return
	}
	return
}
```
`codel算法`的大致思想就是根据延时情况判断丢弃策略，在限流中就是请求的处理时间，算法只需要两个参数，一个是容忍的最大时延，一个是窗口期，当超过容忍时延时，进入一个窗口期的容忍阶段，此阶段不开始丢弃，这样容忍了短时间的突发流量，在超过容忍阶段后如果最近发生过丢弃则进入丢弃状态，否则再容忍一个窗口期，在丢弃状态下只要时延小于容忍时延就停止丢弃，丢弃状态下随着丢弃数量的不断增加丢弃间隔会缩短，也就是丢弃的越来越频繁。
# 4: Vegas
Vegas是TCP中使用的基于时延的拥塞控制算法，通过比较实际吞吐量和期望吞吐量来调节拥塞窗口的大小。

- 当前能处理的请求数量通过的最请求limit（拥塞窗口尺寸）
- 请求最短处理时间 minRTT
- 请求实际处理时间 lastRTT
- 期望通过速率 Expect = limit/minRTT
- 实际通过速率 Actual = limit/lastRTT
- 速率差值 Diff = Expect-Actual
- Diff = queue/minRTT = limit/minRTT - limit/lastRTT = limit(1/minRTT - 1/lastRTT)

最后推导出 queueSize = limit * (1 - minRTT/lastRTT) 也就是等待中的请求大小是queueSize，我们根据queueSize的大小来调整limit

```go
// Vegas tcp vegas.
type Vegas struct {
	limit      int64
	inFlight   int64 //处理中的请求数
	updateTime int64
	minRTT     int64

	sample atomic.Value
	mu     sync.Mutex
	probes int64
}

//记录一个时间段内的统计信息
type sample struct {
	count       int64
	maxInFlight int64
	drop        int64
	// nanoseconds
	totalRTT int64
}

// Acquire No matter success or not,done() must be called at last.
func (v *Vegas) Acquire() (done func(time.Time, rate.Op), success bool) {
	inFlight := atomic.AddInt64(&v.inFlight, 1)
	if inFlight <= atomic.LoadInt64(&v.limit) {
		success = true
	}

	return func(start time.Time, op rate.Op) { //请求处理完后调用
		atomic.AddInt64(&v.inFlight, -1)
		if op == rate.Ignore {
			return
		}
		end := time.Now().UnixNano()
		rtt := end - start.UnixNano()

		s := v.sample.Load().(*sample)
		if op == rate.Drop {
			s.Add(rtt, inFlight, true)
		} else if op == rate.Success {
			s.Add(rtt, inFlight, false)
		}
		if end > atomic.LoadInt64(&v.updateTime) && s.Count() >= 16 {
			v.mu.Lock()
			defer v.mu.Unlock()
			if v.sample.Load().(*sample) != s { //其它协程已经重置sample，进入下一个窗口期，直接返回
				return
			}
			v.sample.Store(&sample{}) //重置sample

			lastRTT := s.RTT() //统计窗口期的平均处理时间
			if lastRTT <= 0 {
				return
			}
			updateTime := end + lastRTT*5 //设置当前窗口期大小
			if lastRTT*5 < _minWindowTime {
				updateTime = end + _minWindowTime
			} else if lastRTT*5 > _maxWindowTime {
				updateTime = end + _maxWindowTime
			}
			atomic.StoreInt64(&v.updateTime, updateTime)
			limit := atomic.LoadInt64(&v.limit)
			queue := float64(limit) * (1 - float64(v.minRTT)/float64(lastRTT)) //估算等待中的请求数量
			v.probes-- //每隔一段时间探测minRTT，如果窗口期的流量与限制差距较大则说明minRTT不准确，更新为lastRTT
			if v.probes <= 0 {
				maxFlight := s.MaxInFlight()
				if maxFlight*2 < v.limit || maxFlight <= _minLimit {
					v.probes = 3*limit + rand.Int63n(3*limit)
					v.minRTT = lastRTT
				}
			}
			if v.minRTT == 0 || lastRTT < v.minRTT {
				v.minRTT = lastRTT
			}
			var newLimit float64
			threshold := math.Sqrt(float64(limit)) / 2
			if s.Drop() {
				newLimit = float64(limit) - threshold
			} else if s.MaxInFlight()*2 < v.limit {
				return
			} else {
				if queue < threshold { //扩缩策略
					newLimit = float64(limit) + 6*threshold
				} else if queue < 2*threshold {
					newLimit = float64(limit) + 3*threshold
				} else if queue < 3*threshold {
					newLimit = float64(limit) + threshold
				} else if queue > 6*threshold {
					newLimit = float64(limit) - threshold
				} else {
					return
				}
			}
			newLimit = math.Max(_minLimit, math.Min(_maxLimit, newLimit))
			atomic.StoreInt64(&v.limit, int64(newLimit))
		}
	}, success
}
```
总结一下，vegas算法根据统计出的最小时延和窗口期的平均时延作比较估算出等待中的请求数量，调整限制大小
