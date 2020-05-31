package main

import (
	"github.com/terraform-provider-proxmox/pkg/client"
)

func main() {

	BaseURL := "https://192.168.122.54:8006/api2/json"
	Username := "root@pam"
	Password := "asdqz123"

	t := client.NewClient(BaseURL, Username, Password, true)
	//	id, err := t.NextId()
	//	fmt.Println(err)
	data := client.ConfigLXCUpdate{
		VMID:        "103",
		Node:        "pve",
		Hostname:    "test1",
		Cores:       "1",
		Memory:      "512",
		Description: "xuy",
		Swap:        "0",
	}
	t.ConfigLXCUpdate(data)
}
