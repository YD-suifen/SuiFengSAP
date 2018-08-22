package models

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/astaxie/beego"
	"fmt"
)


type Server struct{
	Id         int       `gorm:"primary_key"`
	Hostname   string    `gorm:"type:varchar(20);not null"`
	Ip         string    `gorm:"type:varchar(25);not null;index:ip_idx"`
	Resource   string    `gorm`
    Createtime time.Time
}

type User struct{
	Id          int      `gorm:"primary_key"`
	Username    string   `gorm:"type:varchar(20);not null;unique"`
	Password    string   `gorm:"type:varchar(30);not null"`
	Pone        string   `gorm:"type:varchar(20);not null;unique"`
	Mail        string   `gorm:"type:varchar(35);not null"`
}


const (

	_DATA_DRIVER = "mysql"
)

func RegisterDB() {

	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlhost := beego.AppConfig.String("mysqlhost")
	mysqldata := beego.AppConfig.String("mysqldata")
	fmt.Println("hello")

	db, err := gorm.Open("mysql", mysqluser + ":" + mysqlpass + "@/" + mysqldata + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()


	if !db.HasTable(&Server{}){
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Server{}){
			panic(err)
		}

	}
	if !db.HasTable(&User{}){
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{}){
			panic(err)
		}
	}


}





