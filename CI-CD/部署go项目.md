## 生成go二进制文件
下面假设我们将本地编译好的 bluebell 二进制文件、配置文件和静态文件等上传到服务器的/data/app/bluebell目录下。
补充一点，如果嫌弃编译后的二进制文件太大，可以在编译的时候加上-ldflags "-s -w"参数去掉符号表和调试信息，一般能减小20%的大小。
```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/bluebell
```
如果还是嫌大的话可以继续使用 upx 工具对二进制可执行文件进行压缩。
我们编译好 bluebell 项目后，相关必要文件的目录结构如下：
```shell
├── bin
│   └── bluebell
├── conf
│   └── config.yaml
├── static
│   ├── css
│   │   └── app.0afe9dae.css
│   ├── favicon.ico
│   ├── img
│   │   ├── avatar.7b0a9835.png
│   │   ├── iconfont.cdbe38a0.svg
│   │   ├── logo.da56125f.png
│   │   └── search.8e85063d.png
│   └── js
│       ├── app.9f3efa6d.js
│       ├── app.9f3efa6d.js.map
│       ├── chunk-vendors.57f9e9d6.js
│       └── chunk-vendors.57f9e9d6.js.map
└── templates
    └── index.html
```
### supervisor
Supervisor 是业界流行的一个通用的进程管理程序，它能将一个普通的命令行进程变为后台守护进程，并监控该进程的运行状态，当该进程异常退出时能将其自动重启。
首先使用 yum 来安装 supervisor：
如果你还没有安装过 EPEL，可以通过运行下面的命令来完成安装，如果已安装则跳过此步骤：
```shell
sudo yum install epel-release
```
安装 supervisor
```shell
sudo yum install supervisor 
```
Supervisor 的配置文件为：/etc/supervisord.conf ，Supervisor 所管理的应用的配置文件放在 /etc/supervisord.d/ 目录中，这个目录可以在 supervisord.conf 中的include配置。
```shell
[include] 
files = /etc/supervisord.d/*.conf 
```
启动supervisor服务：
```shell
sudo supervisord -c /etc/supervisord.conf 
```
我们在/etc/supervisord.d目录下创建一个名为bluebell.conf的配置文件，具体内容如下。
```shell
[program:bluebell]  ;程序名称 
user=root  ;执行程序的用户 
command=/data/app/bluebell/bin/bluebell /data/app/bluebell/conf/config.yaml;执行的命令 
directory=/data/app/bluebell/ ;命令执行的目录 
stopsignal=TERM  ;重启时发送的信号 
autostart=true   
autorestart=true  ;是否自动重启 
stdout_logfile=/var/log/bluebell-stdout.log  ;标准输出日志位置 
stderr_logfile=/var/log/bluebell-stderr.log  ;标准错误日志位置 
```
创建好配置文件之后，重启supervisor服务
```shell
sudo supervisorctl update # 更新配置文件并重启相关的程序 
```
查看bluebell的运行状态：
```shell
sudo supervisorctl status bluebell 
```
输出：
```shell
bluebell                         RUNNING   pid 10918, uptime 0:05:46 
```
最后补充一下常用的supervisr管理命令：
```shell
supervisorctl status       # 查看所有任务状态 
supervisorctl shutdown     # 关闭所有任务 
supervisorctl start 程序名  # 启动任务 
supervisorctl stop 程序名   # 关闭任务
supervisorctl reload       # 重启supervisor 
```
接下来就是打开浏览器查看网站是否正常了。
## 使用yum安装nginx
EPEL 仓库中有 Nginx 的安装包。如果你还没有安装过 EPEL，可以通过运行下面的命令来完成安装：
```shell
sudo yum install epel-release
```
## 安装nginx
```shell
sudo yum install nginx
```
安装完成后，执行下面的命令设置Nginx开机启动：
```shell
sudo systemctl enable nginx
```
启动Nginx
```shell
sudo systemctl start nginx
```
查看Nginx运行状态：
```shell
sudo systemctl status nginx
```
## Nginx配置文件
通过上面的方法安装的 nginx，所有相关的配置文件都在 /etc/nginx/ 目录中。Nginx 的主配置文件是 /etc/nginx/nginx.conf。
默认还有一个nginx.conf.default的配置文件示例，可以作为参考。你可以为多个服务创建不同的配置文件（建议为每个服务（域名）创建一个单独的配置文件），每一个独立的 Nginx 服务配置文件都必须以 .conf结尾，并存储在 /etc/nginx/conf.d 目录中。
## Nginx常用命令
补充几个 Nginx 常用命令。
```shell
nginx -s stop    # 停止 Nginx 服务
nginx -s reload  # 重新加载配置文件
nginx -s quit    # 平滑停止 Nginx 服务
nginx -t         # 测试配置文件是否正确
```
## Nginx反向代理部署
我们推荐使用 nginx 作为反向代理来部署我们的程序，按下面的内容修改 nginx 的配置文件。
```
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    server {
        listen       80;
        server_name  localhost;

        access_log   /var/log/bluebell-access.log;
        error_log    /var/log/bluebell-error.log;

        location / {
            proxy_pass                 http://127.0.0.1:8084;
            proxy_redirect             off;
            proxy_set_header           Host             $host;
            proxy_set_header           X-Real-IP        $remote_addr;
            proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
    }
}
```
执行下面的命令检查配置文件语法：
```
nginx -t
```
执行下面的命令重新加载配置文件：
```
nginx -s reload
```
接下来就是打开浏览器查看网站是否正常了。
当然我们还可以使用 nginx 的 upstream 配置来添加多个服务器地址实现负载均衡。
```
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    upstream backend {
      server 127.0.0.1:8084;
      # 这里需要填真实可用的地址，默认轮询
      #server backend1.example.com;
      #server backend2.example.com;
    }

    server {
        listen       80;
        server_name  localhost;

        access_log   /var/log/bluebell-access.log;
        error_log    /var/log/bluebell-error.log;

        location / {
            proxy_pass                 http://backend/;
            proxy_redirect             off;
            proxy_set_header           Host             $host;
            proxy_set_header           X-Real-IP        $remote_addr;
            proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
    }
}
```
下面继续修改我们的 nginx 的配置文件来实现上述功能。
```
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    server {
        listen       80;
        server_name  bluebell;

        access_log   /var/log/bluebell-access.log;
        error_log    /var/log/bluebell-error.log;

		# 静态文件请求
        location ~.*\.(gif|jpg|jpeg|png|js|css|eot|ttf|woff|svg|otf)$ {
            access_log off;
            expires    1d;
            root       /data/app/bluebell;
        }

        # index.html页面请求
        # 因为是单页面应用这里使用 try_files 处理一下，避免刷新页面时出现404的问题
        location / {
            root /data/app/bluebell/templates;
            index index.html;
            try_files $uri $uri/ /index.html;
        }

      	# 配置多个项目
      	location /admin {
          	alias /data/app/bluebell/templates/admin;
          	index index.html;
          	try_files $uri $uri/ /index.html;
        }
      	
		# API请求
        location /api {
            proxy_pass                 http://127.0.0.1:8084;
            proxy_redirect             off;
            proxy_set_header           Host             $host;
            proxy_set_header           X-Real-IP        $remote_addr;
            proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
    }
}
```
