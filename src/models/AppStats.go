package models

/*
AppStats struct
represents various statistics
about the app
Its cross platform too ;)
*/
type AppStats struct {
	Requests uint64
	Latency  float32
	MemUsed  string
	Uptime   Uptime
}
