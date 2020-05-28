package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
)

func dataSourceIPAMPrefixes() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceIPAMPrefixesRead,
		Schema: dataSourcePrefixesSchema(),
	}

}

func dataSourceIPAMPrefixesRead(d *schema.ResourceData, meta interface{}) error {

	c := meta.(*Client).ipamClient

	var par = ipam.NewIpamPrefixesListParams()
	res, err := c.Ipam.IpamPrefixesList(par, nil)

	if err == nil {
		fmt.Printf("PREFIXES COUNT %d", res.Payload.Count)
	} else {
		return err
	}

	return nil

}

func dataSourcePrefixesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"prefixes_count": {Type: schema.TypeInt},
	}
}
