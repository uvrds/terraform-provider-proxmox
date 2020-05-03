package model

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/terraform-provider-proxmox/pkg/service"
	"io/ioutil"
	"os"
)

type AuthStruct struct {
	Username string
	Password string
	Url      string
	Insecure bool
}

var pack = "proxmox"

type Cookie struct {
	Data struct {
		CSRFPreventionToken string `json:"CSRFPreventionToken"`
		Ticket              string `json:"ticket"`
		Username            string `json:"username"`
	} `json:"data"`
}

func Auth() (Cookie, error) {
	var b bool
	ins := os.Getenv("INSECURE")
	if ins == "" {
		b = false
	} else if ins == "true" {
		b = true
	} else {
		b = false
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

	resp, err := service.Req(rq, au.Insecure)
	if err != nil {
		log.WithFields(log.Fields{
			"package":  pack,
			"function": "Auth",
			"error":    err,
			"data":     rq,
		}).Fatal("Response", err)
	}
	defer resp.Close()
	content, err := ioutil.ReadAll(resp)
	if err != nil {
		log.WithFields(log.Fields{
			"package":  pack,
			"function": "Auth",
			"error":    err,
			"data":     rq,
		}).Fatal("Response", err)
	}
	log.WithFields(log.Fields{
		"package":  pack,
		"function": "Auth",
		"data":     rq,
	}).Debug("Response", string(content))

	var respCookie Cookie
	err = json.Unmarshal(content, &respCookie)
	if err != nil {
		log.Fatal(err)
	}

	return respCookie, nil
}
