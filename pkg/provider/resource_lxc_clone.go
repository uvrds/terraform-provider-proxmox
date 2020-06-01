package provider

import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-proxmox/pkg/client"
)

func resourceLxcClone() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of lxc container",
			},
			"vmid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of lxc container",
			},
			"taget_node": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The full copy storage",
			},
			"vm_id_template": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Template for lxc container",
			},
			"storage": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Storage lxc container",
			},
			"node": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The node of lxc container",
			},
			"cores": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The cores of lxc container",
			},
			"memory": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The memory of lxc container",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description lxc container",
			},
			"full": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The full copy storage",
			},
			"swap": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The swap lxc container",
			},
			"searchdomain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The searchdomain of lxc container",
			},
		},
		Create: resourceCloneCreate,
		Read:   resourceLxcRead,
		Update: resourceLxcUpdate,
		Delete: resourceLxcDelete,
	}
}

func resourceCloneCreate(d *schema.ResourceData, m interface{}) error {
	var err error

	apiClient := m.(*Client).proxmox
	apiClient.Cond.L.Lock()
	node := d.Get("node").(string)
	if node == "" {

	}
	targetNode := d.Get("taget_node").(string)
	if targetNode == "" {

	}
	vmid := d.Get("vmid").(string)
	if vmid == "" {
		vmid, err = apiClient.NextId()
		if err != nil {
			logger.Fatalf(" id not get %s", err)
		}

	}
	var full string
	f := d.Get("full").(bool)
	if f {
		full = "1"
	} else {
		full = "0"
	}

	data := client.LxcClone{
		VMID:         d.Get("vm_id_template").(string),
		NEWID:        vmid,
		Storage:      d.Get("storage").(string),
		Node:         node,
		TargetNode:   targetNode,
		Hostname:     d.Get("hostname").(string),
		Description:  d.Get("description").(string),
		Full:         full,
		Cores:        d.Get("cores").(string),
		Memory:       d.Get("memory").(string),
		Swap:         d.Get("swap").(string),
		Searchdomain: d.Get("searchdomain").(string),
	}

	d.SetId(vmid)
	err = apiClient.CloneLxc(data)
	if err != nil {
		return err
	}
	apiClient.Cond.L.Unlock()
	return resourceLxcRead(d, m)
}
