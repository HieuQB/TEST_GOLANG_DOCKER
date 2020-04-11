package models

import (
	"fmt"
	"github.com/go-pg/pg"
)

type Flydata struct {
	//ID            int
	Originlat     float64
	Originlong    float64
	Origincity    string
	Originairport string
	Deslat        float64
	Deslong       float64
	Descity       string
	Desairport    string
}

func (u *Flydata) TableName() string {
	return "flydata"

}

func ConnectPG()  (*pg.DB){
	// init db connection
	opt, err := pg.ParseURL("postgres://postgres:123456@localhost/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)
	return db
}

func DisConnectPG(db *pg.DB)  {
	_ = db.Close()
}

func SaveToPostgres(db *pg.DB, dt Flydata)  (numRow int, err error) {
	sql := "insert into flydata (originlat, originlong, origincity, originairport, deslat, deslong, descity, desairport) values (?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := db.Exec(sql,
		dt.Originlat, dt.Originlong,
		dt.Origincity,dt.Originairport,
		dt.Deslat,dt.Deslong,
		dt.Descity, dt.Desairport)
	fmt.Println(err)
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