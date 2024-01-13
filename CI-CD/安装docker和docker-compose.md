## 安装docker
```shell
curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
```
### 设置开机启动docker
```shell
 # 启动docker
systemctl start docker
 # 设置开机自动启动
systemctl enable docker
```
### 进入阿里云
![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1660872911921-73a1d98a-9de3-42cb-8703-e42391a1c95b.png#clientId=u53fae2f6-9790-4&from=paste&height=945&id=uf714162b&originHeight=945&originWidth=1917&originalType=binary&ratio=1&rotation=0&showTitle=false&size=131843&status=done&style=none&taskId=u66c2569e-83cf-406f-87e8-f0ab2781fc0&title=&width=1917)
![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1660872962324-7939a286-96c2-4e17-9536-e7859572f4fe.png#clientId=u53fae2f6-9790-4&from=paste&height=828&id=uc2d5ed4a&originHeight=828&originWidth=893&originalType=binary&ratio=1&rotation=0&showTitle=false&size=69078&status=done&style=none&taskId=u480d9643-f068-4827-a23a-1cc73f8dcc8&title=&width=893)
```shell
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://ka3hytik.mirror.aliyuncs.com"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker
```
配置完阿里云镜像后需要重启docker
```shell
sudo systemctl start docker
```
### 测试是否成功
```shell
sudo docker run hello-world
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1660873294599-61e8f24f-42da-4f28-87f4-1c7c735ef86c.png#clientId=u53fae2f6-9790-4&from=paste&height=403&id=ud2cecead&originHeight=403&originWidth=584&originalType=binary&ratio=1&rotation=0&showTitle=false&size=28673&status=done&style=none&taskId=u4fc5da0b-312e-40f3-af16-fe5cf5f0fbc&title=&width=584)
### docker删除容器
```shell
# 查询容器id
docker ps -a
# 删除容器
docker rm [id]
```
## 安装docker-compose
```shell
curl -L https://get.daocloud.io/docker/compose/releases/download/1.25.0/docker-compose-`uname -s`-`uname -m`> /usr/local/bin/docker-compose

sudo chmod +x /usr/local/bin/docker-compose
```
### 测试
```shell
docker-compose -v
```
