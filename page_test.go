package gorm_page

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

/**
 * @Author caicaijian
 * @Date: 2023/4/23 01:47:00
 * @Desc:
 */
var dsn = "root:123456@tcp(192.168.47.166:3306)/shared_usercenter?charset=utf8mb4&parseTime=True&loc=Local"
var db = getDB()

// 初始化DB的方法
func getDB() *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("err: %v", err)
	}

	return db
}

func TestPage(t *testing.T) {
	p := Page[User]{
		PageNo:   1,
		PageSize: 10,
	}
	wrapper := make(map[string]interface{}, 0)

	wrapper["status"] = 1

	err := p.SelectPages(db.Debug())
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	fmt.Printf("%#v", p)
}
