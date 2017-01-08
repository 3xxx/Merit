//这个作废，被拆成组织结构和定义价值
//链接数据库也要移走————
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
	List     string    `orm:"null"`                                              //选择项
	ListMark string    `orm:"null"`                                              //每个选项对应的分数——这种驼峰式命名似乎在后续添加中不能自动建表
	Created  time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index","auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Category)) //, new(Article)
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/merit.db", 10)
}

func AddCategory(pid int64, title, mark, list, listmark string) (id int64, err error) {
	//重复性检查
	o := orm.NewOrm()
	cate := &Category{
		ParentId: pid,
		Title:    title,
		Mark:     mark,
		List:     list,
		ListMark: listmark,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	id, err = o.Insert(cate)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//由父级id得到所有下级
func GetPids(pid int64) ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	var err error
	//这里进行过滤
	_, err = qs.Filter("ParentId", pid).All(&cates)
	return cates, err
}

//取到所有的价值结构
func GetAllCategory() ([]*Category, error) {
	o := orm.NewOrm()
	category := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&category)
	if err != nil {
		return nil, err
	}
	return category, err
}

//由id取得分类
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

//由分院名称取得分院
func GetCategorybyname(title string) (*Category, error) {
	o := orm.NewOrm()
	// cate := &Category{Id: id}
	category := new(Category)
	qs := o.QueryTable("category")
	err := qs.Filter("title", title).One(category)
	if err != nil {
		return nil, err
	}
	return category, err
}

//由分院id和科室 名称取得科室
func GetCategorybyidtitle(id int64, title string) (*Category, error) {
	o := orm.NewOrm()
	// cate := &Category{Id: id}
	category := new(Category)
	qs := o.QueryTable("category")
	err := qs.Filter("parentid", id).Filter("title", title).One(category)
	if err != nil {
		return nil, err
	}
	return category, err
}

//修改category
func Modifyjson(id int64, title, mark, url, list, listmark string) error {
	o := orm.NewOrm()
	meritcategory := &Category{Id: id}
	var err error
	if o.Read(meritcategory) == nil {
		meritcategory.Title = title
		meritcategory.Mark = mark
		meritcategory.List = list
		meritcategory.ListMark = listmark
		meritcategory.Updated = time.Now()
		_, err = o.Update(meritcategory)
		if err != nil {
			return err
		}
	}
	return err
}

//删除价值结构
func Deletejson(id int64) error { //应该在controllers中显示警告
	o := orm.NewOrm()
	meritcategory := Category{Id: id}
	if o.Read(&meritcategory) == nil {
		_, err := o.Delete(&meritcategory) //删除分院
		if err != nil {
			return err
		}
	}
	//查询下级
	var categories []Category
	_, err := o.QueryTable("Category").Filter("parentid", id).All(&categories, "Id")
	if err != nil {
		return err
	} else {
		_, err = o.QueryTable("Category").Filter("parentid", id).Delete() //删除科室
		if err != nil {
			return err
		}
		for _, v := range categories {
			var categories1 []Category
			_, err = o.QueryTable("Category").Filter("parentid", v.Id).All(&categories1, "Id")
			if err != nil {
				return err
			} else {
				_, err = o.QueryTable("Category").Filter("parentid", v.Id).Delete() //删除价值分类
				if err != nil {
					return err
				}
				for _, w := range categories1 {
					var categories2 []Category
					_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&categories2, "Id")
					if err != nil {
						return err
					} else {
						_, err = o.QueryTable("Category").Filter("parentid", w.Id).Delete() //删除价值内容
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return err
}
