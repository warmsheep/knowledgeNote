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

1. 开始安装，执行 
   ```
   ./postgresql-9.6.2-1-linux-x64.run
   ```
2. Installation Directory \[/opt/PostgreSQL/9.6\]
   安装目录，回车
3. Data Directory \[/opt/PostgreSQL/9.6/data\]，数据目录，回车
4. Please provide a password for the database superuser \(postgres\). A locked Unix user account \(postgres\) will be created if not present.，输入密码
5. Please select the port number the server should listen on. Port \[5432\]，端口，回车
6. Select the locale to be used by the new database cluster. Locale，区域，回车
7. Setup is now ready to begin installing PostgreSQL on your computer.Do you want to continue? \[Y/n\]，输入Y，回车
8. Setup has finished installing PostgreSQL on your computer.最后看到提示，安装完成

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



