provider "proxmox" {
  address = "https://192.168.122.54:8006/" #address proxmox
  user = "root@pam" #user and @ type auth (pve, pam, ldap)
  insecure = true #defaul false
  password = "asdqz123"
}

