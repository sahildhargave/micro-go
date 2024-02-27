package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:postgres123@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource) // Use the global variable instead of declaring a new one
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
