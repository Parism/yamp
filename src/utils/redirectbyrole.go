package utils

import "net/http"

/*
RedirectByRole is a helper function
that returns a url string according to the role of the caller
*/
func RedirectByRole(r *http.Request) string {
	role := GetSessionValue(r, "role")
	isAuthenticated := GetSessionValue(r, "isAuthenticated")
	if isAuthenticated == "false" || role == "" {
		return "/"
	}
	if role == "admin" {
		return "/secretadmin"
	}
	return "secret"
}
