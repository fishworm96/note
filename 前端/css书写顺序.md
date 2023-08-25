css书写顺序

1、布局定位属性：display / position / float / clear / visibility / overflow（建议 display 第一个写，毕竟关系到模式）

2、自身属性：width / height / margin / padding / border / background

3、文本属性：color / font / text-decoration / text-align / vertical-align / white- space / break-word

4、其他属性（css3）：content / cursor / border-radius / box-shadow / text-shadow / background:linear-gradient …

```
.jdc {
    display: block;
    position: relative;
    float: left;
    width: 100px;
    height: 100px;
    margin: 0 10px;
    padding: 20px 0;
    font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;
    color: #333;
    background: rgba(0,0,0,.5);
    -webkit-border-radius: 10px;
    -moz-border-radius: 10px;
    -o-border-radius: 10px;
    -ms-border-radius: 10px;
    border-radius: 10px;
}
```



```
画三角

{width:0px;

height:0px;

border-style:solid;

border-width:10px;

border-color:red（上三角） transparent transparent transparent

font-size:0px;（照顾低版本）

line-height:0px

}


```

