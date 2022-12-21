// Package dao /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:44
 * @Description: 数据库连接对象层
**/
package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ktaexam/pkg/config"
)

// MySql
var (
	username  = config.Yaml.GetString("db.mysql.username")
	mPassword = config.Yaml.GetString("db.mysql.password")
	host      = config.Yaml.GetString("db.mysql.host")
	port      = config.Yaml.GetString("db.mysql.port")
	dbname    = config.Yaml.GetString("db.mysql.dbname")
	charset   = config.Yaml.GetString("db.mysql.charset")
	parseTime = config.Yaml.GetString("db.mysql.parseTime")
	loc       = config.Yaml.GetString("db.mysql.loc")
)

// Redis
var (
	address    = config.Yaml.GetString("db.redis.address")
	rPassword  = config.Yaml.GetString("db.redis.password")
	rDefaultDb = config.Yaml.GetInt("db.redis.default-db")
)

// Ctx Redis上下文
var Ctx = context.Background()

// 设置Redis的key
var (
	Register = "user:register"
	Code     = "code"
)

// RedisExpire 设置Redis键值的过期时间
var (
	RedisExpire = time.Minute * 5
)

var DB *gorm.DB = InitMysql()
var RDB *redis.Client = InitRedis()

func InitMysql() *gorm.DB {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=%v&loc=%v",
		username,
		mPassword,
		host,
		port,
		dbname,
		charset,
		parseTime,
		loc)

	db, err := gorm.Open(mysql.Open(dsn), config.GormConfig())
	if err != nil {
		panic(err)
	}

	return db
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: rPassword,
		DB:       rDefaultDb,
	})
	return rdb
}
