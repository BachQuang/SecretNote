package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/secretnote/backend/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../../..")

	if err != nil {
		log.Fatal("Can not load the config", err)
	}
	testDB, err = sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
