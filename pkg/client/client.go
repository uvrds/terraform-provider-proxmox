package client

import (
	"encoding/json"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
)

func (api *API) GetStatusVM(node string, id string) error {
	path := "/nodes/" + node + "/qemu/" + id + "/status/current"
	err := api.get(path, nil)
	if err != nil {
		return err
	}
	logger.Infof("status vm %s", string(api.resp))
	return nil
}

type RespData struct {
	Data string `json:"data"`
}

type Lxc struct {
	VMID        string
	Ostemplate  string
	Storage     string
	Node        string
	Hostname    string
	Cores       string
	Memory      string
	Description string
}

func (api *API) CreateLxc(data Lxc) error {

	options := map[string]string{
		"ostemplate":  data.Ostemplate,
		"vmid":        data.VMID,
		"storage":     data.Storage,
		"hostname":    data.Hostname,
		"cores":       data.Cores,
		"memory":      data.Memory,
		"description": data.Description,
	}
	path := "/nodes/" + data.Node + "/lxc"
	err := api.post(path, options)
	if err != nil {
		return err
	}
	logger.Infof("create lxc %s", string(api.resp))
	return nil
}

func (api *API) Delete_lxc(data Lxc) error {
	path := "/nodes/" + data.Node + "/lxc/" + data.VMID + "?purge=1"
	err := api.del(path, nil)
	if err != nil {
		return err
	}
	logger.Infof("delete lxc %s", string(api.resp))
	return nil
}

func (api *API) NextId() (string, error) {

	path := "/cluster/nextid"
	err := api.get(path, nil)
	if err != nil {
		return "", err
	}
	var id RespData
	err = json.Unmarshal(api.resp, &id)
	logger.Infof("get id for vm %s", string(api.resp))
	if err != nil {
		return "", err
	}
	return id.Data, nil
}
