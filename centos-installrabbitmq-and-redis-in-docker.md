# CentOS7 中使用Docker部署RabbitMQ/Redis

### 安装Docker

* 使用yum安装docker

```
yum install docker -y
```

* 启动docker

```
systemctl start docker.service
```

* 查看docker服务运行状态

```
systemctl status docker.service
```

### 安装Docker-Compose

* 下载并安装docker-compose

```
curl -L "https://github.com/docker/compose/releases/download/1.11.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

* 设置执行权限

```
chmod +x /usr/local/bin/docker-compose
```

* 查看Docker-Compose版本

```
docker-compose --version
```

### 安装Redis（单机）

* 下载Redis镜像

```
docker pull redis:3.2.8
```

* 运行Redis镜像（单机）

```
docker run -d --name Redis-Container -p 6379:6379 redis:3.2.8
```

* 开启防火墙端口并重新启动防火墙

```
firewall-cmd --zone=public --add-port=6379/tcp --permanent
firewall-cmd --reload
```

### 安装Redis容器集群版

注意：该方式不能从外部访问Sentinel，获取到的可能是容器内网的IP地址，外部无法访问到，需要客户端也部署在容器内，通过容器内的DNS解析进行访问。通常用到服务编排来实现外部机器的访问。

* 切换到/目录，并创建/docker\_data/redis\_container\_conf/目录

```
cd /
mkdir /docker_data
cd /docker_data
mkdir redis_container_conf
```

* 创建docker-compose.yml文件

```
vi docker-compose.yml
```

```
master:
  image: redis:3
slave:
  image: redis:3
  command: redis-server --slaveof redis-master 6379
  links:
    - master:redis-master
sentinel:
  build: sentinel
  environment:
    - SENTINEL_DOWN_AFTER=5000
    - SENTINEL_FAILOVER=5000
  links:
    - master:redis-master
    - slave
```

* 创建sentinel目录

```
mkdir sentinel
cd sentinel
```

* 创建Dockerfile文件

```
vi Dockerfile
```

```
FROM redis:3

MAINTAINER Lin Xuan <linxuana@chanjet.com>

EXPOSE 26379
ADD sentinel.conf /etc/redis/sentinel.conf
RUN chown redis:redis /etc/redis/sentinel.conf
ENV SENTINEL_QUORUM 2
ENV SENTINEL_DOWN_AFTER 30000
ENV SENTINEL_FAILOVER 180000
COPY sentinel-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/sentinel-entrypoint.sh
ENTRYPOINT ["sentinel-entrypoint.sh"]
```

* 创建sentinel-entrypoint.sh文件

```
vi sentinel-entrypoint.sh
```

```
#!/bin/sh

sed -i "s/\$SENTINEL_QUORUM/$SENTINEL_QUORUM/g" /etc/redis/sentinel.conf
sed -i "s/\$SENTINEL_DOWN_AFTER/$SENTINEL_DOWN_AFTER/g" /etc/redis/sentinel.conf
sed -i "s/\$SENTINEL_FAILOVER/$SENTINEL_FAILOVER/g" /etc/redis/sentinel.conf

exec docker-entrypoint.sh redis-server /etc/redis/sentinel.conf --sentinel
```

* 创建sentinel.conf文件

```
vi sentinel.conf
```

```
# Example sentinel.conf can be downloaded from http://download.redis.io/redis-stable/sentinel.conf

port 26379

dir /tmp

sentinel monitor mymaster redis-master 6379 $SENTINEL_QUORUM

sentinel down-after-milliseconds mymaster $SENTINEL_DOWN_AFTER

sentinel parallel-syncs mymaster 1

sentinel failover-timeout mymaster $SENTINEL_FAILOVER
```

* 使用docker-compose执行构建

```
cd /docker_data/redis_container_conf
docker-compose build
```

* 部署并启动容器

```
docker-compose up -d
```

* 检查容器运行状态

```
docker-compose ps
```

* 开启防火墙

```
firewall-cmd --zone=public --add-port=6379/tcp --permanent
firewall-cmd --zone=public --add-port=26379/tcp --permanent
firewall-cmd --reload
```

### 安装RabbitMQ

* 下载RabbitMQ镜像

```
docker pull frodenas/rabbitmq
```

* 运行RabbitMQ容器

```
docker run -d \
  --name RabbitMQ-Container \
  -p 5672:5672 \
  -p 15672:15672 \
  -e RABBITMQ_USERNAME=rabbitmq \
  -e RABBITMQ_PASSWORD=rabbitmq \
  -e RABBITMQ_VHOST=myvhost \
  frodenas/rabbitmq
```

* 开启防火墙端口并重新启动防火墙

```
firewall-cmd --zone=public --add-port=5672/tcp --permanent
firewall-cmd --zone=public --add-port=15672/tcp --permanent
firewall-cmd --reload
```



