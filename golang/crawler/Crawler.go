package crawler

import (
	"fly_jx/models"
	"fly_jx/utils"
)

func GetListFightID()  (list []string, err error){
	// GET DATA FROM API
	dataAirport, err := utils.CallAPI(utils.URL_AIRPORT)
	if err != nil {
		return nil, err
	}
	// PARSE DATA
	airport := dataAirport["result"]["response"]["airport"].(map[string]interface{})
	pluginData := airport["pluginData"].(map[string]interface{})
	schedule := pluginData["schedule"].(map[string]interface{})
	arrivals := schedule["arrivals"].(map[string]interface{})
	data := arrivals["data"].([]interface{})
	for _, v := range data {
		vMap := v.(map[string]interface{})
		flight := vMap["flight"].(map[string]interface{})
		identification := flight["identification"].(map[string]interface{})
		id := identification["id"]
		if id != nil {
			idS := id.(string)
			list = append(list, idS)
		}
	}
	return
}

func GetFlightDetail() (data []models.DataFlightUnProcess){
	listID, _ := GetListFightID()
	for _, flightID := range listID {
		dataFlight, _ := utils.CallAPI(utils.URL_FIGHT_DETAIL + "flightId=" + flightID)
		// PARSE DATA
		dt := dataFlight["result"]["response"]["data"].(map[string]interface{})
		flight := dt["flight"].(map[string]interface{})
		airport := flight["airport"].(map[string]interface{})
		data = append(data, models.DataFlightUnProcess{
			Origin:      airport["origin"],
			Destination: airport["destination"],
		})
	}
	return
}

func CrawlerAndSaveDataToMongoDB() error {
	client := models.ConnectMongo()
	// CRAWLER DATA
	data := GetFlightDetail()
	// SAVE TO MONGO DB
	models.SaveDataCrawlerToMongo(client, data)
	models.DisconnectMongo(client)
	return nil
}
