package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "math"
	"strconv"
	"time"
)

type Ratio struct {
	Id       int64
	Category string  //成果类型
	Unit     string  //计量单位
	Rationum float64 //折标系数
	// Data     time.Time `orm:"null;auto_now_add;type(datetime)"`
	Created time.Time `orm:"index;auto_now_add;type(datetime)"`
	Updated time.Time `orm:"index;auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Ratio))
}

func SaveRatio(ratio Ratio) (cid int64, err error) {
	//重复提交问题
	ratio1 := Ratio{Category: ratio.Category}
	// 	user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	o := orm.NewOrm()
	err1 := o.Read(&ratio1, "Category")
	if err1 == orm.ErrNoRows {
		cid, err = o.Insert(&ratio)
		return cid, err
	} else {
		return 0, err1
	}

}

func GetRatios() (ratios []*Ratio, err error) {
	ratios = make([]*Ratio, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("ratio")
	_, err = qs.All(&ratios)
	return ratios, err
}

//根据成果类型，查系数
func GetRationumbycategory(category string) (ratio float64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("ratio")
	var ratio1 Ratio
	err = qs.Filter("category", category).One(&ratio1, "Rationum")
	if err != nil {
		return 0, err
	}
	return ratio1.Rationum, nil

}

//用saveratio，下面这个没用？
// func AddRatio(category, unit string) (id int64, err error) {
// 	// cid, err := strconv.ParseInt(ratioid, 10, 64)
// 	o := orm.NewOrm()
// 	ratio := &Ratio{
// 		Category: category,
// 		Unit:     unit,
// 	}
// 	id, err = o.Insert(ratio)
// 	if err != nil {
// 		return id, err //如果文章编号相同，则唯一性检查错误，返回id吗？
// 	}
// 	if id == 0 {
// 		var ratio Ratio
// 		err = o.QueryTable("ratio").Filter("category", category).One(&ratio, "Id")
// 		id = ratio.Id
// 	}
// 	return id, err
// }

func ModifyRatio(cid string, ratio1 Ratio) error {
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	ratio := &Ratio{Id: cidNum}
	if o.Read(ratio) == nil {
		// 指定多个字段
		// o.Update(&user, "Field1", "Field2", ...)这个试验没成功
		ratio.Category = ratio1.Category
		ratio.Unit = ratio1.Unit
		ratio.Rationum = ratio1.Rationum
		ratio.Updated = ratio1.Updated
		// _, err = o.Update(&ratio, "ProjectName", "DesignStage", "Section", "Tnumber", "Name", "Category", "Page", "Count", "Drawn", "Designd", "Checked", "Examined", "Data", "Updated", "Author")
		_, err := o.Update(ratio, "Category", "Unit", "Rationum", "Updated") //这里不能用&ratio
		if err != nil {
			return err
		}
	}
	return err
}

func DeletRatio(cid string) error { //应该在controllers中显示警告
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	ratio := Ratio{Id: cidNum}
	if o.Read(&ratio) == nil {
		_, err = o.Delete(&ratio)
		if err != nil {
			return err
		}
	}
	return err
}

func GetRatio(tid string) (*Ratio, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	ratio := new(Ratio)
	qs := o.QueryTable("ratio")
	err = qs.Filter("id", tidNum).One(ratio)
	if err != nil {
		return nil, err
	}
	// ratio.Views++
	// _, err = o.Update(topic)

	// attachments := make([]*Attachment, 0)
	// attachment := new(Attachment)
	// qs = o.QueryTable("attachment")
	// _, err = qs.Filter("topicId", tidNum).OrderBy("FileName").All(&attachments)
	// if err != nil {
	// 	return nil, err
	// }
	return ratio, err
}
