package helpers

import (
	"fmt"
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
