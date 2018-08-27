variable "ssh_public_key" {
  description = "SSH public key to add to droplets"
  default = "~/.ssh/bcsb_id_rsa.pub"
}

variable "image_id" {
  description = "DigitalOcean Image ID"
}

variable "k8s_token" {
  description = "Kubeadm token to use for joining clusters"
  default = "px42ot.m26596nf0z9livpg"
}