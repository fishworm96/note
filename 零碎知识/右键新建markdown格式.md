在 HKEY_CLASSES_ROOT 下面添加一个 .md 的子项。

用记事本打开MD文件的注册表文件：

```shell
Windows Registry Editor Version 5.00

[HKEY_CLASSES_ROOT\.md]
@="MarkdownFile"
"PerceivedType"="text"
"Content Type"="text/plain"

[HKEY_CLASSES_ROOT\.md\ShellNew]
"NullFile"=""

[HKEY_CLASSES_ROOT\MarkdownFile]
@="Markdown File"

[HKEY_CLASSES_ROOT\MarkdownFile\DefaultIcon]
@="%SystemRoot%\system32\imageres.dll,-102"

[HKEY_CLASSES_ROOT\MarkdownFile\shell]

[HKEY_CLASSES_ROOT\MarkdownFile\shell\open]

[HKEY_CLASSES_ROOT\MarkdownFile\shell\open\command]
@="%SystemRoot%\system32\NOTEPAD.EXE %1"
```

如果装了其他的MD编辑器的话只要把 .md 这个项的值改成对应的类型就好了 比如我装的是 Typora 直接导入下面这个注册表文件

```shell
Windows Registry Editor Version 5.00

[HKEY_CLASSES_ROOT\.md]
@="TyporaMarkdownFile"
"PerceivedType"="text"
"Content Type"="text/plain"

[HKEY_CLASSES_ROOT\.md\ShellNew]
"NullFile"=""
```

```shell
Windows Registry Editor Version 5.00
 
[HKEY_CLASSES_ROOT\.md]
@="Typora.md"
"Content Type"="text/markdown"
"PerceivedType"="text"
 
[HKEY_CLASSES_ROOT\.md\ShellNew]
"NullFile"=""
```