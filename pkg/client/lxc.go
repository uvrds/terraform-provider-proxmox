package client

import (
	"encoding/json"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"reflect"
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
	logger.Infof("status lxc %s", string(api.resp))
	return api.resp, nil
}

type Lxc struct {
	VMID         string
	Ostemplate   string
	Storage      string
	Node         string
	Hostname     string
	Cores        string
	Memory       string
	Description  string
	Start        string
	Password     string
	Swap         string
	Searchdomain string
	Nameserver   string
	Rootfs       string
	Net          interface{}
}

func (api *API) CreateLxc(data Lxc) error {
	time.Sleep(time.Second * 2)

	s := reflect.ValueOf(data.Net)
	if s.Kind() != reflect.Slice {
		logger.Fatalf("err")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	options := map[string]string{
		"ostemplate":   data.Ostemplate,
		"vmid":         data.VMID,
		"storage":      data.Storage,
		"hostname":     data.Hostname,
		"cores":        data.Cores,
		"memory":       data.Memory,
		"description":  data.Description,
		"start":        data.Start,
		"password":     data.Password,
		"swap":         data.Swap,
		"searchdomain": data.Searchdomain,
		"nameserver":   data.Nameserver,
		"rootfs":       data.Rootfs,
		"net0":         ret[0].(string),
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
		b, err := api.checkLxc(data.Node, data.VMID)
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
		b, err := api.checkLxc(node, vmid)
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

func (api *API) checkLxc(node string, vmid string) (bool, error) {
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
	VMID         string
	NEWID        string
	Storage      string
	Node         string
	TargetNode   string
	Hostname     string
	Description  string
	Full         string
	Cores        string
	Memory       string
	Swap         string
	Searchdomain string
	Nameserver   string
	Rootfs       string
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

type ConfigLXC struct {
	Data struct {
		Rootfs       string `json:"rootfs"`
		Swap         int    `json:"swap"`
		Description  string `json:"description"`
		Cores        int    `json:"cores"`
		Hostname     string `json:"hostname"`
		Digest       string `json:"digest"`
		Ostype       string `json:"ostype"`
		Arch         string `json:"arch"`
		Memory       int    `json:"memory"`
		Searchdomain string `json:"searchdomain"`
		Nameserver   string `json:"nameserver"`
		Lock         string `json:"lock"`
		Net0         string `json:"net0"`
	} `json:"data"`
}

func (api *API) ReadConfigLXC(node string, id string) ([]byte, error) {
	path := "/nodes/" + node + "/lxc/" + id + "/config"
	err := api.get(path, nil)
	if err != nil {
		return api.resp, err
	}
	logger.Infof("config lxc %s", string(api.resp))
	return api.resp, nil
}

type ConfigLXCUpdate struct {
	VMID         string
	Node         string
	Hostname     string
	Description  string
	Cores        string
	Memory       string
	Swap         string
	Searchdomain string
	Nameserver   string
	Rootfs       string
}

func (api *API) ConfigLXCUpdate(data ConfigLXCUpdate) error {

	path := "/nodes/" + data.Node + "/lxc/" + data.VMID + "/config" +
		"?hostname=" + data.Hostname +
		"&cores=" + data.Cores +
		"&memory=" + data.Memory +
		"&description=" + data.Description +
		"&swap=" + data.Swap +
		//todo вынести в отдельную ф-ции т.к. требуется перезапуск хоста.
		"&searchdomain=" + data.Searchdomain +
		"&nameserver=" + data.Nameserver

	err := api.put(path, nil)
	if err != nil {
		return err
	}
	logger.Infof("config lxc update %s", string(api.resp))
	err = api.resizeLXC(data)
	if err != nil {
		return err
	}
	return nil
}

func (api *API) resizeLXC(data ConfigLXCUpdate) error {
	path := "/nodes/" + data.Node + "/lxc/" + data.VMID + "/resize"

	options := map[string]string{
		"disk": "rootfs",
		"size": data.Rootfs + "G",
	}
	err := api.put(path, options)
	if err != nil {
		return err
	}
	logger.Infof("disk lxc update %s", string(api.resp))
	return nil
}
