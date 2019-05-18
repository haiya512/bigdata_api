package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// pages的很多model方法一个是为了展示，还有是为了存储方便

const (
	TASK_SUCCESS = 0  // 任务执行成功
	TASK_ERROR   = -1 // 任务执行出错
	TASK_TIMEOUT = -2 // 任务执行超时
)

type Message struct {
	Id          	 	int
	Content     	 	string
	ContentHash	 	string
	Nickname      	 	string
	Wxid     		 	string
	PicUrls   		 	string
	IpadId		  	 	string
	Data     			string
	IsDownloaded		bool
	CreateTime      	time.Time
	IsAddToPushed  		bool
	PushHistory			string
	PushHistory2		string
	Repeated			bool
	Fenlei				string
	Fenleib				string
}

func (t *Message) TableName() string {
	return TableName("message")
}

func MessageQ() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Message))
}

func (t *Message) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

// 内容是否为空
func CheckMessageContent(content string) bool {
	exist := MessageQ().Filter("content", content).Exist()
	return exist
}

// 微信昵称是否为空
func CheckNickname(nickname string) bool {
	exist := MessageQ().Filter("nickname", nickname).Exist()
	return exist
}

// wxid是否为空
func CheckWxid(wxid string) bool {
	exist := MessageQ().Filter("wxid", wxid).Exist()
	return exist
}


//添加message
func MessageAdd(mess Message) Message {
	o := orm.NewOrm()
	o.Insert(&mess)
	return mess
}

func MessageList(page, pageSize int, filters ...interface{}) ([]*Message, int64) {
	offset := (page - 1) * pageSize

	tasks := make([]*Message, 0)

	query := orm.NewOrm().QueryTable(TableName("message"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&tasks)

	return tasks, total
}

// 这个函数还没弄明白
func MessageResetGroupId(groupId int) (int64, error) {
	return orm.NewOrm().QueryTable(TableName("message")).Filter("group_id", groupId).Update(orm.Params{
		"group_id": 0,
	})
}

func MessageGetById(id int) (*Message, error) {
	mess := &Message{
		Id: id,
	}

	err := orm.NewOrm().Read(mess)
	if err != nil {
		return nil, err
	}
	return mess, nil
}

func MessageDel(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("message")).Filter("id", id).Delete()
	return err
}
