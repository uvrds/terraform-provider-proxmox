provider "proxmox" {
  #address proxmox
  address = "https://192.168.122.54:8006/api2/json"

  #user and @ type auth (pve, pam, ldap)
  user = "root@pam"

  #TLS defaul false
  insecure = true
  password = "asdqz123"
}

resource "proxmox_lxc" "test" {
  node         = "pve"
  ostemplate   = "local:vztmpl/ubuntu-18.04-standard_18.04.1-1_amd64.tar.gz"
  storage      = "local-lvm"
  hostname     = "kuber03"
  cores        = "1"
  memory       = "523" # 1024MB
  description  = "test2"
  password     = "asdqz123"
  start        = true #старт машины при создании
  swap         = "0"
  searchdomain = "noprod.srv.crpt.tech"
  nameserver   = "192.168.1.1"
  rootfs       = "22" #size disk 10G
  //  count = 10
}

//resource "proxmox_lxc_clone" "test" {
//  vm_id_template = "100"
//  node = "pve"
//  taget_node = "pve"
//  storage = "local-lvm"
//  hostname = "kuber01"
//  cores = "2"
//  memory = "512" # 1024MB
//  swap = "0"
//  description = "test"
//  full = true
//  searchdomain = "noprod.srv.crpt.tech crpt.tech o.crpt.tech"
//  nameserver = "192.168.1.1 8.8.8.8"
//  rootfs = "6"
//}
