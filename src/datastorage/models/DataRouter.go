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
	_ "github.com/go-sql-driver/mysql" //i need the initialization
)

/*
The DataRouter loads the database configuration files
and creates connections to those databases
*/
type DataRouter struct {
	databaseConfigurations []DatabaseConfig
	databases              map[string]databaseclients.DbClient
}

/*
SetDb function
inserts a key-value pair in the datarouter
*/
func (dr *DataRouter) SetDb(dbID string, db databaseclients.DbClient) {
	dr.databases[dbID] = db
}

/*
GetDb function
returns the appropriate db connection if it exists
if not, will return nil as database and an error
*/
func (dr *DataRouter) GetDb(dbID string) (db databaseclients.DbClient, err error) {
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
	var databasesconf []DatabaseConfig
	confFile, err := os.Open("databases.json")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading databases")
	}
	decoder := json.NewDecoder(confFile)
	err = decoder.Decode(&databasesconf)
	if err != nil {
		log.Println(err)
		log.Fatal("Error parsing database conf json")
	}
	dr.databaseConfigurations = databasesconf
	confFile.Close()
}

/*
OpenDatabaseConnections populates the map holding
the connections to the various databases
*/
func (dr *DataRouter) OpenDatabaseConnections() {
	var client databaseclients.DbClient
	dr.databases = make(map[string]databaseclients.DbClient)
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

/*
GetDatabases function
returns the map containing the database clients
mainly used for tests
*/
func (dr *DataRouter) GetDatabases() map[string]databaseclients.DbClient {
	return dr.databases
}

/*
MakeRedis function
returns a RedisClient object
*/
func MakeRedis(dbconf DatabaseConfig) databaseclients.DbClient {
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
func MakeMysql(dbconf DatabaseConfig) databaseclients.DbClient {
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
