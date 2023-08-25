# 数据库

## 数据库相关概念

|    术语    |                         解释说明                         |
| :--------: | :------------------------------------------------------: |
|  database  |      数据库，mongodb数据库软件中可以建立多个数据库       |
| collection |    集合，一组数据的集合，可以理解为javascript中的数组    |
|  document  |    文档，一条具体的数据，可以理解为javascript中的对象    |
|   field    | 字段，文档中的属性名称，可以理解为javascript中的对象属性 |

```
net stop mongodb关闭服务器
net start mongodb 开启服务器
```

### 数据库连接

使用mongodb提供的connect方法即可连接数据库

```
mongoose.connect('mongodb://localhost/数据库名字')
	.then(() => console.log('数据库连接成功'))
	.catch(err => console.log('数据库连接失败', err));
```

#### 创建集合

一是对集合设定规则，二是创建集合，创建mongoose.Schema构造的实力即可创建集合

```
//设定集合规定
const courseSchema = new mongoose.Schama({
	name: String,
	author: String,
	inPublished: Boolean
});
//创建集合并应用规定
const Course = mongose.model('Course', courseSchema);//courses
```

#### 创建文档

创建文档实际上就是向集合中插入数据。

分两步：

1.创建集合实例

2.调用实例对象下的save方法将数据保存到数据库中。

```
方法1：
//创建集合实例
const course = new Course({
	name: 'Node.js course'.
	author: '黑马僵尸',
	tags: ['node', 'backend'],
	isPublished: true
});
//将数据保存到数据库中
course.save();
方法2：
Course.create({name: 'JavaScript基础', author: '黑马僵尸', inPublish: true}, (err, doc) => {
	//错误对象
	console.log(err)
	//当前插入的文档
	console.log(doc)
});
promise方法：
Course.create({name: 'JavaScript', author: '黑马僵尸', isPublished: true})
	.then(doc => console.log(doc))
	.catch(err => console.log(err))
```

#### mongoDB数据库导入数据

mongoimport -d 数据库名称 -c 集合名称 --file 要导入的数据文件

```
mongoimport -d playground -c users --file ./user.json
```

### MongoDB增删改查

#### 查询文档

方法1：

```
//根据条件查找文档(条件为空则查找所有文档)
Course.find().then(result => console.log(result))

//返回文档集合
[{
	_id: 5c09f267aeb04b22f8460968,
	name: 'node.js'.
	author: '黑马'
}]
```

方法2：

```
//根据条件查找文档
Course.findOne({name : 'node.js'}).then(result => console.log(result))

//返回文档
{
	_id: 5c09f267aeb04b22f8460968,
	name: 'node.js'.
	author: '黑马' 
}
```

方法3：

```
//匹配大于:$gt 小于:$lt	gt和lt固定写法
User.find({age: {$gt: 20, $lt: 50}}).then(result => console.log(result))
```

```
//匹配包含
User.find({hobbies: {$in: ['敲代码']}}).then(result => console.log(result))
```

```
//选择要查询的字段
User.find().select('name email -_id').then(result => console.log(result))
不想查询就在不想查询的前面加'-',如'-_id'
```

```
//将数据按照年龄进行排序
User.find().sort('age').then(result => console.log(result))//升序

User.find().sort('-age').then(result => console.log(result))//降序
```

```
//skip跳过多少条数据 limit 限制查询数量		分页时使用
User.find().skip(2).limit(2).then(result => console.log(result))
```

#### 删除文档

方法1：

```
//删除单个
Course.findOneAndDelete({}).then(result => console.log(result))
```

```
//删除多个	如果是空对象将User对象下所有文档删除
User.deleteMany({}).then(result => console.log(result))
```

#### 更新文档

```
//更新单个
User.updateOne({查询条件},{要修改的值}).then(result => console.log(result))
例：
User.updataOne({ name: '李四' }, { name: '李狗蛋' }).then(result => console.log(result))
```

```
//更新多个
User.updateMany({查询条件}, {要更改的值}).then(result => console.log(result))
```

#### Mongoose验证

在创建集合时，可以设置当前字段的验证规则，验证失败就是如插入失败

required: true	必传字段

minlenth: 3	字符串最小长度

maxlength: 20	字符串最大长度

min: 2	数值最小为2

max: 20	数值最大为20

enum: ['html', 'css']

trim: true 去除字符串两边的空格

validate:自定义验证器

default:默认值

```
例：
const postChema = new mongoose.Schema({
    title: {
        type: String,
        //必选字段[自定义错误信息]
        required: [true, '请传入文章标题'],
        //传入字符串最小长度
        minlength: 2,
        //传入字符串最大长度
        maxlength: 5,
        //去除字符串两边空格
        trim: true
    },
    age: {
        type: Number,
        //数字最小范围
        min: 18,
        //数字最大范围
        max: 100
    },
    publishDate: {
        type: Date,
        //默认值
        default: Date.now
    },
    category: {
        type: String,
        //列举出当前字段可以拥有的值
        enum: ['html', 'css', 'node.js']
    },
    //自定义验证器
    author: {
        type: String,
        //选项
        validate: {
            //函数类型
            validator: v => {
                //返回布尔值，true验证成功，false验证失败，v 要验证的值
                return v && v.lenth > 4
            },
            //自定义错误信息
            message: '传入的值不符合验证规则'
        }
    }
});

post.create({ title: 'a', age: 17, category: 'java', author: 'db' })
    .then(result => console.log(result))
    .catch(error => {
        // 获取错误信息对象
        const err = error.errors;
        // 循环错误信息对象
        for (var attr in err) {
            // 将错误信息打印到控制台中
            console.log(err[attr]['mesage']);
        }
    })
```

#### 集合关联

通常不同集合的数据之间是有关系的，例如文章信息和用户信息储存在不同集合中，但文章是某个用户发表的，要查询文章的所有信息包括发表用户，就需要用到集合关联

使用id对集合进行关联

使用populate方法进行关联集合查询

#### 集合关联实现

```
//用户集合
const User = mongoose.model('User', new mongoose.Schama({ name: { type: String } }));
//文章集合
const Post = mongoose.model('Post', new mongoose.Schema({
    title: { type: String },
    //使用ID将文章集合和作者集合进行关联
    author: { type: mongoose.Schema.Types.ObjectId, ref: 'User' }
}));
//联合查询
Post.find()
    .populate('author')
    .then((err, result) => console.log(result));
```

