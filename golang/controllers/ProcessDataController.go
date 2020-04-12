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
	// GET DATA FROM MONGO
	client := models.ConnectMongo()
	list, _ := models.GetAllFlyJV(client)
	models.DisconnectMongo(client)

	// PROCESS DATA AND SAVE TO POSTGRES
	db := models.ConnectPG()
	dtConverted := utils.ProcessData(list)
	var dtResponse []models.Flydata
	for _, dt := range dtConverted {
		num, _ := models.SaveFlightDetailToPostgres(db,dt)
		if num == 1 {
			tmpTrack := []models.TrackData{}
			for _, itemTrack := range dt.Track {
				numTrackInsert,_ := models.SaveTrackToPostgres(db,itemTrack, dt.IDFlight)
				if numTrackInsert == 1 {
					tmpTrack = append(tmpTrack, itemTrack)
				}
			}
			dt.Track = tmpTrack
			dtResponse = append(dtResponse, dt)
		}
	}
	models.DisConnectPG(db)

	// EXPORT RESPONSE
	pr.Data["json"] = dtResponse
 	pr.ServeJSON()
}
