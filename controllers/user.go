package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"book/models"
)

type UserController struct {
	BaseController
}

func (this *UserController) Login() {
	var params models.User
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	params.Password = Sha256(params.Password)
	u, err := models.GetUser(params.UserName, params.Password)
	if err != nil {
		this.EchoError(err.Error())
		return
	}
	this.Echo(u)
}

func Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
