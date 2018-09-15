package datastorage

import (
	"datastorage/models"
	"log"
)

var dataRouter *models.DataRouter

func init() {
	dataRouter = &models.DataRouter{}
	dataRouter.LoadDatabases()
	log.Println("Initialization of database layer")
}

/*
GetDataRouter returns the data router for use by the caller
*/
func GetDataRouter() *models.DataRouter {
	return dataRouter
}
