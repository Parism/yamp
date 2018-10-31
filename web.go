package main

import (
	"datastorage"
	"log"
	"net/http"
	"os"
	"os/signal"
	"variables"
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
	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		log.Println("Closing database connections")
		datastorage.GetDataRouter().StopDbs()
		log.Println("Exiting")
		os.Exit(0)
	}()
	log.Println("Server started..")
	log.Println("Start time:", variables.StartTime)
	http.ListenAndServe(":8000", views.GetMux())
}
