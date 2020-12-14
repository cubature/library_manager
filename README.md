# Library_manager

## 主要工具
```
Golang 1.15
Beego 1.12.3
Go-Micro.v2 2.9.1
```

## 主要结构
```
└─library_manager
    ├─ routers     // 对外的Restful接口声明，路由到对应控制器
    |
    ├─ controllers // 服务调用控制器
    |
    ├─ protos      // RPC通信的数据序列化定义
    |
    ├─ services    // 服务的业务逻辑处理，建立服务注册中心
    |
    ├─ models      // ORM、数据库处理
    ...
```

controllers与services之间通过Go-Micro微服务框架进行服务的调用。
services中实现服务的业务逻辑，配置及建立使用etcd的微服务rpc服务器，将服务注册到注册中心。Go-Micro服务之间使用Protobuf进行RPC通信。
models中使用了beego自带的orm进行数据库连接及设置连接池等。

### 后续思路
当前初步搭建了使用Beego的MVC及Go-Micro的微服务的应用框架，功能上只实现了从数据库中读取所有图书信息，后续开发思路有：
- 基础功能方面：
    - 加入管理（admin），在controllers.baseControllers.Prepare()中加入授权验证，在services中加入对应登录服务以及管理对书本的CRUD管理。
    - 加入用户，借书还书之类的功能。
- 项目结构方面：
    - 将Go-Micro注册中心部分单独抽出成为一个独立子项目，从而能够分开运行对外的网关服务器及对内的RPC服务器，大致将项目分为：
    ```
    └─library_manager
        ├─ gateway
        |   ├─ routers     // 对外的Restful接口声明，路由到对应控制器
        |   |
        |   ├─ controllers // 服务调用控制器
        |   |
        |   └─ main.go     // web服务器入口（启动web服务）
        |
        ├─ protos          // RPC通信的数据序列化定义, gateway及rpc_server公共部分
        |
        └─ rpc_server
            |
            ├─ services    // 服务的业务逻辑处理，建立服务注册中心
            |
            ├─ models      // ORM、数据库处理
            |
            └─ main.go     // rpc服务器入口（启动rpc服务器）
    ```
- 其他：
    - 加入入口反向代理、缓存(Redis)、log、test等
