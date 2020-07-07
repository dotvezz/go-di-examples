package main

import (
	"database/sql"
	"github.com/dotvezz/go-di-examples/decoupled/subscriptionCurry/subscription"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "/")
	if err != nil {
		log.Fatalf("unable to run daily billing: %s", err)
	}

	selectTokens := subscription.NewTokenSelector(db)
	dailyProcess := subscription.NewDailyProcessor(selectTokens, subscription.ProcessToken, log.Printf, time.Now)

	err = dailyProcess()
	if err != nil {
		log.Fatalf("unable to run daily billing: %s", err)
	}
}
