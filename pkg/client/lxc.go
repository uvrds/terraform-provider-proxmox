package client

import (
	"encoding/json"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"time"
)

type RespData struct {
	Data string `json:"data"`
}

type LxcStatus struct {
	Data struct {
		Status string `json:"status"`
	} `json:"data"`
}

func (api *API) statusLXC(node string, id string) ([]byte, error) {
	path := "/nodes/" + node + "/lxc/" + id + "/status/current"
	err := api.get(path, nil)
	if err != nil {
		return api.resp, err
	}
	logger.Infof("status vm %s", string(api.resp))
	return api.resp, nil
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

	time.Sleep(time.Second * 2)

	path = "/nodes/" + data.Node + "/lxc/" + data.VMID + "/status/start"
	err = api.post(path, nil)
	if err != nil {
		return err
	}
	var st = true
	for st {
		resp, err := api.statusLXC(data.Node, data.VMID)
		if err != nil {
			return err
		}
		var stat LxcStatus
		err = json.Unmarshal(resp, &stat)
		if err != nil {
			return err
		}
		if stat.Data.Status == "running" {
			st = false
			logger.Infof("start lxc ok %s", string(api.resp))
		}

	}
	return nil
}

func (api *API) Deletelxc(data Lxc) error {

	path := "/nodes/" + data.Node + "/lxc/" + data.VMID + "/status/stop"
	err := api.post(path, nil)
	if err != nil {
		return err
	}
	var st = true
	for st {
		resp, err := api.statusLXC(data.Node, data.VMID)
		if err != nil {
			return err
		}
		var stat LxcStatus
		err = json.Unmarshal(resp, &stat)
		if err != nil {
			return err
		}
		if stat.Data.Status == "stopped" {
			st = false
		}
		//time.Sleep(time.Second * 2)
	}
	logger.Infof("stop lxc %s", string(api.resp))

	path = "/nodes/" + data.Node + "/lxc/" + data.VMID + "?purge=1"
	err = api.del(path, nil)
	if err != nil {
		return err
	}
	logger.Infof("delete lxc %s", string(api.resp))
	return nil
}

type LxcClone struct {
	VMID        string
	NEWID       string
	Storage     string
	Node        string
	Hostname    string
	Description string
	Full        string
}

func (api *API) CloneLxc(data LxcClone) error {

	options := map[string]string{
		"newid": data.NEWID,
		"vmid":  data.VMID,
	}
	path := "/nodes/" + data.Node + "/lxc/" + data.NEWID + "/clone"
	err := api.post(path, options)
	if err != nil {
		return err
	}
	logger.Infof("clone lxc %s", string(api.resp))

	time.Sleep(time.Second * 2)

	var st = true
	for st {
		path = "/nodes/" + data.Node + "/lxc/" + data.NEWID + "/status/start"
		err = api.post(path, nil)
		if err != nil {
			return nil
		}
		resp, err := api.statusLXC(data.Node, data.VMID)
		if err != nil {
			return nil
		}
		var stat LxcStatus
		err = json.Unmarshal(resp, &stat)
		if err != nil {
			return nil
		}
		if stat.Data.Status == "running" {
			st = false
			logger.Infof("start lxc ok %s", string(api.resp))
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}

//common
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
