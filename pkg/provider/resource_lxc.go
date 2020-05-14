package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-proxmox/pkg/client"
)

func resourceLxc() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of lxc container",
			},
		},
		Create: resourceLxcCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,
	}
}

func resourceLxcCreate(d *schema.ResourceData, m interface{}) error {

	apiClient := m.(*client.API)
	node := d.Get("node").(string)
	d.SetId(node)
	err := apiClient.CreateLxc(node)
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
