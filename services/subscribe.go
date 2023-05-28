package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"

	"github.com/viktoriina/crypto-test-task/helpers"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var email struct {
		Email string `json:"email"`
	}

	json.Unmarshal(reqBody, &email)

	fmt.Println(email.Email)
	if email.Email != "" {
		if _, err := mail.ParseAddress(email.Email); err == nil {
			fmt.Println(email.Email)
			helpers.SaveEmail(email.Email)
		} else {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(""))
			return
		}
	} else {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(""))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
