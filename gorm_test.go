package basegorm

import (
	"fmt"
	"testing"
)

func TestGetDb(t *testing.T) {
	// 初始化测试需要的条件
	InitT()
	db := GetDb()
	fmt.Println(db.DB().Ping())
}
