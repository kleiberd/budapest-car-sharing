resource "digitalocean_firewall" "k8s" {
  name = "22-80"

  droplet_ids = ["${digitalocean_droplet.bcsb-master.id}"]

  inbound_rule = [
    {
      protocol = "tcp"
      port_range = "1-65535"

      source_addresses = [
        "${digitalocean_droplet.bcsb-worker.*.ipv4_address_private}",
        "${digitalocean_droplet.bcsb-worker.*.ipv4_address}",
        "${digitalocean_droplet.bcsb-master.ipv4_address_private}",
        "${digitalocean_droplet.bcsb-master.ipv4_address}",
      ]
    },
    {
      protocol = "udp"
      port_range = "1-65535"

      source_addresses = [
        "${digitalocean_droplet.bcsb-worker.*.ipv4_address_private}",
        "${digitalocean_droplet.bcsb-worker.*.ipv4_address}",
        "${digitalocean_droplet.bcsb-master.ipv4_address_private}",
        "${digitalocean_droplet.bcsb-master.ipv4_address}",
      ]
    },
    {
      protocol = "icmp"
      port_range = "1-65535"

      source_addresses = [
        "${digitalocean_droplet.bcsb-worker.*.ipv4_address_private}",
        "${digitalocean_droplet.bcsb-worker.*.ipv4_address}",
        "${digitalocean_droplet.bcsb-master.ipv4_address_private}",
        "${digitalocean_droplet.bcsb-master.ipv4_address}",
      ]
    },
    {
      protocol = "tcp"
      port_range = "22"
      source_addresses = ["0.0.0.0/0", "::/0"]
    },
    {
      protocol = "tcp"
      port_range = "80"
      source_addresses = ["0.0.0.0/0", "::/0"]
    },
  ]

  outbound_rule = [
    {
      protocol = "tcp"
      port_range = "1-65535"
      destination_addresses = ["0.0.0.0/0", "::/0"]
    },
    {
      protocol = "udp"
      port_range = "1-65535"
      destination_addresses = ["0.0.0.0/0", "::/0"]
    },
    {
      protocol = "icmp"
      port_range = "1-65535"
      destination_addresses = ["0.0.0.0/0", "::/0"]
    },
  ]
}