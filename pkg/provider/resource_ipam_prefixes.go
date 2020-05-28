package provider

import "github.com/hashicorp/terraform/helper/schema"

func resourceIPAMPrefixes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIPAMPrefixesRead,
	}
}
