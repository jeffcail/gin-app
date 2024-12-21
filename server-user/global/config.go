// 公共的全局配置 package
// GoAppConfig 框架应用配置
// Http http服务配置
// Websocket Websocket配置
// Grpc Grpc配置
// Mysql 数据库配置
// Redis redis配置

package global

import (
	"github.com/mazezen/ginframe/server-common/nacosRF"
	"time"
)

var Config *GlobalConfig

// GlobalConfig 框架应用配置
type GlobalConfig struct {
	IsEnableGOMAXPROCS bool                 `json:"is_enable_gomaxprocs"`
	GinAppDebug        string               `json:"gin_app_debug"`
	GinAppName         string               `json:"gin_app_name"`
	Http               Http                 `json:"http"`
	Websocket          Websocket            `json:"websocket"`
	Grpc               Grpc                 `json:"grpc"`
	Nacos              *nacosRF.NacosConfig `json:"nacos"`
	Rpc                Rpc                  `json:"rpc"`
	Mysql              Mysql                `json:"mysql"`
	Redis              Redis                `json:"redis"`
	Mongo              Mongo                `json:"mongo"`
	LevelDb            LevelDb              `json:"level_db"`
	Elastic            Elastic              `json:"elastic"`
	Aes                Aes                  `json:"aes"`
	Jwt                Jwt                  `json:"jwt"`
}

// Http http服务配置
type Http struct {
	BindPort string `json:"bind_port"`
	LogPath  string `json:"log_path"`
}

// Websocket Websocket服务配置
type Websocket struct {
	BindPort string `json:"bind_port"`
}

// Grpc Grpc配置
type Grpc struct {
	Port string `json:"port"`
}

// Rpc Rpc配置
type Rpc struct {
	Listener string `json:"listener"`
}

// Mysql 数据库配置
type Mysql struct {
	DbDsn        string `json:"db_dsn"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

// Redis redis配置
type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

// Mongo mongo配置
type Mongo struct {
	Addr string `json:"addr"`
}

type LevelDb struct {
	Path string `json:"path"`
}

type Elastic struct {
	Url string `json:"url"`
}

type Aes struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type Jwt struct {
	Expire time.Duration `json:"expire"`
	Secret string        `json:"secret"`
}
