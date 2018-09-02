resource "digitalocean_droplet" "bcsb-worker" {
  count = "${var.k8s_workers_count}"
  name = "bcsb-worker-${count.index}"

  image = "${var.image_id}"
  region = "${var.region}"
  size = "${var.size}"
  private_networking = true
  ssh_keys = [
    "${digitalocean_ssh_key.default.fingerprint}"
  ]
}

resource "null_resource" "bcsb-worker" {
  count = "${var.k8s_workers_count}"

  triggers {
    cluster_ids = "${join(",", digitalocean_droplet.bcsb-worker.*.ipv4_address_private)},${null_resource.bcsb-master.id}"
  }

  provisioner "remote-exec" {
    connection {
      host = "${element(digitalocean_droplet.bcsb-worker.*.ipv4_address, count.index)}"
      type = "ssh"
      user = "root"
      private_key = "${file("${var.ssh_private_key}")}"
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