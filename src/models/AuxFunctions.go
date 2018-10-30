package models

import (
	"bytes"
	"strings"
)

/*
DateBuilder function
reverses a string separated with -
used to reverse dates
eg
2018-10-03 becomes 03-10-2018
*/
func DateBuilder(s string) string {
	sarray := strings.Split(s, "-")
	var buffer bytes.Buffer
	for i := len(sarray) - 1; i >= 0; i-- {
		buffer.WriteString(sarray[i])
		if i > 0 {
			buffer.WriteString("-")
		}
	}
	return buffer.String()
}
