# CentOS7 安装 PostgreSQL 9.6

1、前往该地址[https://www.enterprisedb.com/downloads/postgres-postgresql-downloads下载PostggreSQL](https://www.enterprisedb.com/downloads/postgres-postgresql-downloads下载PostggreSQL) 9.6.2版。

![](/assets/postgresql9.6下载界面.jpg)



2、下载完文件名为：postgresql-9.6.2-1-linux-x64.run

3、将该文件上传到CentOS服务器/data目录上。

4、执行chmod +x postgresql-9.6.2-1-linux-x64.run

5、然后开始安装，执行 ./postgresql-9.6.2-1-linux-x64.run

6、安装过程中提示

6.1、Installation Directory \[/opt/PostgreSQL/9.6\]，安装目录，回车

6.2、Data Directory \[/opt/PostgreSQL/9.6/data\]，数据目录，回车

6.3、Please provide a password for the database superuser \(postgres\). A locked Unix

user account \(postgres\) will be created if not present.  
，输入密码

6.4、Please select the port number the server should listen on.  
Port \[5432\]，端口，回车

6.5、Select the locale to be used by the new database cluster.  
Locale，区域，语言，\[763\] zh\_CN.utf8，输入763，回车

6.6、Setup is now ready to begin installing PostgreSQL on your computer.Do you want to continue? \[Y/n\]，输入Y，回车

6.7、最后看到提示Setup has finished installing PostgreSQL on your computer.  
安装完成

7、在CentOS中打开端口firewall-cmd --add-port=5432/tcp

8、vi /var/lib/pgsql/9.6/data/pg\_hba.conf

在该配置文件的host all all 127.0.0.1/32 md5行下添加以下配置，或者直接将这一行修改为以下配置

host    all    all    0.0.0.0/0    md5

9、查看服务运行状态:systemctl status postgresql-9.6

