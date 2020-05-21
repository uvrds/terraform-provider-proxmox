package model

import "github.com/terraform-provider-proxmox/pkg/model/vmtype"

type virtualMachine struct {
	name   string
	vmType vmtype.Alias
}
