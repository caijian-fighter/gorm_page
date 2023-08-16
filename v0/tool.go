package v0

import (
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

/**
 * @Author caicaijian
 * @Date: 2023/4/16 14:49:00
 * @Desc:
 */

/**
 * @Author caicaijian
 * @Date: 2023/4/16 18:37:00
 * @Desc:
 */

func Post(body io.Reader, url string) (string, error) {

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}

// 参考 https://gorm.io/zh_CN/docs/scopes.html  gorm 官方分页实例 详情
func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
