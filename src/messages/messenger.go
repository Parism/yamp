package messages

import (
	"auth/models"
	"datastorage"
	"net/http"
	"time"
)

/*
SetMessage function
store a message to the session object
*/
func SetMessage(r *http.Request, msg string) {
	cookie, _ := r.Cookie("sessionid")
	sessionid := cookie.Value
	rc, _ := datastorage.GetDataRouter().GetDb("sessions")
	redisclient := rc.GetRedisClient()
	res, _ := redisclient.Get(sessionid).Result()
	session := &models.Session{}
	session.Sessionmap = make(map[string]string)
	session.FromJSON(res)
	session.SetKey("message", msg)
	redisclient.Set(sessionid, session.ToJSON(), 2*time.Minute)
}

/*
GetMessage function
fetches the message stored for a session and then deletes it
*/
func GetMessage(r *http.Request) string {
	cookie, _ := r.Cookie("sessionid")
	sessionid := cookie.Value
	rc, _ := datastorage.GetDataRouter().GetDb("sessions")
	redisclient := rc.GetRedisClient()
	res, _ := redisclient.Get(sessionid).Result()
	session := &models.Session{}
	session.Sessionmap = make(map[string]string)
	session.FromJSON(res)
	msg := session.GetKey("message")
	session.DeleteKey("message")
	redisclient.Set(sessionid, session.ToJSON(), 20*time.Minute)
	return msg
}

/*
FetchSessionMessage function
*/
func FetchSessionMessage(sessionid string) string {
	return "test message"
}
