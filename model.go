package gorm_page

/**
 * @Author caicaijian
 * @Date: 2023/4/16 19:49:00
 * @Desc:
 */

type User struct {
	UserId int64  `gorm:"user_id"`
	Email  string `gorm:"email"`
	Mobile string `gorm:"mobile"`
}

func (u *User) TableName() string {
	return "user"
}
