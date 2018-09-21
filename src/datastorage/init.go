package datastorage

import (
	"datastorage/models"
)

var dataRouter *models.DataRouter

func init() {
	dataRouter = &models.DataRouter{}
	dataRouter.BuildStatements()
	dataRouter.LoadDatabases()
	dataRouter.OpenDatabaseConnections()
}

/*
GetDataRouter returns the data router for use by the caller
*/
func GetDataRouter() *models.DataRouter {
	return dataRouter
}
