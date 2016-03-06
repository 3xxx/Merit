package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"time"
)

type Category struct {
	Id       int64     `form:"-"`
	ParentId int64     `orm:"null"`
	Title    string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Mark     string    `orm:"null"`                                              //设置分数
	Url      string    `orm:"null"`
	List     string    `orm:"null"` //选择项
	ListMark string    `orm:"null"` //每个选项对应的分数——这种驼峰式命名似乎在后续添加中不能自动建表
	Created  time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index","auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Category)) //, new(Article)
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/merit.db", 10)
}

func AddCategory(pid int64, title, mark, url, list, listmark string) (id int64, err error) {
	o := orm.NewOrm()
	cate := &Category{
		ParentId: pid,
		Title:    title,
		Mark:     mark,
		Url:      url,
		List:     list,
		ListMark: listmark,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	// qs := o.QueryTable("category") //不知道主键就用这个过滤操作
	id, err = o.Insert(cate)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetPids(pid int64) ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	var err error
	//这里进行过滤
	_, err = qs.Filter("ParentId", pid).All(&cates)
	// _, err = qs.OrderBy("-created").All(&cates)
	// _, err := qs.All(&cates)
	return cates, err
}

func GetCategory(id int64) (*Category, error) {
	o := orm.NewOrm()
	// cate := &Category{Id: id}
	category := new(Category)
	qs := o.QueryTable("category")
	err := qs.Filter("id", id).One(category)
	if err != nil {
		return nil, err
	}
	return category, err
}
