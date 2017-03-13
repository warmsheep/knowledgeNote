# PostgreSQL9.6 流式复制

1、主机规划

  1.1、主机1：192.168.1.27，Master，端口5432

           主机2：192.168.1.28，Slave，端口5432



2、设置Host，Master和Slave都需要

vi /etc/hosts

192.168.1.27 master

192.168.1.28 slave



3、初始化Master数据库







