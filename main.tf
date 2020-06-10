provider "proxmox" {
  #address proxmox
  address = "https://192.168.122.54:8006/api2/json"

  #user and @ type auth (pve, pam, ldap)
  user = "root@pam"

  #TLS defaul false
  insecure = true
  password = "asdqz123"
}

//resource "proxmox_lxc" "test" {
//  node = "pve"
//  ostemplate = "local:vztmpl/ubuntu-18.04-standard_18.04.1-1_amd64.tar.gz"
//  storage = "local-lvm"
//  hostname = "kuber02"
//  cores = "1"
//  # 1024MB
//  memory = "512"
//  description = "test"
//  password = "asdqz123"
//  #старт машины при создании
//  #start = "1" default = 1
//  swap = "0"
//  searchdomain = "noprod.srv.crpt.tech"
//  nameserver = "192.168.1.1"
//  #size disk 10G
//  rootfs = "22"
//  network {
//    name = "eth0"
//    bridge = "vmbr0"
//    gw = "192.169.122.1"
//    ip = "192.169.122.80/24"
//    firewall = "1"
//  }
//    //count = 10
//}

resource "proxmox_lxc_clone" "test" {
  vm_id_template = "100"
  node = "pve"
  taget_node = "pve"
  storage = "local-lvm"
  hostname = "kuber04"
  cores = "1"
  memory = "1024"
  # 1024MB
  swap = "0"
  description = "clone"
#  full = "1"
  searchdomain = "vk.ru"
  nameserver = "192.168.1.1"
  rootfs = "25"
  network {
    name = "eth0"
    bridge = "vmbr0"
    gw = "192.169.122.1"
    ip = "192.169.122.88/24"
    firewall = "1"
  }
}
