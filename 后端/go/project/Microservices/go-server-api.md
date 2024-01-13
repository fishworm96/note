### Birthday类型处理
从数据库取数据直接返回的类型是byte类型，现在转换为日期类型
```go
// BirthDay在proto中定义的类型是int64
data := make(map[string]interface{})

data["id"] = value.Id
data["name"] = value.NickName
data["birthday"] = time.Time(time.Unix(int64(value.BirthDay), 0))
data["gender"] = value.Gender
data["mobile"] = value.Mobile

result = append(result, data)
```
使用结构体的方法返回
```go
// 定义结构体
// response/user.go
type UserResponse struct {
	Id       int32  `json:"id"`
	NickName string `json:"name"`
	Birthday time.Time `json:"birthday"`
	Gender string `json:"gender"`
	Mobile string `json:"mobile"`
}
// api/user.go
user := response.UserResponse{
	Id: value.Id,
	NickName: value.NickName,
	Birthday: time.Time(time.Unix(int64(value.BirthDay), 0)),
	Gender: value.Gender,
	Mobile: value.Mobile,
}

result = append(result, user)
```
修改结构体BirthDay的类型来格式化日期
```go
// response/user.go
type UserResponse struct {
	Id       int32  `json:"id"`
	NickName string `json:"name"`
	Birthday string `json:"birthday"`
	Gender string `json:"gender"`
	Mobile string `json:"mobile"`
}
// api/user.go
user := response.UserResponse{
    Id: value.Id,
	NickName: value.NickName,
	Birthday: time.Time(time.Unix(int64(value.BirthDay), 0)).Format("2006-01-02"),
	Gender: value.Gender,
	Mobile: value.Mobile,
}
```
添加处理函数来格式化BirthDay的日期
```go
// response/user.go
type JsonTime time.Time
// 重写方法在初始化结构体Birthday字段自动调用方法
func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

type UserResponse struct {
	Id       int32  `json:"id"`
	NickName string `json:"name"`
	Birthday JsonTime `json:"birthday"`
	Gender string `json:"gender"`
	Mobile string `json:"mobile"`
}

// api/user.go
user := response.UserResponse{
    Id: value.Id,
	NickName: value.NickName,
	Birthday: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
	Gender: value.Gender,
	Mobile: value.Mobile,
}
```
## 在gin中使用viper
### 定义ymal
```yaml
// 根目录下config-debug.yaml
name: 'user-web'
port: 8021
user_srv:
  host: '127.0.0.1'
  prot: 50051
```
现在通过系统的环境变量来控制是否是开发模式，在系统变量中添加 变量名：**MXSHOP_DEBUG**变量值**：true**
### 类型配置文件
```go
// config/config.go
package config

type UserSrvConfig struct {
    Host string `mapstructure:"host"`
    Port int `mapstructure:"port"`
}

type ServerConfig struct {
    Name string `mapstructure:"name"`
    Port int `mapstructure:"port"`
    UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
}
```
### 定义全局环境变量
```go
// global/global.go
package global

import  "mxshop-api/user-web/config"

var (
    ServerConfig *config.ServerConfig = &config.ServerConfig{}
)
```
### 初始化变量
```go
// initialize
package initialize

import (
    "fmt"
    "mxshop-api/user-web/global"
    
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
    "go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
    viper.AutomaticEnv()
    return viper.GetBool(env)
}

func InitConfig() {
    debug := GetEnvInfo("MXSHOP_DEBUG")
    configFilePrefix := "config"
    configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
    if debug {
        configFileName = fmt.Sprintf("user-web/%s-debug.yaml", configFilePrefix)
    }
    
    v := viper.New()
    // 文件的路径如何设置
    v.SetConfigFile(configFileName)
    if err := v.ReadInConfig(); err != nil {
        panic(err)
    }
    // 其他文件中使用全局变量
    if err := v.Unmarshal(global.ServerConfig); err != nil {
        panic(err)
    }
    zap.S().Infof("配置信息: %v", global.ServerConfig)
    
    // viper的功能-动态监控变化
    v.WatchConfig()
    v.OnConfigChange(func (e fsnotify.Event) {
        zap.S().Infof("配置文件产生变化：%s", e.Name)
        _ = v.ReadInConfig()
        _ = v.Unmarshal(&global.ServerConfig)
    })
}
```
### 在启动文件中使用
```go
// 根目录下main.go
package main

import (
	"fmt"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"

	"go.uber.org/zap"
)

func main() {
	// 初始化全局logger
	initialize.InitLogger()

	// 初始化配置文件
	initialize.InitConfig()
	
	// 初始化routers
	Router := initialize.Routers()

	/* 
		1.S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2.日志是分级别的，debug，info，warn，error，fetal
		3.S函数和L函数很有用，提供了一个全局的安全访问logger的途径
	*/

	zap.S().Debugf("启动服务器,端口: %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}
}
```
## 表单验证
### 定义参数
```go
// forms/user.go

package forms

type PassWordListForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"` // 手机号格式有管饭可循 自定义validator
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}
```
### 定义全局变量
```go
// global/global.go

package global

import (
	"mxshop-api/user-web/config"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans        ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)

```
### 定义api接口
```go
// api/user.go

package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 错误处理函数
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

// 接口函数
func PassWordLogin(c *gin.Context) {
	// 表单验证
	passwordLoginForm := forms.PassWordListForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": removeTopStruct(errs.Translate(global.Trans)),
		})
		return
	}
}
```
### 注册路由
```go
// router/user.go

package router

import (
	"mxshop-api/user-web/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")

	{
		UserRouter.GET("/list", api.GetUserList)
		UserRouter.POST("/pwd_login", api.PassWordLogin)
	}
}
```
### 全局错误处理初始化
```go
// initialize/validator.go

package initialize

import (
	"fmt"
	"mxshop-api/user-web/global"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			en_translations.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}
	return
}

```
### 使用全局错误处理
```go
// mian.go

package main

import (
	"fmt"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"

	"go.uber.org/zap"
)

func main() {
	// 初始化全局logger
	initialize.InitLogger()

	// 初始化配置文件
	initialize.InitConfig()

	// 初始化routers
	Router := initialize.Routers()

	// 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	/*
		1.S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2.日志是分级别的，debug，info，warn，error，fetal
		3.S函数和L函数很有用，提供了一个全局的安全访问logger的途径
	*/

	zap.S().Debugf("启动服务器,端口: %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}
}

```
测试：使用postman api路径：**localhost:8021/u/v1/user/pwd_login**
### 抽离错误处理函数
```go
// api/user.go

package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 错误处理函数
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

// 抽离函数
func HandleValidator(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
}


// 接口函数
func PassWordLogin(c *gin.Context) {
	// 表单验证
	passwordLoginForm := forms.PassWordListForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidator(c, err)
		return
	}
}
```
## 手机号码校验
### 校验手机号码
```go
// validator/validator.go
package Validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	// 使用正则判断是否合法
	ok, _ := regexp.MatchString(`^(1[38][0-9]14[579]|5[^4]|16[6]|7[1-35-8]|9[189]\d{8})$`, mobile)
	if !ok {
		return false
	}
	return true
}
```
### 使用校验器
```go
// main.go

package main

import (
	"fmt"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"

	myvalidator "mxshop-api/user-web/validator"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func main() {
	// 初始化全局logger
	initialize.InitLogger()

	// 初始化配置文件
	initialize.InitConfig()

	// 初始化routers
	Router := initialize.Routers()

	// 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	// 注册校验器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func (ut ut.Translator) error  {
			return ut.Add("mobile", "{0}手机号非法", true)
		}, func (ut ut.Translator, fe validator.FieldError) string  {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	/*
		1.S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2.日志是分级别的，debug，info，warn，error，fetal
		3.S函数和L函数很有用，提供了一个全局的安全访问logger的途径
	*/

	zap.S().Debugf("启动服务器,端口: %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}
}
```
最后需要在tag中required后面添加mobile
```go
// forms/user.go
package forms

type PassWordListForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"` // 手机号格式有管饭可循 自定义validator
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}
```
## 完善登录逻辑
```go
// api/user.go

func PassWordLogin(c *gin.Context) {
	// 表单验证
	passwordLoginForm := forms.PassWordListForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidator(c, err)
		return
	}

	// 拨号连接用户grpc服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [用户服务失败]", "msg", err.Error())
	}
	// 生成grcp的client并调用接口
	userSrvClient := proto.NewUserClient(userConn)

	// 登录逻辑
	if rsp, err := userSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登录失败",
				})
			}
			return
		}
	} else {
		// 只是查询到用户了，并没有检查密码
		if passRsp, pasErr := userSrvClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:          passwordLoginForm.PassWord,
			EncryptedPassword: rsp.PassWord,
		}); pasErr != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"password": "登录失败",
			})
		} else {
			if passRsp.Success {
				c.JSON(http.StatusOK, map[string]string{
					"msg": "登录成功",
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "账号或密码错误",
				})
			}
		}
	}
}

```
## 使用jwt做登录认证
### 定义jwt模型
```go
// models/request.go
package models

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID uint
	NickName string
	AuthorityId uint
	jwt.StandardClaims
}
```
### 定义jwt函数
```go
// middlewares/jwt.go
package middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/models"
	"net/http"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg":"请登录",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				if err == TokenExpired {
					c.JSON(http.StatusUnauthorized, map[string]string{
						"msg":"授权已过期",
					})
					c.Abort()
					return
				}
			}

			c.JSON(http.StatusUnauthorized, "未登陆")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.ServerConfig.JWTInfo.SigningKey), //可以设置过期时间
	}
}

// 创建一个token
func (j *JWT) CreateToken(claims models.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
```
### 在全局配置中添加jwt字段
```go
// config/config.go
package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int `mapstructure:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
	JWTInfo JWTConfig `mapstructure:"jwt"`
}
```
### 在登录接口中使用
```go
// api/user.go
func PassWordLogin(c *gin.Context) {
	// 表单验证
	passwordLoginForm := forms.PassWordListForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidator(c, err)
		return
	}

	// 拨号连接用户grpc服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [用户服务失败]", "msg", err.Error())
	}
	// 生成grcp的client并调用接口
	userSrvClient := proto.NewUserClient(userConn)

	// 登录逻辑
	if rsp, err := userSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登录失败",
				})
			}
			return
		}
	} else {
		// 只是查询到用户了，并没有检查密码
		if passRsp, pasErr := userSrvClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:          passwordLoginForm.PassWord,
			EncryptedPassword: rsp.PassWord,
		}); pasErr != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"password": "登录失败",
			})
		} else {
			if passRsp.Success {
				// 生成token
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID: uint(rsp.Id),
					NickName: rsp.NickName,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(), // 签名时间
						ExpiresAt: time.Now().Unix() + 60*60*24*30, // 30天过期
						Issuer: "imooc",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"id": rsp.Id,
					"nick_name": rsp.NickName,
					"token": token,
					"expired_at": (time.Now().Unix() + 60*60*24*30)*1000,
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "账号或密码错误",
				})
			}
		}
	}
}
```
## 权限认证
### 获取登录用户的信息
```go
在jwt中通过c.Set("claims", claims)来保存登录用户的信息
// api/user.go
func GetUserList(ctx *gin.Context) {
	// other code
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户：%d", currentUser.ID)

    // other code
}
```
现在能获取到登录用户的id
### 权限认证
```go
// middlewares/admin.go
package middlewares

import (
	"mxshop-api/user-web/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
```
## 配置跨域
### 定义跨域函数
```go
// middlewares/cors.go
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
### 使用
```go
// initialize/router.go
package initialize

import (
	"mxshop-api/user-web/middlewares"
	router "mxshop-api/user-web/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Routers() *gin.Engine {
	Router := gin.Default()
    // 使用跨域中间件
	Router.Use(middlewares.Cors())
	
	zap.S().Info("配置用户相关的url")
	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)

	return Router
}
```
## 添加验证码校验
### 定义验证码函数
```go
// api/captcha.go
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.S().Error("生成验证码错误", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath": b64s,
	})
}
```
### 生成验证码路由
```go
// router/base.go
package router

import (
	"mxshop-api/user-web/api"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
	}
}
```
### 注册验证码路由
```go
// initialize/router.go
package initialize

import (
	"mxshop-api/user-web/middlewares"
	router "mxshop-api/user-web/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors())

	zap.S().Info("配置用户相关的url")
	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)
    // 注册生成验证码路由
	router.InitBaseRouter(ApiGroup)

	return Router
}
```
### 定义验证码校验参数
```go
// forms/user.go
package forms

type PassWordListForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"` // 手机号格式有管饭可循 自定义validator
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
    // 验证码字段
	Captcha string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}
```
### 验证与机械验证码
```go
// api/user.go

func PassWordLogin(c *gin.Context) {
	// 表单验证
	passwordLoginForm := forms.PassWordListForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidator(c, err)
		return
	}

	if store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		c.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}
    // other code
}
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1655739517129-e57b707f-eedf-40b0-8625-ab6308d3df1a.png#clientId=u85e5c758-4f40-4&from=paste&height=832&id=u5a177e1c&originHeight=1040&originWidth=1655&originalType=binary&ratio=1&rotation=0&showTitle=false&size=171266&status=done&style=none&taskId=u14ee5fb2-bed4-42a3-a5f0-580cdc878c3&title=&width=1324)
## gin集成consul
### 添加consul配置
```go
// config/config.go

package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Expire int `mapstructure:"expire"`
}

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int `mapstructure:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
	JWTInfo JWTConfig `mapstructure:"jwt"`
	RedisInfo RedisConfig `mapstructure:"redis"`
	ConsulInfo ConsulConfig `mapstructure:"consul"`
}
```
### 配置yaml
```yaml
// config-debug.yaml

name: 'user-web'
port: 8021

user_srv:
  host: '192.168.0.103'
  port: 50051
  name: 'user-srv'

jwt:
  key: '12345'

redis:
  host: '192.168.0.50'
  port: 6379

consul:
  host: '192.168.0.50'
  port: 8500
```
### 注册consul
```go
func GetUserList(ctx *gin.Context) {
    ------
	// 从注册中心获取到用户服务的信息
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvInfo.Name))
	// data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"captcha": "用户服务不可达",
		})
	}

	// 拨号连接用户grpc服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [用户服务失败]", "msg", err.Error())
	}
    ------
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户：%d", currentUser.ID)

	// 生成grcp的client并调用接口
	userSrvClient := proto.NewUserClient(userConn)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	psizeInt, _ := strconv.Atoi(pSize)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(psizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 [用户列表] 失败")
		HandlerGrpcErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		// data := make(map[string]interface{})

		user := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			// Birthday: time.Time(time.Unix(int64(value.BirthDay), 0)).Format("2006-01-02"),
			Birthday: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		// data["id"] = value.Id
		// data["name"] = value.NickName
		// data["birthday"] = time.Time(time.Unix(int64(value.BirthDay), 0))
		// data["gender"] = value.Gender
		// data["mobile"] = value.Mobile

		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
	zap.S().Debug("获取用户列表")
}
```
## 抽离grpc服务到全局共享
### 添加全局变量
```go
// global/global.go

package global

import (
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/proto"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans        ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	UserSrvClient proto.UserClient
)

```
### 将用户服务抽离到初始化
```go
//initialize/src_conn.go

package initialize

import (
	"fmt"

	"mxshop-api/user-web/global"
	"mxshop-api/user-web/proto"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {
	// 从注册中心获取到用户服务的信息
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvInfo.Name))
	// data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 [用户服务失败]")
		return
	}

	// 拨号连接用户grpc服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [用户服务失败]", "msg", err.Error())
	}
	// 生成grcp的client并调用接口
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}

```
### 在全局中使用
```go
// main.go

package main

import (
	"fmt"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"

	myvalidator "mxshop-api/user-web/validator"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func main() {
	// 初始化全局logger
	initialize.InitLogger()

	// 初始化配置文件
	initialize.InitConfig()

	// 初始化routers
	Router := initialize.Routers()

	// 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
------
	// 初始化srv连接
	initialize.InitSrvConn()
------
	// 注册校验器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func (ut ut.Translator) error  {
			return ut.Add("mobile", "{0}手机号非法", true)
		}, func (ut ut.Translator, fe validator.FieldError) string  {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	/*
		1.S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2.日志是分级别的，debug，info，warn，error，fetal
		3.S函数和L函数很有用，提供了一个全局的安全访问logger的途径
	*/

	zap.S().Debugf("启动服务器,端口: %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}
}

```
### 修改需要连接的用户服务
```go
// api/user.go

func GetUserList(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户：%d", currentUser.ID)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	psizeInt, _ := strconv.Atoi(pSize)
    ------
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(psizeInt),
	})
    ------
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 [用户列表] 失败")
		HandlerGrpcErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		// data := make(map[string]interface{})

		user := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			// Birthday: time.Time(time.Unix(int64(value.BirthDay), 0)).Format("2006-01-02"),
			Birthday: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		// data["id"] = value.Id
		// data["name"] = value.NickName
		// data["birthday"] = time.Time(time.Unix(int64(value.BirthDay), 0))
		// data["gender"] = value.Gender
		// data["mobile"] = value.Mobile

		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
	zap.S().Debug("获取用户列表")
}

func PassWordLogin(c *gin.Context) {
	// 表单验证
	passwordLoginForm := forms.PassWordListForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidator(c, err)
		return
	}

	if store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		c.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	// 登录逻辑
    ------
	if rsp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
    ------
		Mobile: passwordLoginForm.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登录失败",
				})
			}
			return
		}
	} else {
		// 只是查询到用户了，并没有检查密码
		if passRsp, pasErr := global.UserSrvClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:          passwordLoginForm.PassWord,
			EncryptedPassword: rsp.PassWord,
		}); pasErr != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"password": "登录失败",
			})
		} else {
			if passRsp.Success {
				// 生成token
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.NickName,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),               // 签名时间
						ExpiresAt: time.Now().Unix() + 60*60*24*30, // 30天过期
						Issuer:    "imooc",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.NickName,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "账号或密码错误",
				})
			}
		}
	}
}

```
## 动态获取可用端口
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

// other code
// 初始化srv连接
initialize.InitSrvConn()
------
viper.AutomaticEnv()
// 如果本地开发环境端口号固定，线上环境自动获取端口号
debug := viper.GetBool("MXSHOP_DEBUG")
if !debug {
    port, err := utils.GetFreePort()
    if err == nil {
        global.ServerConfig.Port = port
    }
}
------
// 注册校验器
if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
    _ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
    _ = v.RegisterTranslation("mobile", global.Trans, func (ut ut.Translator) error  {
        return ut.Add("mobile", "{0}手机号非法", true)
    }, func (ut ut.Translator, fe validator.FieldError) string  {
        t, _ := ut.T("mobile", fe.Field())
        return t
    })
}
 // other code
```
## 配置负载均衡
### 配置负载均衡
```go
// initialize/srv_conn.go

package initialize

import (
    "fmt"
    
    "mxshop-api/user-web/global"
    "mxshop-api/user-web/proto"
    
    "github.com/hashicorp/consul/api"
    "go.uber.org/zap"
    "google.golang.org/grpc"
    // 必须导入这个包，不然无法完成负载均衡
    _ "github.com/mbobakov/grpc-consul-resolver"
)

func InitSrvConn() {
    consulInfo := global.ServerConfig.ConsulInfo
    userConn, err := grpc.Dial(
        fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
        grpc.WithInsecure(),
        grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
    )
    if err != nil {
        zap.S().Fatal("[InitSrvConn] 连接 [用户服务失败]")
    }
    
    userSrvClient := proto.NewUserClient(userConn)
    global.UserSrvClient = userSrvClient
}
```
## go中配置nacos
### 修改yaml配置文件
```yaml
// debug-config.yaml

host: "192.168.0.50"
port: 8848
namespace: "74011732-0b67-434c-b788-c44ebe811137"
user: "nacos"
password: "nacos"
dataid: "user-web.json"
group: "dev"
```
### 添加Nacos配置并修改原来配置为json
```go
// config/config.go

package config

type UserSrvConfig struct {
    Host string `mapstructure:"host" json:"host"`
    Port int    `mapstructure:"port" json:"port"`
    Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
    SigningKey string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
    Host string `mapstructure:"host" json:"host"`
    Port int    `mapstructure:"port" json:"port"`
}

type RedisConfig struct {
    Host   string `mapstructure:"host" json:"host"`
    Port   int    `mapstructure:"port" json:"port"`
    Expire int    `mapstructure:"expire" json:"expire"`
}

type ServerConfig struct {
    Name        string        `mapstructure:"name" json:"name"`
    Port        int           `mapstructure:"port" json:"port"`
    UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
    JWTInfo     JWTConfig     `mapstructure:"jwt" json:"jwt"`
    RedisInfo   RedisConfig   `mapstructure:"redis" json:"redis"`
    ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
}

type NacosConfig struct {
    Host      string `mapstructure:"host"`
    Port      uint64 `mapstructure:"port"`
    Namespace string `mapstructure:"namespace"`
    User      string `mapstructure:"user"`
    Password  string `mapstructure:"password"`
    DataId    string `mapstructure:"dataid"`
    Group     string `mapstructure:"group"`
}

```
### 在global中引用
```go
// global/global.go

package global

import (
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/proto"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans        ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	---
	NacosConfig *config.NacosConfig = &config.NacosConfig{}
    ---
	UserSrvClient proto.UserClient
)

```
### 在初始化配置中使用
[nacos中文文档](https://github.com/nacos-group/nacos-sdk-go/blob/master/README_CN.md)
```go
// initialize/config.go

package initialize
import (
    // other code
    
    // 导入包
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func InitConfig() {
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("user-web/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	// 文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 其他文件中使用全局变量
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息: %v", global.NacosConfig)

	// 从nacos中读取配置信息
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})

	if err != nil {
		panic(err)
	}
	// 想要将一个json字符串转为struct，需要取设置这个struct的tag
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Fatalf("读取nacos配置失败: %s", err.Error())
	}
	fmt.Println(&global.ServerConfig)
}
```
