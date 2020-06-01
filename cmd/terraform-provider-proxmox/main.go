package main

import (
	"encoding/json"
	"fmt"
	"github.com/terraform-provider-proxmox/pkg/client"
	"regexp"
	"strings"
)

func main() {

	BaseURL := "https://192.168.122.54:8006/api2/json"
	Username := "root@pam"
	Password := "asdqz123"
	//BaseURL := "*"
	//Username := "vfrolov@pam"
	//Password := "*"

	t := client.NewClient(BaseURL, Username, Password, true)
	//id, err := t.NextId()
	//fmt.Println(err)
	//data := client.Lxc{
	//	VMID:         id,
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
	//t.CreateLxc(data)

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

	//read
	resp, err := t.ConfigLXC("pve", "102")
	if err != nil {
		fmt.Println(err)
	}
	var stat client.ConfigLXC
	err = json.Unmarshal(resp, &stat)
	if err != nil {
		fmt.Println(err)
	}

	spl := strings.Split(stat.Data.Rootfs, "=")
	re := regexp.MustCompile("[0-9]+")
	te := re.FindAllString(spl[1], -1)
	fmt.Println(te)

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
