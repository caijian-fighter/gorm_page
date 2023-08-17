package do

import (
	"errors"
	"github.com/caijian-fighter/gorm_page/dal"
	"github.com/caijian-fighter/gorm_page/model"
	"gorm.io/gorm"
	"testing"
)

/**
 * @Author caicaijian
 * @Date: 2023/8/17 16:40:00
 * @Desc:
 */

func TestPage(t *testing.T) {
	p := new(Page[model.LhSendingOrder])
	p.PageNo = req.Page
	p.PageSize = req.Limit
	orderQuery := dal.GetQuery().LhSendingOrder
	q := orderQuery.WithContext(c).
		Where(orderQuery.UserID.Eq(userID)).
		Where(orderQuery.ExamStatus.Eq(req.ExamStatus)).DO // 里面嵌套的DO类型

	err = p.SelectPages(q)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {

	}

	if err != nil {

	}
}
