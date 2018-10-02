package utils

import (
	"database/sql"
	"datastorage"
	"models"
)

/*
GetProswpikoList function
returns a sorted list of object models.Proswpiko
*/
func GetProswpikoList(label int) []models.Proswpiko {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var res *sql.Rows
	if label == 1 {
		res, _ = dbc.Query("SELECT id,surname,pname,rank,perigrafi FROM proswpiko_sorted")
	} else {
		res, _ = dbc.Query("SELECT id,surname,pname,rank,perigrafi FROM proswpiko_sorted where iid=? || pid=?", label, label)
	}

	var proswpikoArray []models.Proswpiko
	var proswpiko models.Proswpiko
	for res.Next() {
		_ = res.Scan(
			&proswpiko.ID,
			&proswpiko.Surname,
			&proswpiko.Name,
			&proswpiko.Rank,
			&proswpiko.Label,
		)
		proswpikoArray = append(proswpikoArray, proswpiko)
	}
	return proswpikoArray
}
