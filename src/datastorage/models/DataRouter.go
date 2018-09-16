package models

import (
	"database/sql"
	"datastorage/models/databaseclients"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"logger"
	"os"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

/*
The DataRouter loads the database configuration files
and creates connections to those databases
*/
type DataRouter struct {
	databaseConfigurations []DatabaseConf
	databases              map[string]databaseclients.Storage
}

/*
SetDb function
inserts a key-value pair in the datarouter
*/
func (dr *DataRouter) SetDb(dbID string, db databaseclients.Storage) {
	dr.databases[dbID] = db
}

/*
GetDb function
returns the appropriate db connection if it exists
if not, will return nil as database and an error
*/
func (dr *DataRouter) GetDb(dbID string) (db databaseclients.Storage, err error) {
	if _, exists := dr.databases[dbID]; exists {
		return dr.databases[dbID], nil
	}
	return nil, errors.New("Database does not exist in DataRouter")
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
	dr.databaseConfigurations = databases
	confFile.Close()
}

/*
OpenDatabaseConnections populates the map holding
the connections to the various databases
*/
func (dr *DataRouter) OpenDatabaseConnections() {
	var client databaseclients.Storage
	dr.databases = make(map[string]databaseclients.Storage)
	for _, databaseconf := range dr.databaseConfigurations {
		switch databaseconf.Type {
		case "redis":
			client = MakeRedis(databaseconf)
		case "mysql":
			client = MakeMysql(databaseconf)
		}
		if client.CheckConnection() {
			log.Printf("%s OK\n", databaseconf.ID)
			dr.SetDb(databaseconf.ID, client)
		}

	}
}

func (dr *DataRouter) GetDatabases() map[string]databaseclients.Storage {
	return dr.databases
}

/*
MakeRedis function
returns a RedisClient object
*/
func MakeRedis(dbconf DatabaseConf) databaseclients.Storage {
	var rc = &databaseclients.RedisClient{}
	rc.SetClient(redis.NewClient(&redis.Options{
		Addr:     dbconf.Link,
		Password: dbconf.Password,
		DB:       0,
	}))
	return rc
}

/*
MakeMysql function
returns a MysqlClient object
*/
func MakeMysql(dbconf DatabaseConf) databaseclients.Storage {
	var mc = &databaseclients.MysqlClient{}
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		dbconf.Username,
		dbconf.Password,
		dbconf.Link,
		dbconf.Database)
	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		logger.CheckErrFatal("error loading database", err)
	} else {
		mc.SetClient(db)
	}
	return mc
}
