# css样式

```
<div style="color:颜色;" font-size:"像素"; /div>
```



# 字体

## 1.font-size:大小

作用：

font-size属性用于设置字号

```
p {  
    font-size:20px; 
}
```

单位：

- 可以使用相对长度单位，也可以使用绝对长度单位。
- 相对长度单位比较常用，推荐使用像素单位px，绝对长度单位使用较少。

| 相对长度单位 | 说明                           |
| ------------ | ------------------------------ |
| em           | 相对于当前对象内文本的字体尺寸 |
| px           | 像素，最常用，推荐使用         |
| 绝对长度单位 | 说明                           |
| in           | 英寸                           |
| cm           | 厘米                           |
| mm           | 毫米                           |
| pt           | 点                             |

**注意：**

* 我们文字大小以后，基本就用px了，其他单位很少使用
* 谷歌浏览器默认的文字大小为16px
* 但是不同浏览器可能默认显示的字号大小不一致，我们尽量给一个明确值大小，不要默认大小。一般给body指定整个页面文字的大小

## 2.font-family:字体

作用：

font-family属性用于设置哪一种字体。

```
p{ font-family:"微软雅黑";}
```

- 网页中常用的字体有宋体、微软雅黑、黑体等，例如将网页中所有段落文本的字体设置为微软雅黑
- 可以同时指定多个字体，中间以逗号隔开，表示如果浏览器不支持第一个字体，则会尝试下一个，直到找到合适的字体， 如果都没有，则以我们电脑默认的字体为准。

```
p{font-family: Arial,"Microsoft Yahei", "微软雅黑";}
```

1. 各种字体之间必须使用英文状态下的逗号隔开。
2. 中文字体需要加英文状态下的引号，英文字体一般不需要加引号。当需要设置英文字体时，英文字体名必须位于中文字体名之前。
3. 如果字体名中包含空格、#、$等符号，则该字体必须加英文状态下的单引号或双引号，例如font-family: "Times New Roman";。
4. 尽量使用系统默认字体，保证在任何用户的浏览器中都能正确显示。

### CSS Unicode字体

- 为什么使用 Unicode字体

  - 在 CSS 中设置字体名称，直接写中文是可以的。但是在文件编码（GB2312、UTF-8 等）不匹配时会产生乱码的错误。
  - xp 系统不支持 类似微软雅黑的中文。

- 解决：

  - 方案一： 你可以使用英文来替代。 比如` font-family:"Microsoft Yahei"`。

  - 方案二： 在 CSS 直接使用 Unicode 编码来写字体名称可以避免这些错误。使用 Unicode 写中文字体名称，浏览器是可以正确的解析的。

```
font-family: "\5FAE\8F6F\96C5\9ED1";   表示设置字体为“微软雅黑”。
```

| 字体名称    | 英文名称        | Unicode 编码         |
| ----------- | --------------- | -------------------- |
| 宋体        | SimSun          | \5B8B\4F53           |
| 新宋体      | NSimSun         | \65B0\5B8B\4F53      |
| 黑体        | SimHei          | \9ED1\4F53           |
| 微软雅黑    | Microsoft YaHei | \5FAE\8F6F\96C5\9ED1 |
| 楷体_GB2312 | KaiTi_GB2312    | \6977\4F53_GB2312    |
| 隶书        | LiSu            | \96B6\4E66           |
| 幼园        | YouYuan         | \5E7C\5706           |
| 华文细黑    | STXihei         | \534E\6587\7EC6\9ED1 |
| 细明体      | MingLiU         | \7EC6\660E\4F53      |
| 新细明体    | PMingLiU        | \65B0\7EC6\660E\4F53 |

为了照顾不同电脑的字体安装问题，我们尽量只使用宋体和微软雅黑中文字体

## 3.font-weight:字体粗细

- 在html中如何将字体加粗我们可以用标签来实现
  - 使用 b  和 strong 标签是文本加粗。
- 可以使用CSS 来实现，但是CSS 是没有语义的。

| 属性值  | 描述                                                      |
| ------- | :-------------------------------------------------------- |
| normal  | 默认值（不加粗的）                                        |
| bold    | 定义粗体（加粗的）                                        |
| 100~900 | 400 等同于 normal，而 700 等同于 bold  我们重点记住这句话 |

提倡：

  我们平时更喜欢用数字来表示加粗和不加粗。

## 4.font-style:字体风格

- 在html中如何将字体倾斜我们可以用标签来实现
  - 字体倾斜除了用 i  和 em 标签，
- 可以使用CSS 来实现，但是CSS 是没有语义的
- font-style属性用于定义字体风格，如设置斜体、倾斜或正常字体，其可用属性值如下：

| 属性   | 作用                                                    |
| ------ | :------------------------------------------------------ |
| normal | 默认值，浏览器会显示标准的字体样式  font-style: normal; |
| italic | 浏览器会显示斜体的字体样式。                            |

```
平时我们很少给文字加斜体，反而喜欢给斜体标签（em，i）改为普通模式。
```

## 5 font:综合设置字体样式 (重点)

font属性用于对字体样式进行综合设置

- 基本语法格式如下：

```
font:font-style字体样式 
font-weight:字体粗细 
font-size/line-height:字号(100-900) 
font-famliy:字体
```

注意：

- 使用font属性时，必须按上面语法格式中的顺序书写，不能更换顺序，各个属性以**空格**隔开。

- 其中不需要设置的属性可以省略（取默认值），但必须保留font-size和font-family属性，否则font属性将不起作用。

- ```
  .title{font: italic(倾斜) 700 20px "微软雅黑"}
  ```

## font总结

| 属性        | 表示     | 注意点                                                       |
| :---------- | :------- | :----------------------------------------------------------- |
| font-size   | 字号     | 我们通常用的单位是px 像素，一定要跟上单位                    |
| font-family | 字体     | 实际工作中按照团队约定来写字体                               |
| font-weight | 字体粗细 | 记住加粗是 700 或者 bold  不加粗 是 normal 或者  400  记住数字不要跟单位 |
| font-style  | 字体样式 | 记住倾斜是 italic     不倾斜 是 normal  工作中我们最常用 normal |
| font        | 字体连写 | 1. 字体连写是有顺序的  不能随意换位置 2. 其中字号 和 字体 必须同时出现 |

# CSS外观属性

## 1 color:文本颜色

- 作用：

  color属性用于定义文本的颜色，

- 其取值方式有如下3种：

| 表示表示       | 属性值                                  |
| :------------- | :-------------------------------------- |
| 预定义的颜色值 | red，green，blue，还有我们的御用色 pink |
| 十六进制       | #FF0000，#FF6600，#29D794               |
| RGB代码        | rgb(255,0,0)或rgb(100%,0%,0%)           |

```
color：16进制#00（红色）00（绿色）00（蓝色）color:rgb(000.000.000)  rgb
```

注意

我们实际工作中， 用 16进制的写法是最多的，而且我们更喜欢简写方式比如  #f00 代表红色

## 2 text-align:文本水平对齐方式

- 作用：

  text-align属性用于设置文本内容的水平对齐，相当于html中的align对齐属性

- 其可用属性值如下：

| 属性   |       解释       |
| ------ | :--------------: |
| left   | 左对齐（默认值） |
| right  |      右对齐      |
| center |     居中对齐     |

- ```
  text-align:"left,center,right";对齐方式！
  ```

- 注意：

  是让盒子里面的内容水平居中， 而不是让盒子居中对齐

## 3 line-height:行间距

- 作用：

  line-height属性用于设置行间距，就是行与行之间的距离，即字符的垂直间距，一般称为行高。

- 单位：

  - line-height常用的属性值单位有三种，分别为像素px，相对值em和百分比%，实际工作中使用最多的是像素px

```
一般情况下，行距比字号大7.8像素左右就可以了。line-height: 24px;line-height:"像素";行间距 
```

## 4 text-indent:首行缩进

- 作用：

  text-indent属性用于设置首行文本的缩进，

- 属性值

  - 其属性值可为不同单位的数值、em字符宽度的倍数、或相对于浏览器窗口宽度的百分比%，允许使用负值,
  - 建议使用em作为设置单位。

**1em 就是一个字的宽度   如果是汉字的段落， 1em 就是一个汉字的宽度**

```
p {      /*行间距*/      line-height: 25px;      /*首行缩进2个字  em  */      text-indent: 2em;   } line-height:"像素";行间距  text-indent:"像素";首行缩进 1em就是1个字的距离2em就是2个字的距离
```

## 5 text-decoration 文本的装饰

text-decoration   通常我们用于给链接修改装饰效果

| 值           | 描述                                                  |
| ------------ | ----------------------------------------------------- |
| none         | 默认。定义标准的文本。 取消下划线（最常用）           |
| underline    | 定义文本下的一条线。下划线 也是我们链接自带的（常用） |
| overline     | 定义文本上的一条线。（不用）                          |
| line-through | 定义穿过文本下的一条线。（不常用）                    |

```
text-decoration:文本的装饰    	超链接自带下划线{none:取消下划线underline:添加下划线}blink:闪烁
```

## 6.CSS外观属性总结

| 属性            | 表示     | 注意点                                                  |
| :-------------- | :------- | :------------------------------------------------------ |
| color           | 颜色     | 我们通常用  十六进制   比如 而且是简写形式 #fff         |
| line-height     | 行高     | 控制行与行之间的距离                                    |
| text-align      | 水平对齐 | 可以设定文字水平的对齐方式                              |
| text-indent     | 首行缩进 | 通常我们用于段落首行缩进2个字的距离   text-indent: 2em; |
| text-decoration | 文本修饰 | 记住 添加 下划线  underline  取消下划线  none           |

# 复合选择器

## 1 后代选择器

- 概念：

  后代选择器又称为包含选择器

- 作用：

  用来选择元素或元素组的**子孙后代**

- 其写法就是把外层标签写在前面，内层标签写在后面，中间用**空格**分隔，先写父亲爷爷，在写儿子孙子。 

```
父级 子级{属性:属性值;属性:属性值;}
```

语法：

```
.class h3{color:red;font-size:16px;}div strong{子孙元素选测器，后代选择器}
```

- 当标签发生嵌套时，内层标签就成为外层标签的后代。
- 子孙后代都可以这么选择。 或者说，它能选择任何包含在内 的标签。

## 2 子元素选择器

- 作用：

  子元素选择器只能选择作为某元素**子元素(亲儿子)**的元素。

- 其写法就是把父级标签写在前面，子级标签写在后面，中间跟一个 `>` 进行连接

- 语法：

```
.class>h3{color:red;font-size:14px;}div>strong{子元素选择器,只选亲儿子}
```

## 3 交集选择器



条件

交集选择器由两个选择器构成，找到的标签必须满足：既有标签一的特点，也有标签二的特点。

其中第一个为标签选择器，第二个为class选择器，两个选择器之间**不能有空格**，如h3.special

```
p.red{交集选择器，即是p标签，又是.red类选择器。中间不能包含空格}<p class="red"></p>
```



## 4 并集选择器

- 应用：
  - 如果某些选择器定义的相同样式，就可以利用并集选择器，可以让代码更简洁。
- 并集选择器（CSS选择器分组）是各个选择器通过`,`连接而成的，通常用于集体声明。
- 语法：

任何形式的选择器（包括标签选择器、class类选择器id选择器等），都可以作为并集选择器的一部分。

```
p，span，red{color:red;}  “，”和的意思，通常用于集体声明
```



## 5  链接伪类选择器

作用：

用于向某些选择器添加特殊的效果。比如给链接添加特殊效果， 比如可以选择 第1个，第n个元素。

因为伪类选择器很多，比如链接伪类，结构伪类等等。我们这里先给大家讲解链接伪类选择器。

```
a:link{}  未访问过的链接a:visited{} 已访问过的链接a:hover{} 鼠标移动到链接上a:active{} 选定的链接按照l v h a的顺序编写
```

```
a {   /* a是标签选择器  所有的链接 */			font-weight: 700;			font-size: 16px;			color: gray;}a:hover {   /* :hover 是链接伪类选择器 鼠标经过 */			color: red; /*  鼠标经过的时候，由原来的 灰色 变成了红色 */}
```

```
:focus	焦点选择器主要针对表单元素鼠标点击表单框后显示填写的特定元素
```

```
例：.total input {  border: 1px solid #ccc;  height: 30px;  width: 40px;  transition: all .5s;}/*这个input 获得了焦点*/.total input:focus {  width: 80px;  border: 1px solid skyblue;}
```



## 6 复合选择器总结

| 选择器         | 作用                     | 特征                 | 使用情况 | 隔开符号及用法                          |
| -------------- | ------------------------ | -------------------- | -------- | --------------------------------------- |
| 后代选择器     | 用来选择元素后代         | 是选择所有的子孙后代 | 较多     | 符号是**空格** .nav a                   |
| 子代选择器     | 选择 最近一级元素        | 只选亲儿子           | 较少     | 符号是**>**   .nav>p                    |
| 交集选择器     | 选择两个标签交集的部分   | 既是 又是            | 较少     | **没有符号**  p.one                     |
| 并集选择器     | 选择某些相同样式的选择器 | 可以用于集体声明     | 较多     | 符号是**逗号** .nav, .header            |
| 链接伪类选择器 | 给链接更改状态           |                      | 较多     | 重点记住 a{} 和 a:hover  实际开发的写法 |

# 标签显示模式

- 标签的三种显示模式

- 三种显示模式的特点以及区别

- 理解三种显示模式的相互转化

- 什么是标签的显示模式？

  标签以什么方式进行显示，比如div 自己占一行， 比如span 一行可以放很多个

- 作用： 

  我们网页的标签非常多，再不同地方会用到不同类型的标签，以便更好的完成我们的网页。

- 标签的类型(分类)

  HTML标签一般分为块标签和行内标签两种类型，它们也称块元素和行内元素。

## 1行块元素(block-level)

```
常见的块元素有<h1>~<h6>、<p>、<div>、<ul>、<ol>、<li>等，其中<div>标签是最典型的块元素。
```

- 块级元素的特点

（1）比较霸道，自己独占一行

（2）高度，宽度、外边距以及内边距都可以控制。

（3）宽度默认是容器（父级宽度）的100%

（4）是一个容器及盒子，里面可以放行内或者块级元素。

- 注意：
  - 只有 文字才 能组成段落  因此 p  里面不能放块级元素，特别是 p 不能放div 
  - 同理还有这些标签h1,h2,h3,h4,h5,h6,dt，他们都是文字类块级标签，里面不能放其他块级元素。

## 2行内元素(inline-level)

```
常见的行内元素有<a>、<strong>、<b>、<em>、<i>、<del>、<s>、<ins>、<u>、<span>等，其中<span>标签最典型的行内元素。有的地方也成内联元素
```

- 行内元素的特点：

（1）相邻行内元素在一行上，一行可以显示多个。

（2）高、宽直接设置是无效的。

（3）默认宽度就是它本身内容的宽度。

（4）**行内元素只能容纳文本或则其他行内元素。**

- 链接里面不能再放链接。
- 特殊情况a里面可以放块级元素，但是给a转换一下块级模式最安全。

## 3行内块元素（inline-block）

```
在行内元素中有几个特殊的标签——<img />、<input />、<td>，可以对它们设置宽高和对齐属性，有些资料可能会称它们为行内块元素。
```

行内块元素的特点：

（1）和相邻行内元素（行内块）在一行上,但是之间会有空白缝隙。一行可以显示多个
（2）默认宽度就是它本身内容的宽度。
（3）高度，行高、外边距以及内边距都可以控制。

## 4三种模式总结区别

| 元素模式   | 元素排列               | 设置样式               | 默认宽度         | 包含                     |
| ---------- | ---------------------- | ---------------------- | ---------------- | ------------------------ |
| 块级元素   | 一行只能放一个块级元素 | 可以设置宽度高度       | 容器的100%       | 容器级可以包含任何标签   |
| 行内元素   | 一行可以放多个行内元素 | 不可以直接设置宽度高度 | 它本身内容的宽度 | 容纳文本或则其他行内元素 |
| 行内块元素 | 一行放多个行内块元素   | 可以设置宽度和高度     | 它本身内容的宽度 |                          |

## 5标签显示模式转换 display

```
display:inline;块转行内display:block;行内转块display:inline-block;块、行内元素转换行内块
```

## 单行文本垂直居中

 行高我们利用最多的一个地方是： 可以让单行文本在盒子中垂直居中对齐。

> **文字的行高等于盒子的高度。**

这里情况些许复杂，开始学习，我们可以先从简单地方入手学会。

行高   =  上距离 +  内容高度  + 下距离 

上距离和下距离总是相等的，因此文字看上去是垂直居中的。

**行高和高度的三种关系**

- 如果 行高 等 高度  文字会 垂直居中
- 如果行高 大于 高度   文字会 偏下 
- 如果行高小于高度   文字会  偏上 

# CSS 背景(background)

## 1 背景颜色(color)

语法：

```
background-color:颜色值;   默认的值是 transparent  透明的
```

## 2 背景图片(image)

语法： 

```
background-image : none | url (url) 
```

| 参数 |              作用              |
| ---- | :----------------------------: |
| none |       无背景图（默认的）       |
| url  | 使用绝对或相对地址指定背景图像 |

```
background-image : url(images/demo.png);background-imgae:url(); 1,必须加url 2,url里的地址不要加""引号。
```

小技巧：  我们提倡 背景图片后面的地址，url不要加引号。

## 3 背景平铺（repeat）

语法： 

```
background-repeat : repeat | no-repeat | repeat-x | repeat-y 
```

| 参数      |                 作用                 |
| --------- | :----------------------------------: |
| repeat    | 背景图像在纵向和横向上平铺（默认的） |
| no-repeat |            背景图像不平铺            |
| repeat-x  |         背景图像在横向上平铺         |
| repeat-y  |          背景图像在纵向平铺          |

## 4 背景位置(position) 

语法： 

```
background-position : length || lengthbackground-position : position || position background-position;x坐标 y坐标；背景位置background-position:right top 右上角； left bottom左下角；center center水平居中垂直居中；两个值顺序没有关系
```

| 参数     |                              值                              |
| -------- | :----------------------------------------------------------: |
| length   |         百分数 \| 由浮点数字和单位标识符组成的长度值         |
| position | top \| center \| bottom \| left \| center \| right   方位名词 |

注意：

- 必须先指定background-image属性
- position 后面是x坐标和y坐标。 可以使用方位名词或者 精确单位。
- 如果指定两个值，两个值都是方位名字，则两个值前后顺序无关，比如left  top和top  left效果一致
- 如果只指定了一个方位名词，另一个值默认居中对齐。
- 如果position 后面是精确坐标， 那么第一个，肯定是 x  第二的一定是y
- 如果只指定一个数值,那该数值一定是x坐标，另一个默认垂直居中
- 如果指定的两个值是 精确单位和方位名字混合使用，则第一个值是x坐标，第二个值是y坐标

## 5 背景附着

- 背景附着就是解释背景是滚动的还是固定的

- 语法： 

```
background-attachment : scroll | fixed background-attachment:soroll默认背景滚动  fixed背景固定
```

| 参数   |           作用           |
| ------ | :----------------------: |
| scroll | 背景图像是随对象内容滚动 |
| fixed  |       背景图像固定       |

## 6 背景简写

- background：属性的值的书写顺序官方并没有强制标准的。为了可读性，建议大家如下写：
- background: 背景颜色 背景图片地址 背景平铺 背景滚动 背景位置;
- 语法：

```
background: transparent url(image.jpg) repeat-y  scroll center top ;
```

## 7 背景透明(CSS3)

语法：

```
background: rgba(0, 0, 0, 0.3);
```

- 最后一个参数是alpha 透明度  取值范围 0~1之间
- 我们习惯把0.3 的 0 省略掉  这样写  background: rgba(0, 0, 0, .3);
- 注意：  背景半透明是指盒子背景半透明， 盒子里面的内容不受影响
- 因为是CSS3 ，所以 低于 ie9 的版本是不支持的。

```
示例：background:#ccc url(地址) no-repeat fixed;backgroudn:-webkit-linear-gradient(left,颜色，颜色);线性渐变background: rgba(0, 0, 0, .2)背景透明a代表透明度百分比background-size:背景图片宽度 背景图片高度；单位：长度｜百分比｜cover|contain;只写一个单位另一个单位等比例缩放。cover把背景图片扩展至足够大，以使背景图片完全覆盖区域。宽度高度同时扩展直至同时填满盒子contain把背景图片扩展至足够到，宽度或高度其中一个先到边缘停止扩展。
```

## 8 背景总结

| 属性                  | 作用             | 值                                                           |
| --------------------- | :--------------- | :----------------------------------------------------------- |
| background-color      | 背景颜色         | 预定义的颜色值/十六进制/RGB代码                              |
| background-image      | 背景图片         | url(图片路径)                                                |
| background-repeat     | 是否平铺         | repeat/no-repeat/repeat-x/repeat-y                           |
| background-position   | 背景位置         | length/position    分别是x  和 y坐标， 切记 如果有 精确数值单位，则必须按照先X 后Y 的写法 |
| background-attachment | 背景固定还是滚动 | scroll/fixed                                                 |
| 背景简写              | 更简单           | 背景颜色 背景图片地址 背景平铺 背景滚动 背景位置;  他们没有顺序 |
| 背景透明              | 让盒子半透明     | background: rgba(0,0,0,0.3);   后面必须是 4个值              |

# CSS 三大特性

## 1 CSS层叠性

- 概念：

  所谓层叠性是指多种CSS样式的叠加。

  是浏览器处理冲突的一个能力,如果一个属性通过两个相同选择器设置到同一个元素上，那么这个时候一个属性就会将另一个属性层叠掉

- 原则：

  - 样式冲突，遵循的原则是**就近原则。** 那个样式离着结构近，就执行那个样式。
  - 样式不冲突，不会层叠

  ## 2 CSS继承性

  - 概念：

    子标签会继承父标签的某些样式，如文本颜色和字号。

     想要设置一个可继承的属性，只需将它应用于父元素即可。

  简单的理解就是：  子承父业。

  - **注意**：
    - 恰当地使用继承可以简化代码，降低CSS样式的复杂性。比如有很多子级孩子都需要某个样式，可以给父级指定一个，这些孩子继承过来就好了。
    - 子元素可以继承父元素的样式（**text-，font-，line-这些元素开头的可以继承，以及color属性**）

  ## 3 CSS优先级

  类选择器、属性选择器、伪类选择器，权重为10

  内联》ID选择器》伪类=属性选择器=类选择器》元素选择器【p】》通配符选择器(*)》继承的样式

  内联样式1000》id选择器100》class选择器10》标签选择器1

  内联式 > 嵌入式 > 外部式

  #### 1). 权重计算公式

  关于CSS权重，我们需要一套计算公式来去计算，这个就是 CSS Specificity（特殊性）

  | 标签选择器             | 计算权重公式 |
  | ---------------------- | ------------ |
  | 继承或者 *             | 0,0,0,0      |
  | 每个元素（标签选择器） | 0,0,0,1      |
  | 每个类，伪类           | 0,0,1,0      |
  | 每个ID                 | 0,1,0,0      |
  | 每个行内样式 style=""  | 1,0,0,0      |
  | 每个!important  重要的 | ∞ 无穷大     |

  - 值从左到右，左面的最大，一级大于一级，数位之间没有进制，级别之间不可超越。 
  - 关于CSS权重，我们需要一套计算公式来去计算，这个就是 CSS Specificity（特殊性）
  - div {
        color: pink!important;  
    }


  #### 2). 权重叠加

  我们经常用交集选择器，后代选择器等，是有多个基础选择器组合而成，那么此时，就会出现权重叠加。

  就是一个简单的加法计算

  - div ul  li   ------>      0,0,0,3
  - .nav ul li   ------>      0,0,1,2
  - a:hover      -----—>   0,0,1,1
  - .nav a       ------>      0,0,1,1

   <img src="J:/黑马/01-03 前端开发基础/02-CSS资料/02-CSS资料/CSS-Day02/笔记/media/w.jpg" /> 注意： 

    1. 数位之间没有进制 比如说： 0,0,0,5 + 0,0,0,5 =0,0,0,10 而不是 0,0, 1, 0， 所以不会存在10个div能赶上一个类选择器的情况。

  #### 3). 继承的权重是0

  这个不难，但是忽略很容易绕晕。其实，我们修改样式，一定要看该标签有没有被选中。

  1） 如果选中了，那么以上面的公式来计权重。谁大听谁的。 
  2） 如果没有选中，那么权重是0，因为继承的权重为0.

# 盒子

## 1.盒子边框（border） 

语法：

```
border : border-width || border-style || border-color 
```

| 属性         |          作用          |
| ------------ | :--------------------: |
| border-width | 定义边框粗细，单位是px |
| border-style |       边框的样式       |
| border-color |        边框颜色        |

边框的样式：

- none：没有边框即忽略所有边框的宽度（默认值）
- solid：边框为单实线(最为常用的)
- dashed：边框为虚线  
- dotted：边框为点线



```
border边框；padding内边距；margin外边距；border-width:像素border-style:solid实线；dashed虚线；dotted点线；border-color:颜色
```

边框粗细 边框样式 边框颜色

padding

1个值上下左右相等

2个值上下和左右

3个值上 左右 下

4个值上 右 下 左顺序

```
margin: auto 居中
```

2个值代表上下和左右

```
*{padding:0px; margin:0px}清楚元素默认内外边距
```

padding在没有宽度时不会撑开盒子

margin只继承最高数值的像素，左右边距不继承。只在上下相邻的元素里出现。

margin对嵌套会出现塌陷（嵌套关系，垂直外边距合并），解决方法：

```
1、border-top:1px solid transparent；给父级定义一个上边框2、padding-top:1px;给父级定义一个padding值3、给父元素添加overflow:hidden;
```

width>padding>margin优先使用

4、圆角矩形

按照顺时针的顺序书写

```
Border-radius:length圆角边框
```

```
border-top-left-radius:像素border-top-right-radius:像素border-bottom-right-radius:像素border-bottom-left-radius:像素border-radius:左上 右上 右下 左下（四个角）
```

border-radius:50%圆形 border-radius:10px圆角长条

box-shadow:水平阴影 垂直阴影 模糊距离（虚实） 阴影尺寸（影子大小） 阴影颜色 内/外阴影

```
box-shadow: 2px 2px 2px 2px #000 inset/outset
```

# 浮动

## CSS 布局的三种机制

CSS 提供了 **3 种机制**来设置盒子的摆放位置，分别是**普通流**（标准流）、**浮动**和**定位**，其中： 

1. **普通流**（标准流）
   - **块级元素**会独占一行，**从上向下**顺序排列；
     - 常用元素：div、hr、p、h1~h6、ul、ol、dl、form、table
   - **行内元素**会按照顺序，**从左到右**顺序排列，碰到父元素边缘则自动换行；
     - 常用元素：span、a、i、em等
2. **浮动**
   - 让盒子从普通流中**浮**起来,主要作用让多个块级盒子一行显示。
3. **定位**
   - 将盒子**定**在浏览器的某一个**位**置——CSS 离不开定位，特别是后面的 js 特效。

## 浮动

#### 语法

在 CSS 中，通过 `float`  中文，  浮 漏 特    属性定义浮动，语法如下：

选择器 { float: 属性值; }

| 属性值    | 描述                     |
| --------- | ------------------------ |
| **none**  | 元素不浮动（**默认值**） |
| **left**  | 元素向**左**浮动         |
| **right** | 元素向**右**浮动         |

浮动会改变display属性，把块级元素改为行内块元素

**float** —— **浮漏特**

| 特点 | 说明                                                         |
| ---- | ------------------------------------------------------------ |
| 浮   | 加了浮动的盒子**是浮起来**的，漂浮在其他标准流盒子的上面。   |
| 漏   | 加了浮动的盒子**是不占位置的**，它原来的位置**漏给了标准流的盒子**。 |
| 特   | **特别注意**：浮动元素会改变display属性， 类似转换为了行内块，但是元素之间没有空白缝隙 |

## 清除浮动的条件

1、父级没有高度

2、子盒子又浮动了

3、影响下面布局，我们就应该清除浮动。

- 语法：

```
选择器{clear:属性值;}   clear 清除  
```

| 属性值 | 描述                                       |
| ------ | ------------------------------------------ |
| left   | 不允许左侧有浮动元素（清除左侧浮动的影响） |
| right  | 不允许右侧有浮动元素（清除右侧浮动的影响） |
| both   | 同时清除左右两侧浮动的影响                 |

但是我们实际工作中， 几乎只用 clear: both;

### 1、额外标签法

在浮动的最后一个元素增加一个元素

例：

```
.clear{clear:both}"<div class="clear"></div>"
```

### 2、给父级添加overflow属性方法

例：

```
.father{overflow:hidden,auto（会加上下滚动条），scroll（会加上下左右滚动条）}
```

### 3、使用after伪元素清楚浮动

:after 方式为空元素额外标签法的升级版，好处是不用单独加标签了

```
例：.clearfix:after{				content:"";​								display:block;​								height:0;​								visibility:hidden;​								clear:both;}.clearfix{*zoom:1;“清除ie6 ie7的浮动”}
```

#### 4、使用双伪元素清除浮动



```
例：.clearfix::befor，.clearfix::after	{content:"“​				display:table;}.clearfix::after	{clear:both;}.clearfix	{*zoom:1;	“清除ie6 ie7的浮动”}
```



# 定位模式

## 1 边偏移

简单说， 我们定位的盒子，是通过边偏移来移动位置的。

在 CSS 中，通过 `top`、`bottom`、`left` 和 `right` 属性定义元素的**边偏移**：（方位名词）

| 边偏移属性 | 示例           | 描述                                                     |
| ---------- | :------------- | -------------------------------------------------------- |
| `top`      | `top: 80px`    | **顶端**偏移量，定义元素相对于其父元素**上边线的距离**。 |
| `bottom`   | `bottom: 80px` | **底部**偏移量，定义元素相对于其父元素**下边线的距离**。 |
| `left`     | `left: 80px`   | **左侧**偏移量，定义元素相对于其父元素**左边线的距离**。 |
| `right`    | `right: 80px`  | **右侧**偏移量，定义元素相对于其父元素**右边线的距离**   |

定位的盒子有了边偏移才有价值。 一般情况下，凡是有定位地方必定有边偏移。

## 2定位(position)

在 CSS 中，通过 `position` 属性定义元素的**定位模式**，语法如下：

```
选择器 { position: 属性值; }
```

定位模式是有不同分类的，在不同情况下，我们用到不同的定位模式。

| 值         |     语义     |
| ---------- | :----------: |
| `static`   | **静态**定位 |
| `relative` | **相对**定位 |
| `absolute` | **绝对**定位 |
| `fixed`    | **固定**定位 |

绝对定位不能通过auto进行居中

解决方法：例：

1.left:50%px;让盒子的左侧移动到父级元素的水平中心位置；

2.margin-left:-100px;让盒子像左移动自身宽度的一般。

### 2.1 静态定位(static) - 了解

- **静态定位**是元素的默认定位方式，无定位的意思。它相当于 border 里面的none， 不要定位的时候用。
- 静态定位 按照标准流特性摆放位置，它没有边偏移。
- 静态定位在布局时我们几乎不用的 

### 2.2 相对定位(relative) - 重要

- **相对定位**是元素**相对**于它  原来在标准流中的位置 来说的。（自恋型）相对定位的特点：（务必记住）
  - 相对于 自己原来在标准流中位置来移动的
  - 原来**在标准流的区域继续占有**，后面的盒子仍然以标准流的方式对待它。

#### 2.3 绝对定位(absolute) - 重要  

**绝对定位**是元素以带有定位的父级元素来移动位置 （拼爹型）

1. **完全脱标** —— 完全不占位置；  
2. **父元素没有定位**，则以**浏览器**为准定位（Document 文档）。
3. **父元素要有定位**
   * 将元素依据最近的已经定位（绝对、固定或相对定位）的父元素（祖先）进行定位。

绝对定位的特点：（务必记住）

- 绝对是以带有定位的父级元素来移动位置 （拼爹型） 如果父级都没有定位，则以浏览器文档为准移动位置
- 不保留原来的位置，完全是脱标的。

因为绝对定位的盒子是拼爹的，所以要和父级搭配一起来使用。

##### 定位口诀 —— 子绝父相

刚才咱们说过，绝对定位，要和带有定位的父级搭配使用，那么父级要用什么定位呢？

**子绝父相** —— **子级**是**绝对**定位，**父级**要用**相对**定位。

#### 2.4 固定定位(fixed) - 重要

**固定定位**是**绝对定位**的一种特殊形式： （认死理型）   如果说绝对定位是一个矩形 那么 固定定位就类似于正方形

1. **完全脱标** —— 完全不占位置；
2. 只认**浏览器的可视窗口** —— `浏览器可视窗口 + 边偏移属性` 来设置元素的位置；
   * 跟父元素没有任何关系；单独使用的
   * 不随滚动条滚动。



## 堆叠顺序

z-index:

属性值：

1.正整数、负整数或者0 默认值为0，数值越大，盒子越靠上

2.如果属性值相同，则按照书写顺序，后写居上；

3.数字后面不能加单位；

只能应用与相对定位、绝对定位和固定定位的元素，其他标准流、浮动和静态定位无效；

# 显示与隐藏

## 1、display

```
display：none 隐藏对象display:block 除了转化为块级元素之外，同时还有显示元素的意思。
```

隐藏，但不占据位置

## 2、visibility

```
visibility:visible； 对象可视visibility:hidden； 对象隐藏
```

特点：隐藏后，继续保留原有位置。

## 3、overflow溢出(重点)

visible；不剪切内容也不添加滚动条

hidden；不显示超过对象尺寸的内容，超出的部分隐藏掉

scroll；不管超出内容否，总是显示滚动条

auto；超出自动显示滚动条，不超出不显示滚动条

## 对溢出部分显示省略号的写法

```
white-space: normal;	默认处理方式white-space: nowrap;	强制一行内显示文字（除非遇到br标签才能换行）text-overflow: clip;	不显示省略号（……），而是简单的剪裁text-overflow: ellipsis;	当对象内文本溢出时隙那好似省略号（……）
```

一定要首先强制一行内显示，再次和overflow属性配合

```
/*1.先强制一行内显示文本*/​	white-space: nowrap;/*2.超出的部分隐藏*/​	overflow: hidden;/*3.文字用省略号替代超出的部分*/​	text-overflow: ellipsis;
```



# 鼠标样式cursor

```
style:cursordefault;默认鼠标样式pointer;小手 连接中应用move；移动  图片中应用text；文本  索引中应用not-allowed；禁止  
```

# 轮廓线outline

outline:outline-color;  outline-style;  outline-width

大多数情况都去掉

outline:0;outline:none;却掉写法

```
<input type="text" style="outline:0;"/>
```

# 文本域防拖拽

```
<textarea style="resize:none;"></textarea>>
```

# vertical-align垂直对齐

```
vertical-align:baseline基线|top顶线|middle中线|bottom底线
```

只能应用于行内元素和行内块元素

## 溢出的文字省略号显示

### 1.white-space

white-space设置或检索对象内文本显示方式。

```
white-space: normal;	默认处理方式white-space: nowrap;	强制在同一行内显示所有文本，知道文本结束或遭遇br标签对象才换行
```

### 2.text-overflow文字溢出

设置或检索是否使用一个省略标记...标示对象内文本的溢出

```
text-overflow: clip;	不显示省略标记...，而是简单的剪裁text-overflow: ellipsis; 当对象内文本溢出时显示省略号标记...
```

一定要首先强制一行内显示，再次和overflow属性搭配使用

```
1.强制一行内显示文本white-space: nowrap;2.超出的部分隐藏overflow: hidden;3.文字用省略号代替超出的部分text-overflow: ellipsis;
```

## 解决图片有空白缝隙问题

1、只要不让图片和基线对齐都能解决图片底边空白缝隙

2、给img添加display：block转换成块级元素

# 自定义字体样式

```
@font-face {   font-family: 'icomoon';   src:  url('fonts/icomoon.eot?7kkyc2');   src:  url('fonts/icomoon.eot?7kkyc2#iefix') format('embedded-opentype'),    url('fonts/icomoon.ttf?7kkyc2') format('truetype'),    url('fonts/icomoon.woff?7kkyc2') format('woff'),    url('fonts/icomoon.svg?7kkyc2#icomoon') format('svg');   font-weight: normal;   font-style: normal;}
```



# 过度

通常配合:hover,如果有多组属性配合逗号隔开

属性就是你想要变化的css属性，宽度高度背景颜色内外边距都可以。如果想要所有属性都变化过度，写一个all就可以。过度写在本身上

```
transition:要过度的属性 花费时间 运动曲线 何时开始；transition:width        1s     ease默认    0s默认
```

transition-property	规定应用过度的css属性的名称

transition-duration	定义过渡效果花费的时间。默认0

transition-timing-function	规定过渡效果的时间曲线。默认是"ease"

transition-delay	规定过渡效果何时开始。默认是0

transitionend		监听结束

例：

```
div {			width: 200px;			height: 100px;			background-color: pink;			/* transition: 要过渡的属性  花费时间  运动曲线  何时开始; */			transition: width 0.6s ease 0s, height 0.3s ease-in 1s;			/* transtion 过渡的意思  这句话写到div里面而不是 hover里面 */}div:hover {  /* 鼠标经过盒子，我们的宽度变为400 */			width: 600px;			height: 300px}transition: all 0.6s;  /* 所有属性都变化用all 就可以了  后面俩个属性可以省略 */
```

# CSS3

## 属性选择器

E[att]选择具有att属性的E元素

E[att="val"]选择具有att属性且属性值等于val的E元素

E[att^="val"]匹配具有att属性、且值以val“开头”的E元素

E[att$="val"]匹配具有att属性、且值以val“结尾”的E元素

E[att*="val"]匹配具有att属性、且值“中含”有val的E元素

```
例1：input[type="search"]{} <input type="search">例2：input[type^="icon"]{}<input type="icon1">例3：input[type$="icon"]<input type="absicon">例4：input[type*="icon"]<input type="icon"><input type="icon1"><input type="absicon">
```

## 伪类选择器

E:first-child匹配父元素中的第一个子元素E

E:last-child匹配父元素中的最后一个子元素E

E:nth-child(n)匹配父元素中的第n个子元素E；匹配父元素中的第 n 个子元素，元素类型没有限制。

```
例1：ul li:first-child{} /*第一个li*/ul li:last-child{} /*最后一个li*/ul li:nth-child(n){} /*第n个li*/<ul> <li>1</li> </ul><ul> <li>2</li> </ul><ul> <li>3</li> </ul>例2：div span:nth-child(1){}  /*第一个子元素为p标签，将无法显示描述属性。第二个字元							素为span，选择2将执行描述属性。*/<div>	<p></p>	<span></span></div>
```

nth-child(n)
	n可以是数字、关键字和公式
	n如果是数字，就是选择第几个
	常见的关键词有even 偶数 odd 奇数

​	常见的公式如下(如果n是公式，则从0开始)

​	但是第0个元素或者超出了元素的数会被忽略

​	

```
2n	偶数2n+1	奇数5n	0*5	1*5	2*5n+5	从第5个开始包括第五个到最后-n+5	-1*5	-2*5 ......
```

E:first-of-type指定类型E的第一个

E:last-of-type指定类型E的最后一个

E:nth-of-type指定类型E的第n个；匹配同类型中的第n个同级兄弟元素。

```
div span:nth-of-type(2){}	/*类型将不受父子元素影响，可直接描述类选择器。*/<div>	<p></p>	<span></span></div>
```

## 伪元素选择器

::before	在元素内容的前面插入内容

::after		在元素的内容的后面插入内容

必须搭配content书写

before和after创建一个元素，但是属于行内元素

因为在dom里面看不见刚才创建的元素，所以我们称为伪元素

伪元素和标签选择器一样，权重为1

```
例1：div::before {	content:'我'}例2：P::after{				/*插入字符符号*/	content:'\ea50'	    font-family:'icomoon'}
```

## 2D转换

### 位移

#### `2D` 转换

- `2D` 转换是改变标签在二维平面上的位置和形状
- 移动： `translate`
- 旋转： `rotate`
- 缩放： `scale`

#### translate语法

- x 就是 x 轴上水平移动
- y 就是 y 轴上水平移动

```
transform:translate(x,y)或者分开写transform:translatex(n);transform:translatey(n);
```

定义2D转化中的移动，沿着X和Y轴移动的元素

translate最大的优点：不会影响其他元素的位置

translate中的百分比单位是相对于自身元素的translate:(50%,50%); 盒子自身的百分比

对行内标签没有效果

```
例：div {  background-color: lightseagreen;  width: 200px;  height: 100px;  /* 平移 */  /* 水平垂直移动 100px */  /* transform: translate(100px, 100px); */  /* 水平移动 100px */  /* transform: translate(100px, 0) */  /* 垂直移动 100px */  /* transform: translate(0, 100px) */  /* 水平移动 100px */  /* transform: translateX(100px); */  /* 垂直移动 100px */  transform: translateY(100px)}
```

### 旋转

```
transform:rotate（度数）
```

rotate里面跟度数，单位是deg比如rotate(45deg)

角度为正时，顺时针，负时，为逆时针

默认旋转的中心点是元素的中心点

```
例：img:hover {  transform: rotate(360deg)}
```

### rotate

> 2d旋转指的是让元素在2维平面内顺时针旋转或者逆时针旋转

使用步骤：

1. 给元素添加转换属性 `transform`
2. 属性值为 `rotate(角度)`  如 `transform:rotate(30deg)`  顺时针方向旋转**30度**

```
例：div{      transform: rotate(0deg);}
```

### 三角

#### 转换中心点transform-origin

```
transform-origin: x y;
```

主意后面的参数x和y用空格隔开

x y默认转换的中心点是元素的中心点(50% 50%)

还可以给x y设置 像素 或者 方位名词 (top bottom left right center)

### 缩放

```
transform:scale(x,y);
```

主意其中的x和y用逗号 分隔

transform:scale(1,1) :宽和高都放大一倍，相当于没有放大

transform:scale(2,2):宽和高都放大2倍

transform:scale(2):只写一个参数，第二个参数则和第一个参数一样，相当于scale(2,2)

tramsform:scale(0.5,0.5):缩小

sacle缩放最大的优势：可以设置转换中心点缩放，默认以中心点缩放的，而且不影响其他盒子

### 2D转换中scale

`scale` 的作用

用来控制元素的放大与缩小

```
transform: scale(x, y)
```

知识要点

- 注意，x 与 y 之间使用逗号进行分隔

- `transform: scale(1, 1)`: 宽高都放大一倍，相当于没有放大

- `transform: scale(2, 2)`: 宽和高都放大了二倍

- `transform: scale(2)`: 如果只写了一个参数，第二个参数就和第一个参数一致

- `transform:scale(0.5, 0.5)`: 缩小

  `scale` 最大的优势：可以设置转换中心点缩放，默认以中心点缩放，而且不影响其他盒子

```
例：div:hover {	   /* 注意，数字是倍数的含义，所以不需要加单位 */	   /* transform: scale(2, 2) */   	   /* 实现等比缩放，同时修改宽与高 */	   /* transform: scale(2) */   	   /* 小于 1 就等于缩放*/	   transform: scale(0.5, 0.5)   }
```



### 综合写法

格式：transform:translate()  rotate() scale() 等

其顺序会影响转化的效果。(先旋转会改变坐标轴方向)

当我们同时有位移和其他属性的时候，记得要讲位移放到最先面

```
例：div:hover {  transform: translate(200px, 0) rotate(360deg) scale(1.2)}
```



## 动画

##### 语法格式(定义动画)

```
例：@keyframes 动画名称 {    0% {        width: 100px;    }    100% {        width: 200px    }}
```

##### 语法格式(使用动画)

```
例：div {	/* 调用动画 */    animation-name: 动画名称; 	/* 持续时间 */ 	animation-duration: 持续时间；}
```

##### 动画序列

- 0% 是动画的开始，100 % 是动画的完成，这样的规则就是动画序列
- 在 @keyframs 中规定某项 CSS 样式，就由创建当前样式逐渐改为新样式的动画效果
- 动画是使元素从一个样式逐渐变化为另一个样式的效果，可以改变任意多的样式任意多的次数
- 用百分比来规定变化发生的时间，或用 `from` 和 `to`，等同于 0% 和 100%

```
例：<style>    div {      width: 100px;      height: 100px;      background-color: aquamarine;      animation-name: move;      animation-duration: 0.5s;    }    @keyframes move{      0% {        transform: translate(0px)      }      100% {        transform: translate(500px, 0)      }    }  </style>
```

##### 常见属性

| 属性                      | 描述                                                         |
| ------------------------- | ------------------------------------------------------------ |
| @keyframes                | 关键帧，规定动画                                             |
| animation                 | 所有动画属性的简写属性，除了animation-play-state             |
| animation-name            | 规定@keyframes动画的名称(必须的)                             |
| animation-duration        | 规定动画完成一个周期所花费的秒(s)或毫秒(ms)，默认是0.(必须的) |
| animation-timing-function | 规定动画的速度曲线，默认是"ease"                             |
| animation-delay           | 规定动画何时开始，默认是0.                                   |
| animation-iteration-count | 规定动画被播放的次数，默认是1，infinite重复播放              |
| animation-direction       | 规定动画是否在下一周逆向播放，默认"normal",""alternate"逆播放 |
| animation-play-state      | 规定动画是否正在运行或暂停。默认是"running",还有"paused"     |
| animation-fill-mode       | 规定动画结束后状态，留在结束点forwards，回到起始backwards    |

```
例：div {  width: 100px;  height: 100px;  background-color: aquamarine;  /* 动画名称 */  animation-name: move;  /* 动画花费时长 */  animation-duration: 2s;  /* 动画速度曲线 */  animation-timing-function: ease-in-out;  /* 动画等待多长时间执行 */  animation-delay: 2s;  /* 规定动画播放次数 infinite: 无限循环 */  animation-iteration-count: infinite;  /* 是否逆行播放 */  animation-direction: alternate;  /* 动画结束之后的状态 */  animation-fill-mode: forwards;}div:hover {  /* 规定动画是否暂停或者播放 */  animation-play-state: paused;}
```

##### 动画简写方式

animation:动画名称 持续时间 运动时间 何时开始 播放次数 是否反向 动画起始或者结束的状态；

```
例：animation: myfirst 5s linear 2s infinite alternate;
```

##### 知识要点

简写属性里面不包含animation-play-state

暂停动画：animation-play-state: puase;经常和鼠标经过等其他配合使用

想要动画走回来，而不是直接跳回来：animation-direction : alternate

盒子动画结束后，停在结束位置: animation-fill-mode : forwards

```
animation-timing-function:规定动画的速度曲线，默认"ease"linear	动画从头到尾的速度是相同。匀速ease	默认。动画以低速开始，然后加快，在结束前变慢。ease-in	动画以低速开始。ease-out	动画以低速结束。ease-in-out	动画以低速开始和结束steps()	指定了时间函数中的间隔数量(步长)
```

##  3D转换

##### `3D` 的特点

- 近大远小
- 物体和面遮挡不可见

##### 三维坐标系

- x 轴：水平向右  -- **注意：x 轴右边是正值，左边是负值**

- y 轴：垂直向下  -- **注意：y 轴下面是正值，上面是负值**

- z 轴：垂直屏幕  --  **注意：往外边的是正值，往里面的是负值**

##### 3D位移

translate3d(x,y,z)

​				transform:translateX()

​				transform:translateY()

​				transform:translateZ()

3D旋转：rotate3d(x,y,z,deg)

透视:perspective:像素

透视写再被观察得元素得父盒子上面。透视越小显示越大，透视越大显示越小。

3D呈现transfrom-style

```
transform-style:flat 子元素不开启3D立体空间默认transform-style:preserve-3d;子元素开启立体空间代码写给父级，但是影响得是子合资
```

##### `3D` 移动 `translate3d`

- `3D` 移动就是在 `2D` 移动的基础上多加了一个可以移动的方向，就是 z 轴方向

- `transform: translateX(100px)`：仅仅是在 x 轴上移动

- `transform: translateY(100px)`：仅仅是在 y 轴上移动

- `transform: translateZ(100px)`：仅仅是在 z 轴上移动

- `transform: translate3d(x, y, z)`：其中x、y、z 分别指要移动的轴的方向的距离

- **注意：x, y, z 对应的值不能省略，不需要填写用 0 进行填充**

  ##### 语法

  ```
   transform: translate3d(x, y, z)
  ```

  ```
  例：transform: translate3d(100px, 100px, 100px)/* 注意：x, y, z 对应的值不能省略，不需要填写用 0 进行填充 */transform: translate3d(100px, 100px, 0)
  ```

  ### translateZ

  `translateZ` 与 `perspecitve` 的区别

  - `perspecitve` 给父级进行设置，`translateZ` 给 子元素进行设置不同的大小

### `3D` 旋转`rotateX`

语法

- `transform: rotateX(45deg)` -- 沿着 x 轴正方向旋转 45 度
- `transform: rotateY(45deg)` -- 沿着 y 轴正方向旋转 45 度
- `transform: rotateZ(45deg)` -- 沿着 z 轴正方向旋转 45 度
- `transform: rotate3d(x, y, z, 45deg)` -- 沿着自定义轴旋转 45 deg 为角度

```
例：div {  perspective: 300px;}img {  display: block;  margin: 100px auto;  transition: all 1s;}img:hover {  transform: rotateX(-45deg)}
```

#### 左手准则

- 左手的手拇指指向 x 轴的正方向

- 其余手指的弯曲方向就是该元素沿着 x 轴旋转的方向

- ![img](file://J:/%E9%BB%91%E9%A9%AC/04-06%20%E7%A7%BB%E5%8A%A8Web%E7%BD%91%E9%A1%B5%E5%BC%80%E5%8F%91/01-H5C3%20%E8%BF%9B%E9%98%B6%E8%B5%84%E6%96%99/01-HTML5CSS3_day03/01-%E7%AC%94%E8%AE%B0/images/rotateX.png?lastModify=1617288016)

- ### `3D` 旋转 `rotateY`

```
例：div {  perspective: 500px;}img {  display: block;  margin: 100px auto;  transition: all 1s;}img:hover {  transform: rotateY(180deg)
```

#### 左手准则

- 左手的拇指指向 y 轴的正方向

- 其余的手指弯曲方向就是该元素沿着 y 轴旋转的方向(正值)

- ![img](file://J:/%E9%BB%91%E9%A9%AC/04-06%20%E7%A7%BB%E5%8A%A8Web%E7%BD%91%E9%A1%B5%E5%BC%80%E5%8F%91/01-H5C3%20%E8%BF%9B%E9%98%B6%E8%B5%84%E6%96%99/01-HTML5CSS3_day03/01-%E7%AC%94%E8%AE%B0/images/rotateY.png?lastModify=1617288091)

  ### 3D 旋转 rotateZ

```
例：div {  perspective: 500px;}img {  display: block;  margin: 100px auto;  transition: all 1s;}img:hover {  transform: rotateZ(180deg)}
```

#### rotate3d

- `transform: rotate3d(x, y, z, deg)` -- 沿着自定义轴旋转 deg 为角度
- x, y, z 表示旋转轴的矢量，是标识你是否希望沿着该轴进行旋转，最后一个标识旋转的角度
  - `transform: rotate3d(1, 1, 0, 180deg)` -- 沿着对角线旋转 45deg
  - `transform: rotate3d(1, 0, 0, 180deg)` -- 沿着 x 轴旋转 45deg

```
div {  perspective: 500px;}img {  display: block;  margin: 100px auto;  transition: all 1s;}img:hover {  transform: rotate3d(1, 1, 0, 180deg)}
```

#### `3D` 呈现 transform-style

- 控制子元素是否开启三维立体环境
- `transform-style: flat`  代表子元素不开启 `3D` 立体空间，默认的
- `transform-style: preserve-3d` 子元素开启立体空间
- 代码写给父级，但是影响的是子盒子

## 浏览器私有前缀

-moz-：代表firefox浏览器私有属性

-ms-：代表ie浏览器私有属性

-webkit-：代表safari、chrome私有属性

-o-：代表opera私有属性

## meta视口标签

width					宽度设置得是viewport宽度，可以设置device-width特殊值

initial-scale			初始缩放比，大于0的数字

maximum-scale	最大缩放比，大于0的数字

minimum-scale	最小缩放比，大于0的数字

user-scalable		用户是否可以缩放，yes或no(1或者0)

```
例：<meta name="viewport(理想视口)" content="width=device-width, user-scalable=no,initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0>
```

## css3盒子写法

box-sizing:border-box;将不用计算border和padding多增加的宽度

box-sizing：content-box;传统盒子模型

## 特殊样式

```
/*css3盒子模型*/box-sizing:border-box;-webkit-box-sizing:border-box;/*点击高亮我们需要清楚 设置为transparent 完全透明*/-webkit-tap-highlight-color:transparent;/*在移动端浏览器默认的外观在ios上加上这个属性才恩那个给按钮和输入框自定义样式*/-webkit-appearance:none;/*禁用长按页面时的弹出菜单*/img,a{-webkit-touch-callout:none;}
```

# 移动端技术选择

### 1、单独制作移动端页面(主流)

+ 流式布局（百分比布局）
+ flex 弹性布局（强烈推荐）
+ less+rem+媒体查询布局
+ 混合布局

响应式

+ 媒体查询
+ bootstarp

流式布局：

流式布局，就是百分比布局，也称非固定像素布局。

通过盒子的宽度设置成百分比来根据屏幕的宽度来进行伸缩，不受固定像素的限制，内容向两侧填充。

流式布局方式是移动web开发使用的比较常见的布局方式。

通过

max-width 最大宽度（max-height最大高度）
min-widht 最小宽度（min-height最小高度）

### flex弹性布局(iE9级浏览器适用)

#### flex布局原理

+ flex 是 flexible Box 的缩写，意为"弹性布局"，用来为盒状模型提供最大的灵活性，任何一个容器都可以指定为 flex 布局。
+ 当我们为父盒子设为 flex 布局以后，子元素的 float、clear 和 vertical-align 属性将失效。
+ flex布局又叫伸缩布局 、弹性布局 、伸缩盒布局 、弹性盒布局 
+ 采用 Flex 布局的元素，称为 Flex 容器（flexcontainer），简称"容器"。它的所有子元素自动成为容器成员，称为 Flex 项目（flex
  item），简称"项目"。

**总结**：就是通过给父盒子添加flex属性，来控制子盒子的位置和排列方式

在父级元素中加入flex

```
例：display:flex;{在flex布局中，子元素的float、clear和vertical-align属性将失效}
```

#### 父项常见属性

##### flex-direnction:设置主轴的方向；

+ 在 flex 布局中，是分为主轴和侧轴两个方向，同样的叫法有 ： 行和列、x 轴和y 轴
+ 默认主轴方向就是 x 轴方向，水平向右
+ 默认侧轴方向就是 y 轴方向，水平向下
+ 注意： 主轴和侧轴是会变化的，就看 flex-direction 设置谁为主轴，剩下的就是侧轴。而我们的子元素是跟着主轴来排列的
+ ![img](file://J:/%E9%BB%91%E9%A9%AC/04-06%20%E7%A7%BB%E5%8A%A8Web%E7%BD%91%E9%A1%B5%E5%BC%80%E5%8F%91/02-Flex%20%E4%BC%B8%E7%BC%A9%E5%B8%83%E5%B1%80%E8%B5%84%E6%96%99/02-%E7%A7%BB%E5%8A%A8WEB%E5%BC%80%E5%8F%91_flex%E5%B8%83%E5%B1%80/4-%E7%AC%94%E8%AE%B0/images/1.jpg?lastModify=1617288615)



| row            | 默认值从左到右 |
| -------------- | -------------- |
| row-reverse    | 从右到左       |
| column         | 从上倒下       |
| column-reverse | 从下到上       |

​	

##### justify-content:设置主轴上的子元素排列方式

| flex-start    | 默认值从头部开始 如果主轴是x轴，则从左到右 |
| ------------- | ------------------------------------------ |
| flex-end      | 从尾部开始排列                             |
| center        | 在主轴居中对齐(如果主轴是x轴则 水平居中)   |
| space-around  | 平分剩余空间                               |
| space-between | 先两边贴边，再平分剩余空间(重要)           |

##### flex-warp:设置子元素是否换行；

默认情况下，项目都排在一条线（又称”轴线”）上。flex-wrap属性定义，flex布局中默认是不换行的。

```
nowarp;不换行warp；自动换行
```

##### align-content:设置侧轴上的子元素的排列方式(多行)；

	flex-start;默认值在侧轴的头部开始排列flex-end;在侧轴的尾部开始排列center；在侧轴中间显示space-around;子项在侧轴平分剩余空间space-between;子项在侧轴先分布在两头，在平分剩余空间stretch;设置子项元素高度平分父元素高度


##### align-items:设置侧轴上的子元素排列方式(单行)；

设置子项在侧轴上的排列方式 并且只能用于子项出现 换行 的情况（多行），在单行下是没有效果的。

	flex-start;默认值从上到下flex-end；从下到上center；挤在一起居中(垂直居中)stretch；拉伸

##### align-content 和align-items区别

+ align-items  适用于单行情况下， 只有上对齐、下对齐、居中和 拉伸
+ align-content适应于换行（多行）的情况下（单行情况下无效）， 可以设置 上对齐、下对齐、居中、拉伸以及平均分配剩余空间等属性值。 
+ 总结就是单行找align-items  多行找 align-content

##### flex-flow

:复合属性，相当于同时设置了flex-direction和flex-warp;

##### flex：

让div独占一份

```
div {
	flex: 1
}



##### align-self

:属性允许单个项目有与其他项目不一样的对齐方式，可覆盖align-items属性。
默认值为auto，表示继承父元素的align-items属性，如果没有父元素，则等同于strech。
例：	span:nth-child(2){
		align-self:flex-end;
	}

##### order

属性数值越小，排名越靠前，默认为0；
注意：和z-index不一样

div {

​	//把位置移到第一个盒子前面

​	order: -1

}

### less+rem+媒体查询布局

#### rem基础

##### rem单位

rem (root em)是一个相对单位，类似于em，em是父元素字体大小。

不同的是rem的基准是相对于html元素的字体大小。

比如，根元素（html）设置font-size=12px; 非根元素设置width:2rem; 则换成px表示就是24px。

```
/* 根html 为 12px */
html {
font-size: 12px;
}
/* 此时 div 的字体大小就是 24px */       
div {
font-size: 2rem;
}
```

rem的优势：父元素文字大小可能不一致， 但是整个页面只有一个html，可以很好来控制整个页面的元素大小。

混合布局

### 2、响应式页面兼容移动端(其次)

### 媒体查询

#### 什么是媒体查询

媒体查询（Media Query）是CSS3新语法。

+ 使用 @media查询，可以针对不同的媒体类型定义不同的样式
+ @media 可以针对不同的屏幕尺寸设置不同的样式
+ 当你重置浏览器大小的过程中，页面也会根据浏览器的宽度和高度重新渲染页面 
+ 目前针对很多苹果手机、Android手机，平板等设备都用得到多媒体查询

#### 媒体查询语法规范

+ 用 @media开头 注意@符号
+ mediatype  媒体类型
+ 关键字 and  not  only
+ media feature 媒体特性必须有小括号包含

```
@media mediatype and|not|only (media feature) {    CSS-Code;}
```

1. mediatype 查询类型

       将不同的终端设备划分成不同的类型，称为媒体类型

| 值    | 解释说明                           |
| ----- | ---------------------------------- |
| all   | 用于所有设备                       |
| print | 用于打印机和打印预览               |
| scree | 用于电脑屏幕，平板电脑，智能手机等 |

2. 关键字

       关键字将媒体类型或多个媒体特性连接到一起做为媒体查询的条件。

+ and：可以将多个媒体特性连接到一起，相当于“且”的意思。
+ not：排除某个媒体类型，相当于“非”的意思，可以省略。
+ only：指定某个特定的媒体类型，可以省略。    

3. 媒体特性

   每种媒体类型都具体各自不同的特性，根据不同媒体类型的媒体特性设置不同的展示风格。我们暂且了解三个。

   注意他们要加小括号包含

| 值        | 解释说明                           |
| --------- | ---------------------------------- |
| width     | 定义输出设备中页面可见区域的宽度   |
| min-width | 定义输出设备中页面最小可见区域宽度 |
| max-width | 定义输出设备中页面最大可见区域宽度 |

媒体查询书写规则

注意： 为了防止混乱，媒体查询我们要按照从小到大或者从大到小的顺序来写,但是我们最喜欢的还是从小到大来写，这样代码更简洁

```
例：@media screen and (max-width: 800px) {
	body {
		background-color: red;
	}
}

@media screen and (max-width: 500px) {
	body {
		background-color: purple;
	}
}
```

引入资源

```
<link rel="stylesheet" href="style320.css" media="screen and (min-width: 320px)">
<link rel="stylesheet" href="style320.css" media="screen and (min-width: 640px)">
```

#### Less基础

需要在node中先下载less模块

##### 1.less变量

变量是指没有固定的值，可以改变的。

```
@变量名：值
```

```
例：
在less文件中
@color: red;
body {
	background-color: @color;
}
div {
	color: @color
}
```

##### 2.less编译

配套使用VScode中的easy less，可以生成css文件。

##### 3.less嵌套

父子集嵌套

```
.header {
	width: 200px;
	height: 200px;
	background-color: pink;
	a {
		color: red;
	}
}
```

伪类选择器的用法

```
a {
	color: red;
	//如果有伪类、交集选择器、伪元素选择器 我们内层选择器的前面需要加&
	&:hover {
		color: blue;
	}
}
```

##### 4.less运算

可以直接在less文件中进行运算

```
@border: 5px + 5;
div {
	width: 200px -50;
	height: 200px * 2;
	border: @border solid red;
}
img {
	width: 82 / 50rem
}
```

![img](file://J:/%E9%BB%91%E9%A9%AC/04-06%20%E7%A7%BB%E5%8A%A8Web%E7%BD%91%E9%A1%B5%E5%BC%80%E5%8F%91/03-%E7%A7%BB%E5%8A%A8web%E5%BC%80%E5%8F%91%E8%B5%84%E6%96%99/03-%E7%A7%BB%E5%8A%A8WEB%E5%BC%80%E5%8F%91_rem%E5%B8%83%E5%B1%80/4-%E7%AC%94%E8%AE%B0/images/3.png?lastModify=1617893265)

#### bootstarp

1.container类

响应式布局的容器 固定宽度

大屏(>=1200px) 宽度定为1170px

中屏(>=992px) 宽度定为970px

小屏(>=992px) 宽度定为750px

超小屏(100%)



2.container-fluid类

流式布局容器 百分百宽度

占据全部视口(viewport)的容器

适合单于做移动端开发

