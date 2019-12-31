package main

import (
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"user-srv/basic"
	"user-srv/handler"
	"user-srv/model"
	"github.com/micro/cli"
	"user-srv/basic/config"
	user "user-srv/proto/user"
	"github.com/micro/go-micro/registry"
	"fmt"
)

func main() {
	// 初始化
	//配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
			micro.Action(func(context *cli.Context) {
				model.Init()
				handler.Init()
			}),
		)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
