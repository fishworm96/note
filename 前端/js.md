## JS输入输出语句



alert(msg)	浏览器弹出警示框

console.log(msg)	浏览器控制台打印输出信息

prompt(info)	浏览器弹出输入框，用户可以输入

confirm()	确认框

## 变量

var是JS关键字，用来声明变量

age是定义的变量名

var age;   声明一个名称为age的变量

```
age=10;
```

"="用来把右边的值赋予给左边的变量空间中 此处代表赋值的意思

```
由字母(A-Za-z)、数字(0-9)、下划线(_)、美元符号( $ )组成，如：usrAge, num01, _name
严格区分大小写。 var app; 和 var App; 是两个变量
不能 以数字开头。 18age 是错误的
不能 是关键字、保留字。 例如：var、 for、 while
变量名必须有意义。 MMD BBD nl → age
遵守驼峰命名法。首字母小写，后面单词的首字母需要大写。 myFirstName
推荐翻译网站： 有道 爱词霸
```

## 数据类型

```
Number	数字型，包含整型值和浮点型值，如21、0.21  默认：0
Boolean	布尔值类型，如true、false，等价于1和0	默认：false
string	字符串类型，如“张三”注意咱们js里面，字符串都带引号	默认：''
Undefined var a;声明了变量a但是没有给值，此时a = undefined 默认：undefined
Null	var a = nul;声明了变量a为空值	默认：null
```

### 1.数字型进制

```
1.八进制数字序列范围0～7
var num1 = 07;对应十进制7
var num2 = 019；对应十进制19
var num3 = 08； 对应十进制8
2.十六进制数字序列范围0～9以及A～F
var num= 0xA；
js中八进制前面加0，十六进制前面加0x
```

### 2.字符串string

```
\n	换行符，n是newline的意思
\\	斜杠\
\'	'单引号
\"	"双引号
\t	tab缩进
\b	空格，b是blank的意思
```

### 转换为字符串

```
tostring() 转成字符串	var num= 1; alert(num.toString());
string()强制转换 转成字符串	var num= 1; alert(String(num));
加号拼接字符串 和字符串拼接的结果都是字符串	var num = 1; alert(num+ '空字符串')；
```

### 转换为数字型(重点)

```
parseint(string)函数	将string类型转成整数数值型	parseint('78')
parseFloat(string)函数 将string类型转成浮点数数值型	parseFloat('78.21')
Number()张志转换函数 将string类型转换为数值型	Number('12')
js隐式转换(- * /) 利用算术运算隐式转换为数值型	'12'-0
```

### 转换为布尔型

Boolean()函数	其他类型转成布尔值	boolean('true');

代表空、否定的值会被转换为false ，如 ''、0、 null、 undefined

其余值都会被转换为true

```
console.log(boolean('')) //false
console.log(boolean(0)) //false
console.log(boolean(NaN)) //false
console.log(boolean(nul)) //false
console.log(boolean(undefined)) //false
console.log(boolean('小白')) //true
console.log(boolean(12)) //true
```

# 运算符

## 算数运算符

```
+ 加 1+1=2
- 减 1-1=0
* 乘 1*1=1
/ 除 1/1=1
% 余 3%2=1
```

## 递增和递减

递增(++)和递减(--)

```
例：
var num = 1;
++num; num++;
++先运算再返回   先返回再运算++

```

## 比较运算符

```
< 小于号 1<2 true
> 大于号 1>2 false
>= 大于等于号 2>=2 true
<= 小于等于号 3<=2 false
== 判等号(会转型) 37==37 true 能将字符串转为数字型
!= 不等号 37!=37 false
=== !== 全等要求值和数据类型一致 37==='37' false
```

## 逻辑运算符

```
&& 逻辑“与”	true&&false 返回false
|| 逻辑“或”	true||false 返回true
! 逻辑“非”		!true 返回false
```

## 逻辑中断

```
例
&&表达式1结果为真返回表达式2 表达式1为假返回表达式1
console.log(123 && 456);	返回456
console.log(0 && 456);		返回0
如果有空的或者否定的为假其余都是真的 0 '' null undefined NaN
||表达式1结果为真 返回表达式1 表达式1为假返回表达式2
console.log(123 || 456);	返回123
console.log(123 || )
```

## 赋值运算符

```
= 直接赋值	var useName = '1';
+=、-= 加、减一个数后再赋值	var age = 10; age+=5; //15
*=、 /=、%= 乘、除、取模后再赋值	var age = 2; age*=5; //10
```

## 优先级

```
优先级		运算符		顺序
1		  小括号		()
2		  一元运算符   ++ -- !
3		  算数运算符   先* / % 后+ -
4		  关系运算符   > >= < <=
5		  相等运算符   == !== === !==
6		  逻辑运算符   先&& 后||
7		  赋值运算符   =
8		  逗号运算符   ，
```

# JS流程控制-分支

三元表达式 

num < num1 ? num1 : num;

## if

if(条件){

条件

} else{

}

## else if

if (条件) {

条件

} else if (条件){

条件

} else

## switch

swhitch(条件){

case "条件"：

​	alert(条件)

​	break;

​	default:

​	alert(结束)；

}

# JS循环

## for

for (初始化变量；条件表达式；操作表达式) {		}

```
for (var i = 1; i<= 10; i++){
	console.log(i);
}
```



## 双重for循环

for (外层初始化变量；外层的条件表达式；外层的操作表达式){

​	for (里层初始化变量；里层的条件表达式；里层的操作表达式)

}

```
for (var i = 1; i <= 10; i++){
	for (var j = 1; j <= 10; j++){
		console.log(j);
	}
}
```



## while

var num = 1;

while (num <= 100) {

​	console.log(num);

​	num++;

}

```
例：\
var j = 1
while (j<=10) {
console.log)(j);
j++
}
```

```
var message = prompt('你爱我吗')
while (message !== '我爱你') {
	message = prompt('你爱我吗？')
}
```



## do while

do (

​	循环体代码-——条件表达式true时重复执行循环代码

)while(条件表达式);

```
var i = 1
do {
	console.log(i);
	i++
}	while (i <= 100)
```

```
do {
	条件
}	while (满足条件)
```



## continue break

遇到continue跳出本次继续下一次

```
for (var i = 1; i <= 5;i++){
	if(i == 3) {
	continue;
	}
}
```

遇到break就结束循环

```
for (var i = 1; i <= 5;i++ ){
	if (i == 3) {
		break;
	}
}
```

# 数组

## NEW

var arr = new Array ();  //创建了一个空的数组A要大写

[]中括号，数组表达方式；

var arr = []; //创建了一个空的数组。自变量

```
例：
var arr = [1 , 2, 'pink老师' ,true];	必须逗号分隔。都是数组元素。
索引号 ：    0   1    2         3 
console.log(arr[0]) 返回 1
```

## 遍历

for循环来挨个访问

```
例：
var arr = ['red', 'green', 'blue']
for (var i = 0; i <3; i++) {
	console.log(arr[i]);
}
```

## 数组长度

```
数组名.length
console.log(arr.length);
for (var i = 0; i < arr.length; i++){
	console.log(arr[i]);
}
```

##### 计数器i 数组元素arr[i]

## 追加数组

```
var arr = [1, 2, 3];
arr[3] = 4;追加
arr[3] = 5;替换
不要直接数组元素名赋值，不然就全部变成名字。
```

## 冒泡算法

```
var arr = [5, 4, 3, 2, 1]
for (var i = 0; i <= arr.length -1; i++) {//外层循环管趟数
	for (var j = 0; j <= arr.length -i -1; j++)//里面的循环管每趟的交换次数	if(arr[j]) > arr[j + 1] {
		var temp = arr[j];
		arr[j] = arr[j + 1];
		arr[j + 1] = temp;
	}
}
```

# 函数

## 声明函数



function 函数名 () {

​	函数体

}

调用();

```
例：
function sayHi() {
	console.log('HI')
}
sayHI();
```

## 实参与行参

```
例：
function getSum(str) {	//str是形参
	console.log(str)
}
str(内容)；		//内容是实参
```

## 函数的返回值

```
例：
function 函数名() {
	return 需要返回的结果；
}
函数名();
```

## arguments

arguments存贮了所有的实参；

grauments展示的是一组伪数组；

## 函数命名

function fn() {			//命名函数

​				

}

var fun = function() {		//匿名函数

​	console.log('函数')

}

fun();

## 作用域

## 对象

var obj = {};//创建了一个空的对象

var obj = {

​	name: '张三丰'

​	,age: 18

​	,sex: '难'

​	sayHi: fuction () {

​		console.log('hi');

​	}

}

```
使用方法：调用对象的属性 对象名.属性名

console.log(obj.name);

方法2:对象名['属性名']

console.log(obj['age'])

调用对象的方法

obj.sayHi();
```

### 创建一个新的对象

```
var obj = new Object();

obj.name = '张三丰';
obj.sex = '男';
obj.sayHi = function(){

}
```

### 构造语法(实例化)

```
function 构造函数名() {

​	this.属性 = 值；

​	this.方法 = fuction() {}

}
new 构造函数名();

```

//1.构造函数首字母要大写

//2.构造函数不需要return就可以返回结果

```
例：1
function Star(uname, age, sex) {
    this.name = uname;
    this.age = age;
    this.sex = sex; 
    this.sing = function () {}
}
var ldh = new Star('刘德华', 18, '男')
console.log(ldh.name);
console.log(ldh['sex']);
ldh.sing('冰雨');
```

```
例：2
function hero (name, type, blood) {
    this.name = name;
    this.type = type;
    this.blood = blood;
    this.attack = function(x) {
        console.log(x);
    }
}
var lp = new hero('廉颇', '力量型', '500血量',);
var hy = new hero('后羿', '射手型', '100血量',);
console.log(lp)
lp.attack('近战');
```

### 遍历对象

#### for...in对对象循环

for (变量 in对象) {}  //我们使用for in 里面喜欢变量写k或者key

```
for(var k in obj){
	console.log(k);//k 变量 输出 得到的是属性名
	console.log(obj[k]);//obj[k] 得到的是属性值
}
```

## 内置函数

### 添加删除数组元素方法

#### 1.push

#### 2.unshift

```
var arr = [1, 2, 3];
arr.push(4, 'pink'); //在arr数组末尾添加4和pink
arr.unshift(1, 6); //在arr数组开头添加1和6
Push/unshift是可以给数组追加新的元素
push/unshift() 参数直接写数组元素就可以
push/unshift完毕后，返回的结果是新数组的长度
原数组也会发生变化
```

#### 3.pop

#### **4.shift**

```
arr.pop();//删除数组最后一个元素，括号内不用写数值
arr.shift();//删除数组开头的第一个元素，括号内不写数值
console.log(arr.pop())//返回被删除的元素
原数组也会发生变化
```

```javascript
1.翻转数组
var arr = ['ping', 'red', 'blue'];
arr.reverse();
2.数组排序(冒泡排序)
var arr = [13, 4, 77, 1, 7];
arr1.sort(function(a,b)){
	return a - b; //升序的排序
	return b - a; //降序的排序
}
```

## DOM

创建、增、删、改、查、属性操作、事件操作（注册事件）

```
在document.getElementsByClaseeName /ID里不用区分前缀名利的(. #)直接写
在document.querySelector里
标签类名不用加
class类名前面加.
id类名前面加#
```



```
var str = document.querySelector('')
var str = document.querySelector('.str')
var str = document.querySelector('#str')
```

点击:onclick;

获得焦点:onfocus;

失去焦点:onblur;

鼠标停留:onmouseover;

鼠标离开:onmousout;

鼠标移动触发:onmousemove;

鼠标弹起触发:onmouseup;

鼠标按下触发:onmousedown;

鼠标穿过:onmouseenter;

禁止鼠标右键菜单:contextmenu

禁止鼠标选中:selectstart

### 操作元素内容

（修改普通元素内容）

innerText

innerHTML

### 操作常见元素属性

（修改元素属性）

src、href、title、alt等

### 操作表单元素属性

（修改表单元素）

type、value、disabled（禁用）等

### 操作元素样式属性

（修改元素样式）

element.style

className

```
例
str.style.(比如)colot = 'red';

例
str.calssName = 'str';
```

### 获取属性值

element.属性 获取属性值	//自带属性

element.getAttribute('属性');  //自定义属性

```
修改属性  例:
div.属性='值'
div.id = 'test';
div.setAttribute('属性',值)
div.setAttribute('index',2);
div.setAttribute('class','footer');
```

```
移除属性
div.removeAttribute('index');
```

### H5标准

```
自定义属性都以data-属性为标准 例：getAttribute('data-index');
调用  属性.genAttribute('data-index');
(只能在ie11使用) 可以简化为 div.dateset.index或者div.dateset['index'];
div.getAttrribute('data-list-name') 如果有多个个--连接的单词，获取是用驼峰命名法 改为 div.dateset.listName或div.dateset['listName']
```

## 节点

父节点

node.parentnode

```
例：
<div>
	<span>
	<span>
</div>
span.parentnode只能调用最近的节点
```

子节点

```
例：
<ul>
  <li></li>
  <li></li>
  <li></li>
  <li></li>  
</ul>
ul.childnodes所有的子节点包括 元素节点 文本节点
ul.children 获得所有的子元素节点 非标准(频繁使用)
```

```
ul.firstchild第一个子节点不管文本还是元素
ul.lastchild最后一个子节点不管文本还是元素
ul.firstElementChild返回第一个元素节点  有兼容性问题IE9
ul.lastElementChild返回最后一个元素节点  有兼容性问题IE9
ul.children[o]第一个元素
ul.children[ul.children.length - 1];最后一个元素
```

兄弟节点

```
node.nextSibling下一个兄弟接地那 包含元素节点和文本节点(提倡)

node.previousSibling上一个兄弟接地那 包含元素节点和文本节点

node.nextElementSibling	 元素节点IE9
node.previousElementSibling  元素节点IE9
```

```
封装兄弟节点：
functioin getNextElementSibling(element) {
	var el = element;
	while (el = el.nextSibling) {
		if (el.nodeType === 1) {
			return el;
		}
	}
	return null;
}
```

### 创建节点

```
创建节点
document.createElement('tagName');
例：
var li = document.createElement('li');
```

### 添加节点

```
添加节点
node.appendChild(child);
例：
var ul = document.querySelector('ul');
ul.appendChild(li);
在后面追加节点

插入指定元素的前面
node.insertBefore(child,指定元素)
例：
var lili = document.createElement('li');
ul.insertBefore(lili, ul.children[0]);

```



### 删除节点

```
node.removechild(child0);
例：
<ul>
	<li></li>
	<li></li>
</ul>
 var ul = document.querySelecotr('ul');
 ul.removeChild(ul.children[0]);
```

#### 阻止链接跳转

```
在链接里添加javascript:void(o);或者javascript:;阻止跳转
```

### 复制节点/克隆节点

```
node.cloneNode()

```

### 三种动态创建元素区别

```
document.write()直接将内容写入页面的内容刘，但是文档流执行完毕，则会导致页面全部重绘
例：dcoument.write(<div>123</div>)		
//
innerHTML创建元素
例：inne.innerHTML = '<a href="#">百度</a>'
//创建多个元素效率更高（不要拼接字符串，采用数组形式拼接）
document.createElement()创建元素
var a =document.createElement('a')
//创建多个元素效率稍微低一点点，但是结构更清晰
```

## 注册事件

### 传统注册

```
<button onclick= "alert('hi')"></button>
btn.onchlick = function(){}
注册事件唯一性
```

### 监听注册方式(推荐)

```
addEventListener() IE9以前不支持
代替方法attachEvent()(不建议使用)
evenTarget.addEventListener(type, listener[, useCapture])
type:事件类型字符串，比如 click、mouseover，注意这里不带on
listener:事件处理函数，事件发生时，会调用该监听函数
useCapture:可选参数，是一个布尔值，默认是false。
例：
btns[i].addEventListener('click', function() {
	alret(22);
})
```

```
attachEvent监听(仅了解)
eventTarget.attachEvent(eventName, callback)
eventNasmeWithOn:事件类型字符串,比如onclick、onmouseover,这里要带On
callback:时间处理函数，当目标触发事件时回调函数被调用
```

### 事件对象的常见属性和方法

```
e.target		返回触发事件的对象		标准
e.srcElement	返回触发事件的对象		非标准ie6-8使用
e.type			返回事件的类型 比如click mouseover 不带On
e.cancelBubble	该属性阻止冒泡 非标准 ie6-8使用 e.cancelBubble = true;
e.returnValue	该属性阻止默认事件(默认行为) 非标准 ie6-8使用 比如不让链接跳转
e.preventDefault()	该方法阻止默认时间(默认行为) 标准 比如不让链接跳转
e.stopPropagation() 阻止冒泡  标准
```

### 事件委托原理

不要给每个子节点单独设置事件监听器，而是事件监听器设置在其父节点上，然后利用冒泡原理影响设置每个子节点。

### 鼠标事件对象

```
e.clientX	返回鼠标相当于浏览器窗口可视区的X坐标
e.clientY	返回鼠标相当于浏览器窗口可视区的Y坐标
e.pageX		返回鼠标相当于文档页面的X坐标 IE9支持
e.pageY		返回鼠标相当于文档页面的Y坐标 IE9支持
e.screenX	返回鼠标相对于显示器屏幕的X坐标
e.screenY	返回鼠标相对于显示器屏幕的Y坐标
```

### 常见键盘事件

```
onkeyup		某个键盘按键被松开时触发 执行3
onkeydown	某个键盘按键被按下时触发 执行1
onkeypress 某个键盘按键被按下时触发 但是它不是别功能键比如 ctrl shift箭头等执行2
```

### 移动端添加类

```
classlist
add(class1, class2, ...)	在元素中添加一个或多个类名。
remove(class1, class2, ...)	移除元素中一个或多个类名。
toggle(class, true|false)	在元素中切换类名。
```



## Bom

### 窗口加载事件

```
DOMContentLoaded
document.addEventListener('DOMContentloaded', function() {})
```

### 调整窗口大小事件

```
window.onresize
window.onresize = function(){}
window.addEventListener("resize",function(){});
```

### 重新加载页面触发事件

```
pageshow 重新加载页面
persisted 重新计算缓存 返回 true。这个页面是从缓存取过来的页面也要重新计算刷新。
```



### 定时器

```
setTimeout()	只调用一次
window.setTimeout(调用函数， 延迟的毫秒数);
延迟时间单位是毫秒可以省略，省略默认为0
例：
setTimeout(function(), 3000);
setTimeout('函数名', 3000);

停止clearTimeout()定时器
window.clearTimeout(timeoutId);
```

```
setInterval()	反复调用
window.setInterval(调用函数， 延时毫秒时间)
停止定时器clearInterval();
```

callback回调函数

## 同步与异步

```
异步任务：
1、普通事件,如click、resize等
2、资源加载,如load、error等
3、定时器,包括setInterval、setTimeout等
```

事件循环(event loop)



## location对象

```
location.href		获取或者设置整个URL
location.host		返回主机(域名)
location.port		返回端口号 如果未写返回空字符串
location.pathname	返回路径
location.search		返回参数
location.hash		返回片段 #后面内容 常见于链接 锚点
```

### locatioin对象的方法

```
location.assign()	跟href一样，可以跳转页面(也称为重定向页面)
location.replace()	替换当前页面，因为不记录历史，所以不能后退页面
location.reload()  重新加载页面，相当于刷新按钮f5如果参数为true强制刷新ctrl+f5
```

### navigator对象

```
navigator.userAgent
```

### history对象

```
back()		可以后退功能
forward()	前进功能
go(参数)		前进后退功能 参数如果是1前进1个页面 如果是-1后退1个页面
```

## offset元素偏移量

offset意思就是偏移量，使用offset相关属性可以动态的得到该元素的位置(偏移)、大小等

1、获得元素距离带有定位父元素的位置

2、获得元素自身的大小(宽度高度)

3、注意：返回的数值都不带单位

```
element.offsetParent 返回作为该元素带有定位的父级元素如果父级都没有定位则返回body
element.offsetTop 返回元素相对带有定位父元素上方的偏移
element.offsetLeft 返回元素相对带有定位父元素左边框的偏移
element.offsetWidth 返回自身包括padding、边框、内容区的宽度、返回数值不带单位
element.offsetHeight 返回自身包括padding、边框、内容区的高度，返回数值不带单位
```

元素可视区client系列

```
element.clientTop		返回元素上边框的大小
element.clientLeft		返回元素作边框的大小
element.clientWidth     返回自身包括padding、内容区的宽度，不含边框，返回数值不带单位
element.clientHeight	返回自身包括padding、内容区的高度，不含边框，返回数值不带单位
```

### 立即执行函数

```
(function() {}());第一种写法
(function() {})();第二种写法
立即执行函数最大的作用就是独立创建了一个作用域，里面的都是局部变量不会命名冲突
```

### 元素scroll系列属性

```
element.scrollTop	返回超出的上侧距离，返回数值不带单位  元素
element.scrollLeft	返回超出的左侧距离，返回数值不带单位
element.scrollWidth	返回自身实际的宽度，不含边框，返回数值不带单位
element.scroolHeight返回自身实际的高度，不含边框，返回数值不带单位
```

```
总结
offset系列经常用于获得元素位置 offsetLeft offsetTop
client经常用于获取元素大小 clientWidth clientHeight
scroll经常用于获取滚动距离scrollTop scrollLeft
注意页面滚动的距离通过window.pageXOffset获得
```

## 移动端

### 触屏事件

```
touchstart		手指触摸到一个DOM元素时触发
touchmove		手指在一个DOM元素上滑动时触发
touchend		手指从一个DOM元素上离开时触发
```

### 触摸事件对象

```
touches				正在触摸屏幕的所有手指的一个列表
targetTouches		正在触摸当前DOM元素上的一个手指的一个列表
changedTouches		手指状态发生了改变的列表,从无到有,从又到无变化
```









## 面向对象

### 类

```javascript
类不加小括号	创建一个类class Star{};
对象加小括号	创建一个对象new Star();
构造函数：constructor返回实例对象
//创建类class 创建一个明星类
class Star {
	constructor(uname, age) {
		this.uname = uname;
		this.age = age
	}
}
//利用类创建对象new
var ldh = new Star('刘德华', 18);
console.log(ldh);

类里的函数不用写function
extends:继承	//就近原则，先查找本类里是否存在。
super:调用父类中的构造函数
subtract:减法操作，子类独有的。
class Father {
	constructor(x, y) {
	this.x = x;
	this.y = y;
	}
	sum() {
		console.log(this.x + this.y);
	}
}
calss Son extends Father {		//子类继承父类
	construtor(x, y) {
		//利用super调用了父类中的构造函数
		//super必须在子类this之前调用
		super(x, y);			
		this.x = x;
		this.y = y;
	}
	subtract() {
		console.log(this.x - this.y);
	}
}
var son = new Son(1, 2);
var son1 = new Son(11, 22);
var son2 = new Son(5, 3);
son.sum();//直接调用对象
```

```
element.insertAdjacentHTML(position, text(元素));	将文本插入到dom树中的位置
position:
beforebegin;	插入到父元素前面
afterbegin;		插入到父元素里面的最前面
beforeend;		插入到父元素里面的最后面
afterend;		插入到父元素的后面
text:
是要被解析为HTML或XML元素，并插入到DOM树中的 DOMString。
```



##### 静态成员：在构造函数本身上添加的成员称为静态成员，只能由构造函数本身来访问

```
例：
	function Star(uname, age) {
		this.uname = uname;
		this.age = age;
		this.sing = function() {
			console.log('我会唱歌');
		}
	}
	var ldh = new Star('刘德华', 18);
	Star.sex = '男';	//创建一个静态成员
	console.log(Star.sex)	//调用静态成员
```

##### 实例成员：在构造函数内部创建的对象成员成为实例成员，只能由实例化的对象来访问

```
例：
	function Star(uname, age) {
		this.uname = uname;
		this.age = age;
		this.sing = function() {
			console.log('我会唱歌');
		}
	}
	var ldh = new Star('刘德华', 18);
	//实例成员就是构造函数内部通过this添加的成员 uname age sing 就是实例成员
	//实例成员只能通过实例化的对象来访问
	console.log(ldh.uname);	//实例的属性
	ldh.sing();	//调用实例成员	//实例的方法
	//console.log(Star.uname);	//不可以通过构造函数来访问实例成员
```

#### 构造函数原型

##### prototype：（构造函数）原型对象

##### 原型对象的作用：共享方法

每一个构造函数里面都有一个prototype属性，指向另一个对象。注意这个prototype就是一个对象，这个对象的所有属性和方法，都会被构造函数所拥有。

```
例：
	function Star(uname, age) {
		this.uname = uname;
		this.age = age;
	}
	Star.prototype.sing = function() {
		console.log('我会唱歌')
	}
	var ldh = nwe Star('刘德华', 18);
	ldh.sing();	//通过原型对象来打印
//一般情况下，我们的公共属性定义到构造函数里面，公共的方法我们放到原型对象身上
```

##### __proto__：对象原型

对象都会有一个属性__proto__指向构造函数的prototype原型对象，之所以我们对象可以使用构造函数prototype原型对象的属性和方法，就是因为对象有__proto__原型的存在。

proto对象原型和原型对象prototype是等价的

```
console.log(star.prototype);	//打印原型对象
console.log(ldh.__proto__);		//打印对象原型
```

```
funciton Star() {
	this.uname = uname;
	this.age = age;
}
Star.prototype = {
	//如果我们修改了原来的原型对象，给原型对象赋值的是一个对象，则必须手动的利用constructor指回原来的构造函数
	constructor: Star	//手动指回原来的构造函数
	sing: function() {
		console.log('唱歌');
	}
	movie: function() {
		console.log('演电影');
	}
}
```





#### ES5中的新增方法（函数）

迭代(遍历)：**forEach()**、map()、**fillter()**、**some()**、every();

```
array.forEach(function(currentValue, index, arr))
currentValue:数组当前项的值
index:数组当前项的索引
arr:数组对象本身
例：
var arr=[1, 2, 3];
var sum = 0;
arr.forEach(function(value, index, array)) {
	console.log('每个数组元素' + value);	//第一次返回1
	console.log('每个数组元素的索引号' + index);	//第一次返回0
	console.log('数组本身' + array);	//返回1， 2， 3
	sun += value;	//返回6
}
```

```
array.fillter(function(currentValue, index, arr))
fillter()方法创建一个新的数组，新数组中的元素是通过检查指定数组中符合条件的所有元素，主要用于筛选数组。注意它直接返回一个新数组
currentValue:数组当前项的值
index:数组当前想的索引号
arr:数组对象本身
var arr = [12, 66, 44, 88];
var newArr = arr.fillter(function(value, index)){	//一个新数组
	return value >= 20;	//返回数组中大于20的数
}
```

```
array.some(function(currentValue, index, arr))
some()方法用于检测数组中的元素是否满足指定条件，通俗点查找数组中是否有满足条件的元素
注意它返回值是布尔值，如果查找到这个元素，就返回true，如果查找不到就返回false
如果找到第一个满足条件的元素，则终止循环，不再继续查找
currentValue:数组当前项的值
index:数组当前想的索引号
arr:数组对象本身
var arr = [10, 30, 4];
var flag = arr.some(function(value)) {
	return value >= 20;	//查找返回大于20的值
}
```

#### ES5中新增的对象方法

**Object.keys();**：用于获取对象自身所有的属性

```
Object.keys(obj)
效果类似for...in
返回一个由属性名组成的数组
```

**Object.defineProperty()**：定义对象中新属性或修改原有的属性。

```
Object.defineProperty(obj, prop, descriptor)
obj:必须。目标对象
prop：必须。定义或修改的属性的名称	//有就修改没有就添加
descriptor：必须。目标属性所拥有的特性
第三个参数descriptor说明：以对象形式{}书写
value：设置属性的值 默认为undefined
writable：值是否可以重写。true|false 默认为false
enumerable：目标属性是否可以被枚举(遍历)。true|false默认为false
configurable：目标属性是否可以被删除或是否可以再次修改特性true|false默认为false
例：
	Object.denfineProperty(obj, 'address(属性名)', {
		value: '中国山东找蓝翔',
		wirtable: false;
		enumerable: false;
		configurable: false; //无法再次利用defineProperty修改特性
	})
```

```
delete：删除对象中的属性
delete obj.pname(属性名);
```



### 字符串方法

trim()方法会从一个字符串的两端删除空白字符。

```
str.trim()
例：
var str= '  andy  ';
var str1 = str.trim();
```

trim()方法并不影响原字符串本身，它返回的是一个新的字符串

### 函数

```
1.自定义函数(命名函数)
function fn() {};
2.函数表达式(匿名函数)
var fun = function() {};
3.利用new Function('参数', '参数2', '函数体');构造函数
function fn() {};
var f = new Function('a', 'b', 'console.log(a + b)');
f(1, 2);
4.绑定事件函数
btn.onclick = function () {}; //点击了按钮就可以调用这个函数
5.setInterval(function() {}. 1000);	//这个函数是定时器自动1秒钟调用一次（存在异步）
6.立即执行函数
(function() {};)();	//立即执行函数是自动调用
```

#### this指向

```
1.普通函数this指向window
function fn() {};

2.对象的方法 this指向的是对象
var o = {
	sayHi: function() {}
}

3.构造函数 this指向实力ldh实例对象	原型对象里面的this指向的也是ldh这个实例对象
function Star() {
Star.prototype.sing = funciton(){}
};
var ldh = nwe Star();

4.绑定事件函数 this指向函数的调用者
var btn = document.querySelector('button');
btn.onclick = function() {};

5.定时器函数	this的指向也是window
window.settimeout(function() {}, 1000);

6.立即执行函数 this指向window
(fucntion() {})();
```



#### call()		es5中使用

调用这个函数，并且修改函数运行时的this指向

```
fun.call(thisArg, arg1, arg2, ...)
thisArg:当前调用函数this的指向对象
arg1, arg2 ：传递的其他参数
```

```javascript
functioin fn(x, y) {
	console.log('打印');
}
var o = {
	name: 'andy'
};
//1.call()可以直接调用函数
fn.call();
//2.call()并且改变函数的this的指向，此时这个函数的this就指向了o这个对象
fn.call(o, 1, 2);	

主要作用可以实现继承
function Father(uname, age, sex) {
	this.uname = uname;
	this.age = age;
	this.sex = sex;
}
function Son(uname, age, sex) {
	Father.call(this, uname, age, sex);
}
```

#### apply方法

apply()方法调用一个函数。简单理解为调用函数的方式，但是它可以改变函数的this指向。

```javascript
fun.apply(thisArg, [argsArray])
thisArg：在fun函数运行时指定的this值
arrgsArray：传递的值，必须包含在数组里面
返回值就是函数的返回值，因为它就是调用函数
例：
var o = {
	name: 'andy'
};
function fn() {
    console.log(arr);//'pink' 打印的是字符串，传递的是字符串返回字符串,传递的是数字返回数字。
};
fn.apply(o, ['pink']);
//1.也是调用函数 第二个可以改变函数内部的this指向
//2.但是它的参数必须是数组(伪数组);
//3.apply的主要应用比如说我们可以利用apply借助于数学内置对象求最大值
var arr = [1, 66, 33, 22 44];
var max = Math.max.apply(Math(写调用者), arr); //内置函数求最大值
```

#### bind()方法

bind()方法不会调用函数。但是能改变函数内部this指向

```javascript
fun.bind(thisArg, arg1, arg2, ...)
thisArg：在fun函数运行时指定的this值
arg1， arg2：传递的其他参数
返回由指定的this值和初始化参数改造的原函数拷贝
例：
var o = {
	name: 'andy'
};
function fn(a + b) {
	console.log(this);
};
var f = fn.bind(o, 1, 2);
f();	//返回name: 'andy' 和3
//1.不会调用原来的函数 可以改变原来函数内部this指向
//2.返回的是原函数改变this之后产生的新函数
//3.如果有的函数我们不需要立即调用，但是又想改变这个函数内部的this指向此时用bind
//4.我们有一个按钮，当我们点击了之后，就禁用这个按钮，3秒钟之后开启这个按钮
var btn = document.querySelector('button');
btn.onclick = function() {
	this.disabled(禁用) = true; //这个this 指向的是btn这个按钮
	setTimeout(function() {		//定时器函数里面的this指向的是window
		this.disabled = false;
	}.bind(this), 3000)	//这个this指向的是btn这个对象
}
```

#### 严格模式

```
1.我们的变量名必须先声明在使用
var num = 10;
2.我们不能随意删除已经声明好的变量
delete num;
3.严格模式下全局作用域中函数的this是undefined
4.严格模式下如果构造函数不加new调用，this会报错。
function Star() {
	this.sex = '男';
}
Star();
5.定时器this还是指向window
6.严格模式下函数里面的参数不允许有重名
funciton fn(a, a) {
	console.log(a + a);
}

```

#### 高阶函数

高级函数是对其他函数进行操作的函数，它接收函数作为参数或将函数作为返回值输出

```
例1：
function fn(callback) {
	callback&&callback();
}
fn(function() {alert('hi')})
例2：
function fn() {
	return function() {}
}
fn();
```

#### 闭包

定义：闭包是一个可以访问第一个函数作用域中变量的函数

```
function fn() {
	var num = 10;
	return function() {
		console.log(num);
	}
}
var f = fn();
f();
```

### 递归

定义：一个函数在内部可以调用其本身，那么这个函数就是递归函数。

```
阶乘
function fn(n) {
	if(n == 1) {
		return 1;
	}
	return n * fn(n -1);
}

斐波那契数列（兔子序列）1、1、2、3、5、8
funciton fb(n) {
	if(n === 1 || n === 2) {
		return 1;
	}
	return fb(n-1) + fb(n -2);
}
```

#### 浅拷贝和深拷贝

1.浅拷贝只是拷贝一层，更深层次对象级别的只拷贝引用

2.深拷贝拷贝多层，每一级别的数据都会拷贝

```
ES6中语法糖浅拷贝Object.assign(target,...sources)
var obj = {
	name: 'andy',
	msg: {
		age: 18
	},
	color: ['ping', 'red']
};
var o = {};
object.assign(o, obj);	//将obj对象浅拷贝给o

//封装函数 深拷贝
function deepCopy(newobj, oldobj) {
	for(var k in oldobj) {
		//判断我们的属性值属于哪种数据类型
		//1.获取属性值 oldobj[k]
		var item = oldobj[k];
		//2.判断这个值是否是数组
		//数组也属于对象，判断完对象将跳过数组
		if (item instanceof Array) {
			newobj[k] = [];
			deepCopy(newobj, item)
		} else if (item instanceof Object) {
			//3.判断这个值是否是对象
			newobj[k] = {};
			deepCopy(newobj[k], item)
		} else {
		//4.属于简单数据类型
			newobj[k] = item;
		}
	}
}
deepCopy(o, obj);
```

### 正则表达式

正则表达式(Regular Expression有规律的表达式)适用于匹配字符串中字符组合的模式。在javascript中，正则表达式也是对象。

正则表达式创建方式

```
1.利用RegExp对象来创建正则表达式
var 变量名 = new RegExp(/表达式/)；
var regexp = new RegExp(/123/);
2.利用字面量创建正则表达式
var 变量名 = /表达式/;
var rg = /123/;
```

测试正则表达式test

test()正则表达式方法，用于检测字符串是否符合该规则，该对象会返回true和false，其参数是测试字符串。

```
regexobj.test(str)
1.regexobj是写的正则表达式
2.str我们要测试的文本
3.就是检测str文本是否符合我们写的正则表达式规范
```

##### 边界符

```
^	表示匹配行首的文本(以谁开始)
$	表示匹配行尾的文本(以谁结束)
字符类：[]表示有一系列字符可供选择，只要匹配其中一个就可以	[-]方括号的方位符[a-z]
如果中括号里有^表示取反的意思 
var reg1 = /^[^a-zA-Z0-9-_]/; //表示除了这里的字符
```

量词符

```
*	重复零次或更多次
	*相当于 >=0 可以出现0次或者很多次
	var reg = /^a*$/
+	重复一次或更多次
	+相当于 >= 1可以出现1次或者很多次
	var reg = /^a+$/
?	重复零次或一次
	?相当于1 || 0
	var reg = /^a?$/
{n}	重复n次
{n,}重复n次或者更多次
{n, m}重复n到m次
```

预定义类

```
\d	匹配0-9之间的任一数字，相当于[0-9]
\D	匹配0-9之间的以外数字，相当于[^0-9]
\w	匹配任意的字母、数字和下划线，相当于[a-zA-Z0-9-_]
\W	除了任意的字母、数字和下划线，相当于[^a-zA-Z0-9-_]
\s	匹配空(包括换行符、制表符、空格符)， 相等于[\t\r\n\v\f]
\S	匹配非空格的字符，相当于[^\t\r\n\v\f]
```

##### replace替换

replace()方法可以实现替换字符串操作，用来替换的参数可以是一个字符串或是一个正则表达式。

```
stringObject.replace(regexp/substr,replacement)
1.第一个参数：被替换的字符串或者正则表达式
2.第二个参数：替换为的字符串
3.返回值是一个替换完毕的新字符串
var str = 'andy和red';
var newStr = str.replace(/andy/, 'baby');
返回'andy和red'
```

##### 正则表达式参数

```
/表达式/[switch]
switch(也称为修饰符)按照上面样的模式来匹配有三种值:
g:全局匹配
i:忽略大小写
gi:全局匹配+忽略大小写
结合示范：
	btn.onclick = funciton() {
		div.innerHTML = text.value.replace(/激情|red/g, '**');
	}	//所有输入的激情和red都会被替换为**
```

### ES6

```
let变量
块级作用域
不能重复定义同一个变量名
变量不存在提升
暂时性死区
```

```
const常量
块级作用域
声明常量时必须赋值
常量赋值后，值不能修改
const ary = [100, 200];
ary[0] = 123;
ary = [1, 2] //报错 
console.log(ary);//返回123
```

```
解构赋值
ES6中允许从数组中取值，按照对应位置，对变量赋值。对象也可以实现结构
数组解构
let [a, b, c] = [1, 2, 3];
let ary = [1, 2, 3];
let [a, b, c] = ary;
对象结构
let person = {name: 'lisi', age: 30, sex: '男'};
let { name, age, sex } = person;
console.log(name)	//返回lisi
console.log(age)	//返回30
console.log(sex)	//返回男

let person = { name: 'zhangsan', age:20};
let {name, age } = person;
console.log(name); //'zhangsan'
console.log(age);	//20
let {name: myNam, age: myAge} = person; //myName maAge	属于别名
console.log(myName);	//'zhangsan'
console.log(myAge); //20
```

#### 箭头函数

```
() => {}
const fn = () => {}
函数体中只有一句代码，且代码的执行结果就是返回值，可以省略大括号
function sum(num1, num2) {
	return num1 + num2;
}
const sum = (num1, num2) => num1 + num2;
如果形参只有一个，可以省略小括号
function fn(v) {
	return v;
}
const fn = v => v;
箭头函数不绑定this关键字，箭头函数中的this，指向的是函数定义位置的上下文this
```

#### 剩余参数

剩余参数语法允许我们将一个补丁数量的参数表示为一个数组

```
function sum(first, ...args) {
	console.log(first);	//10
	console.log(args);	//[20, 30]
}
sum(10, 20, 30)
```

剩余参数和解构配合使用

```
let students = ['wangwu', 'zhangsan', 'lisi'];
let [s1, ...s2] = students;
console.log(s1);	//'wangwu'
console.log(s2);	//['zhangsan', 'lisi']
```

#### Array的扩展方法

扩展运算符(展开语法)

扩展运算符可以将数组或者对象转为逗号分隔的参数序列。

```
let ary = [1, 2, 3];
...ary	//1, 2, 3
console.log(...ary)	//1 2 3
```

扩展运算符可以应用于合并数组。

```
//方法一
let ary1 = [1, 2, 3];
let ary2 = [3, 4, 5];
let ary3 = [...ary1, ...ary2];
//方法二
ary1.push(...ary2);
```

将类数组或可遍历对象转换为真正的数组

```
let oDivs = document.getElementsByTagName('div');
oDivs = [...oDivs];	//伪数组
var ary = [...oDivs];	//真数组
目的：为了在数组中追加
ary.push('a');
```

##### 构造函数方法：Arry.from(obj)

```
let arryLike = {
	'0': 'a',
	'1': 'b',
	'2': 'c',
	length: 3
};
let arr2 = Array.from(arrayLike); //['a', 'b', 'c']

var arrayLike = {
	'0': '1',
	'1': '2',
	'length': 2
}
let ary = Array.from(aryLike, item => item *2)
console.log(ary)	//[2, 4]
```

##### 实例方法：find()

用于找出第一个符合条件的数组成员，如果没有找到返回Undefined

```
let ary = [{
	id: 1,
	name: '张三'
}, {
	id: 2,
	name: '李四'
}];
let target = ary.find((item, index) => item.id == 2);
item:当前循环的值
idnex:当前循环的索引
```

##### 实例方法：findIndex()

用于找出**第一个**符合条件的数组成员的位置，如果没有找到返回-1

```
let ary = [1, 5, 10, 15];
let index = ary.findIndex((value, index) => value > 9);
console.log(index); //2
```

##### 实例方法：includex()

表示某个数组是否包含给定的值，返回布尔值

```
[1, 2, 3].includes(2) //true
[1, 2, 3].includes(4) //false
```

#### String的扩展方法

模板字符串

ES6新增的创建字符串的方式，使用反引号。

```
let name = 'zhangsan';
```

模板字符串中可以解析变量

##### ${变量} 相当于 + 变量名

```
let name = '张三';
let sayHello = `hello, my name is ${name}`;	//hello,my name is zhangsan
```

##### 实例方法：startsWith()	和	endsWith()

startsWith():表示参数字符串是否在原字符串的头部，返回布尔值

endsWith():表示参数字符串是否在字符串的尾部，返回布尔值

```
let str = 'Hello world!';
str.startsWith('Hello')	//true
str.endsWith('!')	//true
```

##### 实例方法：repeat()

repeat方法表示将原字符串重复n次，返回一个新字符串

```
'x'.repeat(3)	//'xxx'
'hello'.repeat(2)	//'hellohello'
```

#### Set数据结构

ES6提供了新的数据结构Set。它类似于数组，但是成员的值都是唯一的，没有重复的值。

Set本身是一个构造函数，用来生成Set数据结构。

```
const s = new Set();
console.log(s.size)	//0
```

Set函数可以接收一个数组作为参数，用来初始化。

```
const set = new Set([1, 2, 3, 4, 15]);
console.log(set.size)	//5
```

Set会去除重复的值

```
const s3 = new Set(['a', 'a', 'b', 'b']);
console.log(s3.size)	//2	数组将重复的值去掉
const ary = [...s3];
console.log(art)	//利用Set来完成数组去重
```

```
实例方法
add(value):添加某个值，返回Set结构本身
delete(value):删除某个值，返回一个布尔值，表示删除是否成功
has(value):返回一个布尔值，表示该值是否为Set的成员
clear():清除所有成员，没有返回值
const s = new Set();
s.add(1).add(2).add(3);	//想set结构中添加值
s.delete(2)	//删除set结构中的2值 返回布尔值
s.has(1)	//表示set结构中是否有1这个值 返回布尔值
s.clear()	//清除set结构中的所有值
```

遍历

Set结构的实例和数组一样，也拥有forEach方法，用于对每个成员执行某种操作，没有返回值。

```
s.forEach(value => console.log(value))
```

