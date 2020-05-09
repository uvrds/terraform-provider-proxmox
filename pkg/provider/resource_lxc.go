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
				Description: "The name of the resource, also acts as it's unique ID",
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

	name := d.Get("name").(string)
	d.SetId(name)
	err := apiClient.Authenticate()
	err = apiClient.CreateLxc(name)
	if err != nil {
		return err
	}
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
