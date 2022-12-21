// Package model /*
/*
  @Author:      吴贤权
  @Date:        2022/10/3
  @Description: 用户实体
*/
package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account         string `gorm:"column:account;type:varchar(32);" json:"account"`
	Password        string `gorm:"column:password;type:varchar(32)" json:"password"`
	Salt            string `gorm:"column:salt;type:varchar(32)" json:"salt"`
	UserIdentity    int8   `gorm:"column:user_identity;type:tinyint(1)" json:"status"`
	UserId          string `gorm:"column:user_id;type:varchar(36)" json:"user_id"`
	StudentIdentity string `gorm:"column:student_identity;type:varchar(15)" json:"student_identity"`
	TeacherIdentity string `gorm:"column:teacher_identity;type:varchar(15)" json:"teacher_identity"`
	Gender          int8   `gorm:"column:gender;type:tinyint(1)" json:"gender"`
	Age             int8   `gorm:"column:age;type:tinyint(3)" json:"age"`
	Email           string `gorm:"column:email;type:varchar(100)" json:"email"`
	Institution     string `gorm:"column:institution;type:varchar(255)" json:"institution"`
}

func (u *User) TableName() string {
	return "user"
}
