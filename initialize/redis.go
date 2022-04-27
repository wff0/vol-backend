package initialize

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"vol-backend/global"
)

func Redis() {
	// 建立连接池
	global.VB_REDIS_POOL = &redis.Pool{
		MaxIdle:     2,
		MaxActive:   3,
		IdleTimeout: 200 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", "",
				redis.DialPassword("WFFxhDX520"),
				redis.DialDatabase(0),
				redis.DialConnectTimeout(2*time.Second),
				redis.DialReadTimeout(2*time.Second),
				redis.DialWriteTimeout(2*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}
