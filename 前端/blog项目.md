# blog项目

## 密码加密 bcrypt

哈希加密是单程加密方式： 1234 => abcd

在加密的密码中加入随机字符串可以增加密码被破解的难度

bcrypt依赖的其他环境

1.python 2.x

2.node-gyp

​	npm install -g node-gyp

3.windows-build-tools

​	npm install --global --production windows-build-tools

```
//导入bcrypt模块
const bcrypt = require('bcrypt');
//生成随机字符串 gen => generate 生成 salt 盐
let salt = await bcrypt.genSalt(10);
//使用随机字符串对密码进行加密
let pass = await bcrypt.hash('明文密码', salt);
```

```
//密码比对
let isEqual = await bcrypt.compare('明文密码', '加密密码');
```

## cookie与session

#### cookie:浏览器在电脑硬盘中开辟的一块空间，主要供服务器端存储数据

cookie中的数据是以域名的形式进行区分的。

cookie中的数据是有过期时间的，超过时间数据会被浏览器自动删除。

cookie中的数据会随着请求被自动发送到服务器端。

#### session：实际上就是一个对象，存储在服务器端的内存中，在session对象中也可以存储多条数据，每一条数据都有一个sessionid作为唯一标识。

在node.js中需要借助express-session实现session功能

```
cosnt sessoin = requrie('express-session');
app.use(session({secret: 'secret key'}));
```

### 新增用户

1.为用户列表页的新增用户添加连接

2.添加一个连接对应的路由，在路由处理函数中渲染新增用户模板

3.为新增用户表单指定请求地址、请求方式、为表单项添加name属性

4.增加实现添加用户的功能路由

5.接收到客户端传递过来的请求参数

6.对请求参数的格式进行验证

7.验证当前要注册的邮箱地址是否已经注册过

8.对密码进行加密处理

9.对用户信息添加到数据库中

10.重定向页面到用户列表页面

#### Joi

javascript对象的规则描述语言和验证器

```
const Joi = require('joi');
const Schema = {
	username:Joi.string().alphanum().min(3).max(30).required().error(new Error('错误信息')),
	password: JOi.string().regex(/^[a-zA-Z0-9]{3,30}$/),
	access_token: [Joi.string(), Joi.number()],
	birth: Joi.number().integer().min(1900).max(2020),
	email: Joi.string().email()
};
Joi.validate({username: 'abc', birth: 1994}, schema);
```

### 数据分页

当数据库中的数据非常多时，数据需要分批显示，这时就需要用到数据分页功能。

分页功能核心要素：

1.当前页，用户通过点击上一页或者下一页或者页码产生，客户端通过get参数方式传递到服务器端

2.总页数，根据总页数判断当前页是否最后一个页，根据判断结果做相应操作。

总页数：Math.ceil (总数据条数/每页显示数据条数)

### 数据分页

```
limit(2) //limit限制查询数量 传入每页显示的数据数量
skip(1) //skip跳过多少条数据 传入显示数据的开始位置
```

用户信息修改

1.将要修改的用户ID传递到服务器端

2.建立用户信息修改功能对应的路由

3.接收客户端表单传递过来的请求参数

4.根据id查询用户信息，并将客户端传递过来的密码和数据库中的密码进行比对

5.如果比对失败，对客户端做出响应

6.如果密码比对成功，将用户信息更新到数据库中

### 用户删除信息

\1. 在确认删除框中添加隐藏域用以存储要删除用户的ID值

\2. 为删除按钮添自定义属性用以存储要删除用户的ID值

\3. 为删除按钮添加点击事件，在点击事件处理函数中获取自定义属性中存储的ID值并将ID值存储在表单的隐藏域中

\4. 为删除表单添加提交地址以及提交方式

\5. 在服务器端建立删除功能路由

\6. 接收客户端传递过来的id参数

\7. 根据id删除用户