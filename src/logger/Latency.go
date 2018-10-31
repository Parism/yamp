package logger

import "time"

var count int = 0
var mean float32

/*
UpdateLatency function
a rolling average function
to keep hold of app health
*/
func UpdateLatency(d time.Duration) {
	count++
	differential := (float32(d.Seconds()) - mean) / float32(count)
	mean = mean + differential
}

/*
GetLatency function
returns the latency of the app
*/
func GetLatency() float32 {
	return mean * 1000
}
