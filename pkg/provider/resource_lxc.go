package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-proxmox/pkg/client"
	"log"
)

func resourceLxc() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of lxc container",
			},
			"vmid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of lxc container",
			},
			"ostemplate": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of lxc container",
			},
			"storage": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of lxc container",
			},
			"node": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of lxc container",
			},
		},
		Create: resourceLxcCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,
	}
}

func resourceLxcCreate(d *schema.ResourceData, m interface{}) error {
	var err error
	apiClient := m.(*client.API)
	node := d.Get("node").(string)
	if node == "" {

	}
	vmid := d.Get("vmid").(string)
	if vmid == "" {
		vmid, err = apiClient.NextId()
		if err != nil {
			log.Fatal(err)
		}
	}
	data := client.Lxc{
		VMID:       vmid,
		Ostemplate: d.Get("ostemplate").(string),
		Storage:    d.Get("storage").(string),
		Node:       node,
		Name:       d.Get("name").(string),
	}
	d.SetId(d.Get("vmid").(string))
	err = apiClient.CreateLxc(data)
	if err != nil {
		return err
	}
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
