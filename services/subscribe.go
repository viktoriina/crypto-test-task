package services

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/viktoriina/crypto-test-task/helpers"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	fmt.Println(email)
	if email != "" {
		if _, err := mail.ParseAddress(email); err != nil {
			helpers.SaveEmail(email)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
