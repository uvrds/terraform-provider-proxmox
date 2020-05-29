package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:         providerSchema(),
		DataSourcesMap: providerDataSourcesMap(),
		ResourcesMap:   providerResources(),
		ConfigureFunc:  providerConfigure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ipam_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Hostname of IPAM service",
		},

		"ipam_token": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Token for IPAM connections",
		},
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
	}
}

func providerDataSourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"proxmox_ipam_prefixes": dataSourceIPAMPrefixes(),
	}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"proxmox_ipam_prefixes": resourceIPAMPrefixes(),
		"proxmox_lxc":           resourceLxc(),
		"proxmox_lxc_clone":     resourceLxcClone(),
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := getProviderConfig(d)
	return config.Client()
}

func getProviderConfig(dirtyConfig *schema.ResourceData) *Config {
	return &Config{
		IPAMEndpoint: dirtyConfig.Get("ipam_address").(string),
		IPAMToken:    dirtyConfig.Get("ipam_token").(string),
		Address:      dirtyConfig.Get("address").(string),
		User:         dirtyConfig.Get("user").(string),
		Password:     dirtyConfig.Get("password").(string),
		Insecure:     dirtyConfig.Get("insecure").(bool),
	}
}
