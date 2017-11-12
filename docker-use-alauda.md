# 使用灵雀云加速

```shell
sudo sed -i 's|other_args="|other_args="--registry-mirror=http://warmsheep.m.alauda.cn |g' /etc/sysconfig/docker
sudo sed -i "s|OPTIONS='|OPTIONS='--registry-mirror=http://stmbtfly.m.alauda.cn |g" /etc/sysconfig/docker
sudo sed -i 'N;s|\[Service\]\n|\[Service\]\nEnvironmentFile=-/etc/sysconfig/docker\n|g' /usr/lib/systemd/system/docker.service
sudo sed -i 's|fd://|fd:// $other_args |g' /usr/lib/systemd/system/docker.service
sudo systemctl daemon-reload
sudo service docker restart
```
