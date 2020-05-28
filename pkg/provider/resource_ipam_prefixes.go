package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIPAMPrefixes() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceIPAMPrefixesRead,
		Delete: resourceIPAMPrefixesDelete,
	}
}

func resourceIPAMPrefixesDelete(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("Deleted")

	return nil
}
