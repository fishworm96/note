## package driver

包驱动程序定义了数据库驱动程序要实现的接口，如包 sql 所用。

大多数代码都应使用包 sql。

随着时间的推移，驱动程序接口也在不断演变。驱动程序应实现 Connector 和 DriverContext 接口。Connector.Connect 和 Driver.Open 方法不应返回 ErrBadConn。只有在连接已处于无效（如已关闭）状态时，Validator、SessionResetter 或查询方法才会返回 ErrBadConn。

所有 Conn 实现都应实现以下接口：Pinger、SessionResetter 和 Validator。

如果支持命名参数或上下文，则驱动程序的 Conn 应实现以下接口：ExecerContext、QueryerContext、ConnPrepareContext 和 ConnBeginTx。

要支持自定义数据类型，请实现 NamedValueChecker。通过从 CheckNamedValue 返回 ErrRemoveArgument，NamedValueChecker 还允许查询接受按查询选项作为参数。

如果支持多个结果集，则 Rows 应实现 RowsNextResultSet。如果驱动程序知道如何描述返回结果中的类型，则应实现以下接口：RowsColumnTypeScanType、RowsColumnTypeDatabaseTypeName、RowsColumnTypeLength、RowsColumnTypeNullable 和 RowsColumnTypePrecisionScale。给定的行值也可能返回一个 Rows 类型，该类型可能代表数据库游标值。

连接使用后返回连接池之前，如果已实现，则会调用 IsValid。在连接被重新用于其他查询之前，如果已执行，则会调用 ResetSession。如果连接从未返回连接池，而是立即重新使用，那么在重新使用前会调用 ResetSession，但不会调用 IsValid。

## Index

### Variables

```go
var Bool boolType
```

Bool 是一个 ValueConverter，用于将输入值转换为 bool。

转换规则如下:

- 返回的布尔值不变
- 对于整数类型，1 为真 0 为假，其他整数为错误
- 对于字符串和[]字节，规则与 strconv.ParseBool 相同
- 其他类型均为错误

```go
var DefaultParameterConverter defaultConverter
```

DefaultParameterConverter 是 ValueConverter 的默认实现，当 Stmt 没有实现 ColumnConverter 时就会使用。

如果是 IsValue(arg)，DefaultParameterConverter 会直接返回其参数。否则，如果参数实现了 Valuer，则会使用其 Value 方法返回一个值。作为后备，所提供参数的底层类型会被转换为 Value：底层整数类型会被转换为 int64，浮点类型会被转换为 float64，bool、string 和 []byte 类型会被转换为它们自己。如果参数是一个 nil 指针，ConvertValue 将返回一个 nil Value。如果参数为非 nil 指针，则会取消引用，并递归调用 ConvertValue。其他类型将导致错误。

```go
var ErrBadConn = errors.New("driver: bad connection")
```

ErrBadConn 应由驱动程序返回，以向 sql 包发出信号，表明驱动程序 Conn 处于不良状态（如服务器已提前关闭连接），因此 sql 包应在新连接上重试。

为防止重复操作，如果数据库服务器可能已执行了操作，则不应返回 ErrBadConn。即使服务器返回错误，也不应返回 ErrBadConn。

将使用 errors.Is 检查错误。错误可以封装 ErrBadConn 或实现 Is(error) bool 方法。

```go
var ErrRemoveArgument = errors.New("driver: remove argument from query")
```

ErrRemoveArgument 可能会从 NamedValueChecker 返回，以指示 sql 包不向驱动程序查询接口传递参数。在接受非 SQL 查询参数的查询特定选项或结构时返回。

```go
var ErrSkip = errors.New("driver: skip fast-path; continue as if unimplemented")
```

某些可选接口的方法可能会返回 ErrSkip，以在运行时表明快速路径不可用，sql 包应继续运行，就像该可选接口未被实现一样。ErrSkip 仅在有明确记录的情况下才支持。

```go
var Int32 int32Type
```

Int32 是一种 ValueConverter，它能将输入值转换为 int64，并遵守 int32 值的限制。

```go
var ResultNoRows noRows
```

ResultNoRows 是预定义的结果，供驱动程序在 DDL 命令（如 CREATE TABLE）成功时返回。对于 LastInsertId 和 RowsAffected，它都会返回错误。

```go
var String stringType
```

String 是一个 ValueConverter，用于将输入转换为字符串。如果值已经是字符串或[]字节，则保持不变。如果值是其他类型，则使用 fmt.Sprintf("%v", v) 将其转换为字符串。

### func IsScanValue(v any) bool

IsScanValue 等同于 IsValue。它的存在是为了兼容。

### func IsValue(v any) bool

IsValue 报告 v 是否是有效的 Value 参数类型。

### type ColumnConverter 废除

### type Conn

```go
type Conn interface {
  // Prepare 返回与此连接绑定的预制语句。
  Prepare(query string) (Stmt, error)

  // // 关闭会使当前准备好的语句和事务无效并可能停止，同时将此连接标记为不再使用。
  // 
  // 因为 sql 包会维护一个空闲的连接池，只有当闲置连接过多时才会调用关闭，所以驱动程序不需要自己进行连接缓存。
  // 
  // 驱动程序必须确保 Close 进行的所有网络调用不会无限期阻塞（例如应用超时）。
  Close() error

  // 开始并返回一个新事务。
  // 
  // 已弃用：驱动程序应替代（或额外）实现 ConnBeginTx。
  Begin() (Tx, error)
}
```

Conn 是与数据库的连接。多个程序不会同时使用它。

Conn 假定是有状态的。

### type ConnBeginTx 添加于1.8

```go
type ConnBeginTx interface {
  // BeginTx 启动并返回一个新事务。如果用户取消了上下文，SQL 程序包将在丢弃和关闭连接之前调用 Tx.Rollback。
	//
	// 这必须检查 opts.Isolation，以确定是否设置了隔离级别。如果驱动程序不支持非默认级别且已设置，或者不支持非默认隔离级别，则必须返回错误信息。
	//
	// 还必须检查 opts.ReadOnly，以确定只读值是否为 true，以便在支持的情况下设置只读事务属性，或在不支持的情况下返回错误。
  BeginTx(ctx context.Context, opts TxOptions) (Tx, error)
}
```

ConnBeginTx 通过上下文和 TxOptions 增强 Conn 接口。

### type ConnPrepareContext 添加于1.8

```go
type ConnPrepareContext interface {
  // PrepareContext 返回与此连接绑定的已准备语句。上下文用于准备语句，不得在语句本身中存储上下文。
  PrepareContext(ctx context.Context, query string) (Stmt, error)
}
```

ConnPrepareContext 利用上下文增强 Conn 接口。

### type Connector 添加于1.10

```go
type Connector interface {
  // Connect 返回数据库连接。Connect 可能会返回一个缓存连接（先前关闭的连接），但这样做是不必要的；sql 包会维护一个闲置连接池，以便有效地重新使用。
	//
	// 提供的 context.Context 仅用于拨号目的（参见 net.DialContext），不应存储或用于其他目的。拨号时仍应使用默认超时，因为连接池可能对任何查询异步调用 Connect。
	//
	// 返回的连接一次只能被一个程序使用。
  Connect(context.Context) (Conn, error)

	// Driver 返回连接器的底层驱动程序，主要是为了与 sql.DB.NET 上的 Driver 方法保持兼容。
  Driver() Driver
}
```

一个 Connector 代表一个固定配置的驱动程序，可以创建任意数量的等效 Conn，供多个程序使用。

Connector 可以传递给 sql.OpenDB，允许驱动程序实现自己的 sql.DB 构造函数，也可以由 DriverContext 的 OpenConnector 方法返回，允许驱动程序访问上下文，避免重复解析驱动程序配置。

如果 Connector 实现了 io.Closer，则 sql 包的 DB.Close 方法将调用 Close 并返回错误（如有）。

### type Driver

```go
type Driver interface {
  // Open 返回一个新的数据库连接。名称是一个字符串，格式为特定于驱动程序的格式。
	//
	// Open 可能会返回一个缓存连接（以前关闭的连接），但这样做是不必要的；sql 包会维护一个闲置连接池，以便有效地重复使用。
	//
	// 返回的连接一次只能被一个程序使用。
  Open(name string) (Conn, error)
}
```

驱动程序是数据库驱动程序必须实现的接口。

数据库驱动程序可以实现 DriverContext，以便访问上下文，并对连接池只解析一次名称，而不是对每个连接解析一次。

### type DriverContext 添加于1.10

```go
type DeriverContext interface {
  // // OpenConnector 必须按照 Driver.Open 解析名称参数的格式解析名称。
  OpenConnector(name string) (Connector, error)
}
```

如果驱动程序实现了 DriverContext，那么 sql.DB 将调用 OpenConnector 来获取连接器，然后调用该连接器的 Connect 方法来获取每个所需的连接，而不是为每个连接调用驱动程序的 Open 方法。这两步序列允许驱动程序只解析一次名称，还提供了对每个 Conn 上下文的访问。

### type Execer 废除

### type ExecerContext 添加于1.8

```go
type ExecerContext interface {
  ExecContext(ctx context.Context, query string, args []NamedValue) (Result, error)
}
```

ExecerContext 是一个可选接口，可由 Conn 实现。

如果 Conn 没有实现 ExecerContext，sql 包的 DB.Exec 将返回 Execer；如果 Conn 也没有实现 Execer，DB.Exec 将首先准备查询、执行语句，然后关闭语句。

ExecContext 可能会返回 ErrSkip。

ExecContext 必须遵守上下文超时，并在取消上下文时返回。

### type IsolationLevel 添加于1.8

```go
type IsolationLevel int
```

IsolationLevel 是存储在 TxOptions 中的事务隔离级别。

该类型应被视为与 sql.IsolationLevel 及其定义的任何值相同。

### type NamedValue 添加于1.8

```go
type namedValue struct {
  // 如果 Name 不为空，则应使用它作为参数标识符，而不是序号位置。
  // 
  // 名称不带符号前缀。
  Name string

  // 参数的序数位置，从 1 开始并始终设置。
  Ordinal int

  // Value 是参数值。
  Value Value
}
```

NamedValue 包含值名和值。

### type NamedValueChecker 添加于1.9

```go
type NamedValueChecker interface {
  // CheckNamedValue 在将参数传递给驱动程序之前被调用，并代替任何 ColumnConverter 被调用。CheckNamedValue 必须根据驱动程序的需要进行类型验证和转换。
  CheckNamedValue(*NamedValue) error
}
```

NamedValueChecker 可选择由 Conn 或 Stmt 实现。它为驱动程序提供了更多控制，使其能够处理 Go 和数据库类型，而不局限于允许的默认值类型。

sql 包按照以下顺序检查值校验器，在发现第一个匹配时停止：Stmt.NamedValueChecker、Conn.NamedValueChecker、Stmt.ColumnConverter、DefaultParameterConverter。

如果 CheckNamedValue 返回 ErrRemoveArgument，NamedValue 将不会包含在最终查询参数中。这可用于向查询本身传递特殊选项。

如果返回 ErrSkip，列转换器错误检查路径将用于参数。驱动程序可能希望在用尽自己的特殊情况后返回 ErrSkip。

### type NotNull

```go
type NotNull struct {
  Converter ValueConverter
}
```

NotNull 是一种实现 ValueConverter 的类型，它不允许使用 nil 值，但会委托给另一个 ValueConverter。

#### func (n NotNull) ConvertValue(v any) (Value, error)

### type Null

```go
type Null struct {
  Converter ValueConverter
}
```

Null 是一种实现 ValueConverter 的类型，它允许 nil 值，但会委托给另一个 ValueConverter。

#### func (n Null) ConvertValue(v any) (Value, error)

### Pinger 添加于1.8

```go
type Pinger interface {
  Ping(ctx context.Context)
}
```

Pinger 是一个可选接口，可由 Conn 实现。

如果 Conn 没有实现 Pinger，sql 包的 DB.Ping 和 DB.PingContext 将检查是否至少有一个 Conn 可用。

如果 Conn.Ping 返回 ErrBadConn，DB.Ping 和 DB.PingContext 将从池中移除 Conn。

### type Queryer 废除

### type QueryerContext 添加于1.8

```go
type QueryerContext interface {
  QueryContext(ctx context.Context, query string, args []NamedValue) (Rows, error)
}
```

QueryerContext 是一个可选接口，可由 Conn 实现。

如果 Conn 没有实现 QueryerContext，sql 包的 DB.Query 将退回到 Queryer；如果 Conn 也没有实现 Queryer，DB.Query 将首先准备查询、执行语句，然后关闭语句。

QueryContext 可能会返回 ErrSkip。

QueryContext 必须遵守上下文超时，并在取消上下文时返回。

### type Result

```go
type Result interface {
  // LastInsertId 返回数据库自动生成的 ID，例如 INSERT 到带主键的表后的 ID。
  LastInsertId() (int64, error)

  // RowsAffected 返回受查询影响的记录数。
  RowsAffected() (int64, error)
}
```

Result 是查询执行的结果。

### type Rows

```go
type Rows interface {
  // Columns 返回列的名称。根据切片的长度推断结果的列数。如果不知道某一列的名称，该条目将返回空字符串。
  Columns() []string

  // Close 关闭行迭代器。
  Close() error

  // 调用 Next 是为了将下一行数据填充到提供的片段中。提供的片段将与 Columns() 的宽度大小相同。
	//
	// 当没有更多数据行时，Next 将返回 io.EOF。
	//
	// 不应在 Next 之外写入 dest。关闭 Rows 时应注意不要修改 dest 中的缓冲区。
  Next(dest []Value) error
}
```

行是执行查询结果的迭代器。

### type RowsAffected

```go
type RowsAffected int64
```

RowsAffected 实现了 INSERT 或 UPDATE 操作的结果，这些操作会更改若干行。

#### func (RowsAffected) LastInsertId() (int64, error)

#### func (v RowsAffected) RowsAffected() (int64, error)

### type RowsColumnTypeDatabaseTypeName 添加于1.8

```go
type RowsColumnTypeDatabaseTypeName interface {
  Rows
  ColumnTypeDatabaseTypeName(index int) string
}
```

RowsColumnTypeDatabaseTypeName 可由 Rows 实现。它应返回不含长度的数据库系统类型名称。类型名称应大写。返回类型的示例"varchar", "nvarchar", "varchar2", "char", "text", "decimal", "smallint", "int", "bigint", "bool", "[]bigint", "jsonb", "xml", "timestamp"。

### RowsColumnTypeLength 添加于1.8

```go
type RowsColumnTypeLength interface {
  Rows
  ColumnTypeLength(index int) (length int64, ok bool)
}
```

RowsColumnTypeLength 可由 Rows 实现。如果列是可变长度类型，它应返回列类型的长度。如果列不是可变长度类型，则返回 false。如果长度没有系统限制，则应返回 math.MaxInt64。以下是各种类型返回值的示例：

```go
TEXT (math.MaxInt64, true)
varchar(10) (10, true)
nvarchar(10) (10, true)
decimal (0, false)
int (0, false)
bytea(30) (30, true)
```

### RowsColumnTypeNullable 添加于1.8

```go
type RowsColumnTypeNullable interface {
  Rows
  ColumnTypeNullable(index int) (nullable, ok bool)
}
```

RowsColumnTypeNullable 可由 Rows 实现。如果已知列可能为空，则 nullable 的值应为 true；如果已知列不可为空，则 nullable 的值应为 false。如果列的可空性未知，则 ok 应为 false。

### type RowsColumnTypePrecisionScale 添加于1.8

```go
type RowsColumnTypePrecisionScale interface {
  Rows
  COlumnTypePrecisionScale(index int) (precision, scale int64, ok bool)
}
```

RowsColumnTypePrecisionScale 可由 Rows 实现。它应返回十进制类型的精度和刻度。如果不适用，ok 应为 false。以下是各种类型返回值的示例：

```go
decimal(38, 4) (38, 4, true)
int (0, 0, false)
decimal (math.MaxInt64, math.MaxInt64, true)
```

RowsColumnTypeScanType 可由 Rows 实现。它应返回可用于扫描类型的值类型。例如，数据库列类型为 "bigint "时，它应返回 "reflect.TypeOf(int64(0))"。

### type RowsColumnTypeScanType 添加于1.8

```go
type RowsColumnTypeScanType interface {
  Rows
  ColumnTypeScanType(index int) reflect.TypeOf
}
```

RowsColumnTypeScanType 可由 Rows 实现。它应返回可用于扫描类型的值类型。例如，数据库列类型为 "bigint "时，它应返回 "reflect.TypeOf(int64(0))"。

### RowsNextResultSet 添加于1.8

```go
type RowsNextResultSet interface {
  Rows

  // HasNextResultSet 会在当前结果集结束时调用，并报告在当前结果集之后是否还有另一个结果集。
  HasNextResultSet() bool

  // NextResultSet 将驱动程序推进到下一个结果集中，即使当前结果集中还有剩余记录。
	//
	// 当没有更多结果集时，NextResultSet 应返回 io.EOF。
  NextResultSet() error
}
```

RowsNextResultSet 扩展了 Rows 接口，提供了一种向驱动程序发送前进到下一个结果集信号的方法。

### SessionResetter 添加于1.10

```go
type SessionResetter interface {
  // 如果连接曾被使用过，则在连接上执行查询之前会调用 ResetSession。如果驱动程序返回 ErrBadConn，连接将被丢弃。
  ResetSession(ctx context.Context) error
}
```

SessionResetter 可由 Conn 实现，以允许驱动程序重置与连接相关的会话状态，并发出不良连接信号。

### type Stmt

```go
type Stmt interface {
  // Close 关闭语句。
	//
	// 从 Go 1.1 开始，如果一个语句正在被任何查询使用，它将不会被关闭。
	//
	// 驱动程序必须确保由 Close 进行的所有网络调用不会无限期阻塞（例如，应用超时）。
  Close() error

  // NumInput 返回占位符参数的数量。
	//
	// 如果 NumInput 返回 >= 0，sql 包将检查调用者的参数数，并在调用语句的 Exec 或 Query 方法之前将错误返回给调用者。
	//
	// 如果驱动程序不知道占位符的数量，NumInput 也可能返回-1。在这种情况下，sql 包将不会对 Exec 或 Query 参数计数进行正确性检查。
  NumInput() int

  // 执行不返回记录的查询，如 INSERT 或 UPDATE。
	//
	// 过时： 驱动程序应该实现 StmtExecContext，而不是（或额外）实现 StmtExecContext。
  Exec(args []Value) (Result, error)

  // 查询执行可能返回记录的查询，如 SELECT。
	//
	// 过时： 驱动程序应该实现 StmtQueryContext，而不是（或另外）实现 StmtQueryContext。
  Query(args []Value) (Rows, error)
}
```

Stmt 是准备好的语句。它与 Conn 绑定，不会被多个程序同时使用。

### type StmtExecContext 添加于1.8

```go
type StmtExecContext interface {
  // ExecContext 执行不返回记录的查询，如 INSERT 或 UPDATE。
	//
	// ExecContext 必须遵守上下文超时并在取消时返回。
  ExecContext(ctx context.Context, args []NamedValue) (Result, error)
}
```

StmtExecContext 通过为 Exec 提供上下文来增强 Stmt 接口。

### type StmtQueryContext 添加于1.8

```go
type StmtQueryContext interface {
  // QueryContext 执行可能返回记录的查询，如 SELECT。
	//
	// QueryContext 必须遵守上下文超时，并在超时取消时返回。
  QueryContext(ctx context.Context, args []NamedValue) (Rows, error)
}
```

StmtQueryContext 通过为查询提供上下文来增强 Stmt 接口。

### type Tx

```go
type Tx interface {
  Commit() error
  Rollback() error
}
```

Tx 是事务。

### type TxOptions 添加于1.8

```go
type TxOptions struct {
  Isolation IsolationLevel
  ReadOnly bool
}
```

TxOptions 保存交易选项。

该类型应被视为与 sql.TxOptions.TxOptions 相同。

### type Validator 添加于1.15

```go
type Validator interface {
  // 将连接放入连接池之前会调用 IsValid。如果返回 false，连接将被丢弃。
  IsValid() bool
}
```

Conn 可实施验证器，以便让驱动程序发出信号，说明连接是否有效或是否应予以丢弃。

如果实现了该功能，即使连接池应丢弃该连接，驱动程序也可以从查询中返回底层错误。

### type Value

```go
type Value any
```

Value 是驱动程序必须能够处理的值。它要么是 nil，要么是数据库驱动程序的 NamedValueChecker 接口处理的类型，要么是这些类型的实例之一：

```go
int64
float64
bool
[]byte
string
time.Time
```

如果驱动程序支持游标，返回的 Value 也可以实现该包的 Rows 接口。例如，当用户选择游标（如 "select cursor(select * from my_table) from dual"）时，就会使用这个接口。如果来自选择的 Rows 被关闭，游标 Rows 也将被关闭。

### type ValueConverter

```go
type ValueConverter interface {
  // ConvertValue 将一个值转换为一个驱动程序值。
  ConvertValue(v any) (Value, error)
}
```

ValueConverter 是提供 ConvertValue 方法的接口。

驱动程序包提供了 ValueConverter 的各种实现，以便在不同驱动程序之间提供一致的转换实现。ValueConverter 有多种用途：

将 sql 包提供的 Value 类型转换为数据库表的特定列类型，并确保适合，例如确保特定的 int64 适合于表的 uint16 列。

将数据库提供的值转换为驱动程序的值类型之一。

通过 sql 软件包，在扫描中将驱动程序的值类型转换为用户类型。

### type Valuer

```go
type Valuer interface {
  // Value 返回一个驱动程序 Value。
	// Value 不能引起恐慌。
  Value() (Value, error)
}
```

Valuer 是提供 Value 方法的接口。

实现 Valuer 接口的类型可以将自身转换为驱动值。