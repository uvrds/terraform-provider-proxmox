package provider

import (
	api "github.com/netbox-community/go-netbox/netbox"
	"github.com/netbox-community/go-netbox/netbox/client"
	proxmox "github.com/terraform-provider-proxmox/pkg/client"
)

type Config struct {
	IPAMEndpoint string
	IPAMToken    string
	Address      string
	User         string
	Password     string
	Insecure     bool
}

type Client struct {
	ipamClient    *client.NetBox
	configuration *Config
	proxmox       *proxmox.API
}

func (c *Config) Client() (interface{}, error) {
	cli := api.NewNetboxWithAPIKey(c.IPAMEndpoint, c.IPAMToken)
	cliProxmox := proxmox.NewClient(c.Address, c.User, c.Password, c.Insecure)

	conn := Client{
		ipamClient:    cli,
		proxmox:       cliProxmox,
		configuration: c,
	}
	return &conn, nil
}
