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
	statements             map[string]*sql.Stmt
}

/*
StmtExpr struct
is used to aggregate all needed
data to build and index a prepared statement
*/
type StmtExpr struct {
	Query string
	Db    string
	Index string
}

/*
BuildStatements function
prepares the statement
*/
func (dr *DataRouter) BuildStatements() {
	stmtArray := []StmtExpr{
		StmtExpr{
			Db:    "common",
			Query: "INSERT INTO accounts (username,password,role) VALUES(?,?,?);",
			Index: "insert_new_user",
		},
		StmtExpr{
			Db:    "common",
			Query: "DELETE FROM accounts WHERE id=?;",
			Index: "delete_user",
		},
		StmtExpr{
			Db:    "common",
			Query: "UPDATE accounts SET password=? WHERE username=?;",
			Index: "update_password",
		},
		StmtExpr{
			Db:    "common",
			Query: "DELETE FROM typoiadeiwn WHERE id=?;",
			Index: "delete_typos_adeias",
		},
		StmtExpr{
			Db:    "common",
			Query: "INSERT INTO typoiadeiwn (name,category) VALUES(?,?);",
			Index: "create_typos_adeias",
		},
		StmtExpr{
			Db:    "common",
			Query: "UPDATE accounts SET label=? WHERE id=?;",
			Index: "update_user_label",
		},
		StmtExpr{
			Db:    "common",
			Query: "UPDATE proswpiko SET label=? WHERE id=?;",
			Index: "metathesi",
		},
		StmtExpr{
			Db:    "common",
			Query: "INSERT INTO proswpiko (name,surname,rank,label) VALUES(?,?,?,?);",
			Index: "create_proswpiko",
		},
		StmtExpr{
			Db:    "common",
			Query: "DELETE FROM proswpiko WHERE id=?;",
			Index: "delete_proswpiko",
		},
		StmtExpr{
			Db:    "common",
			Query: "INSERT INTO adeies (type,idperson,start,end) VALUES (?,?,?,?)",
			Index: "create_adeia",
		},
		StmtExpr{
			Db:    "common",
			Query: "DELETE FROM adeies where id=?",
			Index: "delete_adeia",
		},
		StmtExpr{
			Db:    "common",
			Query: "INSERT INTO typoiypiresiwn (perigrafi,idmonadas) VALUES (?,?)",
			Index: "create_typos_ypiresias",
		},
		StmtExpr{
			Db:    "common",
			Query: "DELETE FROM typoiypiresiwn where id = ? and idmonadas = ?",
			Index: "delete_typos_ypiresias",
		},
		StmtExpr{
			Db:    "common",
			Query: "INSERT INTO ypiresies (date,personid,typeid) VALUES(CURDATE(),?,?)",
			Index: "create_person_ypiresia",
		},
		StmtExpr{
			Db:    "common",
			Query: "DELETE FROM ypiresies where idypiresies = ? and personid = ?",
			Index: "delete_person_ypiresia",
		},
		StmtExpr{
			Db:    "common",
			Query: "INSERT INTO categories_adeiwn (category) VALUES (?)",
			Index: "create_category_adeias",
		},
		StmtExpr{
			Db:    "common",
			Query: "DELETE FROM categories_adeiwn where id = ?",
			Index: "delete_category_adeias",
		},
		StmtExpr{
			Db:    "common",
			Query: "INSERT INTO aitiseis (perigrafi,idperson,date) VALUES(?,?,CURDATE())",
			Index: "create_aitisi",
		},
		StmtExpr{
			Db:    "common",
			Query: "DELETE FROM aitiseis where id = ? and idperson = ?",
			Index: "delete_aitisi",
		},
		StmtExpr{
			Db:    "common",
			Query: "INSERT INTO anafores (perigrafi,idperson,date) VALUES(?,?,CURDATE())",
			Index: "create_anafora",
		},
		StmtExpr{
			Db:    "common",
			Query: "DELETE FROM anafores where id = ? and idperson = ?",
			Index: "delete_anafora",
		},
		StmtExpr{
			Db:    "common",
			Query: "INSERT INTO ergasies (perigrafi,idperson,date) VALUES (?,?,CURDATE())",
			Index: "create_ergasia",
		},
		StmtExpr{
			Db:    "common",
			Query: "DELETE FROM ergasies where id = ? and idperson = ?",
			Index: "delete_ergasia",
		},
		StmtExpr{
			Db:    "common",
			Query: "INSERT INTO ypografes_aitisewn (iduser,idaitisi,status,signedas,date) VALUES((SELECT id from accounts where username = ?),?,?,?,CURDATE())",
			Index: "sign_aitisi",
		},
	}
	dr.statements = make(map[string]*sql.Stmt)
	var stmt *sql.Stmt
	var err error
	for _, value := range stmtArray {
		db, _ := dr.GetDb(value.Db)
		dbm := db.GetMysqlClient()
		stmt, err = dbm.Prepare(value.Query)
		logger.CheckErrFatal("stmt failed "+value.Index, err)
		dr.statements[value.Index] = stmt
	}
}

/*
GetStmt function
returns an already prepared statement
*/
func (dr *DataRouter) GetStmt(stm string) *sql.Stmt {
	return dr.statements[stm]
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
	defer confFile.Close()
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
		} else {
			log.Fatalln("Could not establish connection to", databaseconf.ID)
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

/*
StopDbs function of DataRouter
closes all connections to databases
of this datarouter
*/
func (dr *DataRouter) StopDbs() {
	for index := range dr.databases {
		dr.databases[index].CloseConnection()
	}
}
