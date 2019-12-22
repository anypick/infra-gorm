package basegorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

func TestGetDb(t *testing.T) {
	// 初始化测试需要的条件
	InitT()
	db := GetDb()
	if menus, err := GormFind(db); err != nil {
		fmt.Println("errr?:", err)
	} else {
		fmt.Println(menus)
	}
}

// 测试查询
func GormFind(db *gorm.DB) (menus []SysMenus, err error) {
	db.Find(&menus)
	db.Update()
	return
}
