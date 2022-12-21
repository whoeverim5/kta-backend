// Package config /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:43
 * @Description: 读取application.yaml配置类
**/
package config

import (
	"github.com/spf13/viper"
)

const (
	configurationName = "application-dev"
	configurationPath = "./resources"
	configurationType = "yml"
)

var Yaml *viper.Viper

func init() {
	Yaml = viper.New()
	Yaml.SetConfigType(configurationType)
	Yaml.SetConfigName(configurationName)
	Yaml.AddConfigPath(configurationPath)
	if err := Yaml.ReadInConfig(); err != nil {
		panic(err)
	}
}
