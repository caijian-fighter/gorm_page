package gorm_page

/**
 * @Author caicaijian
 * @Date: 2023/4/16 18:04:00
 * @Desc:
 */

type PageInfo struct {
	page      string `gorm:"page"`
	page_size string `gorm:"page_size"`
}

type ResponseInfo struct {
	Code int64  `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}
