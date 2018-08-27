package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	Id       int    `gorm:"primary_key"`
	Hostname string `gorm:"type:varchar(20);not null"`
	Ip       string `gorm:"type:varchar(25);not null;index:ip_idx"`
	// Resource   string `gorm`
	Createtime time.Time
}

type User struct {
	Id       int    `gorm:"primary_key"`
	Username string `gorm:"type:varchar(20);not null;unique"`
	Password string `gorm:"type:varchar(30);not null"`
	Pone     string `gorm:"type:varchar(20);not null;unique"`
	Mail     string `gorm:"type:varchar(35);not null"`
}

const (
	_DATA_DRIVER = "mysql"
)

var db *gorm.DB

func init() {

	var err error

	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlhost := beego.AppConfig.String("mysqlhost")
	mysqldata := beego.AppConfig.String("mysqldata")
	fmt.Println("hello")

	db, err = gorm.Open("mysql", mysqluser+":"+mysqlpass+"@"+"("+mysqlhost+")"+"/"+mysqldata+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

}

func Registdata() {

	if !db.HasTable(&Server{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Server{})

	}
	if !db.HasTable(&User{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	}
}

func InsertInto(table interface{}) bool {

	if err := db.Create(table).Error; err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func SelectUserDate() []User {

	userlist := []User{}
	db.Find(&userlist)
	return userlist

}

func SelectHostDate() []Server {
	hostlist := []Server{}
	db.Find(&hostlist)
	return hostlist
}

func LoginUserSelect(username string) bool {
	user := User{}
	if err := db.Where("Username = ?", username).Find(&user).Error; err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func DeleteUserDate(username string) bool {
	if err := db.Where("Username", username).Delete(&User{}).Error; err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func DeleteHostDate(hostip string) bool {
	if err := db.Where("Ip", hostip).Delete(&Server{}).Error; err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
