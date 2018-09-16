package logger

import "log"

/*
CheckErrFatal will print a generic error message
alongside the error
*/
func CheckErrFatal(msg string, err error) {
	if err != nil {
		log.Fatalf("%s\n%s\n", err, msg)
	}
}
