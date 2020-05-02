package model

import (
	"bytes"
	"fmt"
	"github.com/terraform-provider-proxmox/pkg/service"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type AuthStruct struct {
	Username string
	Password string
	Url      string
	Insecure bool
}

func Auth() {
	b, err := strconv.ParseBool(os.Getenv("INSECURE"))
	if err != nil {
		log.Fatal(err)
	}
	au := AuthStruct{
		Username: os.Getenv("USER"),
		Password: os.Getenv("PASS"),
		Url:      os.Getenv("URL_PROXMOX"),
		Insecure: b,
	}
	rq := service.Data{
		Method: "POST",
		Path:   au.Url + "api2/json/access/ticket",
		Body:   bytes.NewBuffer([]byte("username=" + au.Username + "&password=" + au.Password)),
	}

	resp := service.Req(rq, au.Insecure)
	defer resp.Close()
	content, err := ioutil.ReadAll(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

}
