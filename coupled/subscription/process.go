package subscription

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// ProcessDaily runs every day and bills the $9.99 monthly subscriptionCurry to customers if today is their billing day.
// Billing day is the day of the month the subscriptionCurry started, or the 28th, whichever is earliest.
func ProcessDaily(db *sql.DB) error {
	q := "SELECT `token` FROM `subscriptions` WHERE DAY(`startedDate`) = ?"
	day := time.Now().Day()
	if day == 28 {
		q = "SELECT `token` FROM `subscriptions` WHERE DAY(`startedDate`) >= ?"
	} else if day > 28 {
		return nil
	}

	rows, err := db.Query(q, day)
	if err != nil {
		return fmt.Errorf("unable to run query for processing subscriptions: %w", err)
	}

	for rows.Next() {
		var token string
		 _ = rows.Scan(&token)

		err = processSubscription(token)
		if err != nil {
			log.Printf("unable to process token %s\n", token)
			continue
		}
	}

	return nil
}

func processSubscription(token string) error {
	return nil
}