package basegorm

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra-gorm/config"
	"github.com/anypick/infra/base/props/container"
)

func Init() {
	container.Add(&config.GormConfig{Prefix: config.DefaultPrefix})
	infra.Register(&GormStarter{})
}