package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
)

var testQuery *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	var err error
	// config, err := util.LoadConfig("../..")
	// if err != nil {
	// 	log.Fatal("cannot load config")
	// }

	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		panic(err)
	}

	testQuery = New(testDB)

	os.Exit(m.Run())
}
