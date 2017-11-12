# Docker导入导出镜像

* 镜像导出

```shell
docker save -o /data/docker_image_rabbitmq3-management.tar docker.io/rabbitmq:3-management
docker save -o /data/docker_image_postgres10.tar docker.io/postgres:10
docker save -o /data/docker_image_redis3.2.8.tar docker.io/redis:3.2.8
docker save -o /data/docker_image_nexuslastest.tar index.alauda.cn/warmsheep/nexus:latest
```

* 镜像导入

```shell
docker load -i /data/docker_image_rabbitmq3-management.tar
docker load -i /data/docker_image_postgres10.tar
docker load -i /data/docker_image_redis3.2.8.tar
docker load -i /data/docker_image_nexuslastest.tar
```

* 容器导出

```shell

```

* 容器导入

```shell

```
