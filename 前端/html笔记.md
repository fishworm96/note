# 浏览器内核

| 浏览器  | 内核           |
| ------- | -------------- |
| IE      | Trident        |
| firefox | Gecko          |
| Safari  | webkit         |
| chrome  | Chromium/Blink |
| Opera   | blink          |

# html基本模型

#### 基本骨架

```html
<html>   
    <head>     
        <title></title>
    </head>
    <body>
    </body>
</html>
```

#### html骨架标签总结

| 标签名           |    定义    | 说明                                                    |
| ---------------- | :--------: | :------------------------------------------------------ |
| <html></html>    |  HTML标签  | 页面中最大的标签，我们成为  根标签                      |
| <head></head>    | 文档的头部 | 注意在head标签中我们必须要设置的标签是title             |
| <titile></title> | 文档的标题 | 让页面拥有一个属于自己的网页标题                        |
| <body></body>    | 文档的主体 | 元素包含文档的所有内容，页面内容 基本都是放到body里面的 |

## 文档类型<!DOCTYPE>

**用法：**

```
<!DOCTYPE html> 
```

作用：

```
<!DOCTYPE> 声明位于文档中的最前面的位置，处于 <html> 标签之前。此标签可告知浏览器文档使用哪种 HTML 或 XHTML 规范。
```

## 页面语言lang

```
<html lang="en">  指定html 语言种类
```

最常见的2个：

1. `en`定义语言为英语
2. `zh-CN`定义语言为中文

## 字符集

```
<meta charset="UTF-8" />
```

```
字符集(Character set)是多个字符的集合。

计算机要准确的处理各种字符集文字，需要进行字符编码，以便计算机能够识别和存储各种文字。
```

utf-8是目前最常用的字符集编码方式，常用的字符集编码方式还有gbk和gb2312。

* gb2312 简单中文  包括6763个汉字  GUO BIAO
* BIG5   繁体中文 港澳台等用
* GBK包含全部中文字符    是GB2312的扩展，加入对繁体字的支持，兼容GB2312
* UTF-8则基本包含全世界所有国家需要用到的字符
* **这句代码非常关键， 是必须要写的代码，否则可能引起乱码的情况。**

## HTML标签的语义化

作用：

1. 方便代码的阅读和维护
2. 同时让浏览器或是网络爬虫可以很好地解析，从而更好分析其中的内容 
3. 使用语义化标签会具有更好地搜索引擎优化 

## HTML常用标签

### 标题标签h 

 单词缩写：  head   头部. 标题       title  文档标题

**标题标签语义：**  作为标题使用，并且依据重要性递减

其基本语法格式如下：

```
<h1>   标题文本   </h1>
<h2>   标题文本   </h2>
<h3>   标题文本   </h3>
<h4>   标题文本   </h4>
<h5>   标题文本   </h5>
<h6>   标题文本   </h6>
```

### 段落标签p

单词缩写：  paragraph  段落

```
<p>  文本内容  </p>
```

### 水平线标签hr

单词缩写：  horizontal  横线

```
<hr />是单标签
```

### 换行标签br

单词缩写：  break

在HTML中，一个段落中的文字会从左到右依次排列，直到浏览器窗口的右端，然后自动换行。如果希望某段文本强制换行显示，就需要使用换行标签

```
<br />
```

### div 和  span标签

div   span    是没有语义的     是我们网页布局主要的2个盒子   想必你听过  css+div

div 就是  division  的缩写   分割， 分区的意思  其实有很多div 来组合网页。

span   跨度，跨距；范围    

语法格式：

```
<div> 这是头部 </div>    <span>今日价格</span>
```

* div标签  用来布局的，但是现在一行只能放一个div
* span标签  用来布局的，一行上可以放好多个span

### 文本格式化标签

| 标签                      | 显示效果                                 |
| ------------------------- | ---------------------------------------- |
| <b></b><strong></strong>> | 文字以粗体方式显示(XHTML推荐使用strong)  |
| <i></i>和<em></em>        | 文字以斜体方式显示(XHTML推荐使用em)      |
| <s></s>和<del></del>      | 文字以加删除线方式显示(XHTML推荐使用del) |
| <u></u>和<ins></ins>ins>  | 文字以加下划线方式显示(XHTML推荐使用u)   |

**区别：**

 b  只是加粗          strong  除了可以加粗还有 强调的意思，  语义更强烈。

剩下的同理...    

### 图像标签img

单词缩写：   image  图像

```
<img src="图像URL" />
```

### 链接标签

单词缩写：  anchor

语法格式：

```
<a href="跳转目标" target="目标窗口的弹出方式">文本或图像</a>
```

### 注释标签

语法格式：

```
<!-- 注释语句 -->    
```

### 锚点定位

创建锚点链接分为两步：

```
1. 使用相应的id名标注跳转目标的位置。 (找目标)
  <h3 id="two">第2集</h3> 

2. 使用<a href="#id名">链接文本</a>创建链接文本（被点击的）
  <a href="#two">   
```

### base 标签

**语法：**

```
<base target="_blank" />
```

1. base 可以设置整体链接的打开状态   
2. base 写到  <head>  </head>  之间
3. 把所有的连接 都默认添加 target="_blank"

### 预格式化文本pre标签

被包围在 <pre> 标签 元素中的文本通常会保留空格和换行符。而文本也会呈现为等宽字体。

```
<pre>

  此例演示如何使用 pre 标签

  对空行和 空格

  进行控制

</pre>
```

### 特殊字符

```
空格 &nbsp;
< 小于号 &lt;
> 大于号 &gt;
> & 和号 &amp;
> ￥ 人名币 &yen;
> ? 版权 &copy;
> ? 注册商标 &reg;
> ± 正负号 &plusmn;
> × 乘号 &times;
> ÷ 除号 &divide;
> 2 平方2（上标2) &sup2;
> 3 平方3（上标3) &sup3;
```

#### 什么是XHTML

XHTML 是更严格更纯净的 HTML 代码。

- XHTML 指**可扩展超文本标签语言**（EXtensible HyperText Markup Language）。
- XHTML 的目标是取代 HTML。
- XHTML 与 HTML 4.01 几乎是相同的。
- XHTML 是更严格更纯净的 HTML 版本。
- XHTML 是作为一种 XML 应用被重新定义的 HTML。
- XHTML 是一个 W3C 标准。

#### HTML和 XHTML之间有什么区别?

- XHTML 指的是可扩展超文本标记语言
- XHTML 与 HTML 4.01 几乎是相同的
- XHTML 是更严格更纯净的 HTML 版本
- XHTML 是以 XML 应用的方式定义的 HTML
- XHTML 是 2001 年 1 月发布的 W3C 推荐标准
- XHTML 得到所有主流浏览器的支持
- XHTML 元素是以 XML 格式编写的 HTML 元素。XHTML是严格版本的HTML，例如它要求标签必须小写，标签必须被正确关闭，标签顺序必须正确排列，对于属性都必须使用双引号等。

## 表格 table

创建表格的基本语法：

```
<table>
  <tr>
    <td>单元格内的文字</td>
    ...
  </tr>
  ...
</table>
```

1. table用于定义一个表格标签。

2. tr标签 用于定义表格中的行，必须嵌套在 table标签中。

3. td 用于定义表格中的单元格，必须嵌套在<tr></tr>标签中。

4. 字母 td 指表格数据（table data），即数据单元格的内容，现在我们明白，表格最合适的地方就是用来存储数据的。

### 表格属性

表格有部分属性我们不常用，这里重点记住 cellspacing 、 cellpadding。

### 表头单元格标签th

- 作用：
  - 一般表头单元格位于表格的第一行或第一列，并且文本加粗居中
- 语法：
  - 只需用表头标签&lt;th&gt;</th&gt;替代相应的单元格标签&lt;td&gt;</td&gt;即可。 

![img](file://J:/%E9%BB%91%E9%A9%AC/01-03%20%E5%89%8D%E7%AB%AF%E5%BC%80%E5%8F%91%E5%9F%BA%E7%A1%80/01-HTML%E8%B5%84%E6%96%99/02.HTML-Day02/%E7%AC%94%E8%AE%B0/media/th.png?lastModify=1615298796)

案例2

![img](file://J:/%E9%BB%91%E9%A9%AC/01-03%20%E5%89%8D%E7%AB%AF%E5%BC%80%E5%8F%91%E5%9F%BA%E7%A1%80/01-HTML%E8%B5%84%E6%96%99/02.HTML-Day02/%E7%AC%94%E8%AE%B0/media/tht.png?lastModify=1615298810)

```
<table width="500" border="1" align="center" cellspacing="0" cellpadding="0">
		<tr>  
			<th>姓名</th> 
			<th>性别</th>
			<th>电话</th>
		</tr>
		<tr>
			<td>小王</td>
			<td>女</td>
			<td>110</td>
		</tr>
		<tr>
			<td>小明</td>
			<td>男</td>
			<td>120</td>
		</tr>	
	</table>
```

### 表格标题caption

定义和用法

```
<table>
   <caption>我是表格标题</caption>
</table>
```

1. caption 元素定义**表格标题**，通常这个标题会被居中且显示于表格之上。
2. caption 标签必须紧随 table 标签之后。
3. 这个标签只存在 表格里面才有意义。

### 合并单元格

* 跨行合并：rowspan="合并单元格的个数"      
* 跨列合并：colspan="合并单元格的个数"

```
<rowspan="合并几个单元格，从上第一个到下">跨行合并单元格，合并后删除多余单元格
<colspan="合并几个单元格，从左第一个到右">跨列合并单元格，合并后删除多余单元格
```

合并的顺序们按照   先上 后下     先左  后右 的顺序

## 列表标签

### 无序列表 ul

无序列表的各个列表项之间没有顺序级别之分，是并列的。其基本语法格式如下：

```
<ul>
  <li>列表项1</li>
  <li>列表项2</li>
  <li>列表项3</li>
  ......
</ul>
```

### 有序列表 ol

有序列表即为有排列顺序的列表，其各个列表项按照一定的顺序排列定义，有序列表的基本语法格式如下：

```
<ol>
  <li>列表项1</li>
  <li>列表项2</li>
  <li>列表项3</li>
  ......
</ol>
```

### 自定义列表

定义列表常用于对术语或名词进行解释和描述，定义列表的列表项前没有任何项目符号。其基本语法如下：

```
<dl>
  <dt>名词1</dt>
  <dd>名词1解释1</dd>
  <dd>名词1解释2</dd>
  ...
  <dt>名词2</dt>
  <dd>名词2解释1</dd>
  <dd>名词2解释2</dd>
  ...
</dl>
```

## 表单标签

### input 控件

语法：

```
<input type="属性值" value="你好">
```

- input 输入的意思 
- <input /&gt;标签为单标签
- type属性设置不同的属性值用来指定不同的控件类型
- 除了type属性还有别的属性

![img](file://J:/%E9%BB%91%E9%A9%AC/01-03%20%E5%89%8D%E7%AB%AF%E5%BC%80%E5%8F%91%E5%9F%BA%E7%A1%80/01-HTML%E8%B5%84%E6%96%99/02.HTML-Day02/%E7%AC%94%E8%AE%B0/media/input.png?lastModify=1615299369)

#### 1.type 属性

* 这个属性通过改变值，可以决定了你属于那种input表单。
* 比如 type = 'text'  就表示 文本框 可以做 用户名， 昵称等。
* 比如 type = 'password'  就是表示密码框   用户输入的内容 是不可见的。

```
用户名: <input type="text" /> 
密  码：<input type="password" />
```

#### 2.value属性   值

```
用户名:<input type="text"  name="username" value="请输入用户名"> 
```

value 默认的文本值。 有些表单想刚打开页面就默认显示几个文字，就可以通过这个value 来设置。

#### 3.name属性

```
用户名:<input type="text"  name=“username” /> 
```

name表单的名字， 这样，后台可以通过这个name属性找到这个表单。  页面中的表单很多，name主要作用就是用于区别不同的表单。

* name属性后面的值，是我们自己定义的。


* radio  如果是一组，我们必须给他们命名相同的名字 name   这样就可以多个选其中的一个啦

```
<input type="radio" name="sex"  />男
<input type="radio" name="sex" />女
```

#### 4.checked属性

表示默认选中状态。  较常见于 单选按钮和复选按钮。

```
性    别:
<input type="radio" name="sex" value="男" checked="checked" />男
<input type="radio" name="sex" value="女" />女 
```

### label标签

**作用：** 

 用于绑定一个表单元素, 当点击label标签的时候, 被绑定的表单元素就会获得输入焦点。

**如何绑定元素呢？**

1.第一种用法就是用label直接包括input表单。

```
<label> 用户名： <input type="radio" name="usename" value="请输入用户名">   </label>
```

适合单个表单选择

2.第二种用法 for 属性规定 label 与哪个表单元素绑定。

```
<label for="sex">男</label>
<input type="radio" name="sex"  id="sex">
```

### textarea控件

语法：

```
<textarea >
  文本内容
</textarea>
```

作用：

通过textarea控件可以轻松地创建多行文本输入框.

cols="每行中的字符数" rows="显示的行数"  我们实际开发不用

#### 文本框和文本域区别

| 表单              |  名称  |       区别       |                  默认值显示 |             用于场景 |
| :---------------- | :----: | :--------------: | --------------------------: | -------------------: |
| input type="text" | 文本框 | 只能显示一行文本 | 单标签，通过value显示默认值 | 用户名、昵称、密码等 |
| textarea          | 文本域 | 可以显示多行文本 |  双标签，默认值写到标签中间 |               留言板 |

### select下拉列表

![img](file://J:/%E9%BB%91%E9%A9%AC/01-03%20%E5%89%8D%E7%AB%AF%E5%BC%80%E5%8F%91%E5%9F%BA%E7%A1%80/01-HTML%E8%B5%84%E6%96%99/02.HTML-Day02/%E7%AC%94%E8%AE%B0/media/sele.png?lastModify=1615299590)

语法：

```
<select>
  <option>选项1</option>
  <option>选项2</option>
  <option>选项3</option>
  ...
</select>
```

* 注意：

1. &lt;select&gt;  中至少包含一对 option 
2. 在option 中定义selected =" selected "时，当前项即为默认选中项。
3. 但是我们实际开发会用的比较少

### form表单域

**语法: **

```
<form action="url地址" method="提交方式" name="表单名称">
  各种表单控件
</form>
```

**常用属性：**

| 属性   | 属性值   | 作用                                               |
| ------ | :------- | -------------------------------------------------- |
| action | url地址  | 用于指定接收并处理表单数据的服务器程序的url地址。  |
| method | get/post | 用于设置表单数据的提交方式，其取值为get或post。    |
| name   | 名称     | 用于指定表单的名称，以区分同一个页面中的多个表单。 |

## HTML5

### 基本标签

```
<header>:头部标签</header>
<nav>:导航标签</nav>
<article>:内容标签</article>
<section>:块级标签</section>
<aside>:侧边栏标签</aside>
<footer>:尾部标签</footer>
```

在IE9中需要把这些元素转化为块级元素

### 插入音频

```
<audio>音频标签</audio>
autoplay:如果出现该属性，则音频在就绪后马上播放
controls:如果出现该属性，则向用户显示控件，比如播放按钮
loop:如果出现该属性，则每当音频结束时重新播放
src:url 要播放的音频的URL
例：
<audio controls="controls">
	<source src="路径.mp3" type="audio/mpeg">
	<source src="路径.ogg" type="audio/ogg">
```

支持格式
ogg vorbis
mp3
wav

### 插入视频

```
例：
<video>视频标签</video>
<video controls="controls">
	<source src="move.ogg" type="video/ogg"
	<source src="move.mp4" type="video/mp4"
</video>
autoplay：视频就绪自动播放
controls：向用户显示播放控件
width:px 设置播放器宽度
height:px 设置播放器高度
loop:播放完是否继续播放该视屏，循环播放
preload：auto（预先加载视屏）/none（不应加载视屏） 规定是否预加载视频（如果有autoplay就忽略该属性）
src:url 视屏url地址
poster:"图片路径" 加载等待的画面图片
muted:muted 静音播放
```

### 表单

#### 新增input标签

type="email" 限制用户输入必须为email类型

type="url" 限制用户输入必须为rul类型

type="date" 显示用户输入必须为日期类型

type="time" 显示用户输入必须为时间类型

type="month" 限制用户输入必须为月类型

type="week" 限制用户输入必须为周类型

type="number" 限制用户输入必须为数字类型

type="tel" 手机号码

type="search" 搜索框

type="color" 生成一个颜色选择表单

#### 新增表单属性

required:required 表单拥有该属性表示其内容不恩那个为空，必填

placeholder：提示文本 表单的提示信息，存在默认值将不显示

autofocus:autofocus 自动聚焦属性，页面加载完成自动聚焦到指定表单

autucomplete:off/on 当用户在字段开始键入是，浏览器基于之前键入过的值，应该显示出在字段中填写的选项。默认已经打开，如autocomplete="on" 关闭 autoconplete="off" 需要放在表单内同时加上name属性 同时成功提交

multiple:multiple 可以多选文件提交