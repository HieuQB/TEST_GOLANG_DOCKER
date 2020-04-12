package models

import (
	"fmt"
	"github.com/go-pg/pg"
)

type Flydata struct {
	//ID            int
	IDFlight      string
	Originlat     float64
	Originlong    float64
	Origincity    string
	Originairport string
	Deslat        float64
	Deslong       float64
	Descity       string
	Desairport    string
	Track 		  []TrackData
}

type TrackData struct {
	Latitude 	float64
	Longitude 	float64
	TimeStamp	int
}

func (u *Flydata) TableName() string {
	return "flydata"
}

func ConnectPG()  (*pg.DB){
	// init db connection
	opt, err := pg.ParseURL("postgres://postgres:123456@localhost:5433/testdb?sslmode=disable")
	if err != nil {
		panic(err)
		fmt.Println(err.Error())
	}
	db := pg.Connect(opt)
	fmt.Println("Connect PG")
	return db
}

func DisConnectPG(db *pg.DB)  {
	_ = db.Close()
	fmt.Println("Close PG")
}

func SaveFlightDetailToPostgres(db *pg.DB, dt Flydata)  (numRow int, err error) {
	sql := "insert into flydata (idflight, originlat, originlong, origincity, originairport, deslat, deslong, descity, desairport) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := db.Exec(sql,
		dt.IDFlight,
		dt.Originlat, dt.Originlong,
		dt.Origincity,dt.Originairport,
		dt.Deslat,dt.Deslong,
		dt.Descity, dt.Desairport)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return res.RowsAffected(), err
}

func SaveTrackToPostgres(db *pg.DB, dt TrackData, IDFlight string)  (numRow int, err error) {
	sql := "insert into trackdata (idflight, latitude, longitude, timestamp) values (?, ?, ?, ?)"
	res, err := db.Exec(sql,
		IDFlight,
		dt.Latitude, dt.Longitude,
		dt.TimeStamp)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return res.RowsAffected(), err
}

func GetAllDataFly(db *pg.DB)  (res []*Flydata,err error){
	err = db.Model(&res).Column("flydata.*").Select()
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	return
}