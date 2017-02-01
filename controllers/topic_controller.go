package controllers

import (
	"github.com/astaxie/beego"
	"blog/models"
	"fmt"
	"strconv"
)

type TopicController struct {
	beego.Controller
}

func (p *TopicController) Get(){
	p.Data["Topics"] = models.GetAllTopics()
	fmt.Print("1")
	p.TplName = "topic.html"
}

func (p *TopicController) Add(){
	p.TplName = "topic_add.html"
}

func (p *TopicController) Post() {
	fmt.Println(p.Input().Get("title"), p.Input().Get("content"))
	err := models.AddTopic(p.Input().Get("title"), p.Input().Get("content"))
	if err == nil {
		p.Redirect("/topic", 301)
	} else {
		p.Redirect("/home", 301)
	}
}

func (p *TopicController) Update(){
	p.Data["Topic"] = models.GetTopicById(p.Input().Get("tid"))
	p.TplName = "topic_update.html"
}

func (p *TopicController)  Detail(){
	fmt.Println(p.Input().Get("tid"))
	p.Data["Topic"] = models.GetTopicById(p.Input().Get("tid"))
	p.TplName = "topic_detail.html"
}

func (p *TopicController) Modify(){
	id, err := strconv.ParseInt(p.Input().Get("id"), 10, 64)
	if err != nil {

	}
	topic := &models.Topic{
		Id: id,
		Title: p.Input().Get("title"),
		Content: p.Input().Get("content"),
	}
	models.UpdateTopic(topic)
	fmt.Println(11111)
	p.Redirect("/topic", 301)
}