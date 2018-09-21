package main

import (
	_ "datastorage"
	"log"
	"net/http"
	"views"
	_ "views/adminviews"
	_ "views/adminviews/create"
	_ "views/userviews"
)

func main() {
	log.Println("Server started..")
	http.ListenAndServe(":8000", views.GetMux())
}
