#!/bin/bash



systeminfo=`cat /etc/redhat-release`
kernel=`uname -a`

echo "####SystemInfo####"
echo "$systeminfo"

echo "####Kernel Info####"
echo $kernel


echo " "
echo "update ulimit"
echo "root soft nofile 65535
root hard nofile 65535
* soft nofile 65535
* hard nofile 65535" >> /etc/security/limits.conf


echo "####Disable Firewalld####"

systemctl stop firewalld
systemctl disable firewalld
fireinfo=`firewall-cmd --state`
echo $fireinfo

echo "####add history commod Time####"

echo 'export HISTTIMEFORMAT="%Y-%m-%d %H:%M:%S  "' >> /etc/profile
source /etc/profile

echo "####add garbage####"
mkdir /garbage

echo "####disale selins"

sed -i 's/enforcing/disabled/g' /etc/selinux/config


echo "#######system package install"

echo "install..."
yum -y install ntp make cmake gcc gcc-c++ gcc-g77 flex bison file libtool libtool-libs autoconf kernel-devel patch wget crontabs libjpeg libjpeg-devel libpng libpng-devel libpng10 libpng10-devel gd gd-devel libxml2 libxml2-devel zlib zlib-devel glib2 glib2-devel unzip tar bzip2 bzip2-devel libzip-devel libevent libevent-devel ncurses ncurses-devel curl curl-devel libcurl libcurl-devel e2fsprogs e2fsprogs-devel krb5 krb5-devel libidn libidn-devel openssl openssl-devel vim-minimal gettext gettext-devel ncurses-devel gmp-devel pspell-devel unzip libcap diffutils ca-certificates net-tools libc-client-devel psmisc libXpm-devel git-core c-ares-devel libicu-devel libxslt libxslt-devel xz expat-devel libaio-devel rpcgen libtirpc-devel perl lrzsz telnet vim

yum -y install iptables telnet iftop npm git

echo "install over"

echo "Update TCP"


cat >> /etc/sysctl.conf << EOF
vm.overcommit_memory = 1
net.ipv4.tcp_mem = 94500000 915000000 927000000
net.ipv4.tcp_fin_timeout = 30
net.ipv4.tcp_keepalive_time = 1200
net.ipv4.tcp_syncookies = 1
net.ipv4.tcp_tw_reuse = 1
net.ipv4.tcp_tw_recycle = 1
net.ipv4.ip_local_port_range = 1024 65535
net.ipv4.tcp_timestamps = 0
net.ipv4.tcp_synack_retries = 1
net.ipv4.tcp_syn_retries = 1
net.ipv4.tcp_abort_on_overflow = 0
net.core.rmem_max = 16777216
net.core.wmem_max = 16777216
#net.core.netdev_max_backlog = 262144
#net.core.somaxconn = 262144
net.ipv4.tcp_max_orphans = 3276800
net.ipv4.tcp_max_syn_backlog = 262144
net.core.wmem_default = 8388608
net.core.rmem_default = 8388608
net.ipv4.tcp_max_syn_backlog = 8192
net.ipv4.tcp_max_tw_buckets = 5000
EOF

sysctl -p
if [ $? -ne 0 ]; then
    echo "fail 手动执行"
else
    echo "success"
fi





















