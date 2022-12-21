// Package dao /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:45
 * @Description: 用户数据连接
**/
package dao

import (
	"ktaexam/pkg/model"
)

func GetUser(account string) (*model.User, error) {
	user := new(model.User)
	query := "id, password, user_identity, salt, user_id, teacher_identity, student_identity, email"
	err := DB.Select(query).
		Where("account", account).
		Find(&user).
		Error
	return user, err
}

func SetEmailCodeCache(code, email string) error {
	key := Register + ":" + email + ":" + code
	err := RDB.Set(Ctx, key, code, RedisExpire).Err()
	return err
}

func CreateNewUser(user *model.User) error {
	err := DB.Create(user).
		Error
	return err
}

func QueryCodeFromRedis(code, email string) bool {
	key := Register + ":" + email + ":" + code
	value := RDB.Get(Ctx, key).Val()
	if value != code {
		return false
	}
	return true
}

func JudgeUserExisted(account string) (*model.User, error) {
	user := new(model.User)
	query := "id, account"
	err := DB.Select(query).
		Where("account", account).
		Find(&user).
		Error

	return user, err
}
