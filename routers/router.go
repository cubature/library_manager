package routers

import (
	"library_manager/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.WelcomeController{}, "*:ListAll")
	beego.Router("/book", &controllers.WelcomeController{}, "*:ListAll")
}
