package client

import (
	"encoding/json"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"time"
)

type RespData struct {
	Data string `json:"data"`
}

type StatusLXC struct {
	Data struct {
		Netin     int    `json:"netin"`
		Status    string `json:"status"`
		Diskread  int    `json:"diskread"`
		Diskwrite int    `json:"diskwrite"`
		Maxswap   int    `json:"maxswap"`
		Name      string `json:"name"`
		Netout    int    `json:"netout"`
		Template  string `json:"template"`
		Ha        struct {
			Managed int `json:"managed"`
		} `json:"ha"`
		Uptime  int         `json:"uptime"`
		Mem     int         `json:"mem"`
		Lock    string      `json:"lock"`
		CPU     interface{} `json:"cpu"`
		Swap    int         `json:"swap"`
		Maxmem  int         `json:"maxmem"`
		Type    string      `json:"type"`
		Vmid    string      `json:"vmid"`
		Maxdisk interface{} `json:"maxdisk"`
		Cpus    int         `json:"cpus"`
		Disk    interface{} `json:"disk"`
	} `json:"data"`
}

func (api *API) StatusLXC(node string, id string) ([]byte, error) {
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
	Start       string
	Password    string
}

func (api *API) CreateLxc(data Lxc) error {
	//todo возомжно надо отлавливать ошибку о том что не создался lxc и пробывать снова создать. вместо sleep
	time.Sleep(time.Second * 2)
	options := map[string]string{
		"ostemplate":  data.Ostemplate,
		"vmid":        data.VMID,
		"storage":     data.Storage,
		"hostname":    data.Hostname,
		"cores":       data.Cores,
		"memory":      data.Memory,
		"description": data.Description,
		"start":       data.Start,
		"password":    data.Password,
	}
	path := "/nodes/" + data.Node + "/lxc"
	err := api.post(path, options)
	if err != nil {
		return err
	}
	logger.Infof("create lxc %s", string(api.resp))
	return nil
}

func (api *API) Deletelxc(data Lxc) error {
	err := api.stopLxc(data.Node, data.VMID)
	if err != nil {
		return err
	}
	var s = true
	for s {
		path := "/nodes/" + data.Node + "/lxc/" + data.VMID + "?purge=1"
		err = api.del(path, nil)
		if err != nil {
			return err
		}
		b, err := api.сheckLxc(data.Node, data.VMID)
		if err != nil {
			return err
		}
		if b {
			s = false
		}
	}
	logger.Infof("delete lxc %s", data.VMID)
	return nil
}

func (api *API) stopLxc(node string, vmid string) error {
	path := "/nodes/" + node + "/lxc/" + vmid + "/status/stop"
	err := api.post(path, nil)
	if err != nil {
		return err
	}
	var s = true

	for s {
		b, err := api.сheckLxc(node, vmid)
		if err != nil {
			return err
		}
		if b {
			s = false
		}
		resp, err := api.StatusLXC(node, vmid)
		if err != nil {
			return err
		}
		var stat StatusLXC
		err = json.Unmarshal(resp, &stat)
		if err != nil {
			return err
		}
		if stat.Data.Status == "stopped" {
			logger.Infof("stop lxc id:%s %s", vmid, string(api.resp))
			s = false
		}
	}
	return nil
}

type CheckLxc struct {
	Data interface{} `json:"data"`
}

func (api *API) сheckLxc(node string, vmid string) (bool, error) {
	resp, err := api.StatusLXC(node, vmid)
	if err != nil {
		return false, err
	}
	var stat CheckLxc
	err = json.Unmarshal(resp, &stat)
	if err != nil {
		return false, err
	}
	if stat.Data == nil {
		logger.Infof(" no found lxc id:%s", vmid)
		return true, nil
	}

	return false, nil
}

type LxcClone struct {
	VMID        string
	NEWID       string
	Storage     string
	Node        string
	TargetNode  string
	Hostname    string
	Description string
	Full        string
	Cores       string
	Memory      string
}

func (api *API) CloneLxc(data LxcClone) error {

	options := map[string]string{
		"newid":       data.NEWID,
		"full":        data.Full,
		"storage":     data.Storage,
		"hostname":    data.Hostname,
		"description": data.Description,
		"target":      data.TargetNode,
	}
	path := "/nodes/" + data.Node + "/lxc/" + data.VMID + "/clone"
	err := api.post(path, options)
	if err != nil {
		return err
	}
	logger.Infof("clone lxc %s", string(api.resp))
	return nil
}
