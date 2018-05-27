## VPN搭建

yum -y install epel-release
yum -y install openvpn

vi /etc/sysctl.conf
net.ipv4.ip_forward = 1


sysctl -p

wget https://github.com/OpenVPN/easy-rsa/archive/release/2.x.zip
unzip 2.x.zip
cd easy-rsa-release-2.x/easy-rsa/2.0/


vi vars
# which will be placed in the certificate.
# Don't leave any of these fields blank.
export KEY_COUNTRY="CN"
export KEY_PROVINCE="BEIJING"
export KEY_CITY="BEIJING"
export KEY_ORG="CHANPAY"
export KEY_EMAIL="linxuana@yonyou.com"
export KEY_OU="OPS"
# X509 Subject Field
export KEY_NAME="EasyRSA"


source ./vars
./clean-all
./build-ca

./build-key-server server
./build-key client
./build-dh
cd keys/

rpm -ql openvpn | grep server.conf
/usr/share/doc/openvpn-2.4.5/sample/sample-config-files/roadwarrior-server.conf
/usr/share/doc/openvpn-2.4.5/sample/sample-config-files/server.conf
/usr/share/doc/openvpn-2.4.5/sample/sample-config-files/xinetd-server-config


cp /usr/share/doc/openvpn-2.4.5/sample/sample-config-files/server.conf /etc/openvpn/
mkdir /etc/openvpn/keys
cp * /etc/openvpn/keys
cd /etc/openvpn/
vi server.conf


local x.x.x.x # 侦听IP

port 4911

proto tcp # 默认为udp,修改为tcp

dev tun # 使用tun接口,为三层设备,可直接IP路由

ca keys/ca.crt # 证书相关,你懂的

cert keys/server.crt

key keys/server.key  # This file should be kept secret

dh keys/dh2048.pem

topology subnet # 网络拓扑,子网

server 10.8.0.0 255.255.255.0 # vpn 客户端获取的IP段

ifconfig-pool-persist ipp.txt

push "route 172.31.0.0 255.255.0.0" # 推送给客户端的路由,客户端会自动添加一条路由

route 172.31.0.0 255.255.0.0 # 服务端添加一条路由,要不然服务端去172.31不知道咋走,这里使用路由汇总,一条就可以了,反正走 172.31.0.0/16 段的都走 tun0 接口

client-config-dir ccd # vpn客户端配置目录,每个客户端一个配置文件,对应创建客户端证书时取的名字,回头看看

client-to-client # 允许客户端之间的通信

keepalive 10 120

comp-lzo

max-clients 100

persist-key

persist-tun

status openvpn-status.log

log         openvpn.log

log-append  openvpn.log

verb 3
