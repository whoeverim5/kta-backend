// Package test /*
/*
  @Author:      吴贤权
  @Date:        2022/11/25 17:26
  @Description:
*/
package test

import (
	"log"
	"reflect"
	"testing"

	"ktaexam/pkg/dao"
)

func TestQueryCode(t *testing.T) {
	key := dao.Register + ":" + "1192445074@qq.com" + ":" + "732756"
	err := dao.RDB.Get(dao.Ctx, key).Err()
	if err != nil {
		log.Println("err != nil")
		log.Println(err, reflect.TypeOf(err))
	} else {
		log.Println("err == nil")
	}
}
