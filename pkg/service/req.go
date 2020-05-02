package service

import (
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net/http"
)

type Data struct {
	Method string
	Path   string
	Body   *bytes.Buffer
}

func Req(data Data, insecure bool) io.ReadCloser {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
	var body io.Reader
	if data.Body != nil {
		body = data.Body
	} else {
		body = nil
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(data.Method, data.Path, body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return resp.Body

}
