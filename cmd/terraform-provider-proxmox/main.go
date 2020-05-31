package main

import (
	"fmt"
	"github.com/terraform-provider-proxmox/pkg/client"
)

func main() {

	BaseURL := "https://192.168.122.54:8006/api2/json"
	Username := "root@pam"
	Password := "asdqz123"

	t := client.NewClient(BaseURL, Username, Password, true)
	id, err := t.NextId()
	fmt.Println(err)
	data := client.Lxc{
		VMID:        id,
		Ostemplate:  "local:vztmpl/ubuntu-18.04-standard_18.04.1-1_amd64.tar.gz",
		Storage:     "local-lvm",
		Node:        "pve",
		Hostname:    "test1",
		Cores:       "1",
		Memory:      "512",
		Description: "xuy",
		Start:       "1",
		Password:    "asdqz123",
		Swap:        "0",
	}
	t.CreateLxc(data)
}
