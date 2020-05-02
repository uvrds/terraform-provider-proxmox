package model

import (
	"fmt"
	"os"
)

type AuthStruct struct {
	Username string
	Password string
}

func Auth() {
	au := AuthStruct{
		Username: os.Getenv("USER"),
		Password: os.Getenv("PASS"),
	}

	fmt.Println(au)
}
