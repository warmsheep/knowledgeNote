# Docker Swarm部署

### 主机规划

| 主机名 | IP | 角色 |
| :--- | :--- | :--- |
| Docker-Management | 192.168.20.117 | Management、Service |
| Docker-Service | 192.168.20.122 | Service |

### 安装Docker

Docker-Management和Docker-Service都执行下列操作

* 设置yum仓库

```
sudo yum install yum-utils -y
```

```
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
```

```
sudo yum makecache fast
```

* 安装Docker

```
sudo yum -y install docker-ce
```

```
sudo systemctl start docker
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

### 创建Docker Swarm集群

* 初始化Docker Swarm，在Docker-Management机器如下命令：

```
docker swarm init
```

执行完后弹出如下提示

```
Swarm initialized: current node (wumxoarm2fiolv9bavkj20dxe) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join \
    --token SWMTKN-1-4coj9xxdbkrdlehnlmonp1ej9zroj3aa888exof2wf83kmk0gs-3ecah8qxqk77d4w4rzyku5zv9 \
    192.168.20.117:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
```

* 打开防火墙端口

```
firewall-cmd --zone=public --add-port=2377/tcp --permanent
firewall-cmd --reload
```

* Docker-Service加入集群，在Docker-Service执行上述提示的命令

```
docker swarm join \
--token SWMTKN-1-4coj9xxdbkrdlehnlmonp1ej9zroj3aa888exof2wf83kmk0gs-3ecah8qxqk77d4w4rzyku5zv9 \
192.168.20.117:2377
```

* Docker-Service会给出如下提示，提示已成功加入节点

```
This node joined a swarm as a worker.
```

* 在Docker-Management执行docker info命令可以看到整个集群的节点信息

```
docker info
```

```
Containers: 0
 Running: 0
 Paused: 0
 Stopped: 0
Images: 0
Server Version: 17.03.0-ce
Storage Driver: overlay
 Backing Filesystem: xfs
 Supports d_type: false
Logging Driver: json-file
Cgroup Driver: cgroupfs
Plugins: 
 Volume: local
 Network: bridge host macvlan null overlay
Swarm: active
 NodeID: wumxoarm2fiolv9bavkj20dxe
 Is Manager: true
 ClusterID: v8s6tvmubgb35oa6471z2xm9n
 Managers: 1
 Nodes: 2
 Orchestration:
  Task History Retention Limit: 5
 Raft:
  Snapshot Interval: 10000
  Number of Old Snapshots to Retain: 0
  Heartbeat Tick: 1
  Election Tick: 3
 Dispatcher:
  Heartbeat Period: 5 seconds
 CA Configuration:
  Expiry Duration: 3 months
 Node Address: 192.168.20.117
 Manager Addresses:
  192.168.20.117:2377
Runtimes: runc
Default Runtime: runc
Init Binary: docker-init
containerd version: 977c511eda0925a723debdc94d09459af49d082a
runc version: a01dafd48bc1c7cc12bdb01206f9fea7dd6feb70
init version: 949e6fa
Security Options:
 seccomp
  Profile: default
Kernel Version: 3.10.0-327.el7.x86_64
Operating System: CentOS Linux 7 (Core)
OSType: linux
Architecture: x86_64
CPUs: 1
Total Memory: 1.797 GiB
Name: CentOS1
ID: QTAN:H3MO:HTLL:KCEC:F4CF:GWM7:SHWJ:ATOJ:NNVH:Q5QE:VQUY:Q2NH
Docker Root Dir: /var/lib/docker
Debug Mode (client): false
Debug Mode (server): false
Registry: https://index.docker.io/v1/
WARNING: bridge-nf-call-ip6tables is disabled
Experimental: false
Insecure Registries:
 127.0.0.0/8
Live Restore Enabled: false
```

* 如果要离开节点，输入docker swarm leave则可以离开节点

```
docker swarm leave
```



