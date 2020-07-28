package subscription

import (
	"database/sql"
	"fmt"
	"time"
)

type TokenSelector func(day int) ([]string, error)
type TokenProcessor func(token string) error
type Logger func(string, ...interface{})
type TimeFunc func() time.Time

// NewDailyProcessor returns a func set up to run every day and bills the $9.99 monthly subscriptionCurry to
// customers if today is their billing day.
// Billing day is the day of the month the subscriptionCurry started, or the 28th, whichever is earliest.
func NewDailyProcessor(tokensFor TokenSelector, process TokenProcessor, log Logger, now TimeFunc) func() error {
	return func() error {
		ts, err := tokensFor(now().Day())
		if err != nil {
			return fmt.Errorf("unable to process subscriptions: %w", err)
		}

		for _, t := range ts {
			err := process(t)
			if err != nil {
				log("unable to process token %s\n", t)
			}
		}

		return nil
	}
}

func ProcessToken(token string) error {
	return nil
}

func NewTokenSelector(db *sql.DB) TokenSelector {
	return nil
}
