package logger

import (
	"models"
	"time"
	"variables"
)

/*
GetUptime function
fetches the uptime asynchronously
*/
func GetUptime() models.Uptime {
	var uptime models.Uptime
	diff := time.Since(variables.StartTime)
	uptime.Hours = int(diff.Hours())
	uptime.Minutes = int(diff.Minutes()) % 60
	uptime.Seconds = int(diff.Seconds()) % 60 % 60
	return uptime
}
