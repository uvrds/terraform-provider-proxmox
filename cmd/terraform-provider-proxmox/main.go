package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/terraform-provider-proxmox/pkg/model"
	"os"
)

var pack = "main"

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	cookie, err := model.Auth()
	if err != nil {
		log.WithFields(log.Fields{
			"package":  pack,
			"function": "Req",
			"error":    err,
			"data":     nil,
		}).Fatal("Response", err)
	}

	fmt.Println(cookie)
}
