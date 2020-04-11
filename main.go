package main

import (
	"fly_jx/crawler"
	_ "fly_jx/routers"
	"github.com/astaxie/beego"
)

func main() {
	_ = beego.LoadAppConfig("ini", "data/config.yml")
	_ = crawler.CrawlerAndSaveDataToMongoDB()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
