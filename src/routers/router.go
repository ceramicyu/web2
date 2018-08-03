package routers

import (
	"github.com/astaxie/beego"
    "controllers"
	"fmt"
)

type PP struct {

}

func init()  {
	beego.Router("/",&controllers.TplController{},"get:Index")
	beego.Router("/demo.html",&controllers.TplController{},"get:Demo")
	beego.Router("/demo1.html",&controllers.TplController{},"get:Demo1")
	beego.Router("/demo2.html",&controllers.TplController{},"get:Demo2")

}
func P()  {
	fmt.Println(6666)
}
