package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//记录
type Record struct {
	Id       int64 `orm:"column(id);auto;pk" json:"id"`
	DeviceId int64 `orm:"column(device_id)" json:"device_id"`
	DriverId int64 `orm:"column(driver_id)" json:"driver_id"`
	//工作时间
	StartTime time.Time `orm:"column(start_time)" json:"start_time"`
	EndTime   time.Time `orm:"column(end_time)" json:"end_time"`
	//客户
	Customer string `orm:"column(customer);size(100)" json:"customer"`

	//工作地点
	Workplace string `orm:"column(workplace);size(255)" json:"workplace"`
	//记录时间
	Created time.Time `orm:"column(created)" json:"created"`
	//备注
	Remarks string `orm:"column(remarks);size(1024)" json:"remarks"`
}

func (t *Record) TableName() string {
	return "record"
}

func init() {
	orm.RegisterModel(new(Record))
}

type RecordView struct {
	Id       int64 `orm:"column(id);auto" json:"id"`
	DeviceId int64 `orm:"column(device_id)" json:"device_id"`
	DriveId  int64 `orm:"column(driver_id)" json:"driver_id"`
	//工作时间
	StartTime time.Time `orm:"column(start_time)" json:"start_time"`
	EndTime   time.Time `orm:"column(end_time)" json:"end_time"`
	//客户
	Customer string `orm:"column(customer);size(100)" json:"customer"`

	//工作地点
	Workplace string `orm:"column(workplace);size(255)" json:"workplace"`
	//记录时间
	Created time.Time `orm:"column(created)" json:"created"`
	//备注
	Remarks string `orm:"column(remarks);size(1024)" json:"remarks"`

	//设备型号
	DeviceName string `orm:"column(device_name)" json:"device_name"`

	//司机名字
	DriverName string `orm:"column(driver_name)" json:"driver_name"`
}

func AddRecord(obj *Record) error {
	obj.Created = time.Now()
	o := orm.NewOrm()
	id, err := o.Insert(obj)
	if err == nil {
		obj.Id = id
	}
	return err
}

func GetAllRecords() ([]RecordView, error) {
	o := orm.NewOrm()
	records := make([]RecordView, 0)
	_, err := o.Raw("select r.*,de.name as device_name, d.name as driver_name from record r join driver d join device de on r.driver_id=d.id and r.device_id=de.id order by r.id desc").QueryRows(&records)
	return records, err
}
