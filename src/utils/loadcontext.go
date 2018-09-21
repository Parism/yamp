package utils

import (
	"messages"
	"net/http"
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
}

/*
LoadContext function
returns a context object
according to the request object
*/
func LoadContext(r *http.Request) Context {
	csrftoken := GetSessionValue(r, "csrftoken")
	user := GetSessionValue(r, "username")
	message := messages.GetMessage(r)
	var context Context
	context.Csrftoken = csrftoken
	context.Message = message
	context.User = user
	return context
}
