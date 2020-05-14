provider "proxmox" {
  #address proxmox
  address = "https://192.168.122.54:8006/api2/json"
  #user and @ type auth (pve, pam, ldap)
  user = "root@pam"
  #defaul false
  insecure = true
  password = "asdqz123"
}


resource "proxmox_lxc" "test" {
  node = "pve"
  vmid = "603"
  ostemplate = "local:vztmpl/ubuntu-18.04-standard_18.04.1-1_amd64.tar.gz"
  storage = "local-lvm"
  name = "kuber01"
}
