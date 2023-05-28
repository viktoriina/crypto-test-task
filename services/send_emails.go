package services

import (
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/viktoriina/crypto-test-task/helpers"
)

const (
	hostAddress  = "aaazsxdcfvvv@gmail.com"
	hostPassword = "ltupzurhipekhjrv"
	smtpHost     = "smtp.gmail.com"
	smtpPort     = "587"
)

func SendEmails(w http.ResponseWriter, r *http.Request) {
	auth := smtp.PlainAuth("", hostAddress, hostPassword, smtpHost)

	emails_to_send := helpers.ExtractEmails()

	price, err := GetBasePrice()

	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
		return
	}

	for i := 0; i < len(emails_to_send); i++ {
		email := emails_to_send[i]

		msg := []byte("To: " + email + "\r\n" +

			"Subject: Rate\r\n" +

			"\r\n" +

			"Cost of 1 bitcoin in UAH: " + price + "\r\n")
		to := []string{email}

		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, hostAddress, to, msg)

		if err != nil {
			fmt.Println("Couldn't send to email: " + email)
			fmt.Println(err.Error())
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
