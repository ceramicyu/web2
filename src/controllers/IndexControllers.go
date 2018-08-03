package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)
var cpt *captcha.Captcha

type IndexController struct {
	beego.Controller
}


func (this *IndexController)login(w http.ResponseWriter, r *http.Request) {
	sess, _ := beego.GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	 sess.Set("username","dsad")



}
func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}
func (this *IndexController)Get()  {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}
func (this *IndexController)Index()  {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}
func (this *IndexController)Post()  {
	username:=this.GetString("username")
	password:=this.GetString("password")
	verifycode:=this.GetString("verifycode")
	if username != "" && password != "" && verifycode != "" {

	}
	mystruct:=&JCode{ Msg:0,Param:"msgok",Data:nil}
	this.Data["json"] = &mystruct
	this.ServeJSON()

}