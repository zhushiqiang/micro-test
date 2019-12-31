package basic

import (
	"user-srv/basic/config"
	"user-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}