// Package model /*
/*
  @Author:      吴贤权
  @Date:        2022/10/3
  @Description: 学生身份实体
*/
package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(15)" json:"identity"`
	Number   string `gorm:"column:number;type:varchar(10)" json:"number"`
}

func (s *Student) TableName() string {
	return "student"
}
