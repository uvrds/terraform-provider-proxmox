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
		},
		Create: resourceClone,
		Read:   resourceCloneRead,
		Update: resourceCloneUpdate,
		Delete: resourceCloneDelete,
	}
}

func resourceClone(d *schema.ResourceData, m interface{}) error {
	var err error

	apiClient := m.(*client.API)
	apiClient.Cond.L.Lock()
	node := d.Get("node").(string)
	if node == "" {

	}
	vmid := d.Get("vmid").(string)
	if vmid == "" {

		vmid, err = apiClient.NextId()
		if err != nil {
			logger.Fatalf(" id not get %s", err)
		}

	}
	data := client.Lxc{
		VMID:        vmid,
		Ostemplate:  d.Get("ostemplate").(string),
		Storage:     d.Get("storage").(string),
		Node:        node,
		Hostname:    d.Get("hostname").(string),
		Cores:       d.Get("cores").(string),
		Memory:      d.Get("memory").(string),
		Description: d.Get("description").(string),
	}
	d.SetId(vmid)
	err = apiClient.CreateLxc(data)
	if err != nil {
		return err
	}
	apiClient.Cond.L.Unlock()
	return resourceCloneRead(d, m)

}

func resourceCloneRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCloneUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCloneDelete(d *schema.ResourceData, m interface{}) error {

	return nil
}
