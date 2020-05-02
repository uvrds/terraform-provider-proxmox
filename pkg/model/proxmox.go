package model

import "fmt"

type AuthStruct struct {
	Username string
	Password string
}

func Auth(user string, pass string) {
	au := AuthStruct{
		Username: user,
		Password: pass,
	}

	fmt.Println(au)
}
