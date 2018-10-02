package auth

import (
	authmodels "auth/models"
	"datastorage"
	dbmodels "datastorage/models/databaseclients"
	"encoding/json"
	"log"
	"logger"
	"messages"
	"net/http"
	"strconv"
	"time"
	"utils"
	"variables"

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
CheckRoleAndAuth function
contacts the sessions database
searching for the cookie value
and validating the fields isAuthenticated and role
*/
func (gk *Gatekeeper) CheckRoleAndAuth(sessionid string) int {
	redisclient := gk.dbclient.GetRedisClient()
	res, err := redisclient.Get(sessionid).Result()
	if err != nil {
		return 0
	}
	session := &authmodels.Session{}
	err = json.Unmarshal([]byte(res), session)
	if err != nil {
		log.Println("Error unmarshaling retrieved session gatekeeper.go/CheckRole")
		log.Println(err)
		return 0
	}
	if session.GetKey("isAuthenticated") == "false" {
		return 0
	}
	value, err := strconv.Atoi(session.GetKey("role"))
	if err != nil {
		log.Println(err, "error strconv")
		return 0
	}
	return value
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
func (gk *Gatekeeper) StoreSessionToDb(sessionid, role, username, label string, w http.ResponseWriter, r *http.Request) {
	session := &authmodels.Session{}
	session.Sessionmap = make(map[string]string)
	session.SetKey("isAuthenticated", "true")
	session.SetKey("label", label)
	session.SetKey("role", role)
	session.SetKey("username", username)
	session.SetKey("csrftoken", utils.GetRandStringb64())
	rc, _ := datastorage.GetDataRouter().GetDb("sessions")
	redisclient := rc.GetRedisClient()
	redisclient.Set(sessionid, session.ToJSON(), 20*time.Minute)
	redisclient.ExpireAt(sessionid, time.Now().Add(20*time.Minute))
	roleint, _ := strconv.Atoi(role)
	if roleint >= variables.ADMIN {
		http.Redirect(w, r, "/diaxeiristiko", http.StatusMovedPermanently)
		return
	} else if roleint <= variables.CAPTAIN {
		http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
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
func (gk *Gatekeeper) Login(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("sessionid")
	sessionid := cookie.Value
	mc, _ := datastorage.GetDataRouter().GetDb("common")
	mysqlclient := mc.GetMysqlClient()
	res, err := mysqlclient.Query("SELECT password,role,label from accounts where username=?", r.PostFormValue("username"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	var password, role, label string
	if res.Next() {
		err = res.Scan(&password, &role, &label)
		if err != nil {
			log.Println("error fetching password", err)
		}
	}
	res.Close()
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(r.PostFormValue("password")))
	if err != nil {
		messages.SetMessage(r, "Λάθος κωδικός ή όνομα χρήστη")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	gk.StoreSessionToDb(sessionid, role, r.PostFormValue("username"), label, w, r)
}

/*
Logout function of gatekeeper
responsible to logout a user
get his sessionid
fetch the session from the sessions db
change the value isAuthenticated to false
and redirect him to the / page
*/
func (gk *Gatekeeper) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("sessionid")
	sessionid := cookie.Value
	rc, _ := datastorage.GetDataRouter().GetDb("sessions")
	redisclient := rc.GetRedisClient()
	jsonsession, _ := redisclient.Get(sessionid).Result()
	session := &authmodels.Session{}
	session.FromJSON(jsonsession)
	session.SetKey("isAuthenticated", "false")
	redisclient.Set(sessionid, session.ToJSON(), 20*time.Minute)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
