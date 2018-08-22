package main

import (
	_ "SuiFengSAP/routers"

	"SuiFengSAP/models"

	"github.com/astaxie/beego"
)

func main() {

	models.Registdata()
	beego.Run()
}
