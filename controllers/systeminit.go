package controllers

import (
	"github.com/astaxie/beego"
)

type SystemInitController struct {
	beego.Controller
}

func (c *SystemInitController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
