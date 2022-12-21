// Package config /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午3:12
 * @Description: gorm配置
**/
package config

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
}
