package routers

import (
	"Go_authentication/authtutorial/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.LoginController{}, "get:RegisterView;post:Register")
	beego.Router("/login", &controllers.LoginController{}, "get:LoginView;post:Login")
	// beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
}