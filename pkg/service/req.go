package service

import (
	"bytes"
	"crypto/tls"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Data struct {
	Method string
	Path   string
	Body   *bytes.Buffer
}

var pack = "service"

func Req(data Data, insecure bool) (io.ReadCloser, error) {
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
	if err != nil {
		log.WithFields(log.Fields{
			"package":  pack,
			"function": "Req",
			"error":    err,
			"data":     data.Method + " " + data.Path,
		}).Fatal("Response", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{
			"package":  pack,
			"function": "Req",
			"error":    err,
			"data":     req,
		}).Fatal("Response", err)
	}
	return resp.Body, nil
}
