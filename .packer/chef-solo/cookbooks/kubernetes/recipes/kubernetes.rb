#
# Cookbook:: kubernetes
# Recipe:: kubernetes
#
# Copyright:: 2018, David Kleiber, All Rights Reserved.

execute 'echo' do
  command 'echo "Install Kubernetes"'
end

apt_repository 'kubernetes-key' do
  uri 'http://apt.kubernetes.io/'
  components ['main']
  distribution 'kubernetes-xenial'
  keyserver 'https://packages.cloud.google.com/apt/doc/apt-key.gpg'
end

apt_update

package %w(kubelet kubeadm kubectl) do
  version '1.11.2-00'
  options '--allow-unauthenticated'
end

execute 'weave bridge' do
  command 'sysctl net.bridge.bridge-nf-call-iptables=1'
end
