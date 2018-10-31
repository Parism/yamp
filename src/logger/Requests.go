package logger

var requests uint64

/*
IncrementRequests function
does what it says
*/
func IncrementRequests() {
	requests++
}

/*
GetTotalRequests function
returns the requests var representing
the total num of requests served
*/
func GetTotalRequests() uint64 {
	return requests
}
