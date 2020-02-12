package basegorm

import "github.com/sirupsen/logrus"

type GormLogrus struct {
}

func (GormLogrus) Print(v ...interface{}) {
	logrus.Info("Gorm ", v)
}
