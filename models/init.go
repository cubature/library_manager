package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

func init() {
	mysqlhost := beego.AppConfig.String("mysqlhost")
    mysqlport := beego.AppConfig.String("mysqlport")
    mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqldb := beego.AppConfig.String("mysqldb")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	conn := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlhost + ":" + mysqlport + ")/" + mysqldb + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn)

	orm.RegisterModel(new(Book))

	// 自动创建表、更新表
	//orm.RunSyncdb("default", true, true)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	// 连接池
	orm.SetMaxIdleConns("default",1000)
	orm.SetMaxOpenConns("default",2000)

	fmt.Printf("数据库连接成功 %s\n", conn)
}