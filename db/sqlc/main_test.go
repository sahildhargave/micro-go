package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/sahil/simplebank/util"
)



var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err!=nil{
		log.Fatal("cannot load config:",err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource) // Use the global variable instead of declaring a new one
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
