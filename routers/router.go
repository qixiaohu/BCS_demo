package routers

import (
	"bcsdemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/test",&controllers.MainController{},"get:GetValue")
	beego.Router("/test",&controllers.MainController{},"post:SetValue")
}
