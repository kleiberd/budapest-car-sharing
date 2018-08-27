provider "digitalocean" {}

resource "digitalocean_ssh_key" "default" {
  name = "BCSB K8s Cluster"
  public_key = "${file("${var.ssh_public_key}")}"
}

resource "digitalocean_droplet" "bcsb-master" {
  image = "${var.image_id}"
  name = "bcsb-master"
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
      "sudo kubeadm init --token=${var.k8s_token} --apiserver-advertise-address=${digitalocean_droplet.bcsb-master.ipv4_address_private}",

      "mkdir -p $HOME/.kube",
      "sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config",
      "sudo chown $(id -u):$(id -g) $HOME/.kube/config",

      "kubectl get pods --all-namespaces",
    ]
  }
}

output "master_ip" {
  value = "${digitalocean_droplet.bcsb-master.ipv4_address}"
}
output "master_ip_private" {
  value = "${digitalocean_droplet.bcsb-master.ipv4_address_private}"
}