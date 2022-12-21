// Package model /*
/*
  @Author:      吴贤权
  @Date:        2022/10/3
  @Description: 教师身份实体
*/
package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(15)" json:"identity"`
	Number   string `gorm:"column:number;type:varchar(10)" json:"number"`
}

func (t *Teacher) TableName() string {
	return "teacher"
}
