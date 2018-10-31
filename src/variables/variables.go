package variables

import "time"

//each variable defines
//a certain level of clearance
//the biggest the number
//the more the permissions

var ADMIN = 8

var DITIS = 7

var YDITIS = 6

var P1G = 5

var CAPTAIN = 4

var VICE_CAPTAIN = 3

var GMT = 2

//USER var used for user role
var USER = 1

//LOGIN_EXPIRATION var used for session management
//self explanatory
var LOGIN_EXPIRATION = 30

/*
StartTime The time the app started.
Used to calculate uptime
*/
var StartTime time.Time

func init() {
	StartTime = time.Now()
}
