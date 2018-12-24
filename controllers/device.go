package controllers

import (
	"book/models"
	"encoding/json"
)

type  DeviceController struct {
BaseController
}

func (this *DeviceController) Add() {
	var params models.Device
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	err=models.AddDevice(&params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	this.EchoMessage("add success")
}

func (this *DeviceController) ListAll() {
   all,err:=models.GetAllDevices()
   if err!=nil{
	   this.EchoError(err.Error())
	   return
   }
   this.Echo(all)
}