package models

/*
DbStat struct
represents stats fetched from the database
*/
type DbStat struct {
	QueryType  string
	TotalCount uint64
	PerSecond  float64
	PerMinute  float64
	PerHour    float64
	PerDay     float64
	StartDate  string
	EndDate    string
	Duration   string
}
