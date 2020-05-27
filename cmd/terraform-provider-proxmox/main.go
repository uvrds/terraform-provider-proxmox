package main

import (
	"fmt"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"github.com/terraform-provider-proxmox/pkg/client"
)

func main() {

	BaseURL := "https://192.168.122.54:8006/api2/json"
	Username := "root@pam"
	Password := "asdqz123"

	t := client.NewClient(BaseURL, Username, Password, true)
	id, err := t.NextId()
	data := client.LxcClone{
		VMID:        "100",
		NEWID:       id,
		Storage:     "local-lvm",
		Node:        "pve",
		TargetNode:  "pve",
		Hostname:    "kuber02",
		Description: "213",
		Full:        "1",
	}
	err = t.CloneLxc(data)
	if err != nil {
		logger.Fatalf("%s", err)
	}
	fmt.Println(t)
	//t.GetNodes()
}
