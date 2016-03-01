package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
	"time"
)

type Category struct {
	Id              int64 `form:"-"`
	ParentId        int64
	Uid             int64
	Title           string `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Number          string `orm:"unique",form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"`
	Content         string `orm:"sie(5000)"`
	Route           string
	Created         time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated         time.Time `orm:"index","auto_now_add;type(datetime)"`
	Views           int64     `form:"-",orm:"index"`
	Author          string
	TopicCount      int64  //`form:"-"`
	TopicLastUserId int64  //`form:"-"`
	DiskDirectory   string `orm:"null"`
	Url             string `orm:"null"`
	// Type            string `orm:"null"` //项目类型：供水、枢纽、提防、河道、船闸、电站、水闸
}

func init() {
	orm.RegisterModel(new(Category)) //, new(Article)
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/merit.db", 10)
}

func AddCategory(name, number, content, path, route, uname, diskdirectory, url string) (id int64, err error) {
	o := orm.NewOrm()
	cate := &Category{
		Title:         name,
		Number:        number,
		Content:       content,
		Author:        uname,
		Route:         route,
		Created:       time.Now(),
		Updated:       time.Now(),
		DiskDirectory: diskdirectory,
		Url:           url,
	}
	qs := o.QueryTable("category") //不知道主键就用这个过滤操作
	//进行编号唯一性检查
	err = qs.Filter("number", number).One(cate)
	if err == nil {
		return 0, err
	}
	id, err = o.Insert(cate)
	if err != nil {
		return 0, err
	}
	// var id int64
	// err = qs.Filter("title", name).One(id, "Id")
	// id, err := strconv.ParseInt(tid, 10, 64)
	// var cate Category
	// cates := make([]*Category, 0)
	// id, err := qs.Filter("title", name).All(&cates, "Id") //只是返回查询个数
	var post Category
	err = o.QueryTable("Category").Filter("number", number).One(&post, "Id", "Title")
	if err != nil {
		return 0, err
	}
	// 	var post Post
	// o.QueryTable("post").Filter("Content__istartswith", "prefix string").One(&post, "Id", "Title")
	// //进行目录的添加和parentid的设置
	array := strings.Split(path, ",") //字符串切割 [a b c d e]
	// var j int
	// //将path存成3个数组
	// var JieDuan [10]string
	// var ZhuanYe [10]string
	// var ChengGuo [10]string
	// for i, v := range array {
	// 	switch v {
	// 	case "ghj", "xj", "ky":
	// 		// 先定义一个数组

	// 		JieDuan[i] = v
	// 		j = i
	// 	case "gh", "sg", "shg":
	// 		// 先定义一个数组

	// 		ZhuanYe[i-j] = v
	// 		j = i
	// 	case "dwg", "doc", "xls":
	// 		// 先定义一个数组

	// 		ChengGuo[i-j] = v
	// 	}
	// }
	//将数组存入数据库
	// for _, v := range JieDuan {
	// 	// if v == "" {
	// 	// 	break JLoop
	// 	// }
	// 	cate = &Category{
	// 		Title:    v,
	// 		ParentId: post.Id, //这里存入项目的id
	// 		Created:  time.Now(),
	// 		Updated:  time.Now(),
	// 	}
	// 	// JLoop:
	// 	_, err = o.Insert(cate)

	// 	var posts []Category //详见beego手册的All的示例
	// 	_, err = o.QueryTable("Category").Filter("parentid", post.Id).All(&posts, "Id")
	// 	for i, z := range ZhuanYe {
	// 		cate := &Category{
	// 			Title:    z,
	// 			ParentId: posts[i].Id,
	// 			Created:  time.Now(),
	// 			Updated:  time.Now(),
	// 		}
	// 		_, err = o.Insert(cate)

	// 	}
	// }
	var jieduan string
	for _, v := range array {
		switch v {
		case "A":
			jieduan = "规划"
		case "B":
			jieduan = "项目建议书"
		case "C":
			jieduan = "可行性研究"
		case "D":
			jieduan = "初步设计"
		case "E":
			jieduan = "招标设计"
		case "F":
			jieduan = "施工图设计"
		case "G":
			jieduan = "竣工图"
		case "L":
			jieduan = "专题"
		}
		switch v {
		case "A", "B", "C", "D", "E", "F", "G", "L":
			cate = &Category{
				Title:    jieduan,
				ParentId: post.Id, //这里存入项目的id
				Created:  time.Now(),
				Updated:  time.Now(),
				Author:   uname,
				// filepath := ".\\attachment\\" + ProNumber + category.Title + "\\" + ProJieduan
				DiskDirectory: ".\\attachment\\" + number + name + "\\" + v + "\\",
				Url:           "/attachment/" + number + name + "/" + v + "/",
				// Style:         style,
			}
			_, err = o.Insert(cate)
			if err != nil {
				return 0, err
			}
			//建立目录——注意，models中无法建立目录，必须在controllers中才行
		}
	}

	var leixing string
	for _, v := range array {
		switch v {
		case "FB":
			leixing = "技术报告"
		case "FD":
			leixing = "设计大纲"
		case "FG":
			leixing = "设计/修改通知单"
		case "FT":
			leixing = "工程图纸"
		case "FJ":
			leixing = "计算书"
		case "FP":
			leixing = "PDF文件"
		case "Fdiary":
			leixing = "文章/设代日记"
		}
		switch v {
		case "FB", "FD", "FG", "FT", "FJ", "FP", "Fdiary": //文件类型
			//查到阶段的parentid，符合这个项目的，得出阶段id，作为文件类型parentid
			var posts []Category                                                                                    //详见beego手册的All的示例
			_, err = o.QueryTable("Category").Filter("parentid", post.Id).All(&posts, "Id", "DiskDirectory", "Url") //Id没什么用？
			for _, w := range posts {
				cate := &Category{
					Title:         leixing, //v,
					ParentId:      w.Id,
					Created:       time.Now(),
					Updated:       time.Now(),
					Author:        uname,
					DiskDirectory: diskdirectory + w.DiskDirectory + v + "\\",
					Url:           url + w.Url + v + "/",
					// Style:         style,
				}
				_, err = o.Insert(cate)
				if err != nil {
					return 0, err
				}
			}
		}
	}
	var zhuanye string
	for _, v := range array {
		switch v {
		case "1":
			zhuanye = "综合"
		case "2":
			zhuanye = "规划(含水文、经评)"
		case "3":
			zhuanye = "测量"
		case "4":
			zhuanye = "地质(含钻探)"
		case "5":
			zhuanye = "水工(含公路、安全监测)"
		case "6":
			zhuanye = "建筑"
		case "7":
			zhuanye = "机电"
		case "8":
			zhuanye = "征地、环保、水保"
		case "9":
			zhuanye = "施工、工程造价"
		}
		switch v {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9": //专业
			//查到文件类型的parentid，符合这个阶段的，得出文件分类id，作为专业分类的parentid
			var posts []Category                                                                     //详见beego手册的All的示例
			_, err = o.QueryTable("Category").Filter("parentid", post.Id).All(&posts, "Id", "Title") //这里只用到Id？
			for _, w := range posts {
				var postss []Category //详见beego手册的All的示例
				_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&postss, "Id", "DiskDirectory", "Url")
				for _, t := range postss {
					cate := &Category{
						Title:         zhuanye, //v,
						ParentId:      t.Id,
						Created:       time.Now(),
						Updated:       time.Now(),
						Author:        uname,
						DiskDirectory: diskdirectory + t.DiskDirectory + v + "\\",
						Url:           url + t.Url + v + "/",
						// Style:         style,
					}
					_, err = o.Insert(cate)
					if err != nil {
						return 0, err
					}
				}
			}
		}
	}
	return id, nil
}
