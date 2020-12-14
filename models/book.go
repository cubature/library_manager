package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Id     int64  `orm:"pk;auto;column(id)"`
	Name   string 
	Author string `orm:"size(50)"`
}

func (a *Book) Insert() error {
	if _, err := orm.NewOrm().Insert(a); err != nil {
		return err
	}
	return nil
}

func (a *Book) Read(fields ...string) error {
	if err := orm.NewOrm().Read(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *Book) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *Book) Delete() error {
	if _, err := orm.NewOrm().Delete(a); err != nil {
		return err
	}
	return nil
}

func (a *Book) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(a)
}
