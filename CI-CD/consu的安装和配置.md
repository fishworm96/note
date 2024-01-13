# 1.consul的安装和配置
## 1.安装
```go
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp consul consul agent -dev -client=0.0.0.0
// 显示进程
docker -ps -a
// 自动重启
docker container update --restart=always 容器名字
```
## 2.访问
```go
默认http端口8500
默认dns端口8600
服务器ip:consul端口
例// 192.168.0.50:8500
```
## 3.访问dns
consul提供的dns功能，可以让我们通过dig命令来测试，consul默认的dns端口是8600，命令行：
linux下的dig命令安装：
yum install bind-utils
```go
dig @192.168.0.50 -p 8600 consul.service.consul SRV
```
# 2.consul的api接口
## 1.添加服务
[https://www.consul.io/api-docs/agent/service#register-service](https://www.consul.io/api-docs/agent/service#register-service)
## 2.删除服务
[https://www.consul.io/api-docs/agent/service#deregister-service](https://www.consul.io/api-docs/agent/service#register-service)
## 3.设置健康检查
[https://www.consul.io/api-docs/agent/check](https://www.consul.io/api-docs/agent/check)
## 4.获取服务
[https://www.consul.io/api-docs/agent/check#list-checks](https://www.consul.io/api-docs/agent/check#list-checks)
