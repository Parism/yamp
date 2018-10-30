package main

import (
	_ "datastorage"
	"log"
	"net/http"
	"views"
	_ "views/adminviews"
	_ "views/adminviews/create"
	_ "views/adminviews/delete"
	_ "views/adminviews/retrieve"
	_ "views/adminviews/update"
	_ "views/userviews"
	_ "views/userviews/create"
	_ "views/userviews/delete"
)

func main() {
	log.Println("Server started..")
	http.ListenAndServe(":8000", views.GetMux())
}
