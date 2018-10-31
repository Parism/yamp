package adminviews

import (
	"datastorage"
	"fmt"
	"log"
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
	cdbstats := make(chan []models.DbStat)
	go GetDbStats(cdbstats)
	go GetUptime(cuptime)
	datamap := make(map[string]interface{})
	datamap["Latency"] = logger.GetLatency()
	datamap["Requests"] = logger.GetTotalRequests()
	datamap["Uptime"] = <-cuptime
	datamap["DbStats"] = <-cdbstats
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

/*
GetDbStats function
returns asynchronously the database statistics
*/
func GetDbStats(c chan []models.DbStat) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var dbstat models.DbStat
	var stats []models.DbStat
	res, err := dbc.Query("SELECT * FROM _dba_query_stats")
	if err != nil {
		log.Println("Error fetching db stats", err)
		c <- nil
	}
	for res.Next() {
		_ = res.Scan(
			&dbstat.QueryType,
			&dbstat.TotalCount,
			&dbstat.PerSecond,
			&dbstat.PerMinute,
			&dbstat.PerHour,
			&dbstat.PerDay,
			&dbstat.StartDate,
			&dbstat.EndDate,
			&dbstat.Duration,
		)
		stats = append(stats, dbstat)
	}
	c <- stats
}
