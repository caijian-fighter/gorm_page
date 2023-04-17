package gorm_page

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

/**
 * @Author caicaijian
 * @Date: 2023/4/16 19:34:00
 * @Desc:
 */

func TestGormPageReflectTest(t *testing.T) {
	p := Page{
		PageNo:   1,
		PageSize: 10,
	}
	u := User{}

	query := db.Debug().Where("user_id > ?", 1)
	err := u.SelectPageList(&p, query) //err := SelectPage(&p, query, User{})
	if err != nil {
		t.Fatalf("page Err:%v", err)
	}
	fmt.Printf("%#v\n", p)
}

func (u *User) SelectPageList(p *Page, query *gorm.DB) error {
	err := selectPage(p, query, User{})
	return err
}
