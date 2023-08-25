mysql

net start mysql	//mysql是定义的数据库名  打开数据库

net stop mysql	//关闭数据库

mysql -h localhost -P 3306 -uroot -proot	//-h	地址	-P端口号	-u 用户名	-p密码 连接数据库

1.查看当前所有的数据库

show databases;

2.打开指定的库

use test	//test为库名

3.查看当前库的所有表

show tables;

4.查看其它库的所有表

show tables from mysql;	//切换数据库

5.创建表

create table 表名(

​		列名 列类型,

​		列名 列类型,

)

6.查看表结构

desc 表名;

7.查看服务器的版本

方式一

seleft version();	//查看数据库版本

方式二：没有登录mysql服务端

mysql --version或mysq --V

