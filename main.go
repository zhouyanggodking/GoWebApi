package main

import (
	"GoWebApi/cacheSingleton"
	_ "GoWebApi/routers"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	bm := cacheSingleton.Get()
	bm.Put("godking", "godking", 100*time.Second)

	logs.SetLogger(logs.AdapterFile, `{"filename":"godking.log"}`)

	logs.GetLogger("main").Println("main entry started")

	beego.Run()
}
