package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

/*
GetRandStringb64 function
returns a random base64 encoded string
*/
func GetRandStringb64() string {
	var b []byte
	var bdecoded string
	b = make([]byte, 256)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Print(err)
	}
	bdecoded = base64.URLEncoding.EncodeToString(b)
	return bdecoded
}
