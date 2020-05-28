package provider

import (
	api "github.com/netbox-community/go-netbox/netbox"
	"github.com/netbox-community/go-netbox/netbox/client"
)

type Config struct {
	IPAMEndpoint string
	IPAMToken    string
}

type Client struct {
	ipamClient    *client.NetBox
	configuration *Config
}

func (c *Config) Client() (interface{}, error) {
	cli := api.NewNetboxWithAPIKey(c.IPAMEndpoint, c.IPAMToken)

	conn := Client{
		ipamClient:    cli,
		configuration: c,
	}

	return &conn, nil
}
