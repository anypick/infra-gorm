package config

/**
gorm需要的配置
 */


import (
	"fmt"
	"github.com/anypick/infra/base/props/container"
	"github.com/anypick/infra/utils/common"
	"reflect"
	"time"
)

const (
	DefaultPrefix = "gorm"
)

func init() {
	// 将配置添加到Yaml容器中，Prefix字段必须
	container.RegisterYamContainer(&GormConfig{Prefix: DefaultPrefix})
}

type GormConfig struct {
	Prefix          string                                 // yaml配置文件中每一项配置的前缀，用于获取配置
	DriverName      string        `yaml:"driverName"`      // 驱动名称
	IpAddr          string        `yaml:"ipAddr"`          // ip地址
	Port            string        `yaml:"port"`            // 端口
	Username        string        `yaml:"username"`        // 用户名
	Password        string        `yaml:"password"`        // 密码
	Database        string        `yaml:"database"`        // 数据库名称
	MaxOpenConn     int           `yaml:"maxOpenConn"`     // 最大连接数
	MaxIdeConn      int           `yaml:"maxIdeConn"`      // 最大等待连接
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"` // 连接最大存活时间
	Params          string        `yaml:"params"`          // 数据库参数: ?charset=utf8&parseTime=true, mysql的连接方式：user:password@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true
}

// 实现配置装配接口，主要的作用是装配配置，容器初始化的时候会被调用
func (c *GormConfig) ConfigAdd(config map[interface{}]interface{}) {
	c.DriverName = config["driverName"].(string)
	c.IpAddr = config["ipAddr"].(string)
	c.Port = fmt.Sprintf("%v", config["port"])
	c.Username = config["username"].(string)
	c.Password = config["password"].(string)
	c.Database = config["database"].(string)
	c.MaxOpenConn = config["maxOpenConn"].(int)
	c.MaxIdeConn = config["maxIdeConn"].(int)
	c.ConnMaxLifetime = time.Duration(config["connMaxLifetime"].(int))
	c.Params = config["params"].(string)
}


func (m GormConfig) GetStringByDefault(fieldName, defaultValue string) string {
	stringValue := reflect.ValueOf(m).FieldByName(fieldName).Interface().(string)
	if common.StrIsBlank(stringValue) {
		return defaultValue
	}
	return stringValue
}

func (m GormConfig) GetIntByDefault(fieldName string, defaultValue int) int {
	intValue := reflect.ValueOf(m).FieldByName(fieldName).Interface().(int)
	if intValue == 0 {
		return defaultValue
	}
	return intValue
}

func (m GormConfig) GetDurationDefault(fieldName string, defaultValue time.Duration) time.Duration {
	durationValue := reflect.ValueOf(m).FieldByName(fieldName).Interface().(time.Duration)
	if durationValue == 0 {
		return time.Second * defaultValue
	}
	return time.Second * durationValue
}
