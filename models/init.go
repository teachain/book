package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init(){
	orm.Debug = true
	port, _ := beego.AppConfig.Int("dbport")
	format := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local"
	urlStr := fmt.Sprintf(format, beego.AppConfig.String("dbuser"),
		beego.AppConfig.String("dbpassword"),
		beego.AppConfig.String("dbhost"),
		port,
		beego.AppConfig.String("dbname"),
	)
	orm.RegisterDataBase("default", "mysql", urlStr)
	orm.RunSyncdb("default", false, true)
}
