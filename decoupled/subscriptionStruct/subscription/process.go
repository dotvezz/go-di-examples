package subscription

import (
	"database/sql"
	"fmt"
	"time"
)

type TokenSelector func(day int) ([]string, error)
type tokenProcessor func(token string) error
type Logger func(string, ...interface{})

type Processor interface {
	Run() error
}

type processor struct {
	db      *sql.DB
	process tokenProcessor
	log     Logger
	now     func() time.Time
}

func NewProcessor(db *sql.DB, log Logger, now func() time.Time) Processor {
	return &processor{
		db:      db,
		log:     log,
		now:     now,
		process: processToken,
	}
}

// NewProcessDaily returns a func set up to run every day and bills the $9.99 monthly subscriptionCurry to
// customers if today is their billing day.
// Billing day is the day of the month the subscriptionCurry started, or the 28th, whichever is earliest.
func (p *processor) Run() error {
	ts, err := p.tokensFor(p.now().Day())
	if err != nil {
		return fmt.Errorf("unable to process subscriptions: %w", err)
	}

	for _, t := range ts {
		err := p.process(t)
		if err != nil {
			p.log("unable to process token %s\n", t)
		}
	}

	return nil
}

// NewTokenSelector returns a func set up to select subscriptionCurry billing tokens for a given billing day
func (p *processor) tokensFor(day int) ([]string, error) {
	q := "SELECT `token` FROM `subscriptions` WHERE DAY(`startedDate`) = ?"
	if day == 28 {
		q = "SELECT `token` FROM `subscriptions` WHERE DAY(`startedDate`) >= ?"
	} else if day > 28 {
		return make([]string, 0), nil
	}

	rows, err := p.db.Query(q, day)
	if err != nil {
		return nil, fmt.Errorf("unable to run query for processing subscriptions: %w", err)
	}

	var ts []string
	for rows.Next() {
		var token string
		_ = rows.Scan(&token)
		ts = append(ts, token)
	}
	return ts, nil
}

func processToken(token string) error {
	return nil
}
