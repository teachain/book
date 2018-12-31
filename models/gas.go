package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//记录
type Gas struct {
	Id int64 `orm:"column(id);auto;pk" json:"id"`

	DeviceId int64 `orm:"column(device_id)" json:"device_id"`

	//加油时间
	CreateTime time.Time `orm:"column(create_time)" json:"create_time"`

	//数量
	Amount float32 `orm:"column(amount)" json:"amount"`

	//价格
	Price float32 `orm:"column(price)" json:"price"`
	//总价格
	TotalPrice float32 `orm:"column(total_price)" json:"total_price"`

	//记录时间
	Created time.Time `orm:"column(created)" json:"created"`
}

func (t *Gas) TableName() string {
	return "gas"
}

func init() {
	orm.RegisterModel(new(Gas))
}

type GasView struct {
	Id int64 `orm:"column(id);auto" json:"id"`

	//设备型号
	DeviceName string `orm:"column(device_name)" json:"device_name"`

	DeviceId int64 `orm:"column(device_id)" json:"device_id"`

	//加油时间
	CreateTime time.Time `orm:"column(create_time)" json:"create_time"`

	//数量
	Amount float32 `orm:"column(amount)" json:"amount"`

	//价格
	Price float32 `orm:"column(price)" json:"price"`
	//总价格
	TotalPrice string `orm:"column(total_price)" json:"total_price"`

	//记录时间
	Created time.Time `orm:"column(created)" json:"created"`
}

func AddGas(obj *Gas) error {
	obj.Created = time.Now()
	o := orm.NewOrm()
	id, err := o.Insert(obj)
	if err == nil {
		obj.Id = id
	}
	return err
}

func GetGasRecords() ([]GasView, error) {
	o := orm.NewOrm()
	records := make([]GasView, 0)
	_, err := o.Raw("select g.*,de.name as device_name from gas g join device de on  g.device_id=de.id order by g.id desc").QueryRows(&records)
	return records, err
}
