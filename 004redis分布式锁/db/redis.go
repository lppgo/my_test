package db

import (
	"github.com/garyburd/redigo/redis"
)

func GetConn(addr, password string, db ...int) (conn redis.Conn) {
	var err error
	opts := make([]redis.DialOption, 0)
	if password != "" {
		opts = append(opts, redis.DialPassword(password))
	}
	if len(db) != 0 {
		opts = append(opts, redis.DialDatabase(db[0]))
	}
	conn, err = redis.Dial("tcp4", addr, opts...)
	if err != nil {
		panic(err)
	}
	return
}
