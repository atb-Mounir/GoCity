package database

import (
	"database/sql"
	"log"

	_ "https/github.com/denisenkom/go-mssqldb"
)

	//DbConnection global
var DbConnection *sql.DB

func GetDb() (*sql.DB, error) {

	// Connection properties
	const user = "sa"
	const password = "Pa$$w0rd"
	const dbName = "SpringCityDB"

	// Make DNS
	dns := user + ":" + password + "@tcp(localhost:55657)/" + dbName

	// Get à database handle
	var err error
	DbConnection, err = sql.Open("sqlserver", dns)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DbConnection.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return DbConnection, err
}