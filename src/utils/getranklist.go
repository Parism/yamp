package utils

import (
	"database/sql"
	"datastorage"
	"log"
)

/*
GetRankList function
returns the unique ranks existing in a group
*/
func GetRankList(label int) []string {
	var rankArray []string
	var res *sql.Rows
	var err error
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	if label == 1 {
		res, err = dbc.Query("select distinct proswpiko_sorted.rank,ranks.id from proswpiko_sorted join ranks on proswpiko_sorted.rank = ranks.rank ORDER BY ranks.id DESC")
	} else {
		res, err = dbc.Query("select distinct proswpiko_sorted.rank,ranks.id from proswpiko_sorted join ranks on proswpiko_sorted.rank = ranks.rank where iid=? || pid=? ORDER BY ranks.id DESC", label, label)
	}
	if err != nil {
		log.Fatalln(err)
	}
	var rank string
	var rankid int
	for res.Next() {
		_ = res.Scan(
			&rank,
			&rankid,
		)
		rankArray = append(rankArray, rank)
	}
	return rankArray
}
