package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"nway/nway_400/app/controllers"
	"nway/nway_400/app/service"
	"time"
)

const VERSION = "0.0.1"

func main() {
	service.Init()

	beego.AppConfig.Set("version", VERSION)
	if beego.AppConfig.String("runmode") == "dev" {
		beego.SetLevel(beego.LevelDebug)
	} else if beego.AppConfig.String("runmode") == "online" {
		beego.SetLevel(beego.LevelError)
		orm.Debug = false
	} else {
		beego.SetLevel(beego.LevelInformational)
		beego.SetLogger("file", `{"filename":"`+beego.AppConfig.String("log_file")+`"}`)
		beego.BeeLogger.DelLogger("console")
	}

	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.AutoRouter(&controllers.MainController{})
	beego.Router("/sm/index", &controllers.SmController{}, "*:Index")
	beego.AutoRouter(&controllers.SmController{})
	beego.AutoRouter(&controllers.CdrController{})

	// 记录启动时间
	beego.AppConfig.Set("up_time", fmt.Sprintf("%d", time.Now().Unix()))

	beego.AddFuncMap("i18n", i18n.Tr)

	beego.SetStaticPath("/assets", "assets")
	beego.SetStaticPath("/assets", "assets")
	beego.Run()
}
