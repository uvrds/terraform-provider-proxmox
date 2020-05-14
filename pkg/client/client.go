package client

import (
	"fmt"
)

func (api *API) GetStatusVM(node string, id string) error {
	path := "/nodes/" + node + "/qemu/" + id + "/status/current"
	err := api.get(path, nil)
	if err != nil {
		return err
	}
	fmt.Println(api.resp)
	return nil
}

func (api *API) CreateLxc(node string) error {

	options := map[string]string{
		"ostemplate": "local:vztmpl/ubuntu-18.04-standard_18.04.1-1_amd64.tar.gz",
		"vmid":       "601",
		"storage":    "local-lvm",
	}

	path := "/nodes/" + node + "/lxc"
	err := api.post(path, options)
	if err != nil {
		return err
	}
	return nil
}
