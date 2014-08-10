VAGRANTFILE_API_VERSION = "2"

$build_script = <<SCRIPT
#!/bin/bash
yum install -y nano wget screen
wget http://rpms.famillecollet.com/enterprise/remi-release-6.rpm
rpm -Uvh remi-release-6*.rpm 
yum -y --enablerepo=remi,remi-php55 install zeromq3.x86_64 zeromq3-devel.x86_64 php-zmq.x86_64 golang hg git 
echo 'export GOPATH=/vagrant/gopath' >> /home/vagrant/.bash_profile
SCRIPT

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.box     = "centos65-x86_64-20140116"
  config.vm.box_url = "https://github.com/2creatives/vagrant-centos/releases/download/v6.5.3/centos65-x86_64-20140116.box"

  config.vm.hostname = "modler.local.dev"

  config.vm.provider :virtualbox do |vb|
    vb.customize ["modifyvm", :id, "--memory", "1024"]
    vb.customize ["modifyvm", :id, "--ioapic", "on"]
    vb.customize ["modifyvm", :id, "--cpus", "4"]
  end

  config.vm.provision :shell, inline: $build_script

end
