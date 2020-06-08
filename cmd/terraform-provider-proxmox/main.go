package main

import (
	"encoding/json"
	"fmt"
	"github.com/terraform-provider-proxmox/pkg/client"
)

func main() {

	//test shema

	//
	BaseURL := "https://192.168.122.54:8006/api2/json"
	Username := "root@pam"
	Password := "asdqz123"
	//BaseURL := "*"
	//Username := "vfrolov@pam"
	//Password := "*"

	t := client.NewClient(BaseURL, Username, Password, true)
	/*	id, err := t.NextId()
		fmt.Println(err)*/
	//Netmap := map[string]string{
	//	"name":   "eth0",
	//	"bridge": "vmbr0",
	//	"gw":     "192.168.122.1",
	//	"ip":     "192.168.122.80/24",
	//}
	//data := client.Lxc{
	//	VMID:         "100",
	//	Ostemplate:   "local:vztmpl/ubuntu-18.04-standard_18.04.1-1_amd64.tar.gz",
	//	Storage:      "local-lvm",
	//	Node:         "pve",
	//	Hostname:     "test",
	//	Cores:        "1",
	//	Memory:       "512",
	//	Description:  "xeq",
	//	Start:        "1",
	//	Password:     "asdqz123",
	//	Swap:         "0",
	//	Searchdomain: "noprod.srv.crpt.tech crpt.tech o.crpt.tech",
	//	Nameserver:   "10.73.70.141 10.73.69.11",
	//	Rootfs:       "10",
	//}
	//fmt.Println(Netmap)
	////t.CreateLxc(data)
	//t.ConfigLXCUpdateNetwork(data)

	//update

	//data := client.ConfigLXCUpdate{
	//	VMID:         "102",
	//	Node:         "pve",
	//	Hostname:     "kuber",
	//	Description:  "test",
	//	Cores:        "2",
	//	Memory:       "513",
	//	Swap:         "0",
	//	Searchdomain: "vk.com",
	//	Rootfs:       "23",
	//	Nameserver:   "192.168.1.1",
	//}
	//t.ConfigLXCUpdate(data)

	//read

	resp, err := t.ReadConfigLXC("pve", "100")
	if err != nil {
		fmt.Println(err)
	}
	var stat client.ReadConfigLXC
	err = json.Unmarshal(resp, &stat)
	if err != nil {
		fmt.Println(err)
	}

	//delete

	//data := client.Lxc{
	//	VMID:         "102",
	//	Ostemplate:   "local:vztmpl/ubuntu-18.04-standard_18.04.1-1_amd64.tar.gz",
	//	Storage:      "local-lvm",
	//	Node:         "pve",
	//	Hostname:     "test",
	//	Cores:        "1",
	//	Memory:       "512",
	//	Description:  "xeq",
	//	Start:        "1",
	//	Password:     "asdqz123",
	//	Swap:         "0",
	//	Searchdomain: "noprod.srv.crpt.tech crpt.tech o.crpt.tech",
	//	Nameserver:   "10.73.70.141 10.73.69.11",
	//	Rootfs: "10",
	//}
	//t.Deletelxc(data)

}
