package utils

import (
	"auth/models"
	"datastorage"
	"encoding/json"
	"net/http"
)

/*
GetSessionValue function
will be used by the views
The caller provides the request object and a key
the function will then ask the redis client what the value of the key is
and return it to the caller
*/
func GetSessionValue(r *http.Request, key string) string {
	/*
		fix here.
		nil pointer dereference
	*/
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		return ""
	}
	sessionid := cookie.Value
	rc, _ := datastorage.GetDataRouter().GetDb("sessions")
	redisclient := rc.GetRedisClient()
	res, err := redisclient.Get(sessionid).Result()
	if err != nil {
		return ""
	}
	session := &models.Session{}
	_ = json.Unmarshal([]byte(res), session)
	return session.GetKey(key)
}
