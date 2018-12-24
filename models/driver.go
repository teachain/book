package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//司机
type Driver struct{
	Id int64  `orm:"column(id);auto;pk" json:"id"`
	Name string  `orm:"column(name);size(100)" json:"name"`
	Telephone string `orm:"column(telephone);size(11)" json:"telephone"`
	//记录时间
	Created  time.Time `orm:"column(created)" json:"created"`
}
func (t *Driver) TableName() string {
	return "driver"
}

func init() {
	orm.RegisterModel(new(Driver))
}

func AddDriver(obj *Driver)error{
	obj.Created=time.Now()
	o:=orm.NewOrm()
	id,err:=o.Insert(obj)
	if err==nil{
		obj.Id=id
	}
	return err
}

func GetAllDrivers()([]Driver,error){
	o:=orm.NewOrm()
	drivers:=make([]Driver,0)
	_,err:=o.Raw("select * from driver").QueryRows(&drivers)
	return drivers,err
}
