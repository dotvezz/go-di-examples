package subscription

import (
	"fmt"
	"log"
	"time"
)

// ProcessDaily runs every day and bills the $9.99 monthly subscriptionCurry to customers if today is their billing day.
// Billing day is the day of the month the subscriptionCurry started, or the 28th, whichever is earliest.
func ProcessDaily() error {
	tokens, err := tokensForBillingDay(time.Now().Day())
	if err != nil {
		return fmt.Errorf("unable to get tokens to process: %w", err)
	}

	for _, token := range tokens {
		err = processSubscription(token)
		if err != nil {
			log.Printf("unable to process token %s\n", token)
		}
	}

	return nil
}

func processSubscription(token string) error {
	panic("demo code")
}

func tokensForBillingDay(day int) ([]string, error) {
	ProcessDaily()
	panic("demo code")
}
