package main

import (
	_ "datastorage"
	"log"
	"net/http"
	"views"
)

func main() {
	log.Println("Server started..")
	http.ListenAndServe(":8000", views.GetMux())
}
