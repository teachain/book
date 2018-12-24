package controllers

import (
	"book/models"
	"encoding/json"
)

type ModelController struct {
	BaseController
}

func (this *ModelController) Add() {
	var params models.Model
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	err=models.AddModel(&params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	this.EchoMessage("add model success")
}

func (this *ModelController) ListAll() {
	all,err:=models.GetAllModels()
	if err!=nil{
		this.EchoError(err.Error())
		return
	}
	this.Echo(all)
}
