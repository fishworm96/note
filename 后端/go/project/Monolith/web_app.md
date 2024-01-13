## 初始化设置
### 初始化
```go
go mod init bluebell
go mod tidy
```
### main
```go
package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/routes"
	"bluebell/setting"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// Go Web开发较通用的脚手架模板

func main() {
	// 1. 加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("init setting failed, err:%v\n", err)
		return
	}
	fmt.Println(setting.Conf)
	fmt.Println(setting.Conf.LogConfig == nil)
	// 2. 初始化日志
	if err := logger.Init(setting.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	// 3. 初始化MySQL连接
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	// 4. 初始化Redis连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init InitTrans failed, err:%v\n", err)
		return
	}
	// 5. 注册路由
	r := routes.Setup()
	// 6. 启动服务（优雅关机）
	fmt.Println(setting.Conf.Port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}

```
### setting
```go
// setting/setting.go


package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port int `mapstructure:"port"`
	*LogConfig `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
	MaxSize int `mapstructure:"max_size"`
	MaxAge int `mapstructure:"max_age"`
	MaxBackups int `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host string `mapstructure:"host"`
	User string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName string `mapstructure:"dbname"`
	Port int `mapstructure:"port"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port int `mapstructure:"port"`
	DB int `mapstructure:"db"`
	PoolSize int `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() faild, err:%v\n", err)
		return
	}
	// 把读取到的配置信息反序列化到conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal faild, err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func (in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal faild, err:%v\n", err)
		}
	})
	return
}
```
### logger
```go
// logger/logger.go
package logger

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"bluebell/setting"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init(cfg *setting.LogConfig) (err error) {
	writeSyncer := getLogWriter(
		cfg.Filename,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
	)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg := zap.New(core, zap.AddCaller())
	// 替换zap库中的logger
	zap.ReplaceGlobals(lg)
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

```
### mysql
```go
// dao/mysql/mysql.go


package mysql

import (
	"fmt"

	"bluebell/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}

```
### redis
```go
// dao/redis/redis.go


package redis

import (
	"fmt"

	settings "bluebell/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
```
### router
```go
// routes/routes.go


package routes

import (
	"net/http"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}
```
### config
```yaml
name: "bluebell"
mode: "dev"
port: 8080
machine_id: 1

log:
  level: "debug"
  filename: "bluebell.log"
  max_size: 200
  max_age: 30
  max_backups: 7
mysql:
  host: "127.0.0.1"
  port: 3306
  user: "root"
  password: "root"
  dbname: "sql_demo"
  max_open_conns: 200
  max_idle_conns: 50
redis:
  host: "127.0.0.1"
  port: 6379
  password: ""
  db: 0
  pool_size: 100
```
## 创建user表
```go

CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```
## 雪花算法
### 添加雪花算法
```go
// pkg/snowflake/snowflake.go


package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}

```
### 修改配置文件
```go
// setting/setting.go

// other code
type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port int `mapstructure:"port"`
    // 添加时间
	StartTime string `mapstructure:"start_time"`
	MachineID int64 `mapstaructure:"machine_id"`
	*LogConfig `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
// other code
```
### 使用
```go
// main.go


// other code
if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err%v\n", err)
		return
	}
	defer redis.Close()
// 初始化
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
// other code
```
### config
```yaml
name: "bluebell"
mode: "dev"
port: 8080
version: "v0.1.4"
// 添加时间
start_time: "2022-07-01"
machine_id: 1
```
## 热重载
在windows平台使用air来完成热重载。
创建conf目录，将config.yaml放到里面
### 修改mian和setting
```go
// main.go

// os.Args会输出一个终端中包含内容的数组，比如运行mian.go 会输出当前目录的数组[D:xx/xx/main.go]
// 比如运行mian.go 1 2 3，os.Args会输出 [D:xx/xx/main.go, 1, 2,3]用它来读配置文件路径
	if len(os.Args) < 2 {
		fmt.Println(os.Args)
		fmt.Println("need config file.eg: bluebell config.yaml")
		return
	}
	// 1. 加载配置
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("init setting failed, err:%v\n", err)
		return
	}

// setting\setting.go

// 读配置文件路径
func Init(filePath string) (err error) {
	// 方式1：直接指定配置文件路径（相对路径或者绝对路径）
	// 相对路径：相对执行的可执行文件的相对路径
	//viper.SetConfigFile("./conf/config.yaml")
	// 绝对路径：系统中实际的文件路径
	//viper.SetConfigFile("/Users/liwenzhou/Desktop/bluebell/conf/config.yaml")

	// 方式2：指定配置文件名和配置文件的位置，viper自行查找可用的配置文件
	// 配置文件名不需要带后缀
	// 配置文件位置可配置多个
	//viper.SetConfigName("config") // 指定配置文件名（不带后缀）
	//viper.AddConfigPath(".") // 指定查找配置文件的路径（这里使用相对路径）
	//viper.AddConfigPath("./conf")      // 指定查找配置文件的路径（这里使用相对路径）

	// 基本上是配合远程配置中心使用的，告诉viper当前的数据使用什么格式去解析
	//viper.SetConfigType("json")
	viper.SetConfigFile(filePath)
```
### 安装air
[https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)
With go 1.16 or higher:
```go
go install github.com/cosmtrek/air@latest
```
### 使用
进入项目中，创建一个.air.conf或者.air.toml文件
添加配置内容
```go
# [Air](https://github.com/cosmtrek/air) TOML 格式的配置文件

# 工作目录
# 使用 . 或绝对路径，请注意 `tmp_dir` 目录必须在 `root` 目录下
root = "."
tmp_dir = "tmp"

[build]
# 只需要写你平常编译使用的shell命令。你也可以使用 `make`
cmd = "swag init && go build -o ./tmp/main.exe ."
# 由`cmd`命令得到的二进制文件名
bin = "tmp/main.exe"
# 自定义的二进制，可以添加额外的编译标识例如添加 GIN_MODE=release
# full_bin = "./tmp/main.exe"
# 监听以下文件扩展名的文件.
include_ext = ["go", "tpl", "tmpl", "html", "yaml"]
# 忽略这些文件扩展名或目录
exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
# 监听以下指定目录的文件
include_dir = []
# 排除以下文件
exclude_file = []
# 如果文件更改过于频繁，则没有必要在每次更改时都触发构建。可以设置触发构建的延迟时间
delay = 1000 # ms
# 发生构建错误时，停止运行旧的二进制文件。
stop_on_error = true
# air的日志文件名，该日志文件放置在你的`tmp_dir`中
log = "air_errors.log"
# Add additional arguments when running binary (bin/full_bin). Will run './tmp/main hello world'.
args_bin = ["./conf/config.yaml"]

[log]
# 显示日志时间
time = true

[color]
# 自定义每个部分显示的颜色。如果找不到颜色，使用原始的应用程序日志。
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# 退出时删除tmp目录
clean_on_exit = true
```
## 添加用户
### 添加路由
```go
// routes\routes.go


r.POST("/signUp", controller.SignUpHandler)
```
### 参数校验
```go
// controller\validator.go


package controller

import (
    "fmt"
    "reflect"
    "strings"
    
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/locales/en"
    "github.com/go-playground/locales/zh"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    enTranslations "github.com/go-playground/validator/v10/translations/en"
    zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义一个全局翻译器
var trans ut.Translator

func InitTrans(locale string) (err error) {
    // 修改gin框架中的validator引擎属性，实现自定制
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        
        // 注册一个获取json tag的自定义方法
        v.RegisterTagNameFunc(func(fld reflect.StructField) string {
            name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
            if name == "-" {
                return ""
            }
            return name
        })
        // 为SignUpParam注册自定义校验方法
		v.RegisterStructValidation(SignUpParamStructLevelValidation, models.ParamSignUp{})
        
        zhT := zh.New() // 中文翻译器
        enT := en.New() // 英文翻译器
        
        // 第一个参数是备用（fallback）的语言环境
        // 后面的参数是应该支持的语言环境（支持多个）
        // uni := ut.New(zhT, zhT)也是可以的
        uni := ut.New(enT, zhT, enT)
        
        // locale通常取决于http请求头的'Accept-Language'
        var ok bool
        // 也可以使用uni.FindTranslator(...)传入多个locale进行查找
        trans, ok = uni.GetTranslator(locale)
        if !ok {
            return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
        }
        
        // 注册翻译器
        switch locale {
            case "en":
            err = enTranslations.RegisterDefaultTranslations(v, trans)
            case "zh":
            err = zhTranslations.RegisterDefaultTranslations(v, trans)
            default:
            err = enTranslations.RegisterDefaultTranslations(v, trans)
        }
        return
    }
    return
}

// removeTopStruct 去除提示信息中的结构体名称
func removeTopStruct(fields map[string]string) map[string]string {
    res := map[string]string{}
    for field, err := range fields {
        res[field[strings.Index(field, ".")+1:]] = err
    }
    return res
}

// SignUpParamStructLevelValidation 自定义SignUpParam结构体校验函数
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(models.ParamSignUp)

	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}

```
### main中使用
```go
// main.go


// other code
if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init InitTrans failed, err:%v\n", err)
		return
	}
// other code
```
### 定义User和Param
```go
// models\user.go


package models

type User struct {
	UserID int64 `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

// models\params.go


package models

// 定义请求的参数结构体

type ParamSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
```
### User业务
```go
// controller\user.go


package controller

import (
	"net/http"

	"bluebell/logic"
	"bluebell/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 参数有误，直接返回错误信息
		zap.L().Error("SignUp with invalid  param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	// 3。返回相应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
```
### 连接业务和数据
```go
// logic\user.go


package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

// 存放业务逻辑的代码
func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 生成UID
	userID := snowflake.GenID()
	// 构造一个User实例
	user := &models.User{
		UserID: userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进数据库
	return mysql.InsertUser(user)
}
```
### 处理数据
```go
// dao\mysql\user.go


package mysql

import (
	"bluebell/models"

	"crypto/md5"
	"encoding/hex"
	"errors"
)

const secret = "12345"

// CheckUserExist检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 执行sql语句入库
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

// encryptPassword密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
```
## ![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1661652112367-1081ebe2-4ebb-4f77-997c-8a88e52d01f2.png#clientId=ub6a83f44-a341-4&from=paste&height=481&id=u339aa285&originHeight=481&originWidth=1136&originalType=binary&ratio=1&rotation=0&showTitle=false&size=37852&status=done&style=none&taskId=u09177e5b-6041-4abe-be40-d358a948e5b&title=&width=1136)
## 将信息从日志输出改为终端输出
```go
// logger\logger.go


// other code
// 添加mode参数
func Init(cfg *setting.LogConfig, mode string) (err error) {
	writeSyncer := getLogWriter(
		cfg.Filename,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
	)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
    // 修改输出方式
	var core zapcore.Core
	if mode == "dev" {
		// 进入开发模式，日志输出在终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	lg := zap.New(core, zap.AddCaller())
	// 替换zap库中全局的logger
	zap.ReplaceGlobals(lg)
	return
}
// other code
```
### gin也可以修改输出方式
```go
// routes\routes.go

// 传入参数
func Setup(mode string) *gin.Engine {
    // 判断模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, setting.Conf.Version)
	})

    // 路由分组
	v1 := c.Group("/api/v1")
    
	v1.POST("/signUp", controller.SignUpHandler)
	return r
}
```
### 在main中使用
```go
// main.go


// other code
// 传入mode
if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
// other code
// 传入mode
	r := routes.Setup(setting.Conf.Mode)
// oter code
```
## 登录
### 添加路由
```go
// routes\routes.go


v1.POST("/login", controller.LoginHandler)
```
### 添加登录模型
```go
// models\params.go


// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
```
### 登录处理函数
```go
// controller\user.go


func LoginHandler(c *gin.Context) {
	// 获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回相应
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs , ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误
		})
		return
	}
	// 业务逻辑处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}
	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
```
### 连接业务和数据
```go
// logic\user.go


func Login(p *models.ParamLogin) error {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
```
### 处理数据
```go
// dao\mysql\user.go


func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return errors.New("密码错误")
	}
	return
}
```
## ![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1661652154165-535f18e1-577d-4844-9a8e-51a9eba6d7b9.png#clientId=ub6a83f44-a341-4&from=paste&height=376&id=u1f3756d5&originHeight=376&originWidth=1107&originalType=binary&ratio=1&rotation=0&showTitle=false&size=28398&status=done&style=none&taskId=ufbe02aa3-35e3-4b5f-a6d7-de60ae081ac&title=&width=1107)
## 定义错误码并封装响应方法
### 定义错误码
```go
// controller\code.go


package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserNotExist
	CodeUserExist
	CodeInvalidPassword
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "success",
	CodeInvalidParam: "请求参数错误",
	CodeUserExist: "用户名已存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy: "服务繁忙",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
```
### 定义错误响应
```go
// controller\response.go

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// 默认错误响应
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg: code.Msg(),
		Data: nil,
	})
}

// 自定义错误响应
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg: msg,
		Data: nil,
	})
}

// 成功响应
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg: CodeSuccess.Msg(),
		Data: data,
	})
}
```
### 自定义错误信息
```go
// dao\mysql\user.go


// 自定义错误信息
var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

// CheckUserExist检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
        // 使用自定义错误信息
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 执行sql语句入库
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

// encryptPassword密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
        // 使用自定义错误信息
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
        // 使用自定义错误信息
		return ErrorInvalidPassword
	}
	return
}

```
### 使用错误响应
将使用c.JSON返回的数据都改为定义的响应。
```go
// controller\user.go


package controller

import (
	"errors"

	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 其你去参数有误，直接返回相应
		zap.L().Error("SignUp with invalid  param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
	ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
		}
		return
	}
	// 3。返回相应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回相应
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs , ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 业务逻辑处理
    if err != nil {
	zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
	if errors.Is(err, mysql.ErrorUserNotExist) {
		ResponseError(c, CodeUserNotExist)
		return
	}
	ResponseError(c, CodeInvalidPassword)
	return
}
	// 返回响应
	ResponseSuccess(c, nil)
}
```
## 使用jwt
```go
go get github.com/dgrijalva/jwt-go
```
### 生成token
```go
// pkg\jwt\jwt.go


package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("小秘密")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		userID,
		"username", // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer: "bluebell", // 签发人
		},
	}
	// 使用指定的签名方式创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil	
	})
	if err != nil {
		return nil,	err
	}
	if token.Valid {
		// 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
```
### 获取生成token
```go
// logic\user.go

// other code
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return "", err
	}
    // 返回token
	return jwt.GenToken(user.UserID, user.Username)
}
```
### 返回前端token
```go
// controller\user.go


func LoginHandler(c *gin.Context) {
	// 获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回相应
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs , ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserNotExist)
		}
		return
	}
	// 返回响应
	ResponseSuccess(c, token)
}
```
## 封装jwt校验
### 请求处理
```go
// controller\request.go


package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUser 获取当前登录的用户ID
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
```
### 定义错误码
```go
// controller\code.go

// other code
const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserNotExist
	CodeUserExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "success",
	CodeInvalidParam: "请求参数错误",
	CodeUserExist: "用户名已存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy: "服务繁忙",
	CodeNeedLogin: "需要登录",
	CodeInvalidToken: "无效的token",
}
// other code
```
### 封装jwt校验
```go
// middlewares\auth.go


package middlewares

import (
	"bluebell/controller"
	"bluebell/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// parts[1]是获取到tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文上
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next() // 后续的处理请求的函数中，可以用过c.Get(CtxUserIDKey) 来获取当前请求的用户信息
	}

}

```
```go
// 测试接口


r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 如果是登录的用户,判断请求头中是否有 有效的JWT  ？
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
```
## 修改token时间
添加token时间字段到config.yaml
```yaml
name: "web_app"
mode: "dev"
port: 8080
start_time: "2022-07-01"
machine_id: 1

auth:
jwt_expire: 8760
log:
level: "info"
filename: "bluebell.log"
max_size: 200
max_age: 30
max_backups: 7
mysql:
host: "127.0.0.1"
port: 3306
user: "root"
password: "root"
dbname: "sql_demo"
max_open_conns: 200
max_idle_conns: 50
redis:
host: "127.0.0.1"
port: 6379
password: ""
db: 0
pool_size: 100
```
## 获取community
### sql
```sql
DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `community_id` int(10) unsigned NOT NULL,
  `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_community_id` (`community_id`),
  UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


INSERT INTO `community` VALUES ('1', '1', 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10');
INSERT INTO `community` VALUES ('2', '2', 'leetcode', '刷题刷题刷题', '2020-01-01 08:00:00', '2020-01-01 08:00:00');
INSERT INTO `community` VALUES ('3', '3', 'CS:GO', 'Rush B。。。', '2018-08-07 08:30:00', '2018-08-07 08:30:00');
INSERT INTO `community` VALUES ('4', '4', 'LOL', '欢迎来到英雄联盟!', '2016-01-01 08:00:00', '2016-01-01 08:00:00');
```
### 定义模型
```go
// models\community.go


package models

type Community struct {
    Id int64 `json:"id" db:"community_id"`
    Name string `json:"name" db:"community_name"`
}
```
### 定义路由
```go
// routes\routes.go


v1.GET("/community", controller.CommunityHandler)
```
### community控制器
```go
// controller\community.go


package controller

import (
	"bluebell/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
```
### community数据
```go
// logic\community.go


package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() (community []*models.Community, err error) {
	return mysql.GetCommunityList()
}
```
### community数据库
```go
// dao\mysql\community.go


package mysql

import (
	"bluebell/models"
	"database/sql"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id, community_name from community`
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("there is no community in db")
			err = nil
		}
	}
	return
}
```
## ![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1661652180568-95b9735b-62b8-4417-a4f4-8127bb1a111a.png#clientId=ub6a83f44-a341-4&from=paste&height=418&id=n7pRs&originHeight=418&originWidth=1105&originalType=binary&ratio=1&rotation=0&showTitle=false&size=26577&status=done&style=none&taskId=u7e74d8bb-2a06-4147-b65e-42bde827cba&title=&width=1105)
## 根据id获取社区信息
### 添加路由
```go
// routes\routes.go


v1.GET("/community/:id", controller.CommunityDetailHandler)
```
### 定义模型
```go
// models\community.go

type CommunityDetail struct {
	Community
	Introduction string `json:"introduction,omitempty" db:"introduction"` // omitempty表示不返回为空
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
```
### 控制器
```go
// controller\community.go


// 根据id获取社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	idStr := c.Param("id") // 获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 根据id获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() faild", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
```
### 处理数据
```go
// logic\community.go


func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
```
### 抽离错误
将user中定义的错误抽离出来
```go
// dao\mysql\error_code.go


package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
)

```
### 查询数据库
```go
// dao\mysql\community.go


// 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail,err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time from community where community_id = ?`
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}
```
## ![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1661652222249-78e438bf-984c-4e5a-a00b-14475ef3a443.png#clientId=ub6a83f44-a341-4&from=paste&height=530&id=uf6f1caff&originHeight=530&originWidth=1125&originalType=binary&ratio=1&rotation=0&showTitle=false&size=40546&status=done&style=none&taskId=ud84969fd-a3c3-4b75-aab4-357fb3b3c93&title=&width=1125)
## 创建帖子
### sql
```plsql
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_id` bigint(20) NOT NULL COMMENT '帖子id',
  `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `author_id` bigint(20) NOT NULL COMMENT '作者的用户id',
  `community_id` bigint(20) NOT NULL COMMENT '所属社区',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '帖子状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_post_id` (`post_id`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```
### 定义模型
```go
// models\post.go


package models

import "time"

// 内存对齐

type Post struct {
	ID int64	`json:"id" db:"post_id"`
	AuthorID int64 `json:"author_id" db:"author_id"`
	CommunityID int64	`json:"community_id" db:"community_id" binding:"required"`
	Status int32	`json:"status" db:"status"`
	Title string	`json:"title" db:"title" binding:"required"`
	Content string	`json:"content" db:"content" binding:"required"`
	CreateTime	time.Time `json:"create_time" db:"create_time"`
}
```
### 路由
```go
// routes\routes.go


	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.POST("/post", controller.CreatePostHandler)
	}
```
### 控制器
```go
// controller\post.go


package controller

import (
	"bluebell/logic"
	"bluebell/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	// 获取参数及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) err", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 从c取到当前发送请求的用户id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 响应返回
	ResponseSuccess(c, nil)
}
```
### 处理数据
```go
// logic\post.go


package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) error {
	// 生成post id
	p.ID = snowflake.GenID()
	// 保存到数据库
	return mysql.CreatePost(p)
	// 返回
}
```
### 插入数据
```go
// dao\mysql\post.go


package mysql

import "bluebell/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}
```
## ![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1661652301896-970b454e-6532-442e-b216-1a3cb0e54858.png#clientId=ub6a83f44-a341-4&from=paste&height=498&id=u3137dd6f&originHeight=498&originWidth=1168&originalType=binary&ratio=1&rotation=0&showTitle=false&size=40668&status=done&style=none&taskId=ub0d74208-e11f-4103-87da-b21914b4fbe&title=&width=1168)
## 通过id查询帖子详细信息
### 路由
```go
// routes\routes.go


v1.GET("/post/:id", controller.GetPostDetailHandler)
```
### 模型
```go
// models\post.go


// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName string `json:""author_name`
	*Post // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
}
```
### 控制器
```go
// controller\post.go


// 根据id获取帖子信息
func GetPostDetailHandler(c *gin.Context) {
	// 获取参数（从url中获取帖子的id）
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error(" get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 根据id去除帖子数据（差数据库）
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) falied", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回相应
	ResponseSuccess(c, data)
}
```
### 处理数据
```go
// logic\post.go


// GetPostById 根据帖子id查询帖子详情数据
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合我们接口想要的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid)",zap.Int64("pid", pid), zap.Error(err))
		return
	}
	// 根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}
	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
		return
	}
	// 接口数据拼接
	data = &models.ApiPostDetail{
		AuthorName: user.Username,
		Post: post,
		CommunityDetail: community,
	}
	return
}
```
### 查询数据库
```go
// dao\mysql\post.go


// GetPostyById 通过id查询帖子信息
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}


// dao\mysql\user.go

// GetUserById 通过id获取作者信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr :=  `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}


// dao\mysql\community.go


// 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail,err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time from community where community_id = ?`
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}
```
## ![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1661652314763-ea306b07-78b9-4c87-b5a2-07c4bd9c6058.png#clientId=ub6a83f44-a341-4&from=paste&height=633&id=u6a0836fb&originHeight=633&originWidth=1138&originalType=binary&ratio=1&rotation=0&showTitle=false&size=51525&status=done&style=none&taskId=u3eace673-4c6e-45db-be67-9458a1c0659&title=&width=1138)
## 获取帖子列表与分页
### 路由
```go
// routes\routes.go


v1.GET("/posts/", controller.GetPostListHandler)
```
### 控制器
```go
// controller\post.go


// GetPostListHandler 获取帖子列表的处理函数
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList(page, size) falid", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回相应
	ResponseSuccess(c, data)
}
```
### 分页
```go
// controller\request.go


func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var (
		page int64
		size int64
		err error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	if page == 0{
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
```
### 处理数据
```go
// logic\post.go


// GetPostList 获取帖子列表
func GetPostList(page int64, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) faild", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) faild", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName: user.Username,
			Post: post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

```
### 查询数据库
```go
// dao\mysql\post.go


func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post order by create_time desc limit ?, ?`
	posts = make([]*models.Post, 0, 2) // 不要写成make([]*models.Post, 2)
	err = db.Select(&posts, sqlStr, (page - 1) * size, size)
	return
}
```
## ![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1661652323978-352caaf6-0b95-4e06-870f-527ebfb7c737.png#clientId=ub6a83f44-a341-4&from=paste&height=591&id=uf9706244&originHeight=591&originWidth=1121&originalType=binary&ratio=1&rotation=0&showTitle=false&size=44675&status=done&style=none&taskId=u1bca39f7-d504-4938-8872-05ab60bd8fc&title=&width=1121)
## 帖子点赞
### 定义路由
```go
v1.POST("/vote", controller.PostVoteController)
```
### 定义参数
```go
// models\params.go

// other code
// 需要帖子id和点赞 两个参数
type ParamVoteData struct {
	// UserID 从请求中获取当前的用户
	PostID string `json:"post_id" binding:"required"` // 帖子id
    // direction不用设置为必须值，如果设置为必须值穿0值时会valiadtad会解析成没有参数，所以会报错。
	Direction int8 `json:"direction,string" binding:"oneof=1 0 -1"` // 增长票(1)还是反对票（-1）取消票0
}
```
### 控制器
```go
// controller\vote.go


package controller

import (
	"bluebell/logic"
	"bluebell/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 投票

//type VoteData struct {
//	// UserID 从请求中获取当前的用户
//	PostID    int64 `json:"post_id,string"`   // 贴子id
//	Direction int   `json:"direction,string"` // 赞成票(1)还是反对票(-1)
//}

func PostVoteController(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) // 翻译并去除掉错误提示中的结构体标识
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	// 获取当前请求用户的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 具体投票的业务逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost(userID, p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}
```
### 数据处理
```go
// logic\vote.go


package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"
)

// 推荐阅读
// 基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm/

// 本项目使用简化版的投票分数
// 投一票就加432分   86400/200  --> 200张赞成票可以给你的帖子续一天

/* 投票的几种情况：
direction=1时，有两种情况：
	1. 之前没有投过票，现在投赞成票    --> 更新分数和投票记录
	2. 之前投反对票，现在改投赞成票    --> 更新分数和投票记录
direction=0时，有两种情况：
	1. 之前投过赞成票，现在要取消投票  --> 更新分数和投票记录
	2. 之前投过反对票，现在要取消投票  --> 更新分数和投票记录
direction=-1时，有两种情况：
	1. 之前没有投过票，现在投反对票    --> 更新分数和投票记录
	2. 之前投赞成票，现在改投反对票    --> 更新分数和投票记录

投票的限制：
每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了。
	1. 到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2. 到期之后删除那个 KeyPostVotedZSetPF
*/

func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}

```
### 定义redis的key
```go
// dao\redis\keys.go


package redis

// redis key

// redis key注意使用命名空间的方式,方便查询和拆分

const (
	Prefix = "bluebell:" // 项目key前缀
	KeyPostTimeZSet = "post:time" // zset;帖子及发帖时间
	KeyPostScoreZSet = "post:score" // zset;帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" // zset;记录用户及投票类型;参数是post id

	KeyCommunitySetPF = "community:" // set;保存每个分区下帖子的id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
```
### 修改redis的全局变量
```go
// dao\redis\redis.go


var (
	client *redis.Client
)
// other code

```
### redis的投票
```go
// dao\redis\vote.go


package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

// 推荐阅读
// 基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm/

// 本项目使用简化版的投票分数
// 投一票就加432分   86400/200  --> 200张赞成票可以给你的帖子续一天

/* 投票的几种情况：
   direction=1时，有两种情况：
   	1. 之前没有投过票，现在投赞成票    --> 更新分数和投票记录  差值的绝对值：1  +432
   	2. 之前投反对票，现在改投赞成票    --> 更新分数和投票记录  差值的绝对值：2  +432*2
   direction=0时，有两种情况：
   	1. 之前投过反对票，现在要取消投票  --> 更新分数和投票记录  差值的绝对值：1  +432
	2. 之前投过赞成票，现在要取消投票  --> 更新分数和投票记录  差值的绝对值：1  -432
   direction=-1时，有两种情况：
   	1. 之前没有投过票，现在投反对票    --> 更新分数和投票记录  差值的绝对值：1  -432
   	2. 之前投赞成票，现在改投反对票    --> 更新分数和投票记录  差值的绝对值：2  -432*2

   投票的限制：
   每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了。
   	1. 到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
   	2. 到期之后删除那个 KeyPostVotedZSetPF
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote = 432 // 每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
    ErrVoteRepeated = errors.New("不允许重复投票")
)

func CreatePost(postID, communityID int64) error {
	pipeline := client.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score: float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score: float64(time.Now().Unix()),
		Member: postID,
	})
    // 更新：把帖子id加到社区的set
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 判断投票权限
	// 去redis取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix()) - postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2和3需要放到一个pipeline事务中操作

	// 更新帖子的分数
	// 先去查当前帖子的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
    // 判断是否重复投票
    if value == ov {
        return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op * diff * scorePerVote, postID)

	// 记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score: value, // 赞成票还是反对票
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
```
### 在post中创建redis的key
```go
// logic\post.go


func CreatePost(p *models.Post) (err error) {
	// 生成post id
	p.ID = snowflake.GenID()
	// 保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	return
}
// other code
```
### 修改返回值结构体
```go
// controller\response.go

// 定义请求的参数结构体
const (
	OrderTime = "time"
	OrderScore = "score"
)

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
    // 修改返回值，为空时不返回值。
	Data interface{} `json:"data,omitempty"`
}
// other code
```
## ![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1661652338618-b244bab0-aeb1-47f9-9fbc-8f5618699173.png#clientId=ub6a83f44-a341-4&from=paste&height=513&id=u8af82f8d&originHeight=513&originWidth=1133&originalType=binary&ratio=1&rotation=0&showTitle=false&size=40715&status=done&style=none&taskId=u98777a99-907e-4590-be21-f72e209ccfd&title=&width=1133)
## 根据投票时间或投票排行来获取提诶子列表与分页。
### 定义路由
```go
// routes\routes.go


v1.GET("/posts2", controller.GetPostListHandler2)
```
### 定义模型
```go
// models\params.go


// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64 `json:"community_id" form:"community_id"` // 可以为空
	Page int64 `json:"page" form:"page" example:"1"` // 页码
	Size int64`json:"size" form:"size" example:"10"` // 每页数据量
	Order string `json:"order" form:"order" example:"score"` // 排序依据
}
```
### 控制器
```go
// controller\post.go

// GetPostListHandler2 升级版帖子列表接口
// 根据前端传来的参数动态的获取帖子列表
// 按创建时间排序 或者 按照 分数排序
// 1.获取请求的query string参数
// 2.去redis查询id列表
// 3.根据id去数据库查询帖子详细信息
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数(query string): /api/v1/posts2?page=1&size=10&order=tiem
		// 定义默认参数，如果没有传参就用默认参数
    p := &models.ParamPostList{
		Page: 1,
		Size: 10,
		Order: models.OrderTime, // magic string
	}
	//c.ShouldBind()  根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostList2(p)
	// 获取数据
	if err != nil {
		zap.L().Error("logic.GetPostListNew(p) faild", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回相应
	ResponseSuccess(c, data)
}
```
### 修改模型
```go
// models\post.go


// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	VoteNum int64 `json:"vote_num"`
	*Post // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
}
```
### 数据处理
```go
// logic\post.go


// 查询所有
func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return  0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	// 根据id去mysql数据库查询帖子详细信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("posts", posts))
	// 提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) faild",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根绝社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) faild",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
				continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName: user.Username,
			VoteNum: voteData[idx],
			Post: post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
```
### redis中查询
```go
// dao\redis\post.go

package redis

import (
	"bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// 为了后面扩展，这里抽离出Key的启示位置和末尾位置函数
func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	// 3. ZREVRANGE 按分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取id
	// 根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 确定查询的索引起始点
	return getIDsFormKey(key, p.Page, p.Size)
}

// GetPostVoteData 根据ids查询每篇帖子的投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
		//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPF + id)
	//	// 查找key中分数是1的元素的数量->统计每篇帖子的赞成票的数量
	//	v := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}
	// 使用pipeline一次发送多条命令，减少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
```
### mysql查询
```go
// dao\mysql\post.go


func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id in (?) order by FIND_IN_SET(post_id, ?)`

	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
```
## ![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1661652358502-b63c6b64-b2e1-4315-9055-4e70770885a2.png#clientId=ub6a83f44-a341-4&from=paste&height=400&id=u331dc62d&originHeight=400&originWidth=1128&originalType=binary&ratio=1&rotation=0&showTitle=false&size=28557&status=done&style=none&taskId=u4987b844-4d12-40e0-a416-21edc1d146b&title=&width=1128)
## 添加通过社区id来查询帖子列表分页
### 修改一下controller
```go
// controller\post.go

将controller中的data, err := logic.GetPostList2(p)修改为data, err := logic.GetPostListNew(p) // 更新：合二为一
```
### 数据处理层
将逻辑拆成2部分
根据order类型查询不变，添加了根据社区id查询和判断是否有社区id
```go
// logic\post.go

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return  0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	// 根据id去mysql数据库查询帖子详细信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("posts", posts))
	// 提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) faild",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) faild",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
				continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName: user.Username,
			VoteNum: voteData[idx],
			Post: post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetCommunityPostList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 去redis查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetCommunityPostIDsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("redis.GetCommunityPostIDsInOrder(p)", zap.Any("ids", ids))
	// 根据id去mysql数据库查询帖子详细信息
	// 返回的数据还要按照我给定的id的顺序返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	zap.L().Debug("GetCommunityPostList", zap.Any("posts", posts))
	// 提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
			 zap.Int64("author_id", post.AuthorID),
			 zap.Error(err))
			continue
	}
	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", 
		zap.Int64("community_id", post.CommunityID),
		zap.Error(err))
		continue
	}
	postDetail := &models.ApiPostDetail{
		AuthorName: user.Username,
		VoteNum: voteData[idx],
		Post: post,
		CommunityDetail: community,
	}
	data = append(data, postDetail)
	}
	return
}

// GetPostListNew
func GetPostListNew(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 根据请求参数的不同，执行不同的逻辑
	if p.CommunityID == 0 {
		// 查所有
		data, err = GetPostList2(p)
	} else {
		// 根据社区id查询
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetPostListNew faild", zap.Error(err))
		return nil, err
	}
	return
}

```
### redis中查询
```go
// dao\redis\post.go


func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	// 使用zinterstore 把区分的帖子set与帖子分数的 zset 生成一个新的 zset
	// 针对新的zset 及之前得逻辑取数据

	// 社区的key
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))

	// 利用缓存key减少zinterstore执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(key).Val() < 1 {
		// 不存在，需要计算
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey) // zinterstore 计算
		pipeline.Expire(key, 60 * time.Second) // 设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	// 存在的话就直接根据key查询ids
	return getIDsFormKey(key, p.Page, int64(p.Size))
}
```
## 项目中使用swagger
[swagger地址](https://github.com/swaggo/gin-swagger)
### 安装gin-swagger
```go
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/files
go get -u github.com/swaggo/gin-swagger
go install github.com/swaggo/swag/cmd/swag

// 生成文档
swag init
```
在main中添加注释
```go
// main.go


// @title bluebell项目接口文档
// @version 1.0
// @description Go web开发进阶项目实战bluebell

// @contact.name zhou
// @contact.url https://www.fishworm96.github.io

// @host localhost:8080
// @BasePath /api/v1
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1660464759242-eede9555-c0dc-4fc5-97b4-25ce5cf79ee6.png#clientId=uddb37cb8-d65e-4&from=paste&height=218&id=ubcdc5273&originHeight=273&originWidth=554&originalType=binary&ratio=1&rotation=0&showTitle=false&size=16623&status=done&style=none&taskId=u35e4e204-8859-4b5a-9a14-f3102c159b0&title=&width=443.2)
在需要的接口前面添加注释，一般为controller层
```go
// controller\post.go


// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
```
### ![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1660464777694-a1ea67fd-af7f-4161-8ebe-e88b10f0e28e.png#clientId=uddb37cb8-d65e-4&from=paste&height=711&id=u06531ebe&originHeight=889&originWidth=1834&originalType=binary&ratio=1&rotation=0&showTitle=false&size=76601&status=done&style=none&taskId=u1838993a-f66d-4812-afcd-9edb5587305&title=&width=1467.2)
![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1660464799966-cea0fbcc-d03c-4a58-96a1-a7b126f7c8a3.png#clientId=uddb37cb8-d65e-4&from=paste&height=498&id=u48ff5bbc&originHeight=623&originWidth=1770&originalType=binary&ratio=1&rotation=0&showTitle=false&size=44276&status=done&style=none&taskId=u613c40df-a9c1-4955-928a-7a970994ba1&title=&width=1416)
### 定义文档模型
```go
// controller\doc_response_models.go


package controller

import "bluebell/models"

// 专门用来放接口文档用到的model
// 因为我们的接口文档返回的数据格式是一致的，但是具体的data类型不一致

// _ResponsePostList 帖子列表接口响应数据
type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    []*models.ApiPostDetail `json:"data"`    // 数据
}

```
### 添加注释和例子
```go
// models\params.go


// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64 `json:"community_id" form:"community_id"` // 可以为空
	Page int64 `json:"page" form:"page" example:"1"` // 页码
	Size int64`json:"size" form:"size" example:"10"` // 每页数据量
	Order string `json:"order" form:"order" example:"score"` // 排序依据
}
```
### 添加为文档的路由
```go
// routes\routes.go

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"time"

	_ "bluebell/docs" // 千万不要忘了导入把你上一步生成的docs

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
```
```go
使用 swag init生成文档
```
使用[http://localhost:8080/swagger/index.htm](http://localhost:8080/swagger/index.htm)访问
## 单元测试
### 文章controller层单元测试
```go
// controller\post_test.go


package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePostHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)

	body := `{
		"community_id": 1,
		"title": "test",
		"content": "just a test"
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// 判断响应的内容是不是按预期返回了需要登录的错误

	// 方法1：判断响应内容中是不是包含指定的字符串
	//assert.Contains(t, w.Body.String(), "需要登录")

	// 方法2：将响应的内容反序列化到ResponseData 然后判断字段与预期是否一致
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}

```
### 文章dao层mysql单元测试
```go
// dao\mysql\post_test.go


package mysql

import (
	"bluebell/models"
	"bluebell/setting"
	"testing"
)
单元测试中不存在连接数据库，所以这里需要单独连接数据库。
func init() {
	dbCfg := setting.MySQLConfig{
		Host: "127.0.0.1",
		User: "root",
		Password: "root",
		DbName: "sql_demo",
		Port: 3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePostHandler(t *testing.T) {
	post := models.Post{
		ID: 10,
		AuthorID: 123,
		CommunityID: 1,
		Title: "test",
		Content: "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err:%\v\n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
```
## 压力测试
### 安装
```go
go get github.com/adeven/go-wrk
go install github.com/adeven/go-wrk
```
### 使用方法
```go
go-wrk [flags] url
```
### 参数
```go
-H="User-Agent: go-wrk 0.1 bechmark\nContent-Type: text/html;": 由'\n'分隔的请求头
-c=100: 使用的最大连接数
-k=true: 是否禁用keep-alives
-i=false: if TLS security checks are disabled
-m="GET": HTTP请求方法
-n=1000: 请求总数
-t=1: 使用的线程数
-b="" HTTP请求体
-s="" 如果指定，它将计算响应中包含搜索到的字符串s的频率
```
### 例子
```go
go-wrk -t=8 -c=100 -n=10000 "http://127.0.0.1:8080/api/v1/posts?size=10"
```
### 输出结果
```go

==========================BENCHMARK==========================
URL:                            http://127.0.0.1:8080/api/v1/posts?size=10

Used Connections:               100
Used Threads:                   8
Total number of calls:          10000

===========================TIMINGS===========================
Total time passed:              2.74s
Avg time per request:           27.11ms
Requests per second:            3644.53
Median time per request:        26.88ms
99th percentile time:           39.16ms
Slowest time for request:       45.00ms

=============================DATA=============================
Total response body sizes:              340000
Avg response body per request:          34.00 Byte
Transfer rate per second:               123914.11 Byte/s (0.12 MByte/s)
==========================RESPONSES==========================
20X Responses:          10000   (100.00%)
30X Responses:          0       (0.00%)
40X Responses:          0       (0.00%)
50X Responses:          0       (0.00%)
Errors:                 0       (0.00%)
```
## 配置跨域
```go
// middlewares\cors.go


package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, x-token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		c.Header("Access-Control-Allow-Header", "Content-Length, Access-Control-ALlow-Origin, Access-Control-Allow-Header, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
```
### 在路由的中间件中使用
```go
// routes\routes.go


r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.Cors())
```
## 配置令牌桶
为什么要配置令牌桶？为了服务器性能达到上限时不至于宕机，所以做一些限流策略。
### 配置令牌桶
```go
// middlewares\ratelimit.go


package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就返回响应
		// if bucket.Take(1) > 0
        // TakeAvailable为放出令牌的数量
		if bucket.TakeAvailable(1) != 1 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		// 取到令牌后放行
		c.Next()
	}
}
```
### 在路由的中间件中使用
```go
// routes\routes.go


r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.Cors(), middlewares.RateLimitMiddleware(1 * time.Second, 1))
```
## 使用docker
[https://www.liwenzhou.com/posts/Go/how_to_deploy_go_app_using_docker/](https://www.liwenzhou.com/posts/Go/how_to_deploy_go_app_using_docker/)
### 配置docker文件
```go
// Dockerfile 


FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 bluebell_app
RUN go build -o bluebell_app .

###################
# 接下来创建一个小镜像
###################
FROM debian:stretch-slim

# COPY ./wait-for.sh /
# COPY ./templates /templates
# COPY ./static /static
COPY ./conf /conf

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/bluebell_app /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

# 声明服务端口
EXPOSE 8080

# 需要运行的命令
ENTRYPOINT ["/bluebell_app", "conf/config.yaml"]
```
```go
// 构建镜像
docker build . -t bubble_app
// 运行
docker run -p 8888:8888 goweb_app
```
#### 关联容器
又因为我们的项目中使用了MySQL，我们可以选择使用如下命令启动一个MySQL容器，它的别名为mysql8029；root用户的密码为root1234；挂载容器中的/var/lib/mysql到本地的/Users/q1mi/docker/mysql目录；内部服务端口为3306，映射到外部的13306端口。
### 修改config
```go
name: "web_app"
mode: "dev"
port: 8080
start_time: "2022-07-01"
machine_id: 1

auth:
  jwt_expire: 8760
log:
  level: "info"
  filename: "bluebell.log"
  max_size: 200
  max_age: 30
  max_backups: 7
mysql:
  host: mysql:8029
  port: 3306
  user: "root"
  password: "root"
  dbname: "sql_demo"
  max_open_conns: 200
  max_idle_conns: 50
redis:
  host: redis:5014
  port: 6379
  password: ""
  db: 0
  pool_size: 100
```
```go
docker run --name mysql8019 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root1234 -v /Users/q1mi/docker/mysql:/var/lib/mysql -d mysql:8.0.19
// 构建
docker build . -t bubble_app
// 运行关联的容器
docker run --link=mysql8019:mysql8019 -p 8888:8888 bubble_app
```
### 结合docker compose
```go
// Dockerfile 


FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 bluebell_app
RUN go build -o bluebell_app .

###################
# 接下来创建一个小镜像
###################
FROM debian:stretch-slim

# COPY ./wait-for.sh /
# COPY ./templates /templates
# COPY ./static /static
COPY ./conf /conf

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/bluebell_app /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

# 声明服务端口
EXPOSE 8080

# 需要运行的命令
// ENTRYPOINT ["/bluebell_app", "conf/config.yaml"]
```

```go
// docker-compose.yml


# yaml 配置
version: "3.7"
services:
  mysql8029:
    image: "mysql:8.0.29"
    ports:
      - "33061:3306"
    command: "--default-authentication-plugin=mysql_native_password --init-file /data/application/init.sql"
    environment:
      MYSQL_ROOT_PASSWORD: "root"
      MYSQL_DATABASE: "bluebell"
      MYSQL_PASSWORD: "root"
    volumes:
      - ./init.sql:/data/application/init.sql
  redis5014:
    image: "redis:5.0.14"
    ports:
      - "26379:6379"
  bluebell_app:
    build: .
    command: sh -c "./wait-for.sh mysql8029:3306 redis5014:6379 -- ./bluebell_app ./conf/config.yaml"
    depends_on:
      - mysql8029
      - redis5014
    ports:
      - "8888:8080"

```
使用docker-compose up运行
