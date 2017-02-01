package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)


type Category struct {
	Id              int64
	Title           string
	Views           string
	TopicTime       time.Time
	TopicCount      int64
	TopicLastUserId int64
	Created         time.Time
	Updated         time.Time
	IsDeleted     	string
}

func init()  {
	orm.RegisterModel(new(Category))
}

func GetAllCategories() []Category{
	o := orm.NewOrm()
	var categories []Category
	num,err := o.Raw("select * from t01_category").QueryRows(&categories)
	if err != nil {
		fmt.Println(num, categories)
	}
	return categories
}