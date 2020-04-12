package controllers

import (
	"fly_jx/models"
	"fly_jx/utils"
	"github.com/astaxie/beego"
)

type ProcessDataController struct {
	beego.Controller
}

// @router / [get]
func (pr *ProcessDataController) Get() {
	client := models.ConnectMongo()
	list, _ := models.GetAllFlyJV(client)
	models.DisconnectMongo(client)

	db := models.ConnectPG()
	dtConverted := utils.ProcessData(list)
	for _, dt := range dtConverted {
		_, _ = models.SaveToPostgres(db,dt)
	}
	//response, _ := models.GetAllDataFly(db)
	models.DisConnectPG(db)
	pr.Data["json"] = dtConverted
 	pr.ServeJSON()
}
