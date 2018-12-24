package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//挖土机型号
type Model struct{
	Id int64  `orm:"column(id);auto;pk" json:"id"`
	Name string  `orm:"column(name);size(100)" json:"name"`
	//记录时间
	Created  time.Time `orm:"column(created)" json:"created"`
}
func (t *Model) TableName() string {
	return "model"
}

func init() {
	orm.RegisterModel(new(Model))
}

func AddModel(obj *Model)error{
	obj.Created=time.Now()
	o:=orm.NewOrm()
	id,err:=o.Insert(obj)
	if err==nil{
		obj.Id=id
	}
	return err
}

func GetAllModels()([]Model,error){
	o:=orm.NewOrm()
	models:=make([]Model,0)
	_,err:=o.Raw("select * from model").QueryRows(&models)
	return models,err
}
