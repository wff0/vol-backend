package global

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

var (
	VB_DB         *gorm.DB
	VB_REDIS_POOL *redis.Pool
)
