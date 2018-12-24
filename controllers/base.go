package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Echo(data interface{}) {
	code := 0
	Msg := "操作成功"
	this.Data["json"] = map[string]interface{}{"code": code, "msg": Msg, "data": data}
	this.ServeJSON()
}
func (this *BaseController) EchoError(msg string) {
	beego.Error(msg)
	code := -1
	this.Data["json"] = map[string]interface{}{"code": code, "msg": msg}
	this.ServeJSON()
}
func (this *BaseController) EchoMessage(msg string) {
	code := 0
	this.Data["json"] = map[string]interface{}{"code": code, "msg": msg}
	this.ServeJSON()
}