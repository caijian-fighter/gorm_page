package v0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/caijian-fighter/gorm_page"
	"net/http"
	"testing"
)

// 查询实例
func TestGormOpenAndSearch(t *testing.T) {
	user := gorm_page.User{}
	db.Debug().Table("user").Where("user_id = ?", 1).Last(&user)
	fmt.Printf("%#v\n", user)
}

func TestHttpRequest(t *testing.T) {
	url := fmt.Sprintf("http://api.funny96.com")
	data := gorm_page.User{}
	stu, err := json.Marshal(&data)
	reader := bytes.NewReader(stu)
	s, err := Post(reader, url)
	if err != nil {
		t.Fatalf("err:%v", err)
	}
	fmt.Printf("Response:%#v \n", s)

	var resp ResponseInfo
	err = json.Unmarshal([]byte(s), &resp)
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	fmt.Printf("Response:%#v \n", resp)

}

func TestGormPagination(t *testing.T) {

	users := new([]gorm_page.User)
	url := "http://api.funny96.com"
	pageInfo := PageInfo{
		page:      "1",
		page_size: "10",
	}
	pi, _ := json.Marshal(pageInfo)
	r, err := http.NewRequest("POST", url, bytes.NewReader([]byte(pi)))
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	userModel := gorm_page.User{}
	_ = db.Debug().Model(&userModel).Where("user_id > ?", 1).Scopes(Paginate(r)).Last(users)

	fmt.Printf("user: %v\n", users)
}
