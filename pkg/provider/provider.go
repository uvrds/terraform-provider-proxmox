package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: providerSchema(),
		//DataSourcesMap: providerDataSourcesMap(),
		//ResourcesMap:   providerResourcesMap(),
		ConfigureFunc: providerConfigure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ipam_addresses": {
			Type:     schema.TypeList,
			Required: true,
		},
	}
}

//func providerDataSourcesMap() map[string]*schema.Resource {
//}

//func providerResourcesMap() map[string]*schema.Resource {
//}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := getProviderConfig(d)
	return config.Client()
}

func getProviderConfig(dirty_config *schema.ResourceData) *Config {
	return &Config{
		IPAMEndpoint: dirty_config.Get("ipam_address").(string),
		IPAMToken:    dirty_config.Get("ipam_token").(string),
	}
}
