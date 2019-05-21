package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Keyword struct {
	Id         		 	int
	ipad_id				string
	words1				string
	words2				string
	words3				string
	limit_len			int
	get_all				bool
	contact				string
	only_friends		bool
	only_not_friends	bool
	pushing				bool
	User				string
	Token				string
}

func (t *Keyword) TableName() string {
	return TableName("keyword")
}

func (t *Keyword) Update(fields ...string) error {
	if t.ipad_id == "" {
		return fmt.Errorf("ipad_id不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

// 添加关键词
func KeywordAdd(obj *Keyword) (int64, error) {
	if obj.ipad_id == "" {
		return 0, fmt.Errorf("ipad_id 不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

func KeywordGetById(id int) (*Keyword, error) {
	obj := &Keyword{
		Id: id,
	}

	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func KeywordDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("keyword")).Filter("id", id).Delete()
	return err
}

func KeywordGetList(page, pageSize int) ([]*Keyword, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Keyword, 0)

	query := orm.NewOrm().QueryTable(TableName("keyword"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
