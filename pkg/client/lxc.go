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

type ReadLxc struct {
	Data struct {
		Template  string  `json:"template"`
		Cpus      int     `json:"cpus"`
		Pid       string  `json:"pid"`
		CPU       float64 `json:"cpu"`
		Swap      int     `json:"swap"`
		Status    string  `json:"status"`
		Lock      string  `json:"lock"`
		Diskwrite int     `json:"diskwrite"`
		Maxmem    int     `json:"maxmem"`
		Maxswap   int     `json:"maxswap"`
		Disk      string  `json:"disk"`
		Type      string  `json:"type"`
		Ha        struct {
			Managed int `json:"managed"`
		} `json:"ha"`
		Maxdisk  string `json:"maxdisk"`
		Netout   int    `json:"netout"`
		Diskread int    `json:"diskread"`
		Name     string `json:"name"`
		Netin    int    `json:"netin"`
		Vmid     string `json:"vmid"`
		Uptime   int    `json:"uptime"`
		Mem      int    `json:"mem"`
	} `json:"data"`
}

func (api *API) LxcRead(data Lxc) (ReadLxc, error) {
	resp, err := api.statusLXC(data.Node, data.VMID)
	if err != nil {
		return ReadLxc{}, err
	}
	var read ReadLxc
	err = json.Unmarshal(resp, &read)
	if err != nil {
		return ReadLxc{}, err
	}
	return read, nil
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

func (api *API) Delete_lxc(data Lxc) error {

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

type Nodes struct {
	Data []struct {
		Maxmem         int64   `json:"maxmem"`
		Mem            int     `json:"mem"`
		Node           string  `json:"node"`
		Maxcpu         int     `json:"maxcpu"`
		Type           string  `json:"type"`
		SslFingerprint string  `json:"ssl_fingerprint"`
		Uptime         int     `json:"uptime"`
		Disk           int64   `json:"disk"`
		Status         string  `json:"status"`
		ID             string  `json:"id"`
		Maxdisk        int64   `json:"maxdisk"`
		CPU            float64 `json:"cpu"`
		Level          string  `json:"level"`
	} `json:"data"`
}

func (api *API) GetNodes() (Nodes, error) {

	path := "/nodes"
	err := api.get(path, nil)
	if err != nil {
		return Nodes{}, err
	}
	var node Nodes
	err = json.Unmarshal(api.resp, &node)
	logger.Infof("get nodes %s", string(api.resp))
	if err != nil {
		return Nodes{}, err
	}
	return node, nil
}
