package config

import (
	"flag"
	"jianji-server/config/types"
	"log"
	"strings"

	"github.com/spf13/viper"
)

const DefaultConfigFile string = "config.yml"

var (
	Server   types.Server
	Postgres types.Postgres
	Zap      types.Zap
	Redis    types.Redis
	Email    types.Email
	MaxMind  types.MaxMind
)

func init() {
	var configPath string

	v := viper.New()
	v.SetConfigFile(DefaultConfigFile)

	flag.StringVar(&configPath, "c", "", "choose config file.")
	flag.Parse()

	if configPath == "" {
		configPath = DefaultConfigFile
	}

	//也可以通过环境变量加载配置
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) //环境变量使用连字符_，替换为.

	//处理配置文件加载错误
	if err := v.ReadInConfig(); err != nil {
		log.Printf("加载配置文件 %s 失败\n", configPath)
		log.Panicln(err)
	}

	//处理配置文件读取错误
	var configMap = map[string]any{
		"Server":   &Server,
		"Postgres": &Postgres,
		"Zap":      &Zap,
		"Redis":    &Redis,
		"Email":    &Email,
		"MaxMind":  &MaxMind,
	}

	for key, val := range configMap {
		if err := v.UnmarshalKey(key, val); err != nil {
			log.Printf("读取配置文件 %s 失败\n", configPath)
			log.Panicln(err)
		}
	}

	log.Printf("配置文件 %s 加载完成\n", configPath)
}
