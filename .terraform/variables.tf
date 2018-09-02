variable "region" {
  description = "Digitalocean region"
  default = "fra1"
}

variable "size" {
  description = "Digitalocean droplet size"
  default = "2gb"
}

variable "ssh_public_key" {
  description = "SSH public key to add to droplets"
  default = "~/.ssh/bcsb_id_rsa.pub"
}

variable "ssh_private_key" {
  description = "SSH private key to add to droplets"
  default = "~/.ssh/bcsb_id_rsa"
}

variable "image_id" {
  description = "Digitalocean Image ID"
}

variable "k8s_token" {
  description = "Kubeadm token to use for joining clusters"
  default = "px42ot.m26596nf0z9livpg"
}

variable "k8s_workers_count" {
  description = "Number of Kubernetes worker nodes"
  default = 2
}