package logger

import (
	"fmt"
	"runtime"
)

/*
GetMemUsed function returns
the current memory used by the app
*/
func GetMemUsed() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("%.3f\n", float32(m.Sys)/1024/1024)
}
