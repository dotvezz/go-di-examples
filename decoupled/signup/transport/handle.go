package transport

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/dotvezz/go-di-examples/decoupled/signup/subsciption"
	"net/http"
)

func NewHandlerFunc(db *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tx, err := db.BeginTx(request.Context(), nil)
		if err != nil {
			_, _ = fmt.Fprint(writer, "it broke")
		}
		deps := wireDeps(tx)

		buf := new(bytes.Buffer)
		buf.ReadFrom(request.Body)
		name := buf.String()

		err = deps.fullSignup(name)
		if err != nil {
			_ = tx.Rollback()
			_, _ = fmt.Fprint(writer, "it broke")
		} else {
			_ = tx.Commit()
			_, _ = fmt.Fprint(writer, "it worked")
		}
	}
}

func wireDeps(tx *sql.Tx) Deps {
	deps := Deps{}
	deps.insertCustomer = subsciption.NewCustomerInserter(tx)
	deps.insertSub = subsciption.NewSubInserter(tx)
	deps.createCustomer = subsciption.NewCustomerCreator(deps.insertCustomer)
	deps.createSub = subsciption.NewSubCreator(deps.insertSub)
	deps.fullSignup = subsciption.NewFullSignup(deps.createCustomer, deps.createSub)
	return deps
}

type Deps struct {
	insertCustomer subsciption.InsertCustomer
	insertSub      subsciption.InsertSub
	createCustomer subsciption.CreateCustomer
	createSub      subsciption.CreateSub
	fullSignup     subsciption.FullSignup
}
