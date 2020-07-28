package subsciption

import (
	"database/sql"
	"fmt"
	"time"
)

type FullSignup func(name string) error
type CreateCustomer func(name string) (Customer, error)
type CreateSub func(memberId int) (Subscription, error)
type InsertCustomer func (*Customer) error
type InsertSub func (*Subscription) error

func NewFullSignup(createCustomer CreateCustomer, createSub CreateSub) FullSignup {
	return func(name string) error {
		customer, err := createCustomer(name)
		if err != nil {
			return fmt.Errorf("error in signup: %w", err)
		}
		_, err = createSub(customer.ID)
		if err != nil {
			return fmt.Errorf("error in signup: %w", err)
		}

		return nil
	}
}

func NewCustomerCreator(insertCustomer InsertCustomer) CreateCustomer {
	return func(name string) (Customer, error) {
		customer := Customer{
			Name: name,
		}

		err := insertCustomer(&customer)
		if err != nil {
			return customer, fmt.Errorf("error creating customer: %w", err)
		}

		return customer, nil
	}
}

func NewSubCreator(insertSubscription InsertSub) CreateSub {
	return func(customerID int) (Subscription, error) {
		sub := Subscription{
			CustomerID: customerID,
			BillingDay: time.Now().Day(),
		}

		if sub.BillingDay > 28 {
			sub.BillingDay = 28
		}

		err := insertSubscription(&sub)
		if err != nil {
			return sub, fmt.Errorf("error creating subscription: %w", err)
		}
		return sub, nil
	}
}

// Stuff below is just to suppress compile errors
type Subscription struct {
	ID           int
	BillingDay   int
	BillingToken string
	CustomerID   int
}

type Customer struct {
	ID   int
	Name string
}

func NewCustomerInserter(tx *sql.Tx) InsertCustomer {
	panic("demo code")
}

func NewSubInserter(tx *sql.Tx) InsertSub {
	panic("demo code")
}