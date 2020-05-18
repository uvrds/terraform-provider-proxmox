package main

import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"github.com/terraform-provider-proxmox/pkg/client"
)

func main() {

	BaseURL := "https://192.168.122.54:8006/api2/json"
	Username := "root@pam"
	Password := "asdqz123"

	t := client.NewClient(BaseURL, Username, Password, true)

	_, err := t.NextId()
	if err != nil {
		logger.Fatalf("%s", err)
	}

}
