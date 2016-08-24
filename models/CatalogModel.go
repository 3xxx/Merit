package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"math"
	"strconv"
	"time"
)

type Catalog struct {
	Id            int64
	ProjectNumber string    //项目编号
	ProjectName   string    //项目名称
	DesignStage   string    //阶段
	Section       string    //专业
	Tnumber       string    //成果编号
	Name          string    //成果名称
	Category      string    //成果类型
	Page          string    //成果计量单位
	Count         float64   //成果数量
	Drawn         string    //编制、绘制
	Designd       string    //设计
	Checked       string    //校核
	Examined      string    //审查
	Verified      string    //核定
	Approved      string    //批准
	Complex       float64   //难度系数
	Drawnratio    float64   //编制、绘制占比系数
	Designdratio  float64   //设计系数
	Checkedratio  float64   //校核系数
	Examinedratio float64   //审查系数
	Data          time.Time `orm:"null;auto_now_add;type(datetime)"`
	Created       time.Time `orm:"index;auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"index;auto_now_add;type(datetime)"`
	Author        string    //上传者
}

//员工的编制、设计……分值——全部改成float浮点型小数
type Employeeachievement struct {
	Id       string  //员工Id
	Name     string  //员工姓名
	Drawn    float64 //编制、绘制
	Designd  float64 //设计
	Checked  float64 //校核
	Examined float64 //审查
	Verified float64 //核定
	Approved float64 //批准
	Sigma    float64 //合计
}

//分院里各个科室人员结构
type Secofficeachievement struct {
	Id       int64  //科室Id
	Name     string //科室
	Employee []Employeeachievement
}

func init() {
	orm.RegisterModel(new(Catalog))
}

func SaveCatalog(catalog Catalog) (cid int64, err error) {
	orm := orm.NewOrm()
	// fmt.Println(user)
	cid, err = orm.Insert(&catalog) //_, err = o.Insert(reply)
	return cid, err
}

func GetAllCatalogs(cid string) (catalogs []*Catalog, err error) {
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return nil, err
	}
	catalogs = make([]*Catalog, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	_, err = qs.Filter("parentid", cidNum).All(&catalogs)
	return catalogs, err
}

//用savecatalog，下面这个没用？
func AddCatalog(name, tnumber string) (id int64, err error) {
	// cid, err := strconv.ParseInt(categoryid, 10, 64)
	o := orm.NewOrm()
	catalog := &Catalog{
		Name:    name,
		Tnumber: tnumber,
		// Category:   category,
		// CategoryId: cid,
		// Content:    content,
		// Attachment: attachment,
		// Author:     uname,
		// Created:    time.Now(),
		// Updated:    time.Now(),
		// ReplyTime:  time.Now(),
	}
	//	qs := o.QueryTable("category") //不知道主键就用这个过滤操作
	//	err := qs.Filter("title", name).One(cate)
	//	if err == nil {
	//		return err
	//	}
	id, err = o.Insert(catalog)
	if err != nil {
		return id, err //如果文章编号相同，则唯一性检查错误，返回id吗？
	}
	if id == 0 {
		var catalog Catalog
		err = o.QueryTable("catalog").Filter("tnumber", tnumber).One(&catalog, "Id")
		id = catalog.Id
	}
	return id, err
}

func ModifyCatalog(cid string, catalog1 Catalog) error {
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	catalog := &Catalog{Id: cidNum}
	if o.Read(catalog) == nil {
		// 指定多个字段
		// o.Update(&user, "Field1", "Field2", ...)这个试验没成功
		catalog.ProjectNumber = catalog1.ProjectNumber
		catalog.ProjectName = catalog1.ProjectName
		catalog.DesignStage = catalog1.DesignStage
		catalog.Section = catalog1.Section
		catalog.Tnumber = catalog1.Tnumber
		catalog.Name = catalog1.Name
		catalog.Category = catalog1.Category
		catalog.Page = catalog1.Page
		catalog.Count = catalog1.Count
		catalog.Drawn = catalog1.Drawn
		catalog.Designd = catalog1.Designd
		catalog.Checked = catalog1.Checked
		catalog.Examined = catalog1.Examined
		// catalog.Verified     = catalog1.Verified
		// catalog.Approved     = catalog1.Approved
		// catalog.Complex      = catalog1.Complex
		// catalog.Drawnratio   = catalog1.Drawnratio
		// catalog.Designdratio = catalog1.Designdratio
		// catalog.Checkedratio = catalog1.Checkedratio
		// catalog.Examinedratio= catalog1.Examinedratio
		catalog.Data = catalog1.Data
		// catalog.Created      = catalog1.Created
		catalog.Updated = catalog1.Updated
		catalog.Author = catalog1.Author
		// _, err = o.Update(&catalog, "ProjectName", "DesignStage", "Section", "Tnumber", "Name", "Category", "Page", "Count", "Drawn", "Designd", "Checked", "Examined", "Data", "Updated", "Author")
		_, err := o.Update(catalog) //这里不能用&catalog
		if err != nil {
			return err
		}
	}
	return err
}

func DeletCatalog(cid string) error { //应该在controllers中显示警告
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	catalog := Catalog{Id: cidNum}
	if o.Read(&catalog) == nil {
		_, err = o.Delete(&catalog)
		if err != nil {
			return err
		}
	}
	return err
}

func GetCatalog(tid string) (*Catalog, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	catalog := new(Catalog)
	qs := o.QueryTable("catalog")
	err = qs.Filter("id", tidNum).One(catalog)
	if err != nil {
		return nil, err
	}
	// catalog.Views++
	// _, err = o.Update(topic)

	// attachments := make([]*Attachment, 0)
	// attachment := new(Attachment)
	// qs = o.QueryTable("attachment")
	// _, err = qs.Filter("topicId", tidNum).OrderBy("FileName").All(&attachments)
	// if err != nil {
	// 	return nil, err
	// }
	return catalog, err
}

// func GetPids(pid int64) ([]*Category, error) {
// 	o := orm.NewOrm()
// 	cates := make([]*Category, 0)
// 	qs := o.QueryTable("category")
// 	var err error
// 	//这里进行过滤
// 	_, err = qs.Filter("ParentId", pid).All(&cates)
// 	// _, err = qs.OrderBy("-created").All(&cates)
// 	// _, err := qs.All(&cates)
// 	return cates, err
// }

//由用户姓名取得所有编制、设计、校核分值
func Getemployeevalue(uname string, t1, t2 time.Time) (employeevalue []Employeeachievement, err error) {

	catalogs := make([]*Catalog, 0)

	cond := orm.NewCondition()
	cond1 := cond.And("data__gte", t1).And("data__lte", t2)
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	qs = qs.SetCond(cond1)
	//1、查制图工作量
	_, err = qs.Filter("Drawn", uname).All(&catalogs) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	// slice1 := make([]Person, 0)
	var Drawnvalue float64
	var drawn float64
	var Designdvalue float64
	var designd float64
	var Checkedvalue float64
	var checked float64
	var Examinedvalue float64
	var examined float64
	aa := make([]Employeeachievement, 1)
	// var aa *Employeeachievement
	for _, v := range catalogs {
		// Category      string    //成果类型
		// Page          string    //成果计量单位
		// Count         int       //成果数量
		// Drawn         string    //编制、绘制
		// Designd       string    //设计
		// Checked       string    //校核
		// Examined      string    //审查
		// Verified      string    //核定
		// Approved      string    //批准
		// Complex       int       //难度系数
		// Drawnratio    int       //编制、绘制占比系数
		// Designdratio  int       //设计系数
		// Checkedratio  int       //校核系数
		// Examinedratio int       //审查系数
		// mark, err := strconv.Atoi(v.Count)
		// if err != nil {
		// 	return nil, err
		// }

		//成果数量*难度系数*绘制占比
		//成果数量*难度系数*设计占比
		//如果是图纸
		switch v.Category {
		case "图纸":
			// Name     string //员工姓名
			// Drawn    int    //编制、绘制
			// Designd  int    //设计
			// Checked  int    //校核
			// Examined int    //审查
			// Verified int    //核定
			// Approved int    //批准
			// Sigma    int    //合计
			Drawnvalue = v.Count * v.Complex * v.Drawnratio
		case "报告":
			Drawnvalue = v.Count / 5 * v.Complex * v.Drawnratio
		case "大纲":
			Drawnvalue = v.Count / 5 * v.Complex * v.Drawnratio
		case "计算书":
			Drawnvalue = v.Count / 5 * v.Complex * v.Drawnratio
		case "修改单":
			Drawnvalue = v.Count * v.Complex * v.Drawnratio
		default:
			Drawnvalue = v.Count * v.Complex * v.Drawnratio
		}
		drawn = drawn + Drawnvalue
	}
	aa[0].Drawn = Round(drawn, 1)

	//2、查设计工作量
	_, err = qs.Filter("Designd", uname).All(&catalogs) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	for _, v := range catalogs {
		//成果数量*难度系数*绘制占比
		//成果数量*难度系数*设计占比
		//如果是图纸
		switch v.Category {
		case "图纸":
			Designdvalue = v.Count * v.Complex * v.Designdratio
		case "报告":
			Designdvalue = v.Count / 5 * v.Complex * v.Designdratio
		case "大纲":
			Designdvalue = v.Count / 5 * v.Complex * v.Designdratio
		case "计算书":
			Designdvalue = v.Count / 5 * v.Complex * v.Designdratio
		case "修改单":
			Designdvalue = v.Count * v.Complex * v.Designdratio
		default:
			Designdvalue = v.Count * v.Complex * v.Designdratio
		}
		designd = designd + Designdvalue
	}
	aa[0].Designd = Round(designd, 1)

	//3、查校核工作量
	_, err = qs.Filter("Checked", uname).All(&catalogs) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	for _, v := range catalogs {
		//成果数量*难度系数*绘制占比
		//成果数量*难度系数*设计占比
		//如果是图纸
		switch v.Category {
		case "图纸":
			Checkedvalue = v.Count * v.Complex * v.Checkedratio
		case "报告":
			Checkedvalue = v.Count / 5 * v.Complex * v.Checkedratio
		case "大纲":
			Checkedvalue = v.Count / 5 * v.Complex * v.Checkedratio
		case "计算书":
			Checkedvalue = v.Count / 5 * v.Complex * v.Checkedratio
		case "修改单":
			Checkedvalue = v.Count * v.Complex * v.Checkedratio
		default:
			Checkedvalue = v.Count * v.Complex * v.Checkedratio
		}
		checked = checked + Checkedvalue
	}
	aa[0].Checked = Round(checked, 1)
	//4、查审查工作量
	_, err = qs.Filter("Examined", uname).All(&catalogs) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	for _, v := range catalogs {
		//成果数量*难度系数*绘制占比
		//成果数量*难度系数*设计占比
		//如果是图纸
		switch v.Category {
		case "图纸":
			Examinedvalue = v.Count * v.Complex * v.Examinedratio
		case "报告":
			Examinedvalue = v.Count / 5 * v.Complex * v.Examinedratio
		case "大纲":
			Examinedvalue = v.Count / 5 * v.Complex * v.Examinedratio
		case "计算书":
			Examinedvalue = v.Count / 5 * v.Complex * v.Examinedratio
		case "修改单":
			Examinedvalue = v.Count * v.Complex * v.Examinedratio
		default:
			Examinedvalue = v.Count * v.Complex * v.Examinedratio
		}
		examined = examined + Examinedvalue
	}
	aa[0].Examined = Round(examined, 1)

	aa[0].Name = uname //这个是nickname，千万注意
	user1 := GetUserByNickname(uname)
	id := strconv.FormatInt(user1.Id, 10)
	aa[0].Id = id
	aa[0].Sigma = Round(drawn+designd+checked+examined, 1)
	employeevalue = aa
	// employeevalue = append(employeevalue, aa...)
	return employeevalue, err
}

//由用户Id取得所有编制、设计、校核详细catalog，按成果类型排列
func Getcatalogbyuserid(id, category string) (catalogs []*Catalog, err error) {
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	user := GetUserByUserId(Id)
	o := orm.NewOrm()
	aa := make([]*Catalog, 0)
	bb := make([]*Catalog, 0)
	cc := make([]*Catalog, 0)
	dd := make([]*Catalog, 0)
	qs := o.QueryTable("catalog")
	//1、查图纸类型的工作
	_, err = qs.Filter("Drawn", user.Nickname).Filter("Category", category).All(&aa) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, aa...)
	_, err = qs.Filter("Designd", user.Nickname).Filter("Category", category).All(&bb) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, bb...)
	_, err = qs.Filter("Checked", user.Nickname).Filter("Category", category).All(&cc) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, cc...)
	_, err = qs.Filter("Examined", user.Nickname).Filter("Category", category).All(&dd) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, dd...)
	return catalogs, err
}

//由用户Id取得所有成果按时间顺序排列
func Getcatalog2byuserid(id string) (catalogs []*Catalog, err error) {
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	user := GetUserByUserId(Id)
	o := orm.NewOrm()
	aa := make([]*Catalog, 0)
	bb := make([]*Catalog, 0)
	cc := make([]*Catalog, 0)
	dd := make([]*Catalog, 0)
	qs := o.QueryTable("catalog")
	//1、查图纸类型的工作
	_, err = qs.Filter("Drawn", user.Nickname).All(&aa) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, aa...)
	_, err = qs.Filter("Designd", user.Nickname).All(&bb) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, bb...)
	_, err = qs.Filter("Checked", user.Nickname).All(&cc) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, cc...)
	_, err = qs.Filter("Examined", user.Nickname).All(&dd) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, dd...)
	return catalogs, err
}

//四舍五入
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
