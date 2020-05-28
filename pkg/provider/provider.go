package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: providerSchema(),
		//DataSourcesMap: providerDataSourcesMap(),
		ResourcesMap:  providerResources(),
		ConfigureFunc: providerConfigure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ipam_address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Hostname of IPAM service",
		},
	}
}

//func providerDataSourcesMap() map[string]*schema.Resource {
//}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"proxmox_ipam_prefixes": resourceIPAMPrefixes(),
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
	}
}
