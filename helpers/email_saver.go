package helpers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func SaveEmail(email string) {
	file, err := os.OpenFile(".emails", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0744)

	fmt.Println(email)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	file.Write([]byte(email + "\n"))
}

func ExtractEmails() []string {
	file, err := os.OpenFile(".emails", os.O_RDONLY, os.ModeDevice)

	if err != nil {
		return nil
	}

	fileData, _ := ioutil.ReadAll(file)
	fileEmailsBytes := bytes.Split(fileData, []byte("\n"))

	result := make([]string, len(fileEmailsBytes))
	for i := 0; i < len(fileEmailsBytes); i++ {
		fmt.Println(string(fileEmailsBytes[i]))
		result = append(result, string(fileEmailsBytes[i]))
	}

	return result
}
