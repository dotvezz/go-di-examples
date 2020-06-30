package transport

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/dotvezz/go-di-examples/decoupled/signup/signup"
	"net/http"
)

func NewHandlerFunc(db *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tx, err := db.BeginTx(request.Context(), nil)
		if err != nil {
			_, _ = fmt.Fprint(writer, "it broke")
		}

		newMember := signup.NewMemberCreator(db)
		newSubscription := signup.NewSubscriptionCreator(db)
		fullSignup := signup.NewFullSignup(newMember, newSubscription)

		buf := new(bytes.Buffer)
		buf.ReadFrom(request.Body)
		name := buf.String()

		err = fullSignup(name)
		if err != nil {
			_ = tx.Rollback()
			_, _ = fmt.Fprint(writer, "it broke")
		} else {
			_ = tx.Commit()
			_, _ = fmt.Fprint(writer, "it worked")
		}
	}
}