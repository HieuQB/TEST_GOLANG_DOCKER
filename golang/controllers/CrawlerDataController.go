package controllers

import (
	"fly_jx/models"
	"github.com/astaxie/beego"
)

type CrawlerDataController struct {
	beego.Controller
}

// @router / [get]
func (mg *CrawlerDataController) Get() {
	client := models.ConnectMongo()
	list, err := models.GetAllFlyJV(client)
	models.DisconnectMongo(client)
	if err != nil {
		mg.Data["json"] = err.Error()
	} else {
		mg.Data["json"] = list
	}
	mg.ServeJSON()
}