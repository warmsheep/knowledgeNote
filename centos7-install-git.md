# CentOS7源码安装Git

---

* 安装必要的编译工具

```
yum install -y curl-devel expat-devel gettext-devel openssl-devel zlib-devel gcc perl-ExtUtils-MakeMaker
```

* 下载Git包源码

```
wget https://www.kernel.org/pub/software/scm/git/git-1.9.4.tar.gz
```

* 解压Git包

```
tar -zxvf git-1.9.4.tar.gz
```

* 编译

```
cd git-1.9.4/ && make prefix=/usr/local/git all
```

* 安装

```
cd git-1.9.4/ && make prefix=/usr/local/git install
```

* 设置环境变量

```
echo "export PATH=\$PATH:/usr/local/git/bin" >> /etc/profile
```

* 环境变量生效

```
source /etc/profile
```



