package do

import (
	"gorm.io/gen"
)

/**
 * @Author caicaijian
 * @Date: 2023/8/17 16:38:00
 * @Desc:
 */

type Page[T any] struct {
	PageNo   int64 `json:"page_no"`
	PageSize int64 `json:"page_size"`
	Total    int64 `json:"total"` // 总记录数
	Pages    int64 `json:"pages"` // 总页数
	List     []*T  `json:"list"`  // 实际的list数据
}

func (page *Page[T]) SelectPages(do gen.DO) (e error) {
	page.Total, e = do.Count()
	if e != nil {
		return e
	}
	if page.Total == 0 {
		page.List = []*T{}
		return
	}

	e = do.Scopes(Paginate(page)).Scan(&page.List)

	return
}

func Paginate[T any](page *Page[T]) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
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
		return dao.Offset(offset).Limit(int(size))
	}
}
