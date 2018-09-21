package auth

import (
	"auth/models"
	"datastorage"
	"logger"
	"net/http"
	"time"
	"utils"

	"github.com/go-redis/redis"
)

/*
NewCookie function returns
a cookie lasting 120 minutes
*/
func NewCookie() *http.Cookie {
	rc, err := datastorage.GetDataRouter().GetDb("sessions")
	logger.CheckErrFatal("auth/getcookie/GetCookie():21", err)
	redisclient := rc.GetRedisClient()
	var bdecoded string
	for {
		bdecoded = utils.GetRandStringb64()
		_, err = redisclient.Get(bdecoded).Result()
		if err == redis.Nil {
			session := &models.Session{}
			session.Sessionmap = make(map[string]string)
			session.SetKey("role", "none")
			session.SetKey("isAuthenticated", "false")
			session.SetKey("csrftoken", utils.GetRandStringb64())
			redisclient.SetNX(bdecoded, session.ToJSON(), 20*time.Minute)
			break
		}
	}
	var cookie = &http.Cookie{
		Name:     "sessionid",
		Value:    bdecoded,
		Expires:  time.Now().Add(120 * time.Minute),
		HttpOnly: true,
	}
	return cookie
}
