package subscription

import (
	"fmt"
	"time"
)

func FullSignup(name string) error {
	customer, err := createCustomer(name)
	if err != nil {
		return fmt.Errorf("error in subsciption: %w", err)
	}

	_, err = createSubscription(customer.ID)
	if err != nil {
		return fmt.Errorf("error in subsciption: %w", err)
	}

	return nil
}

func createCustomer(name string) (Customer, error) {
	customer := Customer{
		Name: name,
	}

	err := insertCustomer(&customer)
	if err != nil {
		return customer, fmt.Errorf("error creating customer: %w", err)
	}

	return customer, nil
}

func createSubscription(customerID int) (Subscription, error) {
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


// Stuff below is just to suppress compile errors
func insertCustomer(customer *Customer) error {
	panic("demo code")
}

func insertSubscription(sub *Subscription) error {
	panic("demo code")
}

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
