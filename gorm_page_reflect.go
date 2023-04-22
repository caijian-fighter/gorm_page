package gorm_page

import (
	"gorm.io/gorm"
	"reflect"
)

/**
 * @Author caicaijian
 * @Date: 2023/4/17 18:11:00
 * @Desc:
 */

type Page struct {
	PageNo   int64       `json:"page_no"`
	PageSize int64       `json:"page_size"`
	Total    int64       `json:"total"` // 总记录数
	Pages    int64       `json:"pages"` // 总页数
	List     interface{} `json:"list"`  // 实际的list数据
}

func GormPaginate(page *Page) func(db *gorm.DB) *gorm.DB {
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

func selectPage(p *Page, query *gorm.DB, model interface{}) (e error) {
	e = nil
	query.Model(&model).Count(&p.Total)
	if p.Total == 0 {
		p.List = []interface{}{}
		return
	}

	list := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(model)), 0, 0).Interface()
	//list := reflect.Zero(reflect.SliceOf(reflect.TypeOf(model))).Interface()
	e = query.Model(&model).Scopes(GormPaginate(p)).Last(&list).Error
	if e != nil {
		return
	}
	p.List = list

	return
}
