package adminviews

import (
	"fmt"
	"logger"
	"middleware"
	"models"
	"net/http"
	"time"
	"utils"
	"variables"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/epopteia", middleware.WithMiddleware(epopteia,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.NoCache(),
		middleware.IsAdmin(),
	))
}

func epopteia(w http.ResponseWriter, r *http.Request) {
	cuptime := make(chan models.Uptime)
	go GetUptime(cuptime)
	datamap := make(map[string]interface{})
	datamap["Latency"] = logger.GetLatency()
	datamap["Uptime"] = <-cuptime
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	data.Data = datamap
	t, err := utils.LoadTemplates("epopteia",
		"templates/adminviews/epopteia.html",
		"templates/adminviews/navbar.html",
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	t.ExecuteTemplate(w, "epopteia", data)
}

/*
GetUptime function
fetches the uptime asynchronously
*/
func GetUptime(c chan models.Uptime) {
	var uptime models.Uptime
	diff := time.Since(variables.StartTime)
	uptime.Hours = int(diff.Hours())
	uptime.Minutes = int(diff.Minutes()) % 60
	uptime.Seconds = int(diff.Seconds()) % 60 % 60
	c <- uptime
}
