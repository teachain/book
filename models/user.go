package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户
type User struct{
	Id int64  `orm:"column(id);auto;pk" json:"id"`
	UserName string  `orm:"column(username);size(100)" json:"username"`
	Password string  `orm:"column(password);size(100)" json:"password"`
	//记录时间
	Created  time.Time `orm:"column(created)" json:"created"`
}
func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}


func GetUser(username,password string)(*User,error){
	o:=orm.NewOrm()
	u :=&User{UserName:username,Password:password}
	err:=o.Read(u,"UserName","Password")
    return u,err
}


