package utils

import (
	"datastorage"
	"models"
)

/*
GetRanks function returns all
ranks in the db asynchronously
*/
func GetRanks(c chan []models.Rank) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT * FROM ranks")
	var rank models.Rank
	var ranks []models.Rank
	for res.Next() {
		_ = res.Scan(
			&rank.ID,
			&rank.Rank,
		)
		ranks = append(ranks, rank)
	}
	res.Close()
	c <- ranks
}
