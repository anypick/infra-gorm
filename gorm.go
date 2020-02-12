package basegorm

/**
  初始化*gorm.DB
*/

import (
	"fmt"
	"github.com/anypick/infra"
	"github.com/anypick/infra-gorm/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	return db
}

type GormStarter struct {
	infra.BaseStarter
}

func (m *GormStarter) Setup(ctx infra.StarterContext) {
	var (
		conf = ctx.Yaml()[config.DefaultPrefix].(*config.GormConfig)
		err  error
	)
	if db, err = gorm.Open(conf.DriverName, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s",
		conf.Username, conf.Password, conf.IpAddr, conf.Port, conf.Database, conf.Params)); err != nil {
		panic(err)
	}

	db.SetLogger(GormLogrus{})

	// 获取原生db，配置连接池信息
	sqldb := db.DB()
	sqldb.SetConnMaxLifetime(conf.GetDurationDefault("ConnMaxLifetime", time.Duration(7*24*60*60)))
	sqldb.SetMaxOpenConns(conf.GetIntByDefault("MaxOpenConn", 1000))
	sqldb.SetMaxIdleConns(conf.GetIntByDefault("MaxIdeConn", 1000))
}
