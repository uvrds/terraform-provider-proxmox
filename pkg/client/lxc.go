package client

import (
	"encoding/json"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
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
	Net          *schema.Set
}

func (api *API) CreateLxc(data Lxc) error {
	time.Sleep(time.Second * 2)
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
	}
	path := "/nodes/" + data.Node + "/lxc"
	err := api.post(path, options)
	if err != nil {
		return err
	}
	logger.Infof("create lxc %s", string(api.resp))

	err = api.ConfigLXCUpdateNetwork(data.Net, data.Node, data.VMID)
	if err != nil {
		return err
	}
	return nil
}

func (api *API) Deletelxc(data Lxc) error {
	err := api.stopLxc(data.Node, data.VMID)
	if err != nil {
		return err
	}
	var s = true
	for s {
		time.Sleep(time.Second * 5)
		path := "/nodes/" + data.Node + "/lxc/" + data.VMID + "?purge=1"
		err = api.del(path, nil)
		if err != nil {
			return err
		}
		b, err := api.CheckLxc(data.Node, data.VMID)
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
		b, err := api.CheckLxc(node, vmid)
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

func (api *API) startLxc(node string, vmid string) error {

	var s = true

	for s {
		path := "/nodes/" + node + "/lxc/" + vmid + "/status/start"
		err := api.post(path, nil)
		if err != nil {
			return err
		}

		b, err := api.CheckLxc(node, vmid)
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
		if stat.Data.Status == "running" {
			logger.Infof("start lxc id:%s %s", vmid, string(api.resp))
			s = false
		}
	}
	return nil
}

type CheckLxc struct {
	Data interface{} `json:"data"`
}

func (api *API) CheckLxc(node string, vmid string) (bool, error) {
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
	Net          *schema.Set
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
	path := "/nodes/" + data.TargetNode + "/lxc/" + data.VMID + "/clone"
	err := api.post(path, options)
	if err != nil {
		return err
	}
	logger.Infof("clone lxc %s", string(api.resp))
	var wait = true
	for wait {
		resp, err := api.StatusLXC(data.TargetNode, data.NEWID)
		if err != nil {
			return err
		}
		var stat StatusLXC
		err = json.Unmarshal(resp, &stat)
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
		logger.Infof("lock: %s", stat.Data.Lock)
		if stat.Data.Lock == "" {
			wait = false
		}
	}
	err = api.ConfigLXCUpdateNetwork(data.Net, data.TargetNode, data.NEWID)
	if err != nil {
		return err
	}

	DataClone := ConfigLXCUpdate{
		VMID:         data.NEWID,
		Node:         data.TargetNode,
		Hostname:     data.Hostname,
		Description:  data.Description,
		Cores:        data.Cores,
		Memory:       data.Memory,
		Swap:         data.Swap,
		Searchdomain: data.Searchdomain,
		Nameserver:   data.Nameserver,
		Rootfs:       data.Rootfs,
		Net:          data.Net,
	}
	err = api.ConfigLXCUpdate(DataClone)
	if err != nil {
		return err
	}
	err = api.startLxc(data.TargetNode, data.NEWID)
	if err != nil {
		return err
	}
	return nil
}

type ReadConfigLXC struct {
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
	Net          *schema.Set
}

func (api *API) ConfigLXCUpdate(data ConfigLXCUpdate) error {

	path := "/nodes/" + data.Node + "/lxc/" + data.VMID + "/config"
	options := map[string]string{
		"hostname":     data.Hostname,
		"cores":        data.Cores,
		"memory":       data.Memory,
		"swap":         data.Swap,
		"description":  data.Description,
		"searchdomain": data.Searchdomain,
		"nameserver":   data.Nameserver,
	}
	err := api.put(path, options)
	if err != nil {
		return err
	}
	logger.Infof("config lxc update %s", string(api.resp))
	err = api.resizeLXC(data)
	if err != nil {
		return err
	}
	err = api.ConfigLXCUpdateNetwork(data.Net, data.Node, data.VMID)
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

func (api *API) ConfigLXCUpdateNetwork(net *schema.Set, node string, vmid string) error {
	time.Sleep(time.Second * 2)
	var res string
	options := make(map[string]string)
	for k, v := range net.List() {
		for key, value := range v.(map[string]interface{}) {
			if value.(string) == "" {
				continue
			}
			res += key + "=" + value.(string) + ","
		}
		options["net"+strconv.Itoa(k)] += res
		res = ""
	}
	logger.Infof("options: %s", options)
	path := "/nodes/" + node + "/lxc/" + vmid + "/config"
	err := api.put(path, options)
	if err != nil {
		return err
	}
	logger.Infof("config lxc update network %s", string(api.resp))

	return nil
}

//search template id on nodes

func (api *API) GetNodeTemplateLxc(vmid string) (string, error) {

	var nodes []string
	var tnode string
	resp, err := api.Nodes()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range resp.Data {
		nodes = append(nodes, v.Node)
	}
	for _, v := range nodes {
		b, err := api.lxcVmid(v, vmid)
		if err != nil {
			log.Fatal(err)
		}
		if b == true {
			tnode = v
			break
		}
	}
	logger.Infof("node template:%s", tnode)

	return tnode, nil
}

type LxcVmid struct {
	Data []struct {
		Subdir string `json:"subdir"`
	} `json:"data"`
}

func (api *API) lxcVmid(node string, vmid string) (bool, error) {
	path := "/nodes/" + node + "/lxc/" + vmid
	err := api.get(path, nil)
	if err != nil {
		return false, err
	}
	logger.Infof("search lxc id:%s on node:%s", vmid, node)

	var stat LxcVmid
	err = json.Unmarshal(api.resp, &stat)
	if err != nil {
		return false, err
	}
	if stat.Data != nil {
		logger.Infof("found lxc id:%s on node:%s", vmid, node)
		return true, nil
	}
	logger.Infof("no found lxc id:%s on node:%s", vmid, node)
	return false, nil
}

//
// migrate template
func (api *API) LxcMigrate(data LxcClone) error {

	path := "/nodes/" + data.Node + "/lxc/" + data.VMID + "/migrate?target=" + data.TargetNode
	err := api.post(path, nil)
	if err != nil {
		return err
	}
	time.Sleep(time.Second * 5)
	logger.Infof("migrate template on node:%s", data.TargetNode)

	return nil
}
