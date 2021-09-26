package configs

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"github.com/go-redis/redis/v8"
)

var (
	GlobalConfig *Config
)

type Config struct {
	AppConfig   *AppConfig
	HttpConfig  *bm.ServerConfig
	GrpcConfig  *warden.ServerConfig
	RedisConfig *redis.ClusterOptions
}

type AppConfig struct{}

// 初始化相关配置
func Init() {
	paladin.Init()
	//  读取配置文件
}
