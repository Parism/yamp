package auth

import (
	authmodels "auth/models"
	"datastorage"
	dbmodels "datastorage/models/databaseclients"
	"encoding/json"
	"log"
	"logger"
	"net/http"
	"time"
	"utils"

	"golang.org/x/crypto/bcrypt"
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
func (gk *Gatekeeper) CheckRoleAndAuth(sessionid string, role string) bool {
	redisclient := gk.dbclient.GetRedisClient()
	res, err := redisclient.Get(sessionid).Result()
	if err != nil {
		return false
	}
	session := &authmodels.Session{}
	err = json.Unmarshal([]byte(res), session)
	if err != nil {
		log.Println("Error unmarshaling retrieved session gatekeeper.go/CheckRole")
		log.Println(err)
		return false
	}
	if session.GetKey("isAuthenticated") == "true" && session.GetKey("role") == role {
		return true
	}
	return false
}

/*
SessionExists function
checks whether the sessionid exists in the database or not
*/
func (gk *Gatekeeper) SessionExists(sessionid string) bool {
	rc, _ := datastorage.GetDataRouter().GetDb("sessions")
	redisclient := rc.GetRedisClient()
	_, err := redisclient.Get(sessionid).Result()
	if err != nil {
		return false
	}
	return true
}

/*
StoreSessionToDb function
takes a sessionid, creates a session object that is authenticated
*/
func (gk *Gatekeeper) StoreSessionToDb(sessionid string, role string, w http.ResponseWriter, r *http.Request) {
	session := &authmodels.Session{}
	session.SetKey("isAuthenticated", "true")
	session.SetKey("role", role)
	session.SetKey("csrftoken", utils.GetRandStringb64())
	rc, _ := datastorage.GetDataRouter().GetDb("sessions")
	redisclient := rc.GetRedisClient()
	redisclient.Set(sessionid, session.ToJSON(), 5*time.Minute)
	if role == "admin" {
		http.Redirect(w, r, "/secretadmin", http.StatusMovedPermanently)
		return
	} else if role == "user" {
		http.Redirect(w, r, "/secret", http.StatusMovedPermanently)
		return
	}
}

/*
Login function
is a gatekeeper function that checks whether the credentials
provided are valid or not
if they are the session is stored as authenticated
if not the user gets redirected to /login
*/
func (gk *Gatekeeper) Login(sessionid, role string, w http.ResponseWriter, r *http.Request) {
	mc, _ := datastorage.GetDataRouter().GetDb("common")
	mysqlclient := mc.GetMysqlClient()
	res, err := mysqlclient.Query("SELECT password from accounts where username=?", r.PostFormValue("username"))
	if err != nil {
		log.Println(err, "Login gatekeeper function")
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		return
	}
	var password string
	if res.Next() {
		err = res.Scan(&password)
		if err != nil {
			log.Println("error fetching password", err)
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(r.PostFormValue("password")))
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		return
	}
	gk.StoreSessionToDb(sessionid, role, w, r)
}
