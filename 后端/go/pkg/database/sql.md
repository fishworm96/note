## package sql

sql 包为 SQL（或类 SQL）数据库提供了一个通用接口。

sql 软件包必须与数据库驱动程序结合使用。驱动程序列表请参见 https://golang.org/s/sqldrivers。

不支持上下文取消的驱动程序在查询完成后才会返回。

有关使用示例，请参见维基页面 https://golang.org/s/sqlwiki。

## Index

### Variables

```go
var ErrConnDone = errors.New("sql: connection is already closed")
```

对已返回连接池的连接执行的任何操作都会返回 ErrConnDone。

```go
var ErrNoRows = errors.New("sql: no rows in result set")
```

当 QueryRow 没有返回记录时，Scan 会返回 ErrNoRows。在这种情况下，QueryRow 会返回一个占位符 *Row 值，将此错误推迟到 Scan。

```go
var ErrTxDone = errors.New("sql: transaction has already been committed or rolled back")
```

对已提交或回滚的事务执行的任何操作都会返回 ErrTxDone。

### func Drivers() []string 添加于1.4

Drivers 返回已注册司机名称的排序列表。

### func Register(name string, driver driver.Driver)

Register 使用提供的名称创建数据库驱动程序。如果用相同的名称调用 Register 两次，或者驱动程序为空，系统就会崩溃。

### type ColumnType 添加于1.8

```go
type ColumnType stuct {
  // 包含已筛选或未导出字段
}
```

ColumnType 包含列的名称和类型。

#### func (ci *ColumnType) DatabaseTypeName() string 添加于1.8

DatabaseTypeName 返回列类型的数据库系统名称。如果返回空字符串，则表示不支持该驱动程序类型名称。有关驱动程序数据类型的列表，请查阅驱动程序文档。不包括长度规格。常见类型名称包括 "VARCHAR"、"TEXT"、"NVARCHAR"、"DECIMAL"、"BOOL"、"INT "和 "BIGINT"。

#### func (ci *ColumnType) DecimalSize() (precision, scale int64, ok bool) 添加于1.8

DecimalSize 返回十进制类型的比例和精度。如果不适用或不支持 ok，则返回 false。

#### func (ci *ColumnType) Length() (length int64, ok bool) 添加于1.8

对于长度可变的列类型（如文本和二进制字段类型），Length 返回列类型长度。如果类型长度未限定，值将是 math.MaxInt64（任何数据库限制仍然适用）。如果列类型不是可变长度（如 int），或者驱动程序不支持 ok，则返回 false。

#### func (ci *ColumnType) Name() string 添加于1.8

Name 返回列的名称或别名。

#### func (ci *ColumnType) Nullable() (Nullable, ok bool) 添加于1.8

Nullable 报告列是否可以为空。如果驱动程序不支持该属性，ok 将为 false。

#### func (ci *ColumnType) ScanType() reflect.Type 添加于1.8

ScanType 返回适合使用 Rows.Scan 扫描的 Go 类型。如果驱动程序不支持该属性，ScanType 将返回空接口的类型。

### type Conn 添加于1.9

```go
type Conn struct {
  // 包含已筛选或未导出字段
}
```

Conn 代表单个数据库连接，而不是数据库连接池。除非特别需要持续的单一数据库连接，否则最好从 DB 运行查询。

Conn 必须调用 "关闭 "将连接返回到数据库池，并可能与正在运行的查询同时进行。

调用关闭后，连接上的所有操作都会以 ErrConnDone 失败。

#### func (c *Conn) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error) 添加于1.9

BeginTx 启动一个事务。

在事务提交或回滚之前，将使用所提供的上下文。如果取消上下文，sql 包将回滚事务。如果向 BeginTx 提供的上下文被取消，Tx.Commit 将返回错误。

提供的 TxOptions 是可选的，如果使用默认值，则可以为空。如果使用了驱动程序不支持的非默认隔离级别，将返回错误。

#### func (c *Conn) Close() error 添加于1.9

关闭会将连接返回连接池。关闭后的所有操作都将以 ErrConnDone 返回。Close 可以安全地与其他操作同时调用，并且会阻塞直到所有其他操作结束。首先取消任何已使用的上下文，然后直接调用关闭可能会有帮助。

#### func (c *Conn) ExecContext(ctx context.Context, query string, args ...any) (Result, error) 添加于1.9

ExecContext 执行查询，但不返回任何记录。args 用于查询中的任何占位参数。

#### func (c *Conn) PingContext(ctx context.Context) error 添加于1.9

PingContext 验证与数据库的连接是否仍然有效。

#### func (c *Conn) PrepareContext(ctx context.Context, query string) (*Stmt, error) 添加于1.9

PrepareContext 创建一个准备好的语句，供以后查询或执行。返回的语句可同时运行多个查询或执行。当不再需要该语句时，调用者必须调用语句的关闭方法。

所提供的上下文用于准备语句，而不是执行语句。

#### func (c *Conn) QueryContext(ctx context.COntext, query string, args ...any) (*Rows, error) 添加于1.9

QueryContext 执行返回记录的查询，通常是 SELECT。args 用于查询中的任何占位参数。

#### func (c *Conn) QueryRowContext(ctx context.Context, query string, args ...any) *Row 添加于1.9

QueryRowContext 执行预期最多返回一条记录的查询。QueryRowContext 总是返回一个非零值。在调用 Row 的 Scan 方法之前，错误将被延迟。如果查询没有选择记录，*Row's Scan 将返回 ErrNoRows。否则，*Row's Scan 会扫描第一条选中的记录，并丢弃其余记录。

#### func (c *Conn) Raw(f func(driverConn any) error) (err error) 添加于1.13

Raw 在执行 f 的过程中会暴露底层驱动程序连接。

一旦 f 返回且 err 不是 driver.ErrBadConn，Conn 将继续可用，直到 Conn.Close 被调用。

### type DB

```go
type DB struct {
  // 包含已筛选或未导出字段
}
```

DB 是一个数据库句柄，代表零个或多个底层连接池。它可以安全地供多个程序并发使用。

sql 软件包会自动创建和释放连接；它还会维护一个空闲连接池。如果数据库有每个连接状态的概念，就可以在事务 (Tx) 或连接 (Conn) 中可靠地观察到这种状态。一旦调用 DB.Begin，返回的 Tx 将绑定到单个连接。一旦调用了事务的提交或回滚，该事务的连接就会返回 DB 的空闲连接池。可以使用 SetMaxIdleConns 控制池的大小。

#### func Open(driverName, dataSourceName string) (*DB, error)

Open 打开一个数据库，该数据库由其数据库驱动程序名称和特定于驱动程序的数据源名称指定，通常至少包括一个数据库名称和连接信息。

大多数用户会通过返回 *DB 的特定于驱动程序的连接辅助函数来打开数据库。Go 标准库中不包含数据库驱动程序。有关第三方驱动程序的列表，请参见 https://golang.org/s/sqldrivers。

Open 可能只是验证其参数，而不会创建数据库连接。要验证数据源名称是否有效，请调用 Ping。

返回的数据库可安全地供多个程序并发使用，并维护自己的空闲连接池。因此，Open 函数只需调用一次。很少需要关闭数据库。

#### func OpenDB(c driver.Connector) *DB 添加于1.10

OpenDB 使用连接器打开数据库，允许驱动程序绕过基于字符串的数据源名称。

大多数用户会通过特定于驱动程序的连接辅助函数打开数据库，该函数会返回一个 *DB。Go 标准库中不包含数据库驱动程序。有关第三方驱动程序的列表，请参见 https://golang.org/s/sqldrivers。

OpenDB 可能只是验证其参数，而不会创建数据库连接。要验证数据源名称是否有效，请调用 Ping。

返回的数据库可安全地供多个程序并发使用，并维护自己的空闲连接池。因此，OpenDB 函数只需调用一次。很少需要关闭数据库。

#### func (db *DB) Begin() (*Tx, error)

开始启动一个事务。默认隔离级别取决于驱动程序。

Begin 在内部使用 context.Background；要指定上下文，请使用 BeginTx。

#### func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error) 添加于1.8

BeginTx 启动一个事务。

在事务提交或回滚之前，将使用所提供的上下文。如果取消上下文，sql 包将回滚事务。如果向 BeginTx 提供的上下文被取消，Tx.Commit 将返回错误。

提供的 TxOptions 是可选的，如果使用默认值，则可以为空。如果使用了驱动程序不支持的非默认隔离级别，将返回错误。

#### func (db *DB) Close() error

关闭会关闭数据库并阻止启动新的查询。然后，关闭会等待服务器上已开始处理的所有查询结束。

关闭数据库的情况很少见，因为数据库句柄应该是长期存在的，并在许多程序之间共享。

#### func (db *DB) Conn(ctx context.Context) (*Conn, error) 添加于1.9

Conn 通过打开一个新连接或从连接池中返回一个现有连接来返回单个连接。Conn 将阻塞，直到返回连接或取消 ctx。在同一 Conn 上运行的查询将在同一数据库会话中运行。

每个 Conn 在使用后都必须通过调用 Conn.Close 返回数据库池。

#### func (db *DB) Driver(ctx context.Context) (*Conn, error)

驱动程序返回数据库的底层驱动程序。

#### func (db *DB) Exec(query string, args ...any) (Result, error)

Exec 执行查询，不返回任何记录。args 用于查询中的任何占位参数。

Exec 内部使用 context.Background；要指定上下文，请使用 ExecContext。

#### func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (Result, error) 添加于1.8

ExecContext 执行查询，但不返回任何记录。args 用于查询中的任何占位参数。

#### func (db *DB) Ping() error 添加于1.1

Ping 验证与数据库的连接是否仍然有效，必要时建立连接。

Ping 内部使用 context.Background；要指定上下文，请使用 PingContext。

#### func (db *DB) PingContext(ctx context.Context, query string, args ...any) (Result, error) 添加于1.8

PingContext 验证与数据库的连接是否有效，必要时建立连接。

#### func (db *DB) Prepare(query string) (*Stmt, error)

准备 "会创建一条准备好的语句，供以后查询或执行。返回的语句可同时运行多个查询或执行。当不再需要该语句时，调用者必须调用语句的关闭方法。

Prepare 内部使用 context.Background；要指定上下文，请使用 PrepareContext。

#### func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error) 添加于1.8

PrepareContext 创建一个准备好的语句，供以后查询或执行。返回的语句可同时运行多个查询或执行。当不再需要该语句时，调用者必须调用语句的关闭方法。

所提供的上下文用于准备语句，而不是执行语句。

#### func (db *DB) Query(query string, args ...any) (*Rows, error)

查询执行返回记录的查询，通常是 SELECT。args 用于查询中的任何占位参数。

Query 内部使用 context.Background；要指定上下文，请使用 QueryContext。

#### func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error) 添加于1.8

QueryContext 执行返回记录的查询，通常是 SELECT。args 用于查询中的任何占位参数。

#### func (db *DB) QueryRow(query string, args ...any) *Row

QueryRow 执行预期最多返回一条记录的查询。QueryRow 总是返回一个非零值。在调用 Row 的 Scan 方法之前，错误将被延迟。如果查询没有选择记录，*Row 的 Scan 将返回 ErrNoRows。否则，*Row's Scan 会扫描第一条选中的记录，并丢弃其余记录。

QueryRow 内部使用 context.Background；要指定上下文，请使用 QueryRowContext。

#### func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *Row 添加于1.8

QueryRowContext 执行预期最多返回一条记录的查询。QueryRowContext 总是返回一个非零值。在调用 Row 的 Scan 方法之前，错误将被延迟。如果查询没有选择记录，*Row's Scan 将返回 ErrNoRows。否则，*Row's Scan 会扫描第一条选中的记录，并丢弃其余记录。

#### func (db *DB) SetConnMaxIdleTime(d time.Duration) 添加于1.15

SetConnMaxIdleTime 设置连接空闲的最长时间。

过期的连接可在重新使用前被缓慢关闭。

如果 d <= 0，则连接不会因闲置时间而关闭。

#### func (db *DB) SetConnMaxLifetime(d time.Duration) 添加于1.6

SetConnMaxLifetime 设置连接可重复使用的最长时间。

过期的连接可在重复使用前懒散地关闭。

如果 d <= 0，连接不会因连接的年龄而关闭。

#### func (db *DB) SetMaxIdleConns(n int) 添加于1.1

SetMaxIdleConns 设置空闲连接池中连接的最大数量。

如果 MaxOpenConns 大于 0 但小于新的 MaxIdleConns，那么新的 MaxIdleConns 将减少，以符合 MaxOpenConns 限制。

如果 n <= 0，则不会保留空闲连接。

目前默认的最大空闲连接数为 2，但在未来的版本中可能会有所改变。

#### func (db *DB) SetMaxOpenConns(n int) 添加于1.2

SetMaxOpenConns 设置数据库的最大打开连接数。

如果 MaxIdleConns 大于 0，而新的 MaxOpenConns 小于 MaxIdleConns，那么 MaxIdleConns 将减少，以符合新的 MaxOpenConns 限制。

如果 n <= 0，则打开的连接数没有限制。默认值为 0（无限制）。

#### func (db *DB) Stats() DBStats 添加于1.2

Stats 返回数据库统计数据。

### type DBStats 添加于1.5

```go
type DBStats struct {
  MaxOpenConnections int // 数据库的最大打开连接数。

  MaxOpenConnections int // 使用中和闲置的已建立连接数。
  InUse int // 当前使用的连接数。
  Idle int // 空闲连接数。

  WaitCount int64 // 等待的连接总数。
  WaitDuration time.Duration // 等待新连接的总时间。
  MaxIdleClosed int64 // 由于 SetMaxIdleConns 而关闭的连接总数。
  MaxIdleTimeClosed int64 // 由于 SetConnMaxIdleTime 而关闭的连接总数。
  MaxLifetimeClosed int64 // 由于 SetConnMaxLifetime 而关闭的连接总数。
}
```

DBStats 包含数据库统计信息。

### type IsolationLevel 添加于1.8

```go
type IsolationLevel int
```

IsolationLevel 是 TxOptions 中使用的事务隔离级别

```go
const (
  LevelDefault IsolationLevel = iota
  LevelReadUncommitted
  LevelReadCommitted
  LevelWriteCommitted
  LevelRepeatableRead
  LevelSnapshot
  LevelSerializable
  LevelLinearizable
)
```

BeginTx 中驱动程序可能支持的各种隔离级别。如果驱动程序不支持给定的隔离级别，则可能会返回错误信息。

请参见 https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels。

#### func (i IsolationLevel) String() string 添加于1.11

字符串返回事务隔离级别的名称。

### func NamedArg 添加于1.8

```go
type NamedArg struct {
  // Name 是参数占位符的名称。
  // 
  // 如果为空,则将使用参数列表中的序号位置
  // 
  // 名称必须省略任何符号前缀。
  Name string

  // Value 是参数的值。
  // 
  // 它可以分配与查询参数相同的值类型
  Value any
  // 包含已筛选或未导出字段
}
```

NamedArg 是一个命名参数。NamedArg 值可用作 Query 或 Exec 的参数，并与 SQL 语句中相应的命名参数绑定。

有关创建 NamedArg 值的更简洁方法，请参阅 Named 函数。

#### func Named(name string, value any) NamedArg 添加于1.8

Named 提供了一种更简洁的方法来创建 NamedArg 值。

使用示例:

```go
db.ExecContext(ctx, `
    delete from Invoice
    where
        TimeCreated < @end
        and TimeCreated >= @start;`,
    sql.Named("start", startTime),
    sql.Named("end", endTime),
)
```

### type NullBool

```go
type NullBool struct {
  Bool bool
  Valid bool // 如果 Bool 不为 NULL，则有效
}
```

NullBool 表示可能为空的 bool。NullBool 实现了 Scanner 接口，因此可以用作扫描目标，类似于 NullString。

#### func (n *NullBool) Scan(value any) error

Scan 实现了扫描仪接口。

#### func (n NullBool) Value() (driver.Value, error)

Value 实现了驱动程序 Valuer 接口。

### NullByte 添加于1.17

```go
type NullByte struct {
  Byte byte
  Valid bool // 如果 Byte 不是 NULL，则有效
}
```

NullByte 表示可能为空的字节。NullByte 实现了 Scanner 接口，因此可以用作扫描目标，类似于 NullString。

#### func (n *NUllByte) Scan(value any) error 添加于1.17

Scan 实现了扫描仪接口。

#### func (n NUllByte) Value() (driver.Value, error) 添加于1.17

Value 实现了驱动程序 Valuer 接口。

### type NullFloat64

```go
type NullFloat64 struct {
  Float64 float64
  valid bool // 如果 Float64 不为空， 则有效
}
```

NullFloat64 表示可能为空的 float64。NullFloat64 实现了 Scanner 接口，因此可以用作扫描目标，类似于 NullString。

#### func (n *NullFloat64) Scan(value any) error

Scan 实现了扫描仪接口。

#### func (n NUllFloat64) Value() (driver.Value, error)

Value 实现了驱动程序 Valuer 接口。

### NullInt16 添加于1.17

```go
type NullInt16 struct {
  Int16 Int16
  Valid bool // 如果 Int16 不是 NULL，则有效
}
```

NullInt16 表示可能为空的 int16。NullInt16 实现了 Scanner 接口，因此可以用作扫描目标，类似于 NullString。

#### func (n *NullInt16) Scan(value any) error 添加于1.17

Scan 实现了扫描仪接口。

#### func (n NullInt16) Value() (driver.Value, error) 添加于1.17

Value 实现了驱动程序 Valuer 接口。

### type NullInt32 添加于1.13

```go
type NullInt32 struct {
  NullInt32 int32
  Valid bool // 如果 Int32 不是 NULL，则有效
}
```

NullInt32 表示可能为空的 int32。NullInt32 实现了 Scanner 接口，因此可以用作扫描目标，类似于 NullString。

#### func (n *NullInt32) Scan(value any) error 添加于1.13

Scan 实现了扫描仪接口。

#### func (n NullInt32) Value() (driver.Value, error) 添加于1.13

Value 实现了驱动程序 Valuer 接口。

### type NullInt64

```go
type NullInt64 struct {
  int64 int64
  Valid bool // 如果 Int64 不是 NULL，则有效为真
}
```

NullInt64 表示可能为空的 int64。NullInt64 实现了 Scanner 接口，因此可以用作扫描目标，类似于 NullString。

#### func (n *NullInt64) Scan(value any) error

Scan 实现了扫描仪接口。

#### func (n NullInt64) Value() (driver.Value, error)

Value 实现了驱动程序 Valuer 接口。

### NullString

```go
type NullString struct {
  String string
  Valid bool // 如果字符串不是空值，则有效为真
}
```

NullString 表示可能为空的字符串。NullString 实现了扫描仪接口，因此可用作扫描目标：

```go
var s NullString
err := db.QueryRow("SELECT name FROM foo WHERE id = ?", id).Scan(&s)
...
if s.Valid {
  // 使用 s.String
} else {
  // 值为空
}
```

#### func (ns *NullString) Scan(value any) error

Scan 实现了扫描仪接口。

#### func (ns NullString) Value() (driver.Value, error)

Value 实现了驱动程序 Valuer 接口。

### NullTime 添加于1.13

```go
type NullTime struct {
  Time time.Time
  Valid bool // 如果时间不为空，则有效为真
}
```

NullTime 表示可能为空的 time.Time。NullTime 实现了 Scanner 接口，因此可以用作扫描目标，类似于 NullString。

#### func (n *NullTime) Scan(value any) error 添加于1.13

Scan 实现了扫描仪接口。

#### func (n NullTime) Value() (driver.Value, error) 添加于1.13

Value 实现了驱动程序 Valuer 接口。

### type Out 添加于1.9

```go
type Out struct {
  // Dest 是指向将被设置为存储过程的 OUTPUT 参数结果的值的指针。
  Dest any

  // In 表示参数是否为 INOUT 参数。如果是，存储过程的输入值就是 Dest 指针的取消引用值，然后用输出值替换。
  In bool
  // 包含已筛选或未导出字段
}
```

Out 可用于从存储过程中检索 OUTPUT 值参数。

并非所有驱动程序和数据库都支持 OUTPUT 值参数。

使用示例:

```go
var outArg string
_, err := db.ExecContext(ctx, "ProcName", sql.Named("Arg1", sql.Out{Dest: &outArg}))
```

### type RawBytes

RawBytes 是一个字节片，它保存着对数据库本身所拥有内存的引用。扫描到 RawBytes 后，该字节片仅在下一次调用下一步、扫描或关闭之前有效。

### type Result

```go
type Result interface {
  // LastInsertId 返回数据库响应命令时生成的整数。通常情况下，当插入新行时，该整数将来自 "自动递增 "列。并非所有数据库都支持此功能，此类语句的语法也各不相同。
  LastInsertId() (int64, error)

  // // RowsAffected 返回受更新、插入或删除影响的记录数。并非每个数据库或数据库驱动程序都支持此功能。
  RowsAffected() (int64, error)
}
```

结果汇总了已执行的 SQL 命令。

### type Row

```go
type Row struct {
  // 包含已筛选或未导出字段
}
```

Row 是调用 QueryRow 选择单行的结果。

#### func (r *Row) Err() error 添加于1.15

Err 为封装包提供了一种无需调用 Scan 即可检查查询错误的方法。Err 返回运行查询时遇到的错误（如果有）。如果错误不为零，Scan 也会返回该错误。

#### func (r *Row) Scan(dest ...any) error

Scan 将匹配行中的列复制到 dest 指向的值中。详情请查看 Rows.Scan 文档。如果有多条记录与查询匹配，Scan 会使用第一条记录并丢弃其余记录。如果没有记录与查询匹配，Scan 会返回 ErrNoRows。

### type Rows

```go
type Rows struct {
  // 包含已筛选或未导出字段
}
```

行是查询的结果。其光标从结果集的第一行开始。使用 "下一步 "从一行前进到另一行。

#### func (rs *Rows) Close() error

Close 关闭行，阻止进一步枚举。如果调用 Next 返回 false，且没有其他结果集，则行会自动关闭，只需检查 Err.Close 的结果即可。Close 是幂等的，不会影响 Err.Close 的结果。

#### func (rs *Rows) ColumnTypes() ([]*ColumnType, error) 添加于1.8

ColumnTypes 返回列信息，如列类型、长度和可归零性。某些驱动程序可能无法提供某些信息。

#### func (rs *Rows) Columns() ([]string, error)

Columns 返回列名。如果行已关闭，Columns 将返回错误信息。

#### func (rs *Rows) Err() error

Err 返回迭代过程中遇到的错误（如果有）。Err 可以在显式或隐式关闭后调用。

#### func (rs *Rows) Next() bool

Next 使用扫描方法准备读取下一条结果记录。如果成功则返回 true，如果没有下一条结果记录或在准备过程中发生错误则返回 false。要区分这两种情况，应查阅 Err。

每次调用 Scan 之前，即使是第一次调用，也必须先调用 Next。

#### func (rs *Rows) NextResultSet() bool 添加于1.8

NextResultSet 准备读取下一个结果集。它将报告是否还有下一个结果集，如果没有下一个结果集或在向下一个结果集推进时出现错误，则报告 false。要区分这两种情况，应参考 Err 方法。

调用 NextResultSet 后，应始终在扫描前调用 Next 方法。如果有其他结果集，它们的结果集中可能没有记录。

#### func (rs *Rows) Scan(dest ...any) error

Scan 将当前行中的列复制到 dest 指向的值中。dest 中的值数必须与 Rows 中的列数相同。

Scan 会将从数据库读取的列转换为以下常见 Go 类型和 sql 包提供的特殊类型：

```go
*string
*[]byte
*int, *int8, *int16, *int32, *int64
*uint, *uint8, *uint16, *uint32, *uint64
*bool
*float32, *float64
*interface{}
*RawBytes
*Rows (cursor value)
any type implementing Scanner (see Scanner docs)
```

在最简单的情况下，如果来自源列的值的类型是整数、布尔或字符串类型 T，而 dest 的类型是 *T，那么 Scan 只需通过指针赋值即可。

只要不丢失信息，Scan 还可以在字符串和数字类型之间进行转换。Scan 会将从数字数据库列扫描到的所有数字串成 *string，而扫描到数字类型时会检查是否有溢出。例如，值为 300 的 float64 或值为 "300 "的字符串可以扫描到 uint16，但不能扫描到 uint8，尽管 float64(255) 或 "255 "可以扫描到 uint8。一个例外是，将某些 float64 数字扫描为字符串时，可能会丢失字符串信息。一般情况下，将浮点列扫描为 *float64。

如果 dest 参数的类型为 *[]字节，Scan 会在该参数中保存相应数据的副本。该副本归调用者所有，可以修改并无限期保留。使用 *RawBytes 类型的参数可以避免拷贝；有关使用限制，请参阅 RawBytes 文档。

如果参数的类型为 *interface{}，Scan 会复制底层驱动程序提供的值，而不进行转换。从 []byte 类型的源值扫描至 *interface{} 时，会复制片段，结果归调用者所有。

time.Time 类型的源值可被扫描为 *time.Time、*interface{}、*string 或 *[]byte 类型的值。在转换为后两种类型时，使用 time.RFC3339Nano。

bool 类型的源值可扫描为 *bool、*interface{}、*string、*[]byte 或 *RawBytes 类型。

扫描到 *bool 时，源值可能是 true、false、1、0 或字符串输入，可通过 strconv.ParseBool 解析。

Scan 还能将查询返回的游标（如 "select cursor(select * from my_table) from dual"）转换为可扫描的 *Rows 值。如果父查询 *Rows 已关闭，则父选择查询将关闭任何游标 *Rows。

如果执行 Scanner 的第一个参数中的任何一个返回错误，该错误将被封装在返回的错误中。

### type Scanner

```go
type Scanner interface {
  // 扫描从数据库驱动程序中赋值。
  // 
  // src 值将属于以下类型之一：
	//
	//    int64
	//    float64
	//    bool
	//    []byte
	//    string
	//    time.Time
	//    nil - for NULL values
	//
  // 如果无法在不丢失信息的情况下存储该值，则应返回错误信息。
  // 
  // 诸如 []byte 的引用类型只在下一次调用 Scan 之前有效，不应保留。它们的底层内存归驱动程序所有。如果需要保留，请在下一次调用扫描之前复制它们的值。
  Scan(src any) error
}
```

扫描仪是扫描使用的界面。

### type Stmt

```go
type Stmt struct {
  // 包含已筛选或未导出字段
}
```

Stmt 是准备好的语句。Stmt 允许多个程序同时使用。

如果 Stmt 是在 Tx 或 Conn 上编写的，那么它将永远绑定到一个底层连接上。如果 Tx 或 Conn 关闭，则 Stmt 将无法使用，所有操作都将返回错误。如果 Stmt 是在 DB 上准备的，那么它将在 DB 的整个生命周期内保持可用。当需要在新的底层连接上执行该 Stmt 时，它会自动在新连接上做好准备。

#### func (s *Stmt) Close() error

Close 关闭语句。

#### func (s *Stmt) Exec(args ...any) (Result, error)

Exec 使用给定的参数执行准备好的语句，并返回一个总结语句效果的结果。

Exec 在内部使用 context.Background；要指定上下文，请使用 ExecContext。

#### func (s *Stmt) ExecContext(ctx context.Context, args ...any) (Result, error) 添加于1.8

ExecContext 使用给定的参数执行准备好的语句，并返回一个总结语句效果的结果。

#### func (s *Stmt) Query(args ...any) (*Rows, error)

查询使用给定的参数执行准备好的查询语句，并以 *Rows 的形式返回查询结果。

Query 内部使用 context.Background；要指定上下文，请使用 QueryContext。

#### func (s *Stmt) QueryContext(ctx context.Context, args ...any) (*Rows, error) 添加于1.8

QueryContext 使用给定的参数执行准备好的查询语句，并以 *Rows 的形式返回查询结果。

#### func (s *Stmt) QueryRows(args ...any) *Row

QueryRow 使用给定的参数执行准备好的查询语句。如果在执行语句过程中发生错误，将通过调用返回 *Row 的 Scan 返回错误信息，该信息始终为非零。如果查询没有选择记录，*Row 的 Scan 将返回 ErrNoRows。否则，*Row's Scan 将扫描第一条被选中的记录，并丢弃其余记录。

使用示例:

```go
var name string
err := nameByUseridStmt.QueryRow(id).Scan(&name)
```

QueryRow 内部使用 context.Background；要指定上下文，请使用 QueryRowContext。

#### func (s *Stmt) QueryRowContext(ctx context.Context, args ...any) *Row 添加于1.8

QueryRowContext 使用给定的参数执行准备好的查询语句。如果在执行语句期间发生错误，将通过调用对返回的 *Row 的 Scan 来返回该错误，而该 *Row 始终为非零。如果查询没有选择记录，*Row 的 Scan 将返回 ErrNoRows。否则，*Row's Scan 将扫描第一条被选中的记录，并丢弃其余记录。

### type Tx

```go
type Tx struct {
  // 包含已筛选或未导出字段
}
```

Tx 是正在进行的数据库事务。

事务必须以调用提交或回滚结束。

调用提交或回滚后，事务上的所有操作都会以 ErrTxDone 失败。

通过调用事务的 Prepare 或 Stmt 方法为事务准备的语句将通过调用 Commit 或 Rollback 关闭。

#### func (tx *Tx) Commit() error

Commit 提交事务。

#### func (tx *Tx) Exec(query string, args ...any) (Result, error)

Exec 执行不返回记录的查询。例如：INSERT 和 UPDATE。

Exec 在内部使用 context.Background；要指定上下文，请使用 ExecContext。

#### func (tx *Tx) ExecContext(ctx context.Context, query string, args ...any) (Result, error) 添加于1.8

ExecContext 执行不返回记录的查询。例如：INSERT 和 UPDATE。

#### func (tx *Tx) Prepare(query string) (*Stmt, error)

准备创建一份准备好的报表，供交易使用。

返回的语句在事务中运行，并在事务提交或回滚后关闭。

要在该事务中使用现有的预处理语句，请参阅 Tx.Stmt.

Prepare 在内部使用 context.Background；要指定上下文，请使用 PrepareContext。

#### func (tx *Tx) PrepareContext(ctx context.Context, query string) (*Stmt, error) 添加于1.8

PrepareContext 创建准备好的语句，供在事务中使用。

返回的语句在事务中运行，并在事务提交或回滚后关闭。

要在该事务中使用现有的准备语句，请参阅 Tx.Stmt.PrepareContext。

所提供的上下文将用于准备上下文，而不是用于执行返回语句。返回的语句将在事务上下文中运行。

#### func (tx *Tx) Query(query string, args ...any) (*Rows, error)

查询执行返回记录的查询，通常是 SELECT。

Query 内部使用 context.Background；要指定上下文，请使用 QueryContext。

#### func (tx *Tx) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error) 添加于1.8

QueryContext 执行返回记录的查询，通常是 SELECT。

#### func (tx *Tx) QueryRow(query string, args ...any) *Row

QueryRow 执行预期最多返回一条记录的查询。QueryRow 总是返回一个非零值。在调用 Row 的 Scan 方法之前，错误将被延迟。如果查询没有选择记录，*Row 的 Scan 将返回 ErrNoRows。否则，*Row's Scan 会扫描第一条选中的记录，并丢弃其余记录。

QueryRow 内部使用 context.Background；要指定上下文，请使用 QueryRowContext。

#### func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...any) *Row 添加于1.8

QueryRowContext 执行预期最多返回一条记录的查询。QueryRowContext 总是返回一个非零值。在调用 Row 的 Scan 方法之前，错误将被延迟。如果查询没有选择记录，*Row's Scan 将返回 ErrNoRows。否则，*Row's Scan 会扫描第一条选中的记录，并丢弃其余记录。

#### func (tx *Tx) Rollback() error

回滚终止事务。

#### func (tx *Tx) Stmt(stmt *Stmt) *Stmt

Stmt 从现有语句返回特定于事务的准备语句。

示例:

```go
updateMoney, err := db.Prepare("UPDATE balance SET money = money + ? WHERE id = ?")
...
tx, err := db.Begin()
...
res, err := tx.Stmt(updateMoney).Exec(123.45, 98293203)
```

返回的语句在事务中运行，并在事务提交或回滚后关闭。

Stmt 在内部使用 context.Background；要指定上下文，请使用 StmtContext。

#### func (tx *Tx) StmtContext(ctx context.Context, stmt *Stmt) *Stmt 添加于1.8

StmtContext 从现有语句返回特定于事务的准备语句。

示例:

```go
updateMoney, err := db.Prepare("UPDATE balance SET money = money + ? WHERE id = ?")
...
tx, err := db.Begin()
...
res, err := tx.StmtContext(ctx, updateMoney),Exec(123.45, 98293203)
```

所提供的上下文用于准备语句，而不是执行语句。

返回的语句在事务中运行，并将在事务提交或回滚后关闭。

### type TxOptions 添加于1.8

```go
type TxOptions struct {
  // // Isolation 是事务隔离级别。如果为零，则使用驱动程序或数据库的默认级别。
  Isolation Isolation
  ReadOnly bool
}
```

TxOptions 保存将在 DB.BeginTx 中使用的事务选项。