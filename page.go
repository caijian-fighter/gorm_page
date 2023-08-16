package gorm_page

import (
	"gorm.io/gorm"
)

/**
 * @Author caicaijian
 * @Date: 2023/4/23 1:11:00
 * @Desc:
 */

type Page[T any] struct {
	PageNo   int64 `json:"page_no"`
	PageSize int64 `json:"page_size"`
	Total    int64 `json:"total"` // 总记录数
	Pages    int64 `json:"pages"` // 总页数
	List     []T   `json:"list"`  // 实际的list数据
}

func (page *Page[T]) SelectPages(query *gorm.DB) (e error) {
	e = nil
	var model T
	query.Model(&model).Count(&page.Total)
	if page.Total == 0 {
		page.List = []T{}
		return
	}

	e = query.Model(&model).Scopes(Paginate(page)).Last(&page.List).Error

	return
}

func Paginate[T any](page *Page[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.PageNo <= 0 {
			page.PageNo = 0
		}
		switch {
		case page.PageSize > 100:
			page.PageSize = 100
		case page.PageSize <= 0:
			page.PageSize = 10
		}
		page.Pages = page.Total / page.PageSize
		if page.Total%page.PageSize != 0 {
			page.Pages++
		}
		p := page.PageNo
		if page.PageNo > page.Pages {
			p = page.Pages
		}
		size := page.PageSize
		offset := int((p - 1) * size)
		return db.Offset(offset).Limit(int(size))
	}
}
