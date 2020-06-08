package provider

import (
	"encoding/json"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-proxmox/pkg/client"
	"regexp"
	"strconv"
	"strings"
)

func resourceLxc() *schema.Resource {
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
			"ostemplate": {
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
			"swap": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The swap of lxc container",
				Default:     "0",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description lxc container",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The password root",
			},
			"start": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start of lxc container",
				Default:     "1",
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
				Description: "The Rootfs of lxc container",
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
					},
				},
			},
		},

		Create: resourceLxcCreate,
		Read:   resourceLxcRead,
		Update: resourceLxcUpdate,
		Delete: resourceLxcDelete,
	}
}

func resourceLxcCreate(d *schema.ResourceData, m interface{}) error {
	var err error

	apiClient := m.(*Client).proxmox
	apiClient.Cond.L.Lock()
	node := d.Get("node").(string)
	if node == "" {

	}
	vmid := d.Get("vmid").(string)
	if vmid == "" {
		vmid, err = apiClient.NextId()
		if err != nil {
			logger.Fatalf(" id not get %s", err)
		}
	}
	//todo написать 1 ф-цию для подобных конструкций
	var start string
	f := d.Get("start").(bool)
	if f {
		start = "1"
	} else {
		start = "0"
	}
	//

	data := client.Lxc{
		VMID:         vmid,
		Ostemplate:   d.Get("ostemplate").(string),
		Storage:      d.Get("storage").(string),
		Node:         node,
		Hostname:     d.Get("hostname").(string),
		Cores:        d.Get("cores").(string),
		Memory:       d.Get("memory").(string),
		Description:  d.Get("description").(string),
		Start:        start,
		Password:     d.Get("password").(string),
		Swap:         d.Get("swap").(string),
		Searchdomain: d.Get("searchdomain").(string),
		Nameserver:   d.Get("nameserver").(string),
		Rootfs:       d.Get("rootfs").(string),
		Net:          d.Get("network").(*schema.Set).List(),
	}

	d.SetId(vmid)
	err = apiClient.CreateLxc(data)
	if err != nil {
		return err
	}
	apiClient.Cond.L.Unlock()
	return resourceLxcRead(d, m)

}

func resourceLxcRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*Client).proxmox
	apiClient.Cond.L.Lock()
	node := d.Get("node").(string)
	if node == "" {

	}
	resp, err := apiClient.ReadConfigLXC(node, d.Id())
	if err != nil {
		return err
	}
	var stat client.ConfigLXC
	err = json.Unmarshal(resp, &stat)
	if err != nil {
		return err
	}

	err = d.Set("hostname", stat.Data.Hostname)
	if err != nil {
		return err
	}
	coresStr := strconv.Itoa(stat.Data.Cores)
	err = d.Set("cores", coresStr)
	if err != nil {
		return err
	}
	memStr := strconv.Itoa(stat.Data.Memory)
	err = d.Set("memory", memStr)
	if err != nil {
		return err
	}
	swapStr := strconv.Itoa(stat.Data.Swap)
	err = d.Set("swap", swapStr)
	if err != nil {
		return err
	}
	err = d.Set("description", stat.Data.Description)
	if err != nil {
		return err
	}
	err = d.Set("searchdomain", stat.Data.Searchdomain)
	if err != nil {
		return err
	}
	err = d.Set("nameserver", stat.Data.Nameserver)
	if err != nil {
		return err
	}
	if stat.Data.Lock != "create" {
		spl := strings.Split(stat.Data.Rootfs, "=")
		re := regexp.MustCompile("[0-9]+")
		err = d.Set("rootfs", re.FindAllString(spl[1], -1)[0])
		if err != nil {
			return err
		}
	}

	//tmp чтение интерфейса но там маска подсети которую мы не указываем стоит игнорировать этот параметр.
	//var netSlice []interface{}
	//netSlice = append(netSlice,stat.Data.Net0)
	//
	//err = d.Set("net", netSlice)
	//if err != nil {
	//	return err
	//}

	apiClient.Cond.L.Unlock()
	return nil
}

func resourceLxcUpdate(d *schema.ResourceData, m interface{}) error {
	var err error

	apiClient := m.(*Client).proxmox
	apiClient.Cond.L.Lock()
	node := d.Get("node").(string)
	if node == "" {

	}
	data := client.ConfigLXCUpdate{
		VMID:         d.Id(),
		Node:         node,
		Hostname:     d.Get("hostname").(string),
		Description:  d.Get("description").(string),
		Cores:        d.Get("cores").(string),
		Memory:       d.Get("memory").(string),
		Swap:         d.Get("swap").(string),
		Searchdomain: d.Get("searchdomain").(string),
		Nameserver:   d.Get("nameserver").(string),
		Rootfs:       d.Get("rootfs").(string),
	}
	err = apiClient.ConfigLXCUpdate(data)
	if err != nil {
		return err
	}
	apiClient.Cond.L.Unlock()
	return resourceLxcRead(d, m)
}

func resourceLxcDelete(d *schema.ResourceData, m interface{}) error {

	var err error

	apiClient := m.(*Client).proxmox
	apiClient.Cond.L.Lock()

	node := d.Get("node").(string)
	if node == "" {

	}
	data := client.Lxc{
		VMID: d.Id(),
		Node: node,
	}
	d.SetId(d.Id())
	err = apiClient.Deletelxc(data)
	if err != nil {
		return err
	}
	apiClient.Cond.L.Unlock()
	return nil
}
