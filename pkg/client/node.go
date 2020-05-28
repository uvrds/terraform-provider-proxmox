package client

import (
	"encoding/json"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
)

type Nodes struct {
	Data []struct {
		Status         string  `json:"status"`
		SslFingerprint string  `json:"ssl_fingerprint"`
		Maxmem         int64   `json:"maxmem"`
		Mem            int     `json:"mem"`
		Uptime         int     `json:"uptime"`
		Disk           int64   `json:"disk"`
		Level          string  `json:"level"`
		ID             string  `json:"id"`
		Maxcpu         int     `json:"maxcpu"`
		Node           string  `json:"node"`
		Type           string  `json:"type"`
		Maxdisk        int64   `json:"maxdisk"`
		CPU            float64 `json:"cpu"`
	} `json:"data"`
}

func (api *API) Nodes() (Nodes, error) {
	path := "/nodes"
	err := api.get(path, nil)
	if err != nil {
		return Nodes{}, err
	}
	var data Nodes
	err = json.Unmarshal(api.resp, &data)
	logger.Infof("nodes %s", string(api.resp))
	if err != nil {
		return Nodes{}, err
	}
	return data, nil
}

type Node struct {
	Data []struct {
		Name string `json:"name"`
	} `json:"data"`
}

func (api *API) Node(node string) (Node, error) {
	path := "/nodes/" + node
	err := api.get(path, nil)
	if err != nil {
		return Node{}, err
	}
	var data Node
	err = json.Unmarshal(api.resp, &data)
	logger.Infof("node %s: %s", node, string(api.resp))
	if err != nil {
		return Node{}, err
	}
	return data, nil
}

type NodeStatus struct {
	Data struct {
		CPU     int      `json:"cpu"`
		Loadavg []string `json:"loadavg"`
		Ksm     struct {
			Shared int `json:"shared"`
		} `json:"ksm"`
		Kversion string `json:"kversion"`
		Memory   struct {
			Total int64 `json:"total"`
			Free  int64 `json:"free"`
			Used  int   `json:"used"`
		} `json:"memory"`
		Pveversion string      `json:"pveversion"`
		Wait       interface{} `json:"wait"`
		Uptime     int         `json:"uptime"`
		Cpuinfo    struct {
			Cpus    int    `json:"cpus"`
			Mhz     string `json:"mhz"`
			Model   string `json:"model"`
			Sockets int    `json:"sockets"`
			UserHz  int    `json:"user_hz"`
			Hvm     string `json:"hvm"`
			Flags   string `json:"flags"`
		} `json:"cpuinfo"`
		Idle   int `json:"idle"`
		Rootfs struct {
			Avail int64 `json:"avail"`
			Used  int64 `json:"used"`
			Free  int64 `json:"free"`
			Total int64 `json:"total"`
		} `json:"rootfs"`
		Swap struct {
			Used  int   `json:"used"`
			Free  int64 `json:"free"`
			Total int64 `json:"total"`
		} `json:"swap"`
	} `json:"data"`
}

func (api *API) NodeStatus(node string) (NodeStatus, error) {
	path := "/nodes/" + node + "/status"
	err := api.get(path, nil)
	if err != nil {
		return NodeStatus{}, err
	}
	var data NodeStatus
	err = json.Unmarshal(api.resp, &data)
	logger.Infof("node status %s: %s", node, string(api.resp))
	if err != nil {
		return NodeStatus{}, err
	}
	return data, nil
}
