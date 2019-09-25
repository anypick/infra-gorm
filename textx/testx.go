package testx

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra-gorm"
	"github.com/anypick/infra/base/props"
	"github.com/anypick/infra/base/props/container"
	"github.com/jinzhu/gorm"
	"time"
)

type SysBaseDto struct {
	Id             int64     `gorm:"primary_key"`      // 主键ID
	ObjectVersion  int64     `gorm:"object_version"`   // 版本控制
	CreateBy       int64     `gorm:"create_by"`        // 创建者
	LastUpdateBy   int64     `gorm:"last_update_by"`   // 更新者
	LastUpdateDate time.Time `gorm:"last_update_date"` // 更新时间
	CreateTime     time.Time `gorm:"create_date"`      // 创建时间
	Status___      uint                                // 当前实体类的状态，用于更新或者插入， 0更新，1插入
}

type SysMenus struct {
	SysBaseDto
	Code     string `gorm:"code"`
	Name     string `gorm:"name"`
	Priority int64  `gorm:"priority"`
	Icon     string `gorm:"icon"`
	Parent   int64  `gorm:"parent"`
}

func init() {
	infra.Register(&container.YamlStarter{})
	infra.Register(&basegorm.GormStarter{})

	infra.Register(&infra.BaseInitializerStarter{})

	source := props.NewYamlSource("./config.yml")
	app := infra.New(*source)
	app.Start()
}

// 测试查询
func GormC(db *gorm.DB) (menus []SysMenus, err error) {
	db.Find(&menus)
	return
}
