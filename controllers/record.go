package controllers

import (
	"book/models"
	"encoding/json"
)

type RecordController struct {
	BaseController
}

func (this *RecordController) Add() {
	var params models.Record
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	err=models.AddRecord(&params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	this.EchoMessage("add record success")
}

func (this *RecordController) ListAll() {
	all,err:=models.GetAllRecords()
	if err!=nil{
		this.EchoError(err.Error())
		return
	}
	this.Echo(all)
}

