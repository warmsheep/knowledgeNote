# CentOS7安装JDK8

---

* 在CentOS7中安装JDK8执行如下命令即可：

```
 curl -LO --insecure --junk-session-cookies --location --remote-name --silent \
    --header "Cookie: oraclelicense=accept-securebackup-cookie" \
    http://download.oracle.com/otn-pub/java/jdk/8u74-b02/jdk-8u74-linux-x64.rpm && \ 
  yum localinstall -y jdk-8u74-linux-x64.rpm && \
  yum clean all && \
  rm jdk-8u74-linux-x64.rpm
```

* 或者自行下载JDK8的RPM包，然后执行rpm命令安装

```
rpm -ivh jdk-8u74-linux-x64.rpm
```



