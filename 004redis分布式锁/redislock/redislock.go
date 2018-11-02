package redislock

import (
	"errors"

	"github.com/garyburd/redigo/redis"
)

type RedisLock struct {
	LockKey string
	Value   string
}

//保证原子性（redis是单线程），避免del删除了，其他client获得的lock
var delScript = redis.NewScript(1, `
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end`)

// func (this *RedisLock) Lock(rd *redis.Conn, timeout int) error {
// 	res, err := redis.String((*rd).Do("GET", this.LockKey))
// 	if res != "" {
// 		return errors.New("lock exists!")
// 	} else {
// 		result, err2 := redis.String((*rd).Do("SET", this.LockKey, this.Value, "EX", "5"))
// 		if result == "OK" && err2 == nil {
// 			return nil
// 		} else {
// 			err = errors.New("SETEX : " + this.LockKey + " failed ")
// 		}
// 	}
// 	return err
// }

func (this *RedisLock) Lock(rd *redis.Conn, timeout int) error {
	result, err := redis.String((*rd).Do("SET", this.LockKey, this.Value, "EX", timeout, "NX"))
	if err != nil {
		return errors.New("redis fail")
	}
	if result == "OK" {
		return nil
	} else {
		return errors.New("lock fail")
	}
}

func (this *RedisLock) Unlock(rd *redis.Conn) {
	delScript.Do(*rd, this.LockKey, this.Value)
}
