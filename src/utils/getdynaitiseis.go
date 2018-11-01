package utils

import (
	"bytes"
	"datastorage"
	"log"
	"models"
)

/*
GetDynAitiseisAll function
returns asynchronously
the aitiseis object of everyone needed for dynamologio
*/
func GetDynAitiseisAll(d string, c chan []models.Aitisi) {
	var aitisi models.Aitisi
	var aitiseis []models.Aitisi
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT idperson,perigrafi,name,surname,date from aitiseis ")
	buffer.WriteString("JOIN proswpiko on proswpiko.id = aitiseis.idperson ")
	buffer.WriteString("WHERE date = ?")
	res, err := dbc.Query(buffer.String(), d)
	defer res.Close()
	if err != nil {
		log.Println(err)
		c <- nil
	}
	for res.Next() {
		_ = res.Scan(
			&aitisi.IDPerson,
			&aitisi.Perigrafi,
			&aitisi.Name,
			&aitisi.Surname,
			&aitisi.Date,
		)
		aitisi.Date = models.DateBuilder(aitisi.Date)
		aitiseis = append(aitiseis, aitisi)
	}
	c <- aitiseis
}

/*
GetDynAitiseisLabel function
returns all aitiseis object of a certain label
*/
func GetDynAitiseisLabel(d string, label int, c chan []models.Aitisi) {
	var aitisi models.Aitisi
	var aitiseis []models.Aitisi
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT idperson,aitiseis.perigrafi,name,surname,date from aitiseis ")
	buffer.WriteString("JOIN proswpiko on proswpiko.id = aitiseis.idperson ")
	buffer.WriteString("JOIN ierarxia on proswpiko.label = ierarxia.id ")
	buffer.WriteString("WHERE date = ? and ")
	buffer.WriteString("(ierarxia.id = ? || ierarxia.parentid=?) ")
	res, err := dbc.Query(buffer.String(), d, label, label)
	defer res.Close()
	if err != nil {
		log.Println(err)
		c <- nil
	}
	for res.Next() {
		_ = res.Scan(
			&aitisi.IDPerson,
			&aitisi.Perigrafi,
			&aitisi.Name,
			&aitisi.Surname,
			&aitisi.Date,
		)
		aitisi.Date = models.DateBuilder(aitisi.Date)
		aitiseis = append(aitiseis, aitisi)
	}
	c <- aitiseis
}
