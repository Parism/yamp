package models

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
The DataRouter loads the database configuration files
and creates connections to those databases
*/
type DataRouter struct {
	databaseConfigurations []DatabaseConf
}

/*
LoadDatabases loads the configuration file
and creates a map that holds all the connections
along with a name identifier
*/
func (dr *DataRouter) LoadDatabases() {
	var databases []DatabaseConf
	confFile, err := os.Open("databases.json")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading databases")
	}
	decoder := json.NewDecoder(confFile)
	err = decoder.Decode(&databases)
	if err != nil {
		log.Println(err)
		log.Fatal("Error parsing database conf json")
	}
	for _, db := range databases {
		fmt.Printf("->%s\n", db.ID)
	}
	confFile.Close()
}
