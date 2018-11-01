package utils

import (
	"datastorage"
	"log"
	"models"
)

/*
GetTypoiYpiresiwn function
returns all types of services
*/
func GetTypoiYpiresiwn(label int, c chan []models.TyposYpiresias) {
	var typos models.TyposYpiresias
	var typoi []models.TyposYpiresias
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("select typoiypiresiwn.id,typoiypiresiwn.perigrafi from typoiypiresiwn join ierarxia on idmonadas = ierarxia.id where ierarxia.id = ? || ierarxia.parentid = ?", label, label)
	defer res.Close()
	if err != nil {
		log.Println(err)
		c <- nil
		return
	}
	for res.Next() {
		_ = res.Scan(
			&typos.ID,
			&typos.Perigrafi)
		typoi = append(typoi, typos)
	}
	c <- typoi
}
