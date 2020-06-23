package provider

import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-proxmox/pkg/client"
)

func resourceLxcClone() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of lxc container",
			},
			"vmid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of lxc container",
			},
			"vm_id_template": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Template for lxc container",
			},
			"storage": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Storage lxc container",
			},
			"node": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The node of lxc container",
			},
			"cores": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The cores of lxc container",
			},
			"memory": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The memory of lxc container",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description lxc container",
			},
			"full": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The full copy storage",
				Default:     "1",
			},
			"swap": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The swap lxc container",
				Default:     "0",
			},
			"searchdomain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The searchdomain of lxc container",
			},
			"nameserver": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Nameserver of lxc container",
			},
			"rootfs": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The rootfs of lxc container",
			},
			"network": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						"bridge": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"gw": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"firewall": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"gw6": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"hwaddr": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip6": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"mtu": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"rate": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"trunks": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
		Create: resourceCloneCreate,
		Read:   resourceLxcRead,
		Update: resourceLxcUpdate,
		Delete: resourceLxcDelete,
	}
}

func resourceCloneCreate(d *schema.ResourceData, m interface{}) error {
	var err error

	apiClient := m.(*Client).proxmox
	apiClient.Cond.L.Lock()
	nodeT, err := apiClient.GetNodeTemplateLxc(d.Get("vm_id_template").(string))
	if err != nil {
		return err
	}

	targetNode := d.Get("node").(string)
	//Здесь должен выдать нам ноду ресурсменеджер
	if targetNode == "" {

	}

	vmid := d.Get("vmid").(string)
	if vmid == "" {
		vmid, err = apiClient.NextId()
		if err != nil {
			logger.Fatalf(" id not get %s", err)
		}
	}
	data := client.LxcClone{
		VMID:         d.Get("vm_id_template").(string),
		NEWID:        vmid,
		Storage:      d.Get("storage").(string),
		Node:         nodeT,
		TargetNode:   targetNode,
		Hostname:     d.Get("hostname").(string),
		Full:         d.Get("full").(string),
		Net:          d.Get("network").(*schema.Set),
		Description:  d.Get("description").(string),
		Cores:        d.Get("cores").(string),
		Memory:       d.Get("memory").(string),
		Swap:         d.Get("swap").(string),
		Searchdomain: d.Get("searchdomain").(string),
		Nameserver:   d.Get("nameserver").(string),
		Rootfs:       d.Get("rootfs").(string),
	}

	d.SetId(vmid)
	err = apiClient.LxcMigrate(data)
	if err != nil {
		return err
	}
	err = apiClient.CloneLxc(data)
	if err != nil {
		return err
	}
	apiClient.Cond.L.Unlock()
	return resourceLxcRead(d, m)
}
