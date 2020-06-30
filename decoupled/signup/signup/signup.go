package signup

import (
	"database/sql"
	"fmt"
)

type FullSignup func(name string) error
type MemberCreator func(name string) (int, error)
type SubscriptionCreator func(memberId int) (int, error)

func NewFullSignup(member MemberCreator, sub SubscriptionCreator) FullSignup {
	return func(name string) error {
		memberId, err := member(name)
		if err != nil {
			return err
		}
		_, err = sub(memberId)
		if err != nil {
			return err
		}

		return nil
	}
}

func NewMemberCreator(db *sql.DB) MemberCreator {
	return func(name string) (int, error) {
		r, err := db.Exec("INSERT INTO `members` (`name`) VALUES (?)", name)
		if err != nil {
			return 0, fmt.Errorf("unable to create member: %w", err)
		}
		var id int64
		id, err = r.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("unable to create member: %w", err)
		}
		return int(id), nil
	}
}

func NewSubscriptionCreator(db *sql.DB) SubscriptionCreator {
	return func(memberId int) (int, error) {
		r, err := db.Exec("INSERT INTO `subscriptions` (`startedDate`, `memberId`) VALUES (NOW(), ?)", memberId)
		if err != nil {
			return 0, fmt.Errorf("unable to create subscription: %w", err)
		}
		var id int64
		id, err = r.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("unable to create subscription: %w", err)
		}
		return int(id), nil
	}
}
