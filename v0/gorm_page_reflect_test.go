package v0

import (
	"fmt"
	"github.com/caijian-fighter/gorm_page"
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
	u := gorm_page.User{}

	query := gorm_page.db.Debug().Where("user_id > ?", 1)
	err := u.SelectPageList(&p, query) //err := SelectPage(&p, query, User{})
	if err != nil {
		t.Fatalf("page Err:%v", err)
	}
	fmt.Printf("%#v\n", p)
}

func (u *gorm_page.User) SelectPageList(p *Page, query *gorm.DB) error {
	err := SelectPage(p, query, gorm_page.User{})
	return err
}
