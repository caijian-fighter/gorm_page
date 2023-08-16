package v0

import (
	"fmt"
	"github.com/caijian-fighter/gorm_page"
	"reflect"
	"testing"
)

/**
 * @Author caicaijian
 * @Date: 2023/4/16 19:22:00
 * @Desc:
 */

// 反射创建对应的model的数组
func TestReflectModel(t *testing.T) {
	var i interface{}
	model := gorm_page.User{}
	i = model

	ty := reflect.TypeOf(i)      // 获取类型
	fmt.Printf("type: %v\n", ty) // type: gorm_page.User

	tySlice := reflect.SliceOf(ty)    // 获取该类型的slice类型
	fmt.Printf("type: %v\n", tySlice) // type: []gorm_page.User

	zero := reflect.Zero(tySlice)                   // 初始化该类型得到值
	fmt.Printf("value: %#v, type:%T\n", zero, zero) // value: []gorm_page.User(nil), type:reflect.Value

	inter := zero.Interface()                          // 转化成原有类型的 slice
	fmt.Printf("value: %#v, type:%T \n", inter, inter) // value:[]gorm_page.User(nil), type:[]gorm_page.User

	list := reflect.Zero(reflect.SliceOf(reflect.TypeOf(i))).Interface()
	fmt.Printf("list value: %#v, type:%T \n", list, list) // value:[]gorm_page.User(nil), type:[]gorm_page.User

	list = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(i)), 0, 0).Interface()
	fmt.Printf("list value: %#v, type:%T \n", list, list) // value:[]gorm_page.User(nil), type:[]gorm_page.User
}
