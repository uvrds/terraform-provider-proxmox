package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-provider-proxmox/pkg/client"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"insecure": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"proxmox_lxc":       resourceLxc(),
			"proxmox_lxc_clone": resourceLxcClone(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	address := d.Get("address").(string)
	user := d.Get("user").(string)
	password := d.Get("password").(string)
	ins := d.Get("insecure").(bool)
	return client.NewClient(address, user, password, ins), nil
}
