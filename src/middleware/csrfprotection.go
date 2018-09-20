package middleware

import (
	"auth/models"
	"datastorage"
	"encoding/json"
	"net/http"
)

/*
CsrfProtection function
is a middleware that checks whether the request
holds valid csrf value within the post data
*/
func CsrfProtection() Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, _ := r.Cookie("sessionid")
			sessionid := cookie.Value
			rc, _ := datastorage.GetDataRouter().GetDb("sessions")
			redisclient := rc.GetRedisClient()
			jsonsession, _ := redisclient.Get(sessionid).Result()
			session := &models.Session{}
			err := json.Unmarshal([]byte(jsonsession), session)
			if err != nil {
				http.Redirect(w, r, "/csrfdenied", http.StatusMovedPermanently)
				return
			}
			if r.PostFormValue("csrftoken") != session.GetKey("csrftoken") {
				http.Redirect(w, r, "/csrfdenied", http.StatusMovedPermanently)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
