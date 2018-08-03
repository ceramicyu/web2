package controllers

import (
	"github.com/astaxie/beego"
)


type Kill struct {
	Name string `json:"name"`
}

type MainController struct {
	BasicController
}

func (this *MainController)Post()  {

	mystruct:=&Kill{ Name:"hfsghsgfh"}
	this.Data["json"] = &mystruct
	this.ServeJSON()
}



type MainController2 struct {
	beego.Controller
}
func (this *MainController2)Post()  {

	mystruct:=&Kill{ Name:"85514447555"}
	this.Data["json"] = &mystruct
	this.ServeJSON()
}