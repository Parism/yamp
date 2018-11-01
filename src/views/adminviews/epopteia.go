package adminviews

import (
	"datastorage"
	"fmt"
	"log"
	"logger"
	"middleware"
	"models"
	"net/http"
	"strings"
	"utils"
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
	cappstats := make(chan models.AppStats)
	cdbstats := make(chan []models.DbStat)
	credisstats := make(chan models.RedisStats)
	defer close(cappstats)
	defer close(cdbstats)
	defer close(credisstats)
	go GetDbStats(cdbstats)
	go GetRedisStats(credisstats)
	go GetAppStats(cappstats)
	datamap := make(map[string]interface{})
	datamap["App"] = <-cappstats
	datamap["Db"] = <-cdbstats
	datamap["Redis"] = <-credisstats
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

func GetAppStats(c chan models.AppStats) {
	var appstats models.AppStats
	appstats.Latency = logger.GetLatency()
	appstats.Requests = logger.GetTotalRequests()
	appstats.Uptime = logger.GetUptime()
	appstats.MemUsed = logger.GetMemUsed()
	c <- appstats
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
	res.Close()
	c <- stats
}

/*
GeRedisStats function returns
stats about the redis server
*/
func GetRedisStats(c chan models.RedisStats) {
	var stats models.RedisStats
	r, _ := datastorage.GetDataRouter().GetDb("sessions")
	rc := r.GetRedisClient()
	statusMem := rc.Info("Memory")
	statusStats := rc.Info("Stats")
	tempMem := statusMem.String()
	tempStats := statusStats.String()
	memarray := strings.Fields(tempMem)
	statsarray := strings.Fields(tempStats)
	stats.CommandsProcessed = strings.Split(statsarray[5], ":")[1]
	stats.MemoryInUse = strings.Split(memarray[5], ":")[1]
	c <- stats
}
