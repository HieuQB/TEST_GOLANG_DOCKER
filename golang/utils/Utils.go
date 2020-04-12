package utils

import (
	"crypto/tls"
	"encoding/json"
	"fly_jx/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func CallAPI(url string)  (result map[string]map[string]map[string]interface{}, err error){
	// Create New http Transport
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}
	// Create Http Client
	client := &http.Client{Transport: transCfg, Timeout: 100 * time.Second}
	defer client.CloseIdleConnections()
	request, err := http.NewRequest("GET", url, nil)
	if err == nil {
		response, err := client.Do(request)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			return nil, err
		} else {
			defer response.Body.Close()
			data, _ := ioutil.ReadAll(response.Body)
			err = json.Unmarshal(data, &result)
			return result, nil
		}
	} else {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}
}

func ProcessData(dataCrawler []models.DataFlight)  (res []models.Flydata){
	for _, dtCrawler := range dataCrawler {
		res = append(res, models.Flydata{
			IDFlight:	   dtCrawler.IDFlight,
			Originlat:     dtCrawler.Origin.Position.Latitude,
			Originlong:    dtCrawler.Origin.Position.Longitude,
			Origincity:    dtCrawler.Origin.Position.Region.City,
			Originairport: dtCrawler.Origin.Name,
			Deslat:        dtCrawler.Destination.Position.Latitude,
			Deslong:       dtCrawler.Destination.Position.Longitude,
			Descity:       dtCrawler.Destination.Position.Region.City,
			Desairport:    dtCrawler.Destination.Name,
			Track: 		   dtCrawler.Track,
		})
	}
	return
}

