# mysql
## 下载镜像
```shell
docker pull mysql
```
## 创建mysql8配置文件
```shell
vi /etc/my.cnf #编辑MySQL配置文件
```
my.cnf文件内容
```shell
# Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.
#
# This program is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation; version 2 of the License.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program; if not, write to the Free Software
# Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301 USA
#
# The MySQL  Server configuration file.
#
# For explanations see
# http://dev.mysql.com/doc/mysql/en/server-system-variables.html
[mysqld]
pid-file        = /var/run/mysqld/mysqld.pid
socket          = /var/run/mysqld/mysqld.sock
datadir         = /var/lib/mysql
secure-file-priv= NULL
# Disabling symbolic-links is recommended to prevent assorted security risks
symbolic-links=0
# Custom config should go here
!includedir /etc/mysql/conf.d/
```
## 启动镜像
```shell
docker run -p 60306:3306 -e MYSQL_ROOT_PASSWORD=123 -v /etc/my.cnf:/etc/mysql/my.cnf:rw -v /etc/localtime:/etc/localtime:ro --name mysql8 --restart=always -dit mysql
-p 60306:3306 #本机60306端口映射到容器3306端口
-e MYSQL_ROOT_PASSWORD=123 #设置MySQL的root用户密码
-v /etc/my.cnf:/etc/mysql/my.cnf:rw #本机的MySQL配置文件映射到容器的MySQL配置文件
-v /etc/localtime:/etc/localtime:ro #本机时间与数据库时间同步
--name mysql8 #设置容器别名
--restart=always #当重启Docker时会自动启动该容器
-dit mysql #后台运行并可控制台接入
```
## 进入mysql容器
```shell
docker exec -it b6cfb244d0c0 bash #进入MySQL容器
mysql -uroot -p123 #进入MySQL控制台
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'root'; #修改root用户密码
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';
flush privileges;
```
# redis
## 下载镜像
```shell
docker pull redis:latest
```
## 查看容器
```shell
docker images
```
## 运行redis
```shell
docker run -itd --name redis -p 26379:6379 redis
```
## 进入reids
```shell
docker exec -it redis-test /bin/bash
```
## 测试
```shell
redis-cli
set test 1 # 返回ok
```
# Nginx
## 安装镜像
```shell
docker pull nginx:latest
```
## 运行容器
```shell
docker run --name nginx -p 80:80 -d nginx

--name nginx-test：容器名称。
-p 8080:80： 端口进行映射，将本地 8080 端口映射到容器内部的 80 端口。
-d nginx： 设置容器在在后台一直运行。
```
