package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"time"
)

type MeritTopic struct {
	Id       int64     `form:"-"`
	ParentId int64     `orm:"null"`
	Title    string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Choose   string    `orm:"null"`
	Content  string    `orm:"null"`
	Mark     string    `orm:"null"` //设置分数
	Url      string    `orm:"null"`
	Created  time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index","auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(MeritTopic)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/merit.db", 10)
}

func AddMeritTopic(pid int64, title, choose, content, mark string) (id int64, err error) {
	o := orm.NewOrm()
	topic := &MeritTopic{
		ParentId: pid,
		Title:    title,
		Choose:   choose,
		Content:  content,
		Mark:     mark,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	// qs := o.QueryTable("category") //不知道主键就用这个过滤操作
	id, err = o.Insert(topic)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetMeritTopic(pid int64) ([]*MeritTopic, error) {
	o := orm.NewOrm()
	topics := make([]*MeritTopic, 0)
	// category := new(MeritTopic)
	qs := o.QueryTable("merit_topic")                 //这个表名MeritTopic需要用驼峰式，
	_, err := qs.Filter("parentid", pid).All(&topics) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	return topics, err
}
