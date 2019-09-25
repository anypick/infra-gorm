package testx

import (
	"fmt"
	"github.com/anypick/infra-gorm"
	"testing"
)

func TestGormC(t *testing.T) {
	menus, err := GormC(basegorm.GetDb())
	fmt.Println(menus, err)
}


func BenchmarkGormC(b *testing.B) {
	for i:=0; i< b.N; i++ {
		GormC(basegorm.GetDb())
	}
}