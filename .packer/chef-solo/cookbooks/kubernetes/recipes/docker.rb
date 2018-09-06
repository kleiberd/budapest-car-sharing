#
# Cookbook:: kubernetes
# Recipe:: docker
#
# Copyright:: 2018, David Kleiber, All Rights Reserved.

execute 'echo' do
  command 'echo "Install Docker"'
  action :run
end

docker_installation_package 'default' do
  version '17.03.0'
  action :create
end