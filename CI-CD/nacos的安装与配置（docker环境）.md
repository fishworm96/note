# 安装
## 1.docker下载nacos
```yaml
// 需要虚拟机拥有4G的内存
docker run --name nacos-standalone -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -d nacos/nacos-server:latest
```
## 2.访问
```yaml
// 访问地址
http://192.168.0.50:8848/nacos/index.html
// 账号密码
都是:nacos
```
# 配置
## 1.命名空间
可以隔离配置集，将某些配置集放到某一个命名空间之下。命名空间我们一般用来区分微服务。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1657100863693-af2db1a2-3344-412d-83a0-cf1459c78eb1.png#clientId=u9090fd53-c2a0-4&from=paste&height=409&id=ub47f45bb&originHeight=511&originWidth=1880&originalType=binary&ratio=1&rotation=0&showTitle=false&size=31335&status=done&style=none&taskId=u033f3063-dbb0-4061-9190-846d99c5ac2&title=&width=1504)
![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1657100894010-ba7c1354-65d7-4384-afc4-d01ef4816d58.png#clientId=u9090fd53-c2a0-4&from=paste&height=470&id=uf3a159c6&originHeight=588&originWidth=1885&originalType=binary&ratio=1&rotation=0&showTitle=false&size=58312&status=done&style=none&taskId=u87149bc8-7b32-4b99-8f50-7ee2bd21913&title=&width=1508)
## 2.组
用来区别不同微服务是开发环境、测试环境、还是生产环境
![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1657100914065-444f8678-3d46-45a7-b9fa-6565bc5f64c1.png#clientId=u9090fd53-c2a0-4&from=paste&height=546&id=uc29a9dc8&originHeight=683&originWidth=1363&originalType=binary&ratio=1&rotation=0&showTitle=false&size=39514&status=done&style=none&taskId=u2b1d1c1e-757d-4263-96a9-e29447cf17d&title=&width=1090.4)
## 3.dataid-配置集
一个配置集就是一个配置文件，实际上可以更灵活。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1657100944819-af2c7fde-bfc4-4c4a-b8cd-d0c313fcc472.png#clientId=u9090fd53-c2a0-4&from=paste&height=606&id=uc232a272&originHeight=758&originWidth=1552&originalType=binary&ratio=1&rotation=0&showTitle=false&size=50614&status=done&style=none&taskId=u953d4a45-92ce-45ba-83fa-58bb0ed515b&title=&width=1241.6)
```json
{
  "name":"user-web",
  "port":8021,
  "user_srv":{
    "host":"192.168.0.102",
    "port":50051,
    "name":"user-srv"
  },
  "jwt":{
    "key":"12345"
  },
  "redis":{
    "host":"192.168.0.50",
    "port":6379
  },
  "consul":{
    "host":"192.168.0.50",
    "port":8500
  }
}
```
