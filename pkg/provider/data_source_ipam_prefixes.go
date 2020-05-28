package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
)

func dataSourceIPAMPrefixesRead(d *schema.ResourceData, meta interface{}) error {

	c := meta.(*Client).ipamClient

	var par = ipam.NewIpamPrefixesListParams()
	res, err := c.Ipam.IpamPrefixesList(par, nil)

	if err == nil {
		fmt.Printf("%d", res.Payload.Count)
	} else {
		return err
	}

	return nil

}
