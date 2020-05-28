package provider

import (
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
		d.Set("prefixes_count", res.Payload.Count)
	} else {
		return err
	}

	return nil

}

func barePrefixesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"prefixes_count": {Type: schema.TypeInt},
	}
}

func dataSourcePrefixesSchema() map[string]*schema.Schema {
	s := barePrefixesSchema()

	for k, v := range s {
		switch k {
		case "prefixes_count":
			v.Optional = true

		}
	}

	return s
}
