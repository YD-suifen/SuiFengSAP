package main

import (
	_ "SuiFengSAP/routers"

	"SuiFengSAP/models"

	"github.com/astaxie/beego"
)

func main() {

	models.RegisterDB()
	beego.Run()
}
