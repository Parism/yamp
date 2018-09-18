package auth

import (
	"datastorage"
	"datastorage/models/databaseclients"
	"log"
	"logger"
)

/*
Gatekeeper struct
The Gatekeeper holds the functions that will
check whether a request is authenticated or has
sufficient role to perform an action
*/
type Gatekeeper struct {
	dbclient databaseclients.DbClient
}

var gatekeeper *Gatekeeper

func init() {
	gatekeeper = &Gatekeeper{}
	gatekeeper.SetDb("sessions")

}

/*
GetGatekeeper function
returns the gatekeeper object to the caller
*/
func GetGatekeeper() *Gatekeeper {
	return gatekeeper
}

/*
SetDb function sets the db the gatekeeper
will interact with
*/
func (gk *Gatekeeper) SetDb(db string) {
	var err error
	gk.dbclient, err = datastorage.GetDataRouter().GetDb("sessions")
	logger.CheckErrFatal("Gatekeeper cannot load sessions db", err)
	log.Println("Gatekeeper loaded the sessions db")
}

/*
Checkauth function
contacts the sessions database
searching for the cookie value
if it exists it will return true
else will return false
*/
func (gk *Gatekeeper) Checkauth(sessionid string) bool {
	redisclient := gk.dbclient.GetRedisClient()
	res, err := redisclient.Get(sessionid).Result()
	if err != nil {
		return false
	}
	_ = res
	return true
}
