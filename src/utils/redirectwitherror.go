package utils

import (
	"log"
	"messages"
	"net/http"
)

/*
RedirectWithError function
logs the error to the console,
sets an appropriate message
and redirects the user
*/
func RedirectWithError(w http.ResponseWriter, r *http.Request, url, message string, err error) {
	if err != nil {
		log.Println(err)
	}
	messages.SetMessage(r, message)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
	return
}
