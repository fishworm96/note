# 用户服务
## proto生成go
```go
// 在当前目录下解析user.proto生成go文件
protoc --go_out=:. user.proto
// 在当前目录下解析user.proto生成grpc的go文件
protoc --go-grpc_out=:. user.proto
```
## 添加viper和zap
### 添加配置
```go
// config/config.go

package config

type MysqlConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int `mapstructure:"port" json:"port"`
	Name string `mapstructure:"db" json:"db"`
	User string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type ServerConfig struct {
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
}
```
### global中添加ServerConfig变量
```go
// global/global.go


// other code
var (
	DB *gorm.DB
	ServerConfig config.ServerConfig
)

// other code
```
### 配置初始化
```go
// initialize/config.go

package initialize

import (
	"fmt"
	"mxshop_srvs/user_srv/global"

	"github.com/spf13/viper"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user_srv/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("user_srv/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
}
```
### 数据库初始化
```go
// initialize/db.go

package initialize

import (
	"fmt"
	"log"
	"mxshop_srvs/user_srv/config"
	"mxshop_srvs/user_srv/global"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
	ServerConfig config.ServerConfig
)

func InitDB() {
	c := global.ServerConfig.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Name)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// 全局模式
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
```
### logger初始化
```go
// initialize/logger

package initialize

import (
	"go.uber.org/zap"
)

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
```
### 在main中调用
```go
// mian.go

// other code

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
// 使用zap替代fmt
	flag.Parse()
	zap.S().Info("ip：", *IP)
	zap.S().Info("port：", *Port)

// other code
```
## 注册服务健康检查
```go
// main.go
// 导入包
import (
    "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/health"
)
grpc_health_v1.RegisterHealthServer(server, health.NewServer())
```
### grpc的健康检查规范
[grpc健康文档](https://github.com/grpc/grpc/blob/master/doc/health-checking.md)
grpc健康检查要点：
```go
1.check = {
    "GRPC": f'{ip}:{port}',
    "GRPCUseTLS": False,
    "Timeout": "5s",
    "Interval": "5s",
    "DeregisterCriticalServiceAfter": "5s",
}
```
2.一定要确保网络是通畅的
3.一定要确保srv服务监听端口是对外可访问的
4.GRPC一定要自己填写
## 将grpc服务注册到consul中
### 添加consul配置
```go
// config/config.go
// other code
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name string `mapstructure:"name" json:"name"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
}
```
### 修改yaml配置文件
```yaml
name: 'user-srv'
mysql:
  host: '127.0.0.1'
  port: 3306
  user: 'root'
  password: 'root'
  db: 'mxshop_user_srv'

consul:
  host: '192.168.0.50'
  port: 8500
```
### 在main中使用
```go
package main

import (
	"flag"
	"fmt"
	"net"

	"mxshop_srvs/user_srv/global"
	"mxshop_srvs/user_srv/handler"
	"mxshop_srvs/user_srv/initialize"
	"mxshop_srvs/user_srv/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"github.com/hashicorp/consul/api"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("prot", 50051, "端口号")

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	zap.S().Info(global.ServerConfig)

	flag.Parse()
	zap.S().Info("ip：", *IP)
	zap.S().Info("port：", *Port)

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	// 注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 服务注册
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	check := &api.AgentServiceCheck{
		GRPC: fmt.Sprintf("192.168.0.103:50051"),
		Timeout: "5s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.ID = global.ServerConfig.Name
	registration.Port = *Port
	registration.Tags = []string{"imooc", "user", "srv"}
	registration.Address = "192.168.0.103"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc" + err.Error())
	}
}

```
## 获取动态端口
### 添加utils函数
```go
// utils/addr.go

package utils

import "net"

func GetFreePort() (int, error) {
    addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
    if err != nil {
        return 0, err
    }
    
    l, err := net.ListenTCP("tcp", addr)
    if err != nil {
        return 0, err
    }
    defer l.Close()
    return l.Addr().(*net.TCPAddr).Port, nil
}
```
### 使用动态端口
```go
// mian.go

package main

import (
    "flag"
    "fmt"
    "net"
    
    "mxshop_srvs/user_srv/global"
    "mxshop_srvs/user_srv/handler"
    "mxshop_srvs/user_srv/initialize"
    "mxshop_srvs/user_srv/proto"
    "mxshop_srvs/user_srv/utils"
    
    "github.com/hashicorp/consul/api"
    "go.uber.org/zap"
    "google.golang.org/grpc"
    "google.golang.org/grpc/health"
    "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
    IP := flag.String("ip", "0.0.0.0", "ip地址")
    ---// 将默认端口改为0
    Port := flag.Int("prot", 0, "端口号")
    
    initialize.InitLogger()
    initialize.InitConfig()
    initialize.InitDB()
    zap.S().Info(global.ServerConfig)
    
    flag.Parse()
    zap.S().Info("ip：", *IP)
    ------
    if *Port == 0 {
        *Port, _ = utils.GetFreePort()
    }
    ------
    zap.S().Info("port：", *Port)
    
    server := grpc.NewServer()
    proto.RegisterUserServer(server, &handler.UserServer{})
    lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
    if err != nil {
        panic("failed to listen:" + err.Error())
    }
    // 注册服务健康检查
    grpc_health_v1.RegisterHealthServer(server, health.NewServer())
    
    // 服务注册
    cfg := api.DefaultConfig()
    cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
    
    client, err := api.NewClient(cfg)
    check := &api.AgentServiceCheck{
        ---// 使用动态端口
        GRPC: fmt.Sprintf("192.168.0.101:%d", *Port),
        Timeout: "5s",
        Interval: "5s",
        DeregisterCriticalServiceAfter: "10s",
    }
    
    // 生成注册对象
    registration := new(api.AgentServiceRegistration)
    registration.Name = global.ServerConfig.Name
    registration.ID = global.ServerConfig.Name
    registration.Port = *Port
    registration.Tags = []string{"imooc", "user", "srv"}
    registration.Address = "192.168.0.101"
    registration.Check = check
    
    err = client.Agent().ServiceRegister(registration)
    if err != nil {
        panic(err)
    }
    
    err = server.Serve(lis)
    if err != nil {
        panic("failed to start grpc" + err.Error())
    }
}

```
## 添加负载均衡
### 使用uuid向consul添加不同服务
```go
// mian.go

package mina

import (
    // other code
    
    "github.com/satori/go.uuid"
    
    // other code
)
// ohter code
// 生成注册对象
registration := new(api.AgentServiceRegistration)
registration.Name = global.ServerConfig.Name
------
// 使用uuid生成随机服务名称
serviceID := fmt.Sprintf("%s", uuid.NewV4())
registration.ID = serviceID
------
registration.Port = *Port
registration.Tags = []string{"imooc", "user", "srv"}
registration.Address = "192.168.0.102"
registration.Check = check

  // other code
```
### 监听使用ctrl+c退出服务
```go
// main.go

// other code 
{
    // 使用异步解决程序到此处停止的问题，无法监听的退出。
    go func() {
        err = server.Serve(lis)
        if err != nil {
            panic("failed to start grpc" + err.Error())
        }
    }()
    // 接收终止信号
    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    if err = client.Agent().ServiceDeregister(serviceID); err != nil {
        zap.S().Info("注销失败")
    }
    zap.S().Info("注销成功")
 }
```
