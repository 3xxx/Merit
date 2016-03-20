package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	// "strings"
	"time"
)

type MeritTopic struct {
	Id       int64     `form:"-"`
	ParentId int64     `orm:"null"`
	UserId   int64     `orm:"null"`
	Title    string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Choose   string    `orm:"null"`
	Content  string    `orm:"null"`
	Mark     string    `orm:"null"` //设置分数
	Created  time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index","auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(MeritTopic)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/merit.db", 10)
}

//用户添加价值
func AddMeritTopic(pid int64, uname, title, choose, content, mark string) (id int64, err error) {
	//先由uname取得uid
	user := GetUserByUsername(uname)

	o := orm.NewOrm()
	topic := &MeritTopic{
		ParentId: pid,
		UserId:   user.Id,
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

//根据父级价值id和用户id取得所有价值——返回数量和分值
func GetMeritTopic(pid, uid int64) (topics []*MeritTopic, numbers, marks int, err error) {
	o := orm.NewOrm()
	topics = make([]*MeritTopic, 0)
	// category := new(MeritTopic)
	qs := o.QueryTable("merit_topic")                                      //这个表名MeritTopic需要用驼峰式，
	_, err = qs.Filter("parentid", pid).Filter("userid", uid).All(&topics) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, 0, 0, err
	}
	for _, v := range topics {
		mark, err := strconv.Atoi(v.Mark)
		if err != nil {
			return nil, 0, 0, err
		}
		marks = marks + mark
	}
	numbers = len(topics)
	return topics, numbers, marks, err
}

//取得用户id的所有价值
func GetAllMeritTopic(uid int64) ([]*MeritTopic, error) {
	o := orm.NewOrm()
	topics := make([]*MeritTopic, 0)
	// category := new(MeritTopic)
	qs := o.QueryTable("merit_topic")               //这个表名MeritTopic需要用驼峰式，
	_, err := qs.Filter("userid", uid).All(&topics) //而这个字段userid为何又不用呢
	if err != nil {
		return nil, err
	}
	return topics, err
}

//根据topicid取得topic
func GetMeritTopicbyId(tid string) (*MeritTopic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	merittopic := new(MeritTopic)
	qs := o.QueryTable("merit_topic")
	err = qs.Filter("id", tidNum).One(merittopic)
	if err != nil {
		return nil, err
	}
	return merittopic, err
}

//删除merittopic
func DeletMeritTopic(id string) error { //应该在controllers中显示警告
	tidNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	merittopic := MeritTopic{Id: tidNum}
	if o.Read(&merittopic) == nil {
		_, err = o.Delete(&merittopic)
		if err != nil {
			return err
		}
	}
	// _, err = o.Delete(&topic) //这句为何重复？
	return err
}

//修改merittopic
func ModifyMeritTopic(tid, title, choose, content, mark string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	merittopic := &MeritTopic{Id: tidNum}
	if o.Read(merittopic) == nil {
		merittopic.Title = title
		merittopic.Choose = choose
		merittopic.Content = content
		merittopic.Mark = mark
		merittopic.Updated = time.Now()
		_, err = o.Update(merittopic)
		if err != nil {
			return err
		}
	}
	return err
}

//管理员取得所有价值
