package views

import (
	"net/http"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()
}

/*
GetMux function
returns the mux object used by the http listen and serve function
when the app starts
each view registers a path and itself in this object
*/
func GetMux() *http.ServeMux {
	return mux
}
