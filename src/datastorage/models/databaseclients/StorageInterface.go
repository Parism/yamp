package databaseclients

/*
Storage interface
is used to have multiple types of databases
in the datarouter
e.g. mysql,redis and so on
*/
type Storage interface {
	CheckConnection() bool
}
