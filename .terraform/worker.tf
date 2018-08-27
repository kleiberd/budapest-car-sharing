resource "digitalocean_droplet" "bcsb-worker" {
  count = 2
  image = "${var.image_id}"
  name = "bcsb-worker-${count.index}"
  region = "fra1"
  size = "2gb"
  private_networking = true
  ssh_keys = [
    "${digitalocean_ssh_key.default.fingerprint}"
  ]

  provisioner "remote-exec" {
    connection {
      type = "ssh"
      user = "root"
    }
    inline = [
      "sudo kubeadm join --token=${var.k8s_token} --discovery-token-unsafe-skip-ca-verification ${digitalocean_droplet.bcsb-master.ipv4_address_private}:6443",
    ]
  }
}

output "workers_ips" {
  value = "${digitalocean_droplet.bcsb-master.*.ipv4_address}"
}
output "workers_ips_private" {
  value = "${digitalocean_droplet.bcsb-master.*.ipv4_address_private}"
}