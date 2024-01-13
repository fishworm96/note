## 需求
我现在需要在nginx部署多个vite打包的react项目
## vite打包
### 单一项目
![image.png](https://cdn.nlark.com/yuque/0/2023/png/12511308/1681461194841-32dedd48-54d3-44aa-9df0-c264679ea005.png#averageHue=%2323211f&clientId=u295d8233-1c2a-4&from=paste&height=48&id=u820050ec&originHeight=48&originWidth=176&originalType=binary&ratio=1&rotation=0&showTitle=false&size=1840&status=done&style=none&taskId=ubeca724b-6c98-4959-8487-297ff24c1b2&title=&width=176)![image.png](https://cdn.nlark.com/yuque/0/2023/png/12511308/1681461212670-4f31da96-4d01-483a-b0b2-d6516ec9e61f.png#averageHue=%2323211f&clientId=u295d8233-1c2a-4&from=paste&height=102&id=ue9dba761&originHeight=102&originWidth=194&originalType=binary&ratio=1&rotation=0&showTitle=false&size=4380&status=done&style=none&taskId=u565c378d-af1d-4d00-ad56-2410a142504&title=&width=194)
base使用'/'输出目录dist
### 多个项目
这里我的两个项目一个叫blog一个叫admin
需要修改vite的配置
修改vite的base
![image.png](https://cdn.nlark.com/yuque/0/2023/png/12511308/1681461288892-bc9b1c15-ab0b-4f87-bfc8-2169ecebfd4d.png#averageHue=%23282420&clientId=u295d8233-1c2a-4&from=paste&height=76&id=ub217862b&originHeight=76&originWidth=180&originalType=binary&ratio=1&rotation=0&showTitle=false&size=4611&status=done&style=none&taskId=ue4f00c4a-4d59-48ce-8157-7d7f4b7836c&title=&width=180)
![image.png](https://cdn.nlark.com/yuque/0/2023/png/12511308/1681461298998-8af15970-dd2e-4d7b-9fa2-fa8cb78a93a6.png#averageHue=%2321201f&clientId=u295d8233-1c2a-4&from=paste&height=86&id=u55d602e0&originHeight=86&originWidth=212&originalType=binary&ratio=1&rotation=0&showTitle=false&size=3509&status=done&style=none&taskId=u51dae63e-f6b7-4000-989d-5f864933223&title=&width=212)
## 修改Nginx
普通Nginx
```nginx
location /admin {
  root /data/app/blog/templates/admin/dist;
  index index.html;
  try_files $uri $uri/ /index.html;
}
```
部署多个项目
项目路径在/data/app/blog/templates下，分别是blog和admin
详细信息
**/data/app/blog/templates/blog/dist**
**/data/app/blog/templates/blog/dist**
dist是vite打包的目录
```nginx
    	location /blog {
                alias /data/app/blog/templates/blog/dist;
                index index.html;
                try_files $uri $uri/ /blog/index.html;
        }

      location /admin {
                alias /data/app/blog/templates/admin/dist;
                index index.html;
                try_files $uri $uri/ /admin/index.html;
        }

```
## nginx具体内容
```nginx
# For more information on configuration, see:
#   * Official English Documentation: http://nginx.org/en/docs/
#   * Official Russian Documentation: http://nginx.org/ru/docs/

user nginx;
worker_processes auto;
error_log /var/log/nginx/error.log;
pid /run/nginx.pid;

# Load dynamic modules. See /usr/share/doc/nginx/README.dynamic.
include /usr/share/nginx/modules/*.conf;

events {
    worker_connections 1024;
}

http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile            on;
    tcp_nopush          on;
    tcp_nodelay         on;
    keepalive_timeout   65;
    types_hash_max_size 4096;

    include             /etc/nginx/mime.types;
    default_type        application/octet-stream;

    # Load modular configuration files from the /etc/nginx/conf.d directory.
    # See http://nginx.org/en/docs/ngx_core_module.html#include
    # for more information.
    include /etc/nginx/conf.d/*.conf;

    gzip on;
    gzip_min_length 1k;
    gzip_comp_level 2;
    gzip_types text/plain applicationi/javascript application/x-javascript text/css application/xml text/javascript application/x-httped-php image/jpeg image/gif image/png;
   
    gzip_vary on;

    server {
        listen       80;
        listen       [::]:80;
        server_name  www.fishworm.top;
	
        # Load configuration files for the default server block.
        include /etc/nginx/default.d/*.conf;

	#location ~ \.(gif|jpg|jpeg|png|js|css|eot|ttf|woff|svg|otf)$ {
		#access_log off;
		#expires 1d;
		#root /data/app/blog/templates/;
	#}

#blog项目
	location /blog {
		alias /data/app/blog/templates/blog/dist;
		index index.html;
		try_files $uri $uri/ /blog/index.html;
	}

#admin项目
	location /admin {
		alias /data/app/blog/templates/admin/dist;
		index index.html;
		try_files $uri $uri/ /admin/index.html;
	}

#根目录
	location / {
		root /data/app/blog/templates/blog/dist;
		try_files $uri $uri/ /index.html;
	}

	location /api {
            proxy_pass                 http://127.0.0.1:8080;
            proxy_redirect             off;
            proxy_set_header           Host             $host;
            proxy_set_header           X-Real-IP        $remote_addr;
            proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
        }

        error_page 404 /404.html;
        location = /404.html {
        }

        error_page 500 502 503 504 /50x.html;
        location = /50x.html {
        }

    }
}
```
