package utils

import (
	"datastorage"
	"models"
)

/*
GetDimoiries function
returns all the nodes from ierarxia
that have no childs
*/
func GetDimoiries(c chan []models.Ierarxia) {
	ierarxiaArray := []models.Ierarxia{}
	ierarxia := models.Ierarxia{}
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT id,perigrafi from ierarxia;")
	for res.Next() {
		_ = res.Scan(
			&ierarxia.ID,
			&ierarxia.Perigrafi,
		)
		ierarxiaArray = append(ierarxiaArray, ierarxia)
	}
	res.Close()
	c <- ierarxiaArray
}

/*
GetLabels function returns all node childs
of the hierarchy tree
*/
func GetLabels(label int, c chan []models.Ierarxia) {
	ierarxiaArray := []models.Ierarxia{}
	ierarxia := models.Ierarxia{}
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT id,perigrafi from ierarxia where id=? || parentid=?;", label, label)
	for res.Next() {
		_ = res.Scan(
			&ierarxia.ID,
			&ierarxia.Perigrafi,
		)
		ierarxiaArray = append(ierarxiaArray, ierarxia)
	}
	res.Close()
	c <- ierarxiaArray
}
