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
	id, err := t.NextId()
	data := client.Lxc{
		VMID:        id,
		Ostemplate:  "local:vztmpl/ubuntu-18.04-standard_18.04.1-1_amd64.tar.gz",
		Storage:     "local-lvm",
		Node:        "pve",
		Hostname:    "test",
		Cores:       "1",
		Memory:      "512",
		Description: "test",
	}
	t.CreateLxc(data)
	if err != nil {
		logger.Fatalf("%s", err)
	}

}
