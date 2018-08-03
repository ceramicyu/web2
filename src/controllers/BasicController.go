package controllers

import "github.com/astaxie/beego"

type BasicController struct {
	beego.Controller
}

type JCode struct {
	Msg int `json:"msg"`
	Param string `json:"param"`
	Data interface{} `json:"data"`
}