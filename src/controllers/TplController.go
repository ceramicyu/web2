package controllers

import "github.com/astaxie/beego"

type TplController struct {
	beego.Controller
}


func(this *TplController) Index()  {
	this.TplName = "index.html"
}
func(this *TplController) Demo()  {
	this.TplName = "demo.html"
}
func(this *TplController) Demo1()  {
	this.TplName = "demo1.html"
}
func(this *TplController) Demo2()  {
	this.TplName = "demo2.html"
}