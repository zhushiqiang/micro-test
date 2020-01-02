package basic

import (
	"user-web/basic/config"
	"user-web/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}