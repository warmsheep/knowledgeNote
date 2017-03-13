# CentOS7 安装 PostgreSQL 9.6

### 下载、配置文件

* 前往该地址[https://www.enterprisedb.com/downloads/postgres-postgresql-downloads下载PostggreSQL](https://www.enterprisedb.com/downloads/postgres-postgresql-downloads下载PostggreSQL) 9.6.2版。

![](/assets/postgresql9.6下载界面.jpg)

* 下载完文件名为：postgresql-9.6.2-1-linux-x64.run

* 将该文件上传到CentOS服务器/data目录上。

* 新建postgres用户

  ```
  adduser postgres
  ```

* 将文件移动到/home/postgres文件夹下

  ```
  mv postgresql-9.6.2-1-linux-x64.run /home/postgres
  ```

* 增加执行权限

  ```
  chmod +x postgresql-9.6.2-1-linux-x64.run
  ```

### 

### 安装过程

* 开始安装，执行 

```
./postgresql-9.6.2-1-linux-x64.run
```

```
Welcome to the PostgreSQL Setup Wizard.

1、指定PostgreSQL安装目录
----------------------------------------------------------------------------
Please specify the directory where PostgreSQL will be installed.

Installation Directory [/opt/PostgreSQL/9.6]: /home/postgres/PostgreSQL/9.6


2、指定PostgreSQL数据存放目录
----------------------------------------------------------------------------
Please select a directory under which to store your data.

Data Directory [/home/postgres/PostgreSQL/9.6/data]: /home/postgres/PostgreSQL/9.6/data


3、设置PostgreSQL数据库密码
----------------------------------------------------------------------------
Please provide a password for the database superuser (postgres). A locked Unix 
user account (postgres) will be created if not present.

Password :
Retype password :


4、设置PostgreSQL数据库端口
----------------------------------------------------------------------------
Please select the port number the server should listen on.

Port [5432]: 5432


5、设置PostgreSQL区域
----------------------------------------------------------------------------
Advanced Options

Select the locale to be used by the new database cluster.

Locale

[1] [Default locale]
……
……

Please choose an option [1] : 1


6、确认安装
----------------------------------------------------------------------------
Setup is now ready to begin installing PostgreSQL on your computer.

Do you want to continue? [Y/n]: Y


7、安装
----------------------------------------------------------------------------
Please wait while Setup installs PostgreSQL on your computer.

 Installing
 0% ______________ 50% ______________ 100%
 ########################################


8、安装完成
----------------------------------------------------------------------------
Setup has finished installing PostgreSQL on your computer.
```

### 操作防火墙：

* 在防火墙中打开服务

  ```
  firewall-cmd --add-service=postgresql --permanent
  ```

* 在CentOS中打开端口

  ```
  firewall-cmd --zone=public --add-port=5432/tcp --permanent
  ```

* 重载防火墙

  ```
  firewall-cmd --reload
  ```

* 查看占用端口

  ```
  firewall-cmd --list-ports
  ```

* 修改网络配置

  ```
  vi /home/postgres/PostgreSQL/9.6/data/pg_hba.conf
  ```

* 在该配置文件的

* 在该配置文件的host all all 127.0.0.1/32 md5行下添加以下配置，或者直接将这一行修改为以下配置

  ```
  host    all    all    0.0.0.0/0    md5
  ```

* 重启postgresql服务

  ```
  systemctl restart postgresql-9.6
  ```

* 查看服务运行状态

  ```
  systemctl status postgresql-9.6
  ```

### 设置环境变量

* 编辑/etc/profile文件

  ```
  vi /etc/profile
  ```

* 在文件末位添加如下内容

  ```
  PGHOME=/home/postgres/PostgreSQL/9.6
  export PGHOME
  PGDATA=$PGHOME/data
  export PGDATA
  PATH=$PATH:$HOME/.local/bin:$HOME/bin:$PGHOME/bin
  export PATH
  ```



