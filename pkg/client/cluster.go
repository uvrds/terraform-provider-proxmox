package client

import (
	"encoding/json"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
)

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

func (api *API) searchLxc(node string, vmid string) error {

	return nil
}

type Cluster struct {
	Data []struct {
		Name string `json:"name"`
	} `json:"data"`
}

func (api *API) Cluster() (Cluster, error) {

	path := "/cluster"
	err := api.get(path, nil)
	if err != nil {
		return Cluster{}, err
	}
	var data Cluster
	err = json.Unmarshal(api.resp, &data)
	logger.Infof("cluster %s", string(api.resp))
	if err != nil {
		return Cluster{}, err
	}
	return data, nil
}

type ClusterStatus struct {
	Data []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		IP     string `json:"ip"`
		Level  string `json:"level"`
		Nodeid int    `json:"nodeid"`
		Type   string `json:"type"`
		Online int    `json:"online"`
		Local  int    `json:"local"`
	} `json:"data"`
}

func (api *API) ClusterStatus() (ClusterStatus, error) {

	path := "/cluster/status"
	err := api.get(path, nil)
	if err != nil {
		return ClusterStatus{}, err
	}
	var data ClusterStatus
	err = json.Unmarshal(api.resp, &data)
	logger.Infof("status cluster %s", string(api.resp))
	if err != nil {
		return ClusterStatus{}, err
	}
	return data, nil
}

type ClusterResources struct {
	Data []struct {
		Diskread  int         `json:"diskread,omitempty"`
		Node      string      `json:"node"`
		Maxmem    int         `json:"maxmem,omitempty"`
		Diskwrite int         `json:"diskwrite,omitempty"`
		Mem       int         `json:"mem,omitempty"`
		Status    string      `json:"status"`
		ID        string      `json:"id"`
		Netout    int         `json:"netout,omitempty"`
		Uptime    int         `json:"uptime,omitempty"`
		Maxdisk   int64       `json:"maxdisk"`
		Maxcpu    int         `json:"maxcpu,omitempty"`
		Type      string      `json:"type"`
		Template  int         `json:"template,omitempty"`
		CPU       interface{} `json:"cpu,omitempty"`
		Vmid      int         `json:"vmid,omitempty"`
		Name      string      `json:"name,omitempty"`
		Netin     int         `json:"netin,omitempty"`
		Disk      int         `json:"disk"`
		Level     string      `json:"level,omitempty"`
		Storage   string      `json:"storage,omitempty"`
		Shared    int         `json:"shared,omitempty"`
	} `json:"data"`
}

func (api *API) ClusterResources() (ClusterResources, error) {
	path := "/cluster/resources"
	err := api.get(path, nil)
	if err != nil {
		return ClusterResources{}, err
	}
	var data ClusterResources
	err = json.Unmarshal(api.resp, &data)
	logger.Infof("resources  cluster %s", string(api.resp))
	if err != nil {
		return ClusterResources{}, err
	}
	return data, nil
}
