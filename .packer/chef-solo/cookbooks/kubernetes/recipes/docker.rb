#
# Cookbook:: kubernetes
# Recipe:: digitalocean
#
# Copyright:: 2018, David Kleiber, All Rights Reserved.

execute 'echo' do
  command 'echo "Install Docker"'
  action :run
end

execute 'update-upgrade' do
  command 'apt-get update'
  action :run
end

package %w(apt-transport-https ca-certificates curl software-properties-common) do
  action :install
end