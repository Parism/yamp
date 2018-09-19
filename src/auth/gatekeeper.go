package auth

import (
	authmodels "auth/models"
	"datastorage"
	dbmodels "datastorage/models/databaseclients"
	"encoding/json"
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
	dbclient dbmodels.DbClient
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
CheckRole function
contacts the sessions database
searching for the cookie value
and validating the fields isAuthenticated and role
*/
func (gk *Gatekeeper) CheckRole(sessionid string, role string) bool {
	redisclient := gk.dbclient.GetRedisClient()
	res, err := redisclient.Get(sessionid).Result()
	if err != nil {
		return false
	}
	session := &authmodels.Session{}
	err = json.Unmarshal([]byte(res), session)
	if err != nil {
		log.Println("Error unmarshaling retrieved session gatekeeper.go/CheckRole:64")
		log.Println(err)
		return false
	}
	if session.GetKey("isAuthenticated") == "true" && session.GetKey("role") == role {
		return true
	}
	return false
}
