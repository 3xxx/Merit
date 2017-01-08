//部门结构，价值，ip地址段，日历
package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"time"
)

// type AdminCategory struct {
// 	Id       int64     `form:"-"`
// 	ParentId int64     `orm:"null"`
// 	Title    string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
// 	Code     string    `orm:"null"`
// 	Grade    int       `orm:"null"`
// 	Created  time.Time `orm:"index","auto_now_add;type(datetime)"`
// 	Updated  time.Time `orm:"index","auto_now_add;type(datetime)"`
// }
//科室结构
type AdminDepartment struct {
	Id       int64     `form:"-"`
	ParentId int64     `orm:"null"`
	Title    string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Code     string    `orm:"null"`
	Created  time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index","auto_now_add;type(datetime)"`
}

//价值分类
type AdminMerit struct {
	Id       int64     `form:"-"`
	Title    string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Mark     string    `orm:"null"`                                              //设置分数
	List     string    `orm:"null"`                                              //选择项
	ListMark string    `orm:"null"`
	Created  time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index","auto_now_add;type(datetime)"`
}

//ip地址段权限
type AdminIpsegment struct {
	Id      int64     `form:"-"`
	Title   string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	StartIp string    `orm:"not null"`
	EndIp   string    `orm:"null"`
	Iprole  int       `orm:"null"`
	Created time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated time.Time `orm:"index","auto_now_add;type(datetime)"`
}

//日历
type AdminCalenda struct {
	Id        int64     `form:"-"`
	Title     string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(100)"` //orm:"unique",
	starttime time.Time `orm:"not null;type(datetime)"`
	endtime   time.Time `orm:"null;type(datetime)"`
	allday    int8      `orm:"not null;default(0)"`
	color     string    `orm:"null"`
}

// `id` int(11) NOT NULL AUTO_INCREMENT,
//   `title` varchar(100) NOT NULL,
//   `starttime` int(11) NOT NULL,
//   `endtime` int(11) DEFAULT NULL,
//   `allday` tinyint(1) NOT NULL DEFAULT '0',
//   `color` varchar(20) DEFAULT NULL,

func init() {
	orm.RegisterModel(new(AdminDepartment), new(AdminMerit), new(AdminIpsegment)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/engineer.db", 10)
}

//添加部门
func AddAdminDepart(pid int64, title, code string) (id int64, err error) {
	o := orm.NewOrm()
	depart := &AdminDepartment{
		ParentId: pid,
		Title:    title,
		Code:     code,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	id, err = o.Insert(depart)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//修改
func UpdateAdminDepart(cid int64, title, code string) error {
	o := orm.NewOrm()
	category := &AdminDepartment{Id: cid}
	if o.Read(category) == nil {
		category.Title = title
		category.Code = code
		category.Updated = time.Now()
		_, err := o.Update(category)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteAdminDepart(cid int64) error {
	o := orm.NewOrm()
	category := &AdminDepartment{Id: cid}
	if o.Read(category) == nil {
		_, err := o.Delete(category)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据部门id取得所有科室
//如果父级id为0，则取所有部门
func GetAdminDepart(pid int64) (departs []*AdminDepartment, err error) {
	o := orm.NewOrm()
	departs = make([]*AdminDepartment, 0)
	qs := o.QueryTable("AdminDepartment")
	_, err = qs.Filter("parentid", pid).All(&departs)
	if err != nil {
		return nil, err
	}
	return departs, err
}

//根据父级id取得所有
//如果父级id为空，则取所有一级category
// func GetAdminCategory(pid int64) (categories []*AdminDepartment, err error) {
// 	o := orm.NewOrm()
// 	categories = make([]*AdminDepartment, 0)
// 	qs := o.QueryTable("AdminDepartment")
// 	_, err = qs.Filter("parentid", pid).All(&categories)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return categories, err
// }
//根据部门名字title返回自身
func GetAdminDepartName(title string) (AdminDepartment, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminDepartment")
	var cate AdminDepartment
	err := qs.Filter("title", title).One(&cate)
	if err != nil {
		return cate, err
	}
	return cate, err
}

//根据部门名字title查询所有下级科室category
func GetAdminDepartTitle(title string) (categories []*AdminDepartment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminDepartment")
	var cate AdminDepartment
	err = qs.Filter("title", title).One(&cate)
	categories = make([]*AdminDepartment, 0)
	_, err = qs.Filter("parentid", cate.Id).All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, err
}

//根据id查科室
func GetAdminDepartbyId(id int64) (category AdminDepartment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminDepartment")
	err = qs.Filter("id", id).One(&category)
	if err != nil {
		return category, err
	}
	return category, err
}

//由分院id和科室 名称取得科室
func GetAdminDepartbyidtitle(id int64, title string) (*AdminDepartment, error) {
	o := orm.NewOrm()
	// cate := &Category{Id: id}
	category := new(AdminDepartment)
	qs := o.QueryTable("AdminDepartment")
	err := qs.Filter("parentid", id).Filter("title", title).One(category)
	if err != nil {
		return nil, err
	}
	return category, err
}

//*********价值********************
//添加价值
func AddAdminMerit(title, mark, list, listmark string) (id int64, err error) {
	//重复性检查
	o := orm.NewOrm()
	cate := &AdminMerit{
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

//由父级id（科室id）得到所有价值
func GetAdminMeritbyPid(pid int64) ([]*AdminMerit, error) {
	o := orm.NewOrm()
	cates := make([]*AdminMerit, 0)
	qs := o.QueryTable("AdminMerit")
	var err error
	//这里进行过滤
	_, err = qs.Filter("ParentId", pid).All(&cates)
	return cates, err
}

//取到所有的价值结构
func GetAdminMerit() ([]*AdminMerit, error) {
	o := orm.NewOrm()
	merit := make([]*AdminMerit, 0)
	qs := o.QueryTable("AdminMerit")
	_, err := qs.All(&merit)
	if err != nil {
		return nil, err
	}
	return merit, err
}

//由id取得价值
func GetAdminMeritbyId(id int64) (*AdminMerit, error) {
	o := orm.NewOrm()
	// cate := &Category{Id: id}
	merit := new(AdminMerit)
	qs := o.QueryTable("AdminMerit")
	err := qs.Filter("id", id).One(merit)
	if err != nil {
		return nil, err
	}
	return merit, err
}

//修改merit
func UpdateAdminMerit(id int64, title, mark, list, listmark string) error {
	o := orm.NewOrm()
	merit := &AdminMerit{Id: id}
	var err error
	if o.Read(merit) == nil {
		merit.Title = title
		merit.Mark = mark
		merit.List = list
		merit.ListMark = listmark
		merit.Updated = time.Now()
		_, err = o.Update(merit)
		if err != nil {
			return err
		}
	}
	return err
}

//删除价值结构
func DeleteAdminMerit(id int64) error { //应该在controllers中显示警告
	o := orm.NewOrm()
	merit := AdminMerit{Id: id}
	var err error
	if o.Read(&merit) == nil {
		_, err = o.Delete(&merit) //删除分院
		if err != nil {
			return err
		}
	}
	return err
}

//****************IP******************
//添加ip地址段
func AddAdminIpsegment(title, startip, endip string, iprole int) (id int64, err error) {
	o := orm.NewOrm()
	ipsegment := &AdminIpsegment{
		Title:   title,
		StartIp: startip,
		EndIp:   endip,
		Iprole:  iprole,
		Created: time.Now(),
		Updated: time.Now(),
	}
	id, err = o.Insert(ipsegment)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//修改Ip地址段
func UpdateAdminIpsegment(cid int64, title, startip, endip string, iprole int) error {
	o := orm.NewOrm()
	ipsegment := &AdminIpsegment{Id: cid}
	if o.Read(ipsegment) == nil {
		ipsegment.Title = title
		ipsegment.StartIp = startip
		ipsegment.EndIp = endip
		ipsegment.Iprole = iprole
		ipsegment.Updated = time.Now()
		_, err := o.Update(ipsegment)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteAdminIpsegment(cid int64) error {
	o := orm.NewOrm()
	ipsegment := &AdminIpsegment{Id: cid}
	if o.Read(ipsegment) == nil {
		_, err := o.Delete(ipsegment)
		if err != nil {
			return err
		}
	}
	return nil
}

//查询所有Ip地址段
func GetAdminIpsegment() (ipsegments []*AdminIpsegment, err error) {
	o := orm.NewOrm()
	// ipsegments = make([]*AdminIpsegment, 0)

	qs := o.QueryTable("AdminIpsegment") //这个表名AchievementTopic需要用驼峰式，
	// if pid != "" {                      //如果给定父id则进行过滤
	//pid转成64为
	// pidNum, err := strconv.ParseInt(pid, 10, 64)
	// if err != nil {
	// 	return nil, err
	// }
	_, err = qs.All(&ipsegments)
	if err != nil {
		return nil, err
	}

	return ipsegments, err
	// } else { //如果不给定父id（PID=0），则取所有一级
	// _, err = qs.Filter("parentid", 0).All(&categories)
	// if err != nil {
	// 	return nil, err
	// }
	// return categories, err
	// }
}
