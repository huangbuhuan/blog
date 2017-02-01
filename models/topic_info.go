package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Topic struct {
	Id      int64
	Uid     int64
	Title   string
	Content string
	Attachment string
	Created time.Time
	Updated time.Time
	Views int64
	Author string
	ReplyTime time.Time
	ReplyCount int64
	ReplyLastUserId int64
	IsDeleted string
}

func init()  {
	orm.RegisterModel(new(Topic))
}

func AddTopic(title, content string) error{
	o := orm.NewOrm()
	now := time.Now()

	_, err := o.Raw("INSERT INTO t01_topic (title, content, created, updated, is_deleted ) VALUES (?, ?, ?, ?, '0')", title, content, now, now).Exec()

	if err != nil {
		beego.Error(err)
	}

	return  nil
}

func GetTopicById(id string) Topic{
	o := orm.NewOrm()
	var topic Topic
	err := o.Raw("select * from t01_topic where id = ?", id).QueryRow(&topic)
	topic.Views++
	if err != nil {

	}
	UpdateTopic(&topic)
	return topic
}

func GetAllTopics() []Topic {
	o := orm.NewOrm()
	var topics []Topic
	_, err := o.Raw("select * from t01_topic where is_deleted = '0' order by created desc").QueryRows(&topics)

	if err != nil {
		beego.Error(err)
	}

	fmt.Println(topics)
	return topics
}

func UpdateTopic(p *Topic) error {

	o := orm.NewOrm()

	if p.Views != 0 {
		o.Raw("UPDATE t01_topic SET title = ?, views = ? WHERE id =?", p.Title, p.Views, p.Id).Exec()
	} else {
		fmt.Print(p.Id, p.Content, p.Title)
		o.Raw("UPDATE t01_topic SET title = ?, content = ?, updated = ? WHERE id =?", p.Title, p.Content, time.Now(), p.Id).Exec()
	}

	return nil
}