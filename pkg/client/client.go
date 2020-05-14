package client

import (
	"encoding/json"
)

func (api *API) GetStatusVM(node string, id string) error {
	path := "/nodes/" + node + "/qemu/" + id + "/status/current"
	err := api.get(path, nil)
	if err != nil {
		return err
	}
	return nil
}

type Lxc struct {
	VMID       string
	Ostemplate string
	Storage    string
	Node       string
	Hostname   string
	Cores      string
	Memory     string
}

func (api *API) CreateLxc(data Lxc) error {

	options := map[string]string{

		"ostemplate": data.Ostemplate,
		"vmid":       data.VMID,
		"storage":    data.Storage,
		"hostname":   data.Hostname,
		"cores":      data.Cores,
		"memory":     data.Memory,
	}

	path := "/nodes/" + data.Node + "/lxc"
	err := api.post(path, options)
	if err != nil {
		return err
	}
	return nil
}

type Id struct {
	Data string `json:"data"`
}

func (api *API) NextId() (string, error) {

	path := "/cluster/nextid"
	err := api.get(path, nil)
	if err != nil {
		return "", err
	}
	var id Id
	err = json.Unmarshal(api.resp, &id)
	if err != nil {
		return "", err
	}

	return id.Data, nil
}
