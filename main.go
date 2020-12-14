package main

import (
	_ "library_manager/models"
	_ "library_manager/routers"
	_ "library_manager/controllers"
	_ "library_manager/services"
	"github.com/astaxie/beego"	

)

func main() {
	beego.Run()
}

