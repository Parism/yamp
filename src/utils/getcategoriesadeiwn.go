package utils

import (
	"datastorage"
	"models"
)

/*
GetCategoriesAdeiwn function
returns asynchronously all the category rows
*/
func GetCategoriesAdeiwn(c chan []models.CategoryAdeias) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT * from categories_adeiwn")
	var category models.CategoryAdeias
	var categories []models.CategoryAdeias
	for res.Next() {
		_ = res.Scan(
			&category.ID,
			&category.Category,
		)
		categories = append(categories, category)
	}
	res.Close()
	c <- categories
}
