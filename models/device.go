package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//挖土机
type Device struct{
	Id int64  `orm:"column(id);auto;pk" json:"id"`
	Name string  `orm:"column(name);size(100)" json:"name"`
	ModelId int64  `orm:"column(model_id)" json:"model_id"`
	//记录时间
	Created  time.Time `orm:"column(created)" json:"created"`
}
func (t *Device) TableName() string {
	return "device"
}

func init() {
	orm.RegisterModel(new(Device))
}


func AddDevice(obj *Device)error{
	obj.Created=time.Now()
	o:=orm.NewOrm()
	id,err:=o.Insert(obj)
	if err==nil{
		obj.Id=id
	}
	return err
}

func GetAllDevices()([]Device,error){
	o:=orm.NewOrm()
	devices:=make([]Device,0)
	_,err:=o.Raw("select * from device").QueryRows(&devices)
	return devices,err
}