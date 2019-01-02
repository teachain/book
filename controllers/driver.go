package controllers

import (
	"encoding/json"
	"book/models"
)

type DriverController struct {
	BaseController
}

func (this *DriverController) Add() {
	var params models.Driver
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	err = models.AddDriver(&params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	this.EchoMessage("add driver success")
}

func (this *DriverController) ListAll() {
	all, err := models.GetAllDrivers()
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	this.Echo(all)
}
