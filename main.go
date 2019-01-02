package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"book/models"
	_ "book/routers"
	"time"
)

func main() {
	models.Init()
	//testAddModel()
	beego.Run()
}

func testAddModel() {

	model1 := &models.Model{Name: "型号1"}
	model2 := &models.Model{Name: "型号2"}
	model3 := &models.Model{Name: "型号3"}
	model4 := &models.Model{Name: "型号4"}
	models.AddModel(model1)
	models.AddModel(model2)
	models.AddModel(model3)
	models.AddModel(model4)

	device1 := &models.Device{Name: "挖土机1", ModelId: 1}
	device2 := &models.Device{Name: "挖土机2", ModelId: 2}
	device3 := &models.Device{Name: "挖土机3", ModelId: 3}
	device4 := &models.Device{Name: "挖土机4", ModelId: 4}
	models.AddDevice(device1)
	models.AddDevice(device2)
	models.AddDevice(device3)
	models.AddDevice(device4)

	driver1 := &models.Driver{Name: "小李", Telephone: "13912345671"}

	driver2 := &models.Driver{Name: "小赵", Telephone: "13912345672"}
	driver3 := &models.Driver{Name: "小钱", Telephone: "13912345673"}
	driver4 := &models.Driver{Name: "小孙", Telephone: "13912345674"}

	models.AddDriver(driver1)
	models.AddDriver(driver2)
	models.AddDriver(driver3)
	models.AddDriver(driver4)

	record1 := &models.Record{DeviceId: 1, DriverId: 1,
		Workplace: "大港", StartTime: time.Now(), EndTime: time.Now().Add(1000000),
		Customer: "客户1", Created: time.Now(), Remarks: "备注1"}

	record2 := &models.Record{DeviceId: 2, DriverId: 2,
		Workplace: "丹徒", StartTime: time.Now(), EndTime: time.Now().Add(1000000),
		Customer: "客户2", Created: time.Now(), Remarks: "备注2"}

	record3 := &models.Record{DeviceId: 3, DriverId: 3,
		Workplace: "辛丰", StartTime: time.Now(), EndTime: time.Now().Add(1000000),
		Customer: "客户3", Created: time.Now(), Remarks: "备注3"}

	record4 := &models.Record{DeviceId: 4, DriverId: 4,
		Workplace: "京口", StartTime: time.Now(), EndTime: time.Now().Add(1000000),
		Customer: "客户4", Created: time.Now(), Remarks: "备注4"}

	models.AddRecord(record1)
	models.AddRecord(record2)
	models.AddRecord(record3)
	models.AddRecord(record4)

	records, err := models.GetAllRecords()
	if err == nil {
		for _, r := range records {
			fmt.Printf("r=%+v", r)
		}
	}
}
