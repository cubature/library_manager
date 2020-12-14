package services

import (
	etcdv3 "github.com/micro/go-plugins/registry/etcdv3/v2"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"library_manager/protos"
)

func init() {
	// 用的etcd 做为服务发现
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})

	// 初始化服务
	service := micro.NewService(
		micro.Name("library"),
		micro.Registry(reg),
	)

	service.Init()

	// 注册 Handler
	protos.RegisterBookServiceHandler(service.Server(), new(BookService))

	// run server
	go func() {
		if err := service.Run(); err != nil {
			panic(err)
		}
	}()
}
