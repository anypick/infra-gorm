package basegorm

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra/base/props"
	"github.com/anypick/infra/base/props/container"
	"path/filepath"
	"runtime"
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

func InitT() {
	infra.Register(&container.YamlStarter{})
	Init()
	infra.Register(&infra.BaseInitializerStarter{})

	source := props.NewYamlSource(GetCurrentFilePath("./testx/config.yml", 0))
	app := infra.New(*source)
	app.Start()
}


// 获取文件的绝对路径
func GetCurrentFilePath(fileName string, skip int) string {
	_, file, _, _ := runtime.Caller(skip)
	// 解析出文件路径
	dir := filepath.Dir(file)
	return filepath.Join(dir, fileName)
}