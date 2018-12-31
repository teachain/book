package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//记录
type Maintain struct {
	Id int64 `orm:"column(id);auto;pk" json:"id"`

	DeviceId int64 `orm:"column(device_id)" json:"device_id"`

	//加油时间
	CreateTime time.Time `orm:"column(create_time)" json:"create_time"`

	//记录时间
	Created time.Time `orm:"column(created)" json:"created"`
}

func (t *Maintain) TableName() string {
	return "maintain"
}

func init() {
	orm.RegisterModel(new(Maintain))
}

type MaintainView struct {
	Id int64 `orm:"column(id);auto" json:"id"`

	//设备型号
	DeviceName string `orm:"column(device_name)" json:"device_name"`

	DeviceId int64 `orm:"column(device_id)" json:"device_id"`

	//加油时间
	CreateTime time.Time `orm:"column(create_time)" json:"create_time"`

	//记录时间
	Created time.Time `orm:"column(created)" json:"created"`
}

func AddMaintain(obj *Maintain) error {
	obj.Created = time.Now()
	o := orm.NewOrm()
	id, err := o.Insert(obj)
	if err == nil {
		obj.Id = id
	}
	return err
}

func GetAllMaintain() ([]MaintainView, error) {
	o := orm.NewOrm()
	records := make([]MaintainView, 0)
	_, err := o.Raw("select m.*,de.name as device_name from Maintain m join device de on  m.device_id=de.id order by m.id desc").QueryRows(&records)
	return records, err
}
