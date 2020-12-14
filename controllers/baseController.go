package controllers

import (
	"github.com/astaxie/beego"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	etcdv3 "github.com/micro/go-plugins/registry/etcdv3/v2"
	"library_manager/protos"
)

var bookServiceClent protos.BookService

type baseController struct {
	beego.Controller
}

func (this *baseController) Prepare() {}

func init() {
	// etcd做服务发现
    reg := etcdv3.NewRegistry(func(op *registry.Options){
        op.Addrs = []string{
            "http://127.0.0.1:2379",
        }
    })

    // 初始化服务
    service := micro.NewService(
        micro.Registry(reg),
    )
    service.Init()
    bookServiceClent = protos.NewBookService("library", service.Client())
}
