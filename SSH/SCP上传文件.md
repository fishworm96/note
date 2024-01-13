# 上传文件
```shell
scp 本地文件地址 root@serverIp:服务器路径
如
scp /Users/a/a.txt root@10.10.1.1:/www/a/a.txt
如果ssh端口不为22，使用-P指定端口（注意，是大写P）
scp -P 1111 /Users/a/a.txt root@10.10.1.1:/www/a/a.txt
```
# 上传文件夹
```shell
#拷贝文件夹,多加上一个-r参数即可
scp -r file username@ip:filepath
```
# 拷贝大量小文件到远程服务器
```shell
# 使用tar将/home打包,通过管道传输到远程机器解压到/根目录
# 这里是tar的参数
# -cvf 表示只打包不使用压缩
## -c tar的打包参数
## -v 显示详情,进度
## -f 指定文件名称,必须放到所有选项后面
# -xvf 表示解压指定文件
## -x 表示解压
## -C 大C表示解压到指定位置
tar -cvf-/ home | ssh root@192.168.1.1 tar -xvf--C /
```
# 示例
## scp把远端机器上的文件下载到本地
```shell
#把192.168.0.10机器上的source.txt文件拷贝到本地的/home/work目录下
scp work@192.168.0.10:/home/work/source.txt /home/work/ 
```
## scp把远端机器A的文件拷贝到远端机器B
```shell
#把192.168.0.10机器上的source.txt文件拷贝到192.168.0.11机器的/home/work目录下
scp work@192.168.0.10:/home/work/source.txt work@192.168.0.11:/home/work/ 
```
# 递归拷贝-r(会覆盖)
注意，如果本地存在同名文件，会覆盖且无警告提示
```shell
#拷贝文件夹，加-r参数
scp -r /home/work/sourcedir work@192.168.0.10:/home/work/
```
# SCP断点续传
如果你要强调传输的安全性 可以采用rsync + ssh
```shell
## -a, --archive 归档模式，表示以递归方式传输文件，并保持所有文件属性
## -v, --verbose 详细模式输出
## -z, --compress 对备份的文件在传输时进行压缩处理
##-P 断点续传并打印过程
## -e, --rsh=COMMAND 指定使用rsh、ssh方式进行数据同步
# rsync -avzP  -e 'ssh -p port' root@[$Remote_Host]:[$Remote_Dir] [$Local_Dir]
# rsync -avzP  -e 'ssh -p port' root@[$Remote_Host]:[$Remote_Dir] [$Local_Dir]
rsync -avzP --rsh=ssh local_file_pic.tar.gz root@192.168.205.304:/home/Remote_file.tar.gz
```
## 文件断点下载
```shell
# 文件断点下载
##-P 断点续传并打印过程
## -e, --rsh=COMMAND 指定使用rsh、ssh方式进行数据同步
rsync -P --rsh=ssh root@192.168.0.11:/root/large.tar.gz /dounine/targe.tar.gz
```
## 文件断点上传
```shell
# 文件断点上传
##-P 断点续传并打印过程
## -e, --rsh=COMMAND 指定使用rsh、ssh方式进行数据同步
rsync -P --rsh=ssh /dounine/targe.tar.gz root@192.168.0.11:/root/large.tar.gz
```
## 文件目录断点下载
```shell
# 文件目录断点下载
##-P 断点续传并打印过程
## -e, --rsh=COMMAND 指定使用rsh、ssh方式进行数据同步
rsync -P --rsh=ssh -r root@192.168.0.11:/root/storage /dounine
```
## 文件目录断点上传
```shell
# 文件目录断点上传
##-P 断点续传并打印过程
## -e, --rsh=COMMAND 指定使用rsh、ssh方式进行数据同步
rsync -P --rsh=ssh -r /dounine root@192.168.0.11:/root/storage
```
