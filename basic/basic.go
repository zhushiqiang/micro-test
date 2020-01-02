package basic

import (
	"github.com/zhushiqiang/micro-test/basic/config"
	"github.com/zhushiqiang/micro-test/basic/db"
	"github.com/zhushiqiang/micro-test/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}