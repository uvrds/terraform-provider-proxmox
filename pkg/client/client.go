package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type API struct {
	BaseURL  string
	Username string
	Password string
	Client   *http.Client

	ticket              string
	csrfPreventionToken string

	Auth bool
}

func NewClient(baseURL string, username string, password string) *API {
	return &API{
		BaseURL:  baseURL,
		Username: username,
		Password: password,
		Client:   &http.Client{},
	}
}

type Data struct {
	Method string
	Path   string
	Body   *bytes.Buffer
}

type Cookie struct {
	Data struct {
		CSRFPreventionToken string `json:"CSRFPreventionToken"`
		Ticket              string `json:"ticket"`
		Username            string `json:"username"`
	} `json:"data"`
}

func (api *API) req(data Data) error {
	var body io.Reader
	if data.Body != nil {
		body = data.Body
	} else {
		body = nil
	}
	req, err := http.NewRequest(data.Method, api.BaseURL+data.Path, body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Cookie", fmt.Sprintf("PVEAuthCookie=%s", api.ticket))
	req.Header.Set("CSRFPreventionToken", api.csrfPreventionToken)
	if err != nil {
		return err
	}
	resp, err := api.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	if data.Path == "/access/ticket" {
		var respCookie Cookie
		err = json.Unmarshal(content, &respCookie)
		if err != nil {
			return err
		}
		api.csrfPreventionToken = respCookie.Data.CSRFPreventionToken
		api.ticket = respCookie.Data.Ticket
		api.Auth = true
	}
	return nil
}

func (api *API) Authenticate() error {

	rq := Data{
		Method: "POST",
		Path:   "/access/ticket",
		Body:   bytes.NewBuffer([]byte("username=" + api.Username + "&password=" + api.Password)),
	}
	if api.Auth == false {
		err := api.req(rq)
		if err != nil {
			return err
		}
	}
	return nil
}

func (api *API) GetStatus(node string, id string) error {
	rq := Data{
		Method: "GET",
		Path:   "/nodes/" + node + "/qemu/" + id + "/status/current",
		Body:   nil,
	}
	err := api.req(rq)
	if err != nil {
		return err
	}
	return nil
}
