package main

import (
	"fmt"
	"github.com/terraform-provider-proxmox/pkg/client"
)

func main() {

	BaseURL := "https://192.168.122.54:8006/api2/json"
	Username := "root@pam"
	Password := "asdqz123"
	//BaseURL := "https://proxmox03.noprod.srv.crpt.tech:8006/api2/json"
	//Username := "vfrolov@pam"
	//Password := "Vfrfhjyf!564236"

	t := client.NewClient(BaseURL, Username, Password, true)
	id, err := t.NextId()
	fmt.Println(err)
	data := client.Lxc{
		VMID:         id,
		Ostemplate:   "local:vztmpl/ubuntu-18.04-standard_18.04.1-1_amd64.tar.gz",
		Storage:      "local-lvm",
		Node:         "pve",
		Hostname:     "test",
		Cores:        "1",
		Memory:       "512",
		Description:  "xeq",
		Start:        "1",
		Password:     "asdqz123",
		Swap:         "0",
		Searchdomain: "noprod.srv.crpt.tech crpt.tech o.crpt.tech",
		Nameserver:   "10.73.70.141 10.73.69.11",
	}
	t.CreateLxc(data)

	//update
	//data := client.ConfigLXCUpdate{
	//	VMID:         "102",
	//	Node:         "pve",
	//	Hostname:     "kuber",
	//	Description:  "test",
	//	Cores:        "1",
	//	Memory:       "2",
	//	Swap:         "0",
	//	Searchdomain: "test",
	//}
	//t.ConfigLXCUpdate(data)
}
