package routers

import (
	"book/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/device/add", &controllers.DeviceController{}, "post:Add")
	beego.Router("/device/list", &controllers.DeviceController{}, "get:ListAll")
	beego.Router("/driver/add", &controllers.DriverController{}, "post:Add")
	beego.Router("/driver/list", &controllers.DriverController{}, "get:ListAll")
	beego.Router("/model/add", &controllers.ModelController{}, "post:Add")
	beego.Router("/model/list", &controllers.ModelController{}, "get:ListAll")
	beego.Router("/record/add", &controllers.RecordController{}, "post:Add")
	beego.Router("/record/list", &controllers.RecordController{}, "get:ListAll")

	beego.Router("/user/login", &controllers.UserController{}, "post:Login")

	beego.Router("/gas/add", &controllers.GasController{}, "post:Add")
	beego.Router("/gas/list", &controllers.GasController{}, "get:ListAll")

}
