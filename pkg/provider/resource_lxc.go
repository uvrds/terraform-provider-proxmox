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
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
			},
		},
		Create: resourceCreateLxc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateLxc(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.API)

	name := d.Get("name").(string)

	err := apiClient.Authenticate()
	err = apiClient.CreateLxc(name)
	if err != nil {
		return err
	}
	return nil
}
