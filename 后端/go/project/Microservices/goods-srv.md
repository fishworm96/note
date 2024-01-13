# 商品服务
## 定义数据库表字段
商品服务和用户服务基本一样，先把创建一个goods_srv然后把用户服务全都复制过来，把user_srv改成goods_srv。
```go
// model/base.go

package model

import (
    "time"
    
    "gorm.io/gorm"
)

type BaseModel struct {
    // 为什么ID要定义为int32，因为数据库的类型不同会出现错误。type为int类型基本够用，如果数据量大可以定义为bigint
    ID        int32     `gorm:"primarykey;type:int"`
    CreatedAt time.Time `gorm:"column:add_time"`
    UpdatedAt time.Time `gorm:"column:update_time"`
    DeletedAt gorm.DeletedAt
    IsDeleted bool
}
```
## 定义goods
```go
package model

// 在开发中，尽量不要设置null
// https://zhuanlan.zhihu.com/p/73997266
// 分类表
type Category struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null;comment:'商品名称'"`
	// 在类型转换中经常要用int32或者int64为了方便定义为int32
	ParentCategoryID int32     `gorm:"comment:'自关联id'"`
	ParentCategory   *Category `gorm:"comment:'自关联商品'"`
	Level            int32     `gorm:"type:int;not null;default:1;comment:'1代表1级类目，2代表二级类目，3代表三级类目'"`
	IsTab            bool      `gorm:"default:false;not null;comment:'是否在tab栏展示'"`
}

// 品牌表
type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null;comment:'品牌名称'"`
	Logo string `gorm:"type:varchar(200);default:'';not null;comment:'品牌logo图片'"`
}

// 品牌和分类关联表
type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"` // 联合唯一索引，解决一个数据重复添加2次
	Category Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

// 轮播图表
type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url string `gorm:"type:varcahr(200);not null"`
	Index int32 `gorm:"type:int;default:1;not null"`
}

// 商品表
type Goods struct {
	BaseModel

	// 这里的唯一键是由商品id或商品名称来确定
	CategoryID int32 `gorm:"type:int;not null;comment:'分类id'"`
	Category Category
	BrandsID int32 `gorm:"type:int;not null;comment:'商品id'"`
	Brands Brands

	OnSale bool `gorm:"default:false;not null;comment:'是否上架'"`
	ShipFree bool `gorm:"default:false;not null;comment:'是否免运费'"`
	IsNew bool `gorm:"default:false;not null;comment:'是否新品'"`
	IsHot bool `gorm:"default:false;not null;comment:'是否热门商品'"`

	Name string `gorm:"type:varchar(50);not null;comment:'商品名称'"`
	GoodsSn string `gorm:"type:varchar(50);not null;comment:'商品编号'"`
	ClickNum int32 `gorm:"type:int;default:0;not null;comment:'点击数'"`
	SoldNum int32 `gorm:"type:int;default:0;not null;comment:'购买数'"`
	FavNum int32 `gorm:"type:int;default:0;not null;comment:'收藏数'"`
	MarketPrice float32 `gorm:"not null;comment:'市场价'"`
	ShopPrice float32 `gorm:"not null;comment:'本店价格'"`
	GoodsBrief string `gorm:"type:varchar(100);not null;comment:'商品简介'"`
	// 另外再建一张表存储的话，通过join后会有性能问题，所以使用gorm的自定义类型来出来。
	Images GormList `gorm:"type:varchar(1000);not null;comment:'商品展示图片'"`
	DescImages GormList `gorm:"type:varchar(1000);not null;comment:'商品内容图片'"`
	GoodsFrontImage string `gorm:"type:varchar(200);not null;comment:'商品封面图片'"`
}
```
## 添加gorm自定义类型
```go
package model

// 在开发中，尽量不要设置null
// https://zhuanlan.zhihu.com/p/73997266
// 分类表
type Category struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null;comment:'商品名称'"`
	// 在类型转换中经常要用int32或者int64为了方便定义为int32
	ParentCategoryID int32     `gorm:"comment:'自关联id'"`
	ParentCategory   *Category `gorm:"comment:'自关联商品'"`
	Level            int32     `gorm:"type:int;not null;default:1;comment:'1代表1级类目，2代表二级类目，3代表三级类目'"`
	IsTab            bool      `gorm:"default:false;not null;comment:'是否在tab栏展示'"`
}

// 品牌表
type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null;comment:'品牌名称'"`
	Logo string `gorm:"type:varchar(200);default:'';not null;comment:'品牌logo图片'"`
}

// 品牌和分类关联表
type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"` // 联合唯一索引，解决一个数据重复添加2次
	Category Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

// 轮播图表
type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null;"`
	Url string `gorm:"type:varchar(200);not null"`
	Index int32 `gorm:"type:int;default:1;not null"`
}

// 商品表
type Goods struct {
	BaseModel

	// 这里的唯一键是由商品id或商品名称来确定
	CategoryID int32 `gorm:"type:int;not null;comment:'分类id'"`
	Category Category
	BrandsID int32 `gorm:"type:int;not null;comment:'商品id'"`
	Brands Brands

	OnSale bool `gorm:"default:false;not null;comment:'是否上架'"`
	ShipFree bool `gorm:"default:false;not null;comment:'是否免运费'"`
	IsNew bool `gorm:"default:false;not null;comment:'是否新品'"`
	IsHot bool `gorm:"default:false;not null;comment:'是否热门商品'"`

	Name string `gorm:"type:varchar(50);not null;comment:'商品名称'"`
	GoodsSn string `gorm:"type:varchar(50);not null;comment:'商品编号'"`
	ClickNum int32 `gorm:"type:int;default:0;not null;comment:'点击数'"`
	SoldNum int32 `gorm:"type:int;default:0;not null;comment:'购买数'"`
	FavNum int32 `gorm:"type:int;default:0;not null;comment:'收藏数'"`
	MarketPrice float32 `gorm:"not null;comment:'市场价'"`
	ShopPrice float32 `gorm:"not null;comment:'本店价格'"`
	GoodsBrief string `gorm:"type:varchar(100);not null;comment:'商品简介'"`
	// 另外再建一张表存储的话，通过join后会有性能问题，所以使用gorm的自定义类型来出来。
	Images GormList `gorm:"type:varchar(1000);not null;comment:'商品展示图片'"`
	DescImages GormList `gorm:"type:varchar(1000);not null;comment:'商品内容图片'"`
	GoodsFrontImage string `gorm:"type:varchar(200);not null;comment:'商品封面图片'"`
}
```
## gorm表迁移
```go
// model/main/main.go

package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"mxshop_srvs/goods_srv/model"
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&model.Category{}, &model.Brands{}, &model.GoodsCategoryBrand{}, &model.Banner{}, &model.Goods{})
}

```
## 定义proto
```protobuf
// proto/goods.proto

syntax = "proto3";
import "google/protobuf/empty.proto";
option go_package = ".;proto";

service Goods{
//商品接口
rpc GoodsList(GoodsFilterRequest) returns(GoodsListResponse);
//现在用户提交订单有多个商品，你得批量查询商品的信息吧
rpc BatchGetGoods(BatchGoodsIdInfo) returns(GoodsListResponse); //批量获取商品信息
rpc CreateGoods(CreateGoodsInfo) returns (GoodsInfoResponse);
rpc DeleteGoods(DeleteGoodsInfo) returns (google.protobuf.Empty);
rpc UpdateGoods(CreateGoodsInfo) returns (google.protobuf.Empty);
rpc GetGoodsDetail(GoodInfoRequest) returns(GoodsInfoResponse);

//商品分类
rpc GetAllCategorysList(google.protobuf.Empty) returns(CategoryListResponse); //获取所有的分类
//获取子分类
rpc GetSubCategory(CategoryListRequest) returns(SubCategoryListResponse);
rpc CreateCategory(CategoryInfoRequest) returns(CategoryInfoResponse); //新建分类信息
rpc DeleteCategory(DeleteCategoryRequest) returns(google.protobuf.Empty); //删除分类
rpc UpdateCategory(CategoryInfoRequest) returns(google.protobuf.Empty); //修改分类信息

//品牌和轮播图
rpc BrandList(BrandFilterRequest) returns(BrandListResponse); //
rpc CreateBrand(BrandRequest) returns(BrandInfoResponse); //新建品牌信息
rpc DeleteBrand(BrandRequest) returns(google.protobuf.Empty); //删除品牌
rpc UpdateBrand(BrandRequest) returns(google.protobuf.Empty); //修改品牌信息

//轮播图
rpc BannerList(google.protobuf.Empty) returns(BannerListResponse); //获取轮播列表信息
rpc CreateBanner(BannerRequest) returns(BannerResponse); //添加banner图
rpc DeleteBanner(BannerRequest) returns(google.protobuf.Empty); //删除轮播图
rpc UpdateBanner(BannerRequest) returns(google.protobuf.Empty); //修改轮播图

//品牌分类
rpc CategoryBrandList(CategoryBrandFilterRequest) returns(CategoryBrandListResponse); //获取轮播列表信息
//通过category获取brands
rpc GetCategoryBrandList(CategoryInfoRequest) returns(BrandListResponse);
rpc CreateCategoryBrand(CategoryBrandRequest) returns(CategoryBrandResponse); //添加banner图
rpc DeleteCategoryBrand(CategoryBrandRequest) returns(google.protobuf.Empty); //删除轮播图
rpc UpdateCategoryBrand(CategoryBrandRequest) returns(google.protobuf.Empty); //修改轮播图
}

message CategoryListRequest {
int32 id = 1;
int32 level = 2;
}

message CategoryInfoRequest {
int32 id = 1;
string name = 2;
int32 parentCategory = 3;
int32 level = 4;
bool isTab = 5;
}

message DeleteCategoryRequest {
int32 id = 1;
}

message QueryCategoryRequest {
int32 id = 1;
string name = 2;
}

message CategoryInfoResponse {
int32 id = 1;
string name = 2;
int32 parentCategory = 3;
int32 level = 4;
bool isTab = 5;
}

message CategoryListResponse {
int32 total = 1;
repeated CategoryInfoResponse data = 2;
string jsonData = 3;
}

message SubCategoryListResponse {
int32 total = 1;
CategoryInfoResponse info = 2;
repeated CategoryInfoResponse subCategorys = 3;
}

message CategoryBrandFilterRequest  {
int32 pages = 1;
int32 pagePerNums = 2;
}

message FilterRequest  {
int32 pages = 1;
int32 pagePerNums = 2;
}

message CategoryBrandRequest{
int32 id = 1;
int32 categoryId = 2;
int32 brandId = 3;
}
message CategoryBrandResponse{
int32 id = 1;
BrandInfoResponse brand = 2;
CategoryInfoResponse category = 3;
}

message BannerRequest {
int32 id = 1;
int32 index = 2;
string image = 3;
string url = 4;
}

message BannerResponse {
int32 id = 1;
int32 index = 2;
string image = 3;
string url = 4;
}

message BrandFilterRequest {
int32 pages = 1;
int32 pagePerNums = 2;
}

message BrandRequest {
int32 id = 1;
string name = 2;
string logo = 3;
}

message BrandInfoResponse {
int32 id = 1;
string name = 2;
string logo = 3;
}

message BrandListResponse {
int32 total = 1;
repeated BrandInfoResponse data = 2;
}

message BannerListResponse {
int32 total = 1;
repeated BannerResponse data = 2;
}

message CategoryBrandListResponse {
int32 total = 1;
repeated CategoryBrandResponse data = 2;
}



message BatchGoodsIdInfo {
repeated int32 id = 1;
}


message DeleteGoodsInfo {
int32 id = 1;
}

message CategoryBriefInfoResponse {
int32 id = 1;
string name = 2;
}

message CategoryFilterRequest {
int32 id = 1;
bool  isTab = 2;
}

message GoodInfoRequest {
int32 id = 1;
}

message CreateGoodsInfo {
int32 id = 1;
string name = 2;
string goodsSn = 3;
int32 stocks = 7; //库存，
float marketPrice = 8;
float shopPrice = 9;
string goodsBrief = 10;
string goodsDesc = 11;
bool shipFree = 12;
repeated string images = 13;
repeated string descImages = 14;
string goodsFrontImage = 15;
bool isNew = 16;
bool isHot = 17;
bool onSale = 18;
int32 categoryId = 19;
int32 brandId = 20;
}

message GoodsReduceRequest {
int32 GoodsId = 1;
int32 nums = 2;
}

message BatchCategoryInfoRequest {
repeated int32 id = 1;
int32 goodsNums = 2;
int32 brandNums = 3;
}

message GoodsFilterRequest  {
int32 priceMin = 1;
int32 priceMax = 2;
bool  isHot = 3;
bool  isNew = 4;
bool  isTab = 5;
int32 topCategory = 6;
int32 pages = 7;
int32 pagePerNums = 8;
string keyWords = 9;
int32 brand = 10;
}


message GoodsInfoResponse {
int32 id = 1;
int32 categoryId = 2;
string name = 3;
string goodsSn = 4;
int32 clickNum = 5;
int32 soldNum = 6;
int32 favNum = 7;
float marketPrice = 9;
float shopPrice = 10;
string goodsBrief = 11;
string goodsDesc = 12;
bool shipFree = 13;
repeated string images = 14;
repeated string descImages = 15;
string goodsFrontImage = 16;
bool isNew = 17;
bool isHot = 18;
bool onSale = 19;
int64 addTime = 20;
CategoryBriefInfoResponse category = 21;
BrandInfoResponse brand = 22;
}

message GoodsListResponse {
int32 total = 1;
repeated GoodsInfoResponse data = 2;
}
```
## 生成接口
```shell
// 在当前目录下解析goods.proto生成go文件
protoc --go_out=:. goods.proto
// 在当前目录下解析goods.proto生成grpc的go文件
protoc --go-grpc_out=:. goods.proto
```
## 配置nacos
### 在config中添加配置
```go
// config/config.go

// other code
type ServerConfig struct {
    Name string `mapstructure:"name" json:"name"`
    Host string `mapstructure:"host" json:"host"`
    Tags []string `mapstructure:"tags" json:"tags"`
    MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
    ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
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
### 在global中使用
```go
// global/global.go

// other code
var (
    DB           *gorm.DB
    ServerConfig config.ServerConfig
    NacosConfig config.NacosConfig
)
// other code
```
### 在配置初始化
```go
// initialize/config.go


package initialize

import (
    "encoding/json"
    "fmt"
    "mxshop_srvs/goods_srv/global"
    
    "github.com/nacos-group/nacos-sdk-go/clients"
    "github.com/nacos-group/nacos-sdk-go/common/constant"
    "github.com/nacos-group/nacos-sdk-go/vo"
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
    configFileName := fmt.Sprintf("goods_srv/%s-pro.yaml", configFilePrefix)
    if debug {
        configFileName = fmt.Sprintf("goods_srv/%s-debug.yaml", configFilePrefix)
    }
    
    v := viper.New()
    v.SetConfigFile(configFileName)
    if err := v.ReadInConfig(); err != nil {
        panic(err)
    }
    if err := v.Unmarshal(&global.NacosConfig); err != nil {
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
### 修改main中的配置
```go
// main.go


// 将写死的配置修改为nacos中的配置
// other code
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", global.ServerConfig.Host, *Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
// other code
registration.Tags = global.ServerConfig.Tags
registration.Address = global.ServerConfig.Host
// other code
```
### 修改debug文件
```yaml
// config-debug.yaml


host: '192.168.0.50'
port: 8848
namespace: '2d2739b8-b5be-4568-9555-0ef3e80c1922'
user: 'nacos'
password: 'nacos'
dataid: 'goods-srv.json'
group: 'dev'
```
### 在nacos中添加配置
创建goods空间并使用一下配置
![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1657556059114-f4a5b715-ac48-413f-a603-345f79e22671.png#clientId=u6d957314-1bed-4&from=paste&height=330&id=u92e2caab&originHeight=412&originWidth=1636&originalType=binary&ratio=1&rotation=0&showTitle=false&size=17759&status=done&style=none&taskId=u1169054d-58b2-4a02-9025-da001b51601&title=&width=1308.8)
![image.png](https://cdn.nlark.com/yuque/0/2022/png/12511308/1657555621783-e45478f5-f221-42b3-a10a-7de2b1a47569.png#clientId=u6d957314-1bed-4&from=paste&height=574&id=u72cd29d2&originHeight=717&originWidth=1546&originalType=binary&ratio=1&rotation=0&showTitle=false&size=51129&status=done&style=none&taskId=ua8a27722-d74f-475d-ba23-534ed146a52&title=&width=1236.8)
## 品牌列表
### 设置测试数据
```go
// goods_srv/test/barnds.go


package main

import (
    "context"
    "fmt"
    "mxshop_srvs/goods_srv/proto"
    
    "google.golang.org/grpc"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func Init() {
    var err error
    conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
    if err != nil {
        panic(err)
    }
    brandClient = proto.NewGoodsClient(conn)
}

func TestGetBrandList() {
    rsp, err := brandClient.BrandList(context.Background(), &proto.BrandFilterRequest{})
    if err != nil {
        panic(err)
    }
    fmt.Println(rsp.Total)
    for _, brand := range rsp.Data {
        fmt.Println(brand.Name)
    }
}

func main() {
    Init()
    TestGetBrandList()
    conn.Close()
}

```
为了测试，修改固定端口
```go
// goods_srv/main.go

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("prot", 50051, "端口号")
```
### 添加分页
```go
// goods_srv/model/base.go


func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
```
### 品牌列表接口
```go
// goods_srv/handler/brands.go


package handler

import (
	"context"

	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"
)

func (s *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	brandListResponse := proto.BrandListResponse{}

	var brands []model.Brands
	result := global.DB.Scopes(model.Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)
	brandListResponse.Total = int32(total)

	var brandResponse []*proto.BrandInfoResponse
	for _, brand := range brands {
		brandResponse = append(brandResponse, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	brandListResponse.Data = brandResponse
	return &brandListResponse, nil
}

// CreateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*BrandInfoResponse, error)
// DeleteBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
// UpdateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)

```
## 品牌增删改
```go
// goods_srv/handler/brands.go


package handler

import (
	"context"

	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)



func (s *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	if result := global.DB.First(&model.Brands{}); result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌已存在")
	}

	brand := &model.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}
	global.DB.Save(brand)
	return &proto.BrandInfoResponse{Id: brand.ID}, nil
}
func (s *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Brands{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	return &emptypb.Empty{}, nil
}
func (s *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	brand := model.Brands{}
	if result := global.DB.First(&brand); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	if req.Name != "" {
		brand.Name = req.Name
	}
	if req.Logo != "" {
		brand.Logo = req.Logo
	}

	global.DB.Save(&brand)
	return &emptypb.Empty{}, nil
}

```
## 轮播图增删改查
```go
// goods_srv/handler/banner.go


package handler

import (
	"context"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GoodsServer) BannerList(ctx context.Context, req *emptypb.Empty) (*proto.BannerListResponse, error) {
	bannerListResponse := proto.BannerListResponse{}

	var banners []model.Banner
	result := global.DB.Find(&banners)
	bannerListResponse.Total = int32(result.RowsAffected)

	var bannerResponses []*proto.BannerResponse
	for _, banner := range banners {
		bannerResponses = append(bannerResponses, &proto.BannerResponse{
			Id:    banner.ID,
			Image: banner.Image,
			Index: banner.Index,
			Url:   banner.Url,
		})
	}

	bannerListResponse.Data = bannerResponses

	return &bannerListResponse, nil
}
func (s *GoodsServer) CreateBanner(ctx context.Context, req *proto.BannerRequest) (*proto.BannerResponse, error) {
	banner := model.Banner{}

	banner.Image = req.Image
	banner.Index = req.Index
	banner.Url = req.Url

	global.DB.Save(&banner)

	return &proto.BannerResponse{Id: banner.ID}, nil
}

func (s *GoodsServer) DeleteBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Banner{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	var banner model.Banner

	if result := global.DB.First(&banner, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")
	}
	
	if req.Url != "" {
		banner.Url = req.Url
	}
	if req.Image != "" {
		banner.Image = req.Image
	}
	if req.Index != 0 {
		banner.Index = req.Index
	}

	global.DB.Save(&banner)

	return &emptypb.Empty{}, nil
}
```
## 获取品牌列表
### 获取品牌函数
```go
// goods_srv/handler/category.go


func (s *GoodsServer) GetAllCategorysList(context.Context, *emptypb.Empty) (*proto.CategoryListResponse, error) {
	var categorys []model.Category
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	for _, category := range categorys {
		fmt.Println(category.Name)
	}
	b, _ := json.Marshal(&categorys)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}
```
### 修改品牌表
```go
// goods_srv/model/goods.go

type Category struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null;comment:'商品名称' json:"name""`
	// 在类型转换中经常要用int32或者int64为了方便定义为int32
	ParentCategoryID int32     `gorm:"comment:'自关联id'" json:"parent"`
	ParentCategory   *Category `gorm:"comment:'自关联商品'" json:"-"`
	SubCategory []*Category `gorm:"foreignKey: ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32     `gorm:"type:int;not null;default:1;comment:'1代表1级类目，2代表二级类目，3代表三级类目'" json:"level"`
	IsTab            bool      `gorm:"default:false;not null;comment:'是否在tab栏展示'" json:"is_tab"`
}

// goods_srv/model/base.go

type BaseModel struct {
	// 为什么ID要定义为int32，因为数据库的类型不同会出现错误。type为int类型基本够用，如果数据量大可以定义为bigint
	ID        int32     `gorm:"primarykey;type:int" json:"id"`
	CreatedAt time.Time `gorm:"column:add_time" json:"-"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	IsDeleted bool `json:"-"`
}
```
### 修改vscode运行路径
```json
在setting中添加
"code-runner.executorMap": {
  "go": "cd $dir && go run .",
  },
```
### 单元测试
```go
// goods_srv/test/base.go

package main

import (
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc"
)

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	brandClient = proto.NewGoodsClient(conn)
}

func main() {
	Init()
	// TestGetBrandList()
	TestGetCategoryList()
	conn.Close()
}

// goods_srv/test/brands.go

package main

import (
	"context"
	"fmt"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func TestGetBrandList() {
	rsp, err := brandClient.BrandList(context.Background(), &proto.BrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand.Name)
	}
}

// goods_srv/test/category.go

package main

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

func TestGetCategoryList() {
	rsp, err := brandClient.GetAllCategorysList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.JsonData)
	// for _, category := range rsp.Data {
	// 	fmt.Println(category.Name)
	// }
}

```
## 获取子分类
### 获取子分类接口
```go
// goods_srv/handler/category.go


package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	categoryListResponse := proto.SubCategoryListResponse{}

	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	categoryListResponse.Info = &proto.CategoryInfoResponse{
		Id: category.ID,
		Name: category.Name,
		Level: category.Level,
		IsTab: category.IsTab,
	}

	var subCategorys []model.Category
	var subCategoryResponse []*proto.CategoryInfoResponse
	// preloads := "SubCategory"
	// if category.Level == 1 {
	// 	preloads = "SubCategory.SubCategory"
	// }
	// global.DB.Where(&model.Category{ParentCategoryID: req.Id}).Preload(preloads).Find(&subCategorys)
	global.DB.Where(&model.Category{ParentCategoryID: req.Id}).Find(&subCategorys)
	for _, subCategory := range subCategorys {
		subCategoryResponse = append(subCategoryResponse, &proto.CategoryInfoResponse{
			Id: subCategory.ID,
			Name: subCategory.Name,
			Level: subCategory.Level,
			IsTab: subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}
	categoryListResponse.SubCategorys = subCategoryResponse
	return &categoryListResponse, nil
}
```
### 单元测试
```go
// goods_srv/test/category.go


func TestGetSubCategoryList() {
	rsp, err := brandClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: 135200,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.SubCategorys)
}

// goods_srv/test/base.go

func main() {
	Init()
	// TestGetBrandList()
	// TestGetCategoryList()
	TestGetSubCategoryList()
	conn.Close()
}
```
### 分类增删改
```go
// goods_srv/handler/category.go


func (s *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
    category := model.Category{}
    
    category.Name = req.Name
    category.Level = req.Level
    if req.Level != 1 {
        // 去查询父类目是否存在
        category.ParentCategoryID = req.ParentCategory
    }
    category.IsTab = req.IsTab
    
    global.DB.Save(&category)
    
    return &proto.CategoryInfoResponse{Id: int32(category.ID)}, nil
}

func (s *GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
    if result := global.DB.Delete(&model.Category{}, req.Id); result.RowsAffected == 0 {
        return nil, status.Errorf(codes.NotFound, "商品分类不存在")
    }
    return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
    var category model.Category
    
    if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
        return nil, status.Errorf(codes.NotFound, "商品分类不存在")
    }
    
    if req.Name != "" {
        category.Name = req.Name
    }
    if req.ParentCategory != 0 {
        category.ParentCategoryID = req.ParentCategory
    }
    if req.Level != 0 {
        category.Level = req.Level
    }
    if req.IsTab {
        category.IsTab = req.IsTab
    }
    
    global.DB.Save(&category)
    return &emptypb.Empty{}, nil
}
```
## 品牌分类关联增删改查
### 品牌分类列表
```go
// goods_srv/handler/category_brand.go


func (s *GoodsServer) CategoryBrandList(ctx context.Context, req *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
    var categoryBrands []model.GoodsCategoryBrand
    categoryBrandListResponse := proto.CategoryBrandListResponse{}
    
    var total int64
    global.DB.Model(&model.GoodsCategoryBrand{}).Count(&total)
    categoryBrandListResponse.Total = int32(total)
    
    global.DB.Preload("Category").Preload("Brands").Scopes(model.Paginate(int(req.Pages), int(req.PagePerNums))).Find(&categoryBrands)
    
    var categoryResponse []*proto.CategoryBrandResponse
    for _, categoryBrand := range categoryBrands {
        categoryResponse = append(categoryResponse, &proto.CategoryBrandResponse{
            Category: &proto.CategoryInfoResponse{
                Id: categoryBrand.Category.ID,
                Name: categoryBrand.Category.Name,
                Level: categoryBrand.Category.Level,
                IsTab: categoryBrand.Category.IsTab,
                ParentCategory: categoryBrand.Category.ParentCategoryID,
            },
            Brand: &proto.BrandInfoResponse{
                Id: categoryBrand.Brands.ID,
                Name: categoryBrand.Brands.Name,
                Logo: categoryBrand.Brands.Logo,
            },
        })
    }
    
    categoryBrandListResponse.Data = categoryResponse
    return &categoryBrandListResponse, nil
}
```
### 获取品牌分类列表
```go
// goods_srv/handler/category_brand.go
func (s *GoodsServer) GetCategoryBrandList(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {
    brandListResponse := proto.BrandListResponse{}
    
    var category model.Category
    if result := global.DB.Find(&category, req.Id).First(&category); result.RowsAffected == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
    }
    
    var categoryBrands []model.GoodsCategoryBrand
    if result := global.DB.Where(&model.GoodsCategoryBrand{CategoryID: req.Id}).Find(&categoryBrands); result.RowsAffected > 0 {
        brandListResponse.Total = int32(result.RowsAffected)
    }
    
    var brandInfoResponse []*proto.BrandInfoResponse
    for _, categoryBrand := range categoryBrands {
        brandInfoResponse = append(brandInfoResponse, &proto.BrandInfoResponse{
            Id: categoryBrand.Brands.ID,
            Name: categoryBrand.Brands.Name,
            Logo: categoryBrand.Brands.Logo,
        })
    }
    
    brandListResponse.Data = brandInfoResponse
    
    return &brandListResponse, nil
}
```
### 创建商品分类
```go
// goods_srv/handler/category_brand.go


func (s *GoodsServer) CreateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
    var category model.Category
    if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
    }
    
    var brand model.Brands
    if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
    }
    
    categoryBrand := model.GoodsCategoryBrand{
        CategoryID: req.CategoryId,
        BrandsID: req.BrandId,
    }
    
    global.DB.Save(&categoryBrand)
    return &proto.CategoryBrandResponse{Id: categoryBrand.ID}, nil
}
```
### 删除品牌分类
```go
// goods_srv/handler/category_brand.go


func (s *GoodsServer) DeleteCategoryBrand(ctx context.Context,req *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
    if result := global.DB.Delete(&model.GoodsCategoryBrand{}, req.Id); result.RowsAffected == 0 {
        return nil, status.Errorf(codes.NotFound, "品牌分类不存在")
    }
    
    return &emptypb.Empty{}, nil
}
```
### 修改品牌分类
```go
// goods_srv/handler/category_brand.go


func (s *GoodsServer) UpdateCategoryBrand(ctx context.Context,req *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
    var categoryBrand model.GoodsCategoryBrand
    
    if result := global.DB.First(&categoryBrand, req.Id); result.RowsAffected == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "品牌分类不存在")
    }
    
    var category model.Category
    if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
    }
    
    var brand model.Brands
    if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "分类不存在")
    }
    
    categoryBrand.CategoryID = req.CategoryId
    categoryBrand.BrandsID = req.BrandId
    
    global.DB.Save(&categoryBrand)
    return &emptypb.Empty{}, nil
}
```
### 单元测试
```go
// goods_srv/test/category_brand.go


package main

import (
    "context"
    "fmt"
    "mxshop_srvs/goods_srv/proto"
)

func TestGetCategoryBrandList() {
    rsp, err := brandClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{})
    if err != nil {
        panic(err)
    }
    fmt.Println(rsp.Total)
    fmt.Println(rsp.Data)
}
```
## 查询商品列表
```go
// goods_srv/handler/goods.go


package handler

import (
	"context"
	"fmt"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

func ModelToResponse(goods model.Goods) proto.GoodsInfoResponse {
	return proto.GoodsInfoResponse{
		Id: goods.ID,
		CategoryId: goods.CategoryID,
		Name: goods.Name,
		GoodsSn: goods.GoodsSn,
		ClickNum: goods.ClickNum,
		SoldNum: goods.SoldNum,
		FavNum: goods.FavNum,
		MarketPrice: goods.MarketPrice,
		ShopPrice: goods.ShopPrice,
		GoodsFrontImage: goods.GoodsFrontImage,
		IsNew: goods.IsNew,
		IsHot: goods.IsHot,
		OnSale: goods.OnSale,
		DescImages: goods.DescImages,
		Images: goods.Images,
		Category: &proto.CategoryBriefInfoResponse{
			Id: goods.Category.ID,
			Name: goods.Category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id: goods.Brands.ID,
			Name: goods.Brands.Name,
			Logo: goods.Brands.Logo,
		},
	}
}

func (s *GoodsServer) GoodsList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {

	//关键词搜索、查询新品、查询热门商品、通过价格区间筛选， 通过商品分类筛选
	goodsListResponse := &proto.GoodsListResponse{}

	var goods []model.Goods
	localDB := global.DB.Model(model.Goods{})
	if req.KeyWords != "" {
		// 搜索
		localDB = localDB.Where("name LIKE ?", "%"+req.KeyWords+"%")
	}

	if req.IsHot {
		localDB = localDB.Where(model.Goods{IsHot: true})
	}

	if req.IsNew {
		localDB = localDB.Where(model.Goods{IsNew: true})
	}

	if req.PriceMin > 0 {
		localDB = localDB.Where("shop_price>=?", req.PriceMin)
	}

	if req.PriceMax > 0 {
		localDB = localDB.Where("shop_price<=?", req.PriceMax)
	}

	if req.Brand > 0 {
		localDB = localDB.Where("brand_id=?", req.Brand)
	}

	// 通过category去查询商品
	var subQuery string
	if req.TopCategory > 0 {
		var category model.Category
		if result := global.DB.First(&category, req.TopCategory); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}

		if category.Level == 1 {
			subQuery = fmt.Sprintf("select id from category WHERE parent_category_id in (select id from category WHERE parent_category_id=%d)", req.TopCategory)
		} else if category.Level == 2 {
			subQuery = fmt.Sprintf("select id from category WHERE from parent_category_id=%d", req.TopCategory)
		} else if category.Level == 3 {
			subQuery = fmt.Sprintf("select id from category WHERE id=%d", req.TopCategory)
		}

		localDB = localDB.Where(fmt.Sprintf("category_id in (%s)", subQuery))
	}

	var count int64
	localDB.Count(&count)
	goodsListResponse.Total = int32(count)

	result := localDB.Preload("Category").Preload("Brands").Scopes(model.Paginate(int(req.Pages), int(req.PagePerNums))).Find(&goods)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, good := range goods {
		GoodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &GoodsInfoResponse)
	}

	return goodsListResponse, nil
}

```
### 单元测试
```go
// goods_srv/test/goods.go


package main

import (
	"context"
	"fmt"
	"mxshop_srvs/goods_srv/proto"
)

func TestGetGoodsList() {
	rsp, err := brandClient.GoodsList(context.Background(), &proto.GoodsFilterRequest{
		TopCategory: 135200,
		PriceMin: 60,
		// KeyWords: "三都港",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, good := range rsp.Data {
		fmt.Println(good.Name, good.ShopPrice)
	}
}
```
## 获取商品
### 获取商品列表
```go
// goods_srv/handler/goods.go


func (s *GoodsServer) BatchGetGoods(ctx context.Context, req *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
		goodsListResponse := &proto.GoodsListResponse{}
		var goods []model.Goods
		result := global.DB.Where(req.Id).Find(&goods)

		for _, good := range goods {
			goodsInfoResponse := ModelToResponse(good)
			goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
		}
		goodsListResponse.Total = int32(result.RowsAffected)
		return goodsListResponse, nil
	}
```
### 获取商品信息
```go
// goods_srv/handler/goods.go


func (s *GoodsServer) GetGoodsDetail(ctx context.Context, req *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
		var goods model.Goods

		if result := global.DB.First(&goods, req.Id); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品不存在")
		}
		goodsInfoResponse := ModelToResponse(goods)
		return &goodsInfoResponse, nil
	}
```
### 单元测试
```go
// goods_srv/test/goods.go


func TestBatchGetGoods() {
	rsp, err := brandClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: []int32{421, 422, 423},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, good := range rsp.Data {
		fmt.Println(good.Name, good.ShopPrice)
	}
}

func TestGetGoodsDetail() {
	rsp, err := brandClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: 421,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Name)
}
```
### 增加商品
```go
// goods_srv/hadnler/goods.go


func (s *GoodsServer) CreateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
		var category model.Category
		if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
		}

		var brand model.Brands
		if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
		}

		goods := model.Goods{
			Brands: brand,
			BrandsID: brand.ID,
			Category: category,
			CategoryID: category.ID,
			Name: req.Name,
			GoodsSn: req.GoodsSn,
			MarketPrice: req.MarketPrice,
			ShopPrice: req.ShopPrice,
			GoodsBrief: req.GoodsBrief,
			ShipFree: req.ShipFree,
			Images: req.Images,
			DescImages: req.DescImages,
			GoodsFrontImage: req.GoodsFrontImage,
			IsNew: req.IsNew,
			IsHot: req.IsHot,
			OnSale: req.OnSale,
		}

		global.DB.Save(&goods)
		return &proto.GoodsInfoResponse{
			Id: goods.ID,
		}, nil
	}
```
### 删除商品
```go
// goods_srv/hadnler/goods.go


func (s *GoodsServer) DeleteGoods(ctx context.Context, req *proto.DeleteGoodsInfo) (*emptypb.Empty, error) {
		if result := global.DB.Delete(&model.Goods{}, req.Id); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品不存在")
		}
		return &emptypb.Empty{}, nil
	}
```
### 修改商品
```go
// goods_srv/hadnler/goods.go


func (s *GoodsServer) UpdateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*emptypb.Empty, error) {
		var goods model.Goods

		if result := global.DB.First(&goods, req.Id); result.RowsAffected == 0 {
			return nil, status.Error(codes.NotFound, "商品不存在")
		}

		var category model.Category
		if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
		}

		var brand model.Brands
		if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
		}

		goods.Brands = brand
		goods.BrandsID = brand.ID
		goods.Category = category
		goods.CategoryID = category.ID
		goods.Name = req.Name
		goods.GoodsSn = req.GoodsSn
		goods.ShopPrice = req.ShopPrice
		goods.GoodsBrief = req.GoodsBrief
		goods.ShipFree = req.ShipFree
		goods.Images = req.Images
		goods.DescImages = req.DescImages
		goods.IsNew = req.IsNew
		goods.IsHot = req.IsHot
		goods.OnSale = req.OnSale

		global.DB.Save(&goods)
		return &emptypb.Empty{}, nil
	}
```
