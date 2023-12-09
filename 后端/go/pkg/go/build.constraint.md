## package build/constraint

包约束实现了构建约束行的解析和评估。有关构建约束本身的文档，请参阅 https://golang.org/cmd/go/#hdr-Build_constraints。

此软件包既能解析原始的"// +build "语法，也能解析 Go 1.17 中新增的"//go:build "语法。有关"//go:build "语法的详细信息，请参见 https://golang.org/design/draft-gobuild。

## Index

### func GoVersion(x Expr) string 添加于1.21.0

GoVersion 返回给定构建表达式所隐含的最小 Go 版本。如果表达式无需任何 Go 版本标记即可满足要求，GoVersion 返回空字符串。

例如：

```go
GoVersion(linux && go1.22) = "go1.22"
GoVersion((linux && go1.22) || (windows && go1.20)) = "go1.20" => go1.20
GoVersion(linux) = ""
GoVersion(linux || (windows && go1.22)) = ""
GoVersion(!go1.22) = "
```

GoVersion 假定任何标记或否定标记都可能独立为真，因此它的分析可以是纯结构性的，而无需 SAT 求解。因此，"不可能 "子表达式可能会影响结果。

例如:

```go
GoVersion((linux && !linux && go1.20) || go1.21) = "go1.20"
```

### IsGoBuild(line string) bool

IsGoBuild 报告文本行是否为"//go:build "约束。它只检查文本的前缀，而不检查表达式本身是否能解析。

### func IsPlusBuild(line string) bool

IsPlusBuild 报告文本行是否为"// +build "约束。它只检查文本的前缀，而不检查表达式本身是否能解析。

### func PlusBuildLines(x Expr) ([]string, error)

如果表达式过于复杂，无法直接转换为"// +build "行，PlusBuildLines 将返回错误信息。

### type AndExpr

```go
type AndExpr struct {
	X, Y Expr
}
```

AndExpr 表示表达式 X && Y。

#### func (x *AndExpr) Eval(ok func(tag string) bool) bool

#### func (x *AndExpr) String() string

### type Expr

```go
type Expr interface {
	// String 使用 go:build 行中的布尔语法返回表达式的字符串形式。
	String() string

	// Eval 报告表达式是否求值为 true。它会根据需要调用 ok(tag)，以确定当前构建配置是否满足给定的构建标记。
	Eval(ok func(tag string) bool) bool
	// 包含已过滤或未导出的方法
}
```

Expr 是构建标记约束表达式。其基础具体类型是 *AndExpr、*OrExpr、*NotExpr 或 *TagExpr。

#### func Parse(line string) (Expr, error)

Parse 会解析格式为"//go:build ..." 或"//+build ..." 的单行构建约束，并返回相应的布尔表达式。

### type NotExpr

```go
type NotExpr struct {
	X Expr
}
```

NotExpr 表示表达式 !X（X 的否定）。

#### func (x *NotExpr) Eval(ok func(tag string) bool) bool

#### func (x *NotExpr) String() string

### type OrExpr

```go
type OrExpr struct {
	X, Y Expr
}
```

OrExpr 表示表达式 X || Y。

#### func (x *OrExpr) Eval(ok func(tag string) bool) bool

#### func (x *OrExpr) String() string

### type SyntaxError

```go
type SyntaxError struct {
	Offset int // 检测到错误的输入字节偏移量
	Err string // 错误描述
}
```

语法错误（SyntaxError）报告已解析的构建表达式中的语法错误。

#### func (e *SyntaxError) Error() string

### type TagExpr

```go
type TagExpr struct {
	Tag string // 例如，"linux "或 "cgo"
}
```

TagExpr 是单个标签 Tag 的 Expr。

#### func (x *TagExpr) Eval(ok func(tag string) bool) bool

#### func (x *TagExpr) String() string