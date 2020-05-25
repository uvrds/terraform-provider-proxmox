package provider

//func Provider() terraform.ResourceProvider {
//	return &schema.Provider{
//		Schema: map[string]*schema.Schema{
//			"address": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			"user": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			"password": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			"insecure": {
//				Type:     schema.TypeBool,
//				Required: true,
//			},
//		},
//		ResourcesMap: map[string]*schema.Resource{
//			"proxmox_lxc": resourceLxc(),
//		},
//		ConfigureFunc: providerConfigure,
//	}
//}
//
//func providerConfigure(d *schema.ResourceData) (interface{}, error) {
//	address := d.Get("address").(string)
//	user := d.Get("user").(string)
//	password := d.Get("password").(string)
//	ins := d.Get("insecure").(bool)
//	return client.NewClient(address, user, password, ins), nil
//}
