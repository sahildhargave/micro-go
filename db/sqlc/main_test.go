package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:postgres123@localhost:5432/simple_bank?sslmode=disable"
)


var testQueries *Queries
//var testDB *sql.DB
var conn *sql.DB

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
    
	testQueries = New(conn)
	os.Exit(m.Run())
}
