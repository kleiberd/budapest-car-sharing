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
      "kubeadm join --token=${var.k8s_token} ${digitalocean_droplet.bcsb-master.ipv4_address_private}:6443 --discovery-token-unsafe-skip-ca-verification"
    ]
  }
}

output "workers_ips" {
  value = "${digitalocean_droplet.bcsb-worker.*.ipv4_address}"
}
output "workers_ips_private" {
  value = "${digitalocean_droplet.bcsb-worker.*.ipv4_address_private}"
}