package configs

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var AppConfigInstance *AppConfig

type AppConfig struct {
	RedisConfig RedisConfig `mapstructure:"redis" json:"redis" yaml:"redis"`
	MongoConfig MongoConfig `mapstructure:"mongo" json:"mongo" yaml:"redis"`
	StockConfig StockConfig `mapstructure:"stock" json:"stock" yaml:"stock"`
}

const configType = "yaml"

func (appConfig *AppConfig) LoadConfigByViper(paths ...string) *viper.Viper {
	v := viper.New()
	v.SetConfigType(configType)
	for _, path := range paths {
		vTemp := viper.New()
		vTemp.SetConfigFile(path)
		vTemp.SetConfigType(configType)
		err := vTemp.ReadInConfig() //读取配置
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		mergeChildSettings(vTemp, v) //合并配置
		vTemp.WatchConfig()          //监听配置变化
		vTemp.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file changed:", e.Name)
			mergeChildSettings(vTemp, v) //合并配置
			unmarshal(v, appConfig)
			//TODO 根据
		})
	}
	unmarshal(v, appConfig)
	return v
}

func mergeChildSettings(child, parent *viper.Viper) {
	settings := child.AllSettings()
	for k, val := range settings {
		parent.SetDefault(k, val)
	}
}

func unmarshal(v *viper.Viper, bindValue interface{}) {
	if err := v.Unmarshal(bindValue); err != nil {
		fmt.Println(err)
	}
}
