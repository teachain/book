package controllers

import (
	"book/models"
	"encoding/json"
)

type GasController struct {
	BaseController
}

func (this *GasController) Add() {
	var params models.Gas
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	err = models.AddGas(&params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	this.EchoMessage("add success")
}

func (this *GasController) ListAll() {
	all, err := models.GetGasRecords()
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	this.Echo(all)
}
