package main

import (
	"database/sql"
	"github.com/dotvezz/go-di-examples/decoupled/subscriptionStruct/subscription"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "/")
	if err != nil {
		log.Fatalf("unable to open mysql connection: %s", err)
	}

	dailyProcess := subscription.NewProcessor(db, log.Printf, time.Now)

	err = dailyProcess.RunDailyBatch()
	if err != nil {
		log.Fatalf("unable to run daily billing: %s", err)
	}
}
