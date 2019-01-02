package controllers

import (
	"encoding/json"
	"book/models"
)

type MaintainController struct {
	BaseController
}

func (this *MaintainController) Add() {
	var params models.Maintain
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	err = models.AddMaintain(&params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	this.EchoMessage("add success")
}

func (this *MaintainController) ListAll() {
	all, err := models.GetAllMaintain()
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	this.Echo(all)
}
