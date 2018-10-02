package utils

import (
	"messages"
	"net/http"
	"strconv"
)

/*
Context struct
holds common values used by views
such as csrftoken and user
convert to map in the future?
*/
type Context struct {
	Csrftoken string
	Message   string
	User      string
	Label     string
	Role      int
}

/*
LoadContext function
returns a context object
according to the request object
*/
func LoadContext(r *http.Request) Context {
	csrftoken := GetSessionValue(r, "csrftoken")
	user := GetSessionValue(r, "username")
	label := GetSessionValue(r, "label")
	message := messages.GetMessage(r)
	role := GetSessionValue(r, "role")
	roleint, _ := strconv.Atoi(role)
	var context Context
	context.Csrftoken = csrftoken
	context.Message = message
	context.User = user
	context.Label = label
	context.Role = roleint
	return context
}
