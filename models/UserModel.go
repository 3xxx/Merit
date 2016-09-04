package models

import (
	"errors"
	"strconv"
	// "fmt"
	"log"
	"time"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	. "github.com/beego/admin/src/lib"
)

//用户表
type User struct {
	Id            int64  `PK`
	Username      string `orm:"unique"` //这个拼音的简写
	Nickname      string //中文名，注意这里，很多都要查询中文名才行`orm:"unique;size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(2)"`
	Password      string
	Repassword    string    `orm:"-" form:"Repassword" valid:"Required"`
	Email         string    `orm:"size(32)" form:"Email" valid:"Email"`
	Department    string    //分院
	Secoffice     string    //科室,这里应该用科室id，才能保证即时重名也不怕。否则，查看科室必须要上溯到分院才能避免科室名称重复问题
	Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
	Ip            string    //ip地址
	Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
	Role          int
	// Roles         []*Role   `orm:"rel(m2m)"`
}

// Id            int64
// Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)"`
// Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`

func init() {
	orm.RegisterModel(new(User))
}

func SaveUser(user User) (uid int64, err error) {
	o := orm.NewOrm()
	var user1 User
	//判断是否有重名
	err = o.QueryTable("user").Filter("username", user.Username).One(&user1, "Id")
	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		uid, err = o.Insert(&user)
		if err != nil {
			return 0, err
		}
	} else { //应该进行更新操作
		// user1 := &User{Id: user1.Id}
		user1.Username = user.Username
		user1.Nickname = user.Nickname
		user1.Password = user.Password
		user1.Repassword = user.Repassword
		user1.Email = user.Email
		user1.Department = user.Department
		user1.Secoffice = user.Secoffice
		// user1.Remark = user.Remark
		// user1.Ip = user.Ip
		// user1.Status = user.Status
		user1.Lastlogintime = user.Lastlogintime
		user1.Createtime = time.Now()
		user1.Role = user.Role
		_, err = o.Update(&user1)
		if err != nil {
			return 0, err
		}
		uid = user1.Id
	}
	return uid, err
}

func ValidateUser(user User) error {
	orm := orm.NewOrm()
	var u User

	// user = new(User)
	qs := orm.QueryTable("user")
	err := qs.Filter("username", user.Username).Filter("password", user.Password).One(&u)
	if err != nil {
		return err
	}

	// orm.Where("username=? and pwd=?", user.Username, user.Pwd).Find(&u)
	if u.Username == "" {
		return errors.New("用户名或密码错误！")
	}
	return nil
}

func CheckUname(user User) error {
	orm := orm.NewOrm()
	var u User
	// user = new(User)
	qs := orm.QueryTable("user")
	err := qs.Filter("username", user.Username).One(&u)
	if err != nil {
		return err
	}
	// orm.Where("username=? and pwd=?", user.Username, user.Pwd).Find(&u)
	// if u.Username == "" {
	// 	return errors.New("用户名或密码错误！")
	// }
	return nil
}

func GetUname(user User) ([]*User, error) {
	orm := orm.NewOrm()
	users := make([]*User, 0)
	qs := orm.QueryTable("user")
	_, err := qs.Filter("Username__contains", user.Username).All(&users)
	if err != nil {
		return users, err
	}
	return users, err
}

// func SearchTopics(tuming string, isDesc bool) ([]*Topic, error) {
// 	o := orm.NewOrm()
// 	topics := make([]*Topic, 0)
// 	qs := o.QueryTable("topic")
// 	var err error
// 	if isDesc {
// 		if len(tuming) > 0 {
// 			qs = qs.Filter("Title__contains", tuming) //这里取回
// 		}
// 		_, err = qs.OrderBy("-created").All(&topics)
// 	} else {
// 		_, err = qs.Filter("Title__contains", tuming).OrderBy("-created").All(&topics)
// 		//o.QueryTable("user").Filter("name", "slene").All(&users)
// 	}
// 	return topics, err
// }

// func (u *User) TableName() string {
// 	return beego.AppConfig.String("rbac_user_table")
// }

func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

//验证用户信息
func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

/************************************************************/

//get user list
func Getuserlist(page int64, page_size int64, sort string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&users)
	count, _ = qs.Count()
	return users, count
}

func GetAllusers(page int64, page_size int64, sort string) (users []*User, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).All(&users)
	count, _ = qs.Count()
	return users, count
}

//根据分院和科室名称查所有用户
func GetUsersbySec(department, secoffice string) (users []*User, count int, err error) {
	o := orm.NewOrm()
	// cates := make([]*Category, 0)
	qs := o.QueryTable("user")
	//这里进行过滤
	_, err = qs.Filter("Department", department).Filter("Secoffice", secoffice).OrderBy("Username").All(&users)
	if err != nil {
		return nil, 0, err
	}
	// _, err = qs.OrderBy("-created").All(&cates)
	// _, err := qs.All(&cates)
	count = len(users)
	return users, count, err
}

//根据科室id查所有用户
func GetUsersbySecId(secofficeid string) (users []*User, count int, err error) {
	o := orm.NewOrm()
	// cates := make([]*Category, 0)
	qs := o.QueryTable("user")
	//这里进行过滤
	secid, err := strconv.ParseInt(secofficeid, 10, 64)
	if err != nil {
		return nil, 0, err
	}
	//由secid查自身科室名称
	secoffice, err := GetCategory(secid)
	if err != nil {
		return nil, 0, err
	}
	//由secoffice的pid查分院名称
	department, err := GetCategory(secoffice.ParentId)
	if err != nil {
		return nil, 0, err
	}
	//由分院名称和科室名称查所有用户
	_, err = qs.Filter("Department", department.Title).Filter("Secoffice", secoffice.Title).OrderBy("Username").All(&users)
	if err != nil {
		return nil, 0, err
	}
	// _, err = qs.OrderBy("-created").All(&cates)
	// _, err := qs.All(&cates)
	count = len(users)
	return users, count, err
}

//添加用户
func AddUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := new(User)
	user.Username = u.Username
	user.Password = Strtomd5(u.Password)
	user.Nickname = u.Nickname
	user.Email = u.Email
	user.Remark = u.Remark
	user.Status = u.Status

	id, err := o.Insert(user)
	return id, err
}

//更新用户
// func UpdateUser(u *User) (int64, error) {
// 	if err := checkUser(u); err != nil {
// 		return 0, err
// 	}
// 	o := orm.NewOrm()
// 	user := make(orm.Params)
// 	if len(u.Username) > 0 {
// 		user["Username"] = u.Username
// 	}
// 	if len(u.Nickname) > 0 {
// 		user["Nickname"] = u.Nickname
// 	}
// 	if len(u.Email) > 0 {
// 		user["Email"] = u.Email
// 	}
// 	if len(u.Remark) > 0 {
// 		user["Remark"] = u.Remark
// 	}
// 	if len(u.Password) > 0 {
// 		user["Password"] = Strtomd5(u.Password)
// 	}
// 	if u.Status != 0 {
// 		user["Status"] = u.Status
// 	}
// 	if len(user) == 0 {
// 		return 0, errors.New("update field is empty")
// 	}
// 	var table User
// 	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
// 	return num, err
// }
func UpdateUser(userid, nickname, email, password string) error {
	// if err := checkUser(&u); err != nil {
	// 	return err
	// }
	id, err := strconv.ParseInt(userid, 10, 64)
	o := orm.NewOrm()
	// qs := o.QueryTable("user") //不知道主键就用这个过滤操作
	user := User{Id: id}
	// err := qs.Filter("Username", u.Username).One(&u)
	// if err == nil {
	// 	return err
	// }
	// user := User{Username: u.Username}
	// if err := o.Read(&user); err == nil {
	user.Nickname = nickname
	user.Email = email
	if password != "" {
		user.Password = password
		// user1 := make(orm.Params)
		// var table User
		_, err = o.Update(&user, "password", "nickname", "email")
		if err != nil {
			return err
		}
	} else {
		_, err = o.Update(&user, "nickname", "email")
		if err != nil {
			return err
		}
	}
	// } else {
	// return err
	// }
	return nil
}

//更新用户登陆时间
func UpdateUserlastlogintime(username string) error {
	o := orm.NewOrm()
	user := make(orm.Params)
	if len(username) > 0 {
		user["Lastlogintime"] = time.Now()
	}

	if len(username) == 0 {
		return errors.New("update field is empty")
	}
	var table User
	_, err := o.QueryTable(table).Filter("Username", username).Update(user)
	return err
}

func DelUserById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&User{Id: Id})
	return status, err
}

//###*****这里特别注意，这个是用户名，是汉语拼音，不是Nickname！！！！
func GetUserByUsername(username string) (user User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("user") //不知道主键就用这个过滤操作
	//进行编号唯一性检查
	err = qs.Filter("username", username).One(&user)
	if err != nil {
		return user, err
	}
	// user = User{Username: username} //指定字段查询，这样也行
	// o := orm.NewOrm()
	// o.Read(&user,"Username")
	return user, err
}

//根据用户nickname取得用户
func GetUserByNickname(nickname string) (user User) {
	o := orm.NewOrm()
	qs := o.QueryTable("user") //不知道主键就用这个过滤操作
	//进行编号唯一性检查
	qs.Filter("nickname", nickname).One(&user)
	// user = User{Username: username} //指定字段查询，这样也行
	// o := orm.NewOrm()
	// o.Read(&user,"Username")
	return user
}

func GetUserByUserId(userid int64) (user User) {
	user = User{Id: userid}
	o := orm.NewOrm()
	o.Read(&user) //这里是默认主键查询。=(&user,"Id")
	return user
}

// func GetAllReplies(tid string) (replies []*Comment, err error) {
// 	tidNum, err := strconv.ParseInt(tid, 10, 64)
// 	if err != nil {
// 		return nil, err
// 	}
// 	replies = make([]*Comment, 0)

// 	o := orm.NewOrm()
// 	qs := o.QueryTable("comment")
// 	_, err = qs.Filter("tid", tidNum).All(&replies)
// 	return replies, err

// }

func GetRoleByUserId(userid int64) (roles []*Role, count int64) { //*Topic, []*Attachment, error
	roles = make([]*Role, 0)
	o := orm.NewOrm()
	// role := new(Role)
	count, _ = o.QueryTable("role").Filter("Users__User__Id", userid).All(&roles)
	return roles, count
	// 通过 post title 查询这个 post 有哪些 tag
	// var tags []*Tag
	// num, err := dORM.QueryTable("tag").Filter("Posts__Post__Title", "Introduce Beego ORM").All(&tags)

}

func GetRoleByUsername(username string) (roles []*Role, count int64, err error) { //*Topic, []*Attachment, error
	roles = make([]*Role, 0)
	o := orm.NewOrm()
	// role := new(Role)
	count, err = o.QueryTable("role").Filter("Users__User__Username", username).All(&roles)
	return roles, count, err
	// 通过 post title 查询这个 post 有哪些 tag
	// var tags []*Tag
	// num, err := dORM.QueryTable("tag").Filter("Posts__Post__Title", "Introduce Beego ORM").All(&tags)

}
