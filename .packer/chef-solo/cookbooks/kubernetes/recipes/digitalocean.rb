#
# Cookbook:: kubernetes
# Recipe:: digitalocean
#
# Copyright:: 2018, David Kleiber, All Rights Reserved.

execute 'echo' do
  command 'echo "Install DigitalOcean Monitoring"'
  action :run
end

#It works only DigitalOcean servers, TODO Solve with ENV.
#execute 'install' do
#  command 'curl -sSL https://agent.digitalocean.com/install.sh | sh'
#end