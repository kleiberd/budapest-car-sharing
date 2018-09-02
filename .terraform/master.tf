provider "digitalocean" {}

resource "digitalocean_ssh_key" "default" {
  name = "BCSB K8s Cluster"
  public_key = "${file("${var.ssh_public_key}")}"
}

data "template_file" "api" {
  template = "${file("./services/api.yml")}"

  vars {
    external_ip = "${digitalocean_droplet.bcsb-master.ipv4_address}"
  }
}

resource "digitalocean_droplet" "bcsb-master" {
  name = "bcsb-master"
  image = "${var.image_id}"
  region = "${var.region}"
  size = "${var.size}"
  private_networking = true
  ssh_keys = [
    "${digitalocean_ssh_key.default.fingerprint}"
  ]
}

resource "null_resource" "bcsb-master" {
  triggers {
    cluster_instance_ids = "${digitalocean_droplet.bcsb-master.ipv4_address_private}"
  }

  provisioner "file" {
    connection {
      host = "${digitalocean_droplet.bcsb-master.ipv4_address}"
      type = "ssh"
      user = "root"
      private_key = "${file("${var.ssh_private_key}")}"
    }

    content     = "${data.template_file.api.rendered}"
    destination = "/tmp/api.yml"
  }

  provisioner "remote-exec" {
    connection {
      host = "${digitalocean_droplet.bcsb-master.ipv4_address}"
      type = "ssh"
      user = "root"
      private_key = "${file("${var.ssh_private_key}")}"
    }
    inline = [
      "kubeadm init --token=${var.k8s_token}",

      "export KUBECONFIG=/etc/kubernetes/admin.conf",

      "kubectl apply -f \"https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')\"",

      "kubectl apply -f /tmp/api.yml",
    ]
  }
}

output "master_ip" {
  value = "${digitalocean_droplet.bcsb-master.ipv4_address}"
}
output "master_ip_private" {
  value = "${digitalocean_droplet.bcsb-master.ipv4_address_private}"
}