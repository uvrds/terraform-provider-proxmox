//package provider
//
//import (
//	"fmt"
//	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
//	"github.com/hashicorp/terraform/helper/schema"
//	"github.com/terraform-provider-proxmox/pkg/client"
//)
//
//func resourceLxc() *schema.Resource {
//	fmt.Print()
//	return &schema.Resource{
//		Schema: map[string]*schema.Schema{
//			"hostname": {
//				Type:        schema.TypeString,
//				Required:    true,
//				Description: "The name of lxc container",
//			},
//			"vmid": {
//				Type:        schema.TypeString,
//				Optional:    true,
//				Description: "The id of lxc container",
//			},
//			"ostemplate": {
//				Type:        schema.TypeString,
//				Required:    true,
//				Description: "Template for lxc container",
//			},
//			"storage": {
//				Type:        schema.TypeString,
//				Required:    true,
//				Description: "Storage lxc container",
//			},
//			"node": {
//				Type:        schema.TypeString,
//				Required:    true,
//				Description: "The id of lxc container",
//			},
//			"cores": {
//				Type:        schema.TypeString,
//				Required:    true,
//				Description: "The id of lxc container",
//			},
//			"memory": {
//				Type:        schema.TypeString,
//				Required:    true,
//				Description: "The id of lxc container",
//			},
//			"description": {
//				Type:        schema.TypeString,
//				Required:    true,
//				Description: "The id of lxc container",
//			},
//		},
//		Create: resourceLxcCreate,
//		Read:   resourceServerRead,
//		Update: resourceServerUpdate,
//		Delete: resourceServerDelete,
//	}
//}
//
//func resourceLxcCreate(d *schema.ResourceData, m interface{}) error {
//	var err error
//
//	apiClient := m.(*client.API)
//	apiClient.Cond.L.Lock()
//	node := d.Get("node").(string)
//	if node == "" {
//
//	}
//	vmid := d.Get("vmid").(string)
//	if vmid == "" {
//
//		vmid, err = apiClient.NextId()
//		if err != nil {
//			logger.Fatalf(" id not get %s", err)
//		}
//
//	}
//	data := client.Lxc{
//		VMID:        vmid,
//		Ostemplate:  d.Get("ostemplate").(string),
//		Storage:     d.Get("storage").(string),
//		Node:        node,
//		Hostname:    d.Get("hostname").(string),
//		Cores:       d.Get("cores").(string),
//		Memory:      d.Get("memory").(string),
//		Description: d.Get("description").(string),
//	}
//	d.SetId(vmid)
//	err = apiClient.CreateLxc(data)
//	if err != nil {
//		return err
//	}
//	apiClient.Cond.L.Unlock()
//	return resourceServerRead(d, m)
//
//}
//
//func resourceServerRead(d *schema.ResourceData, m interface{}) error {
//	return nil
//}
//
//func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
//	return nil
//}
//
//func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
//	var err error
//
//	apiClient := m.(*client.API)
//	apiClient.Cond.L.Lock()
//	node := d.Get("node").(string)
//	if node == "" {
//
//	}
//	data := client.Lxc{
//		VMID: d.Id(),
//		Node: node,
//	}
//	d.SetId(d.Id())
//	err = apiClient.Delete_lxc(data)
//	if err != nil {
//		return err
//	}
//	apiClient.Cond.L.Unlock()
//	return nil
//}
