package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/validation"
	"strings"
	"web_push/libs/logger"
	"web_push/libs"
)

//时间格式
const (
	TIME_FORMAT = "2006-01-02 15:04:05"

// enabled 字段, 0: 删除, 1: 正常
	ENABLED = 1
	DELETED = 0
)

func init() {
	dbconf, _ := beego.AppConfig.GetSection("db")
	// 注册数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dbconf["username"], dbconf["password"], dbconf["host"], dbconf["port"], dbconf["database"], dbconf["charset"])
	orm.RegisterDataBase(dbconf["alias"], dbconf["driver"], dsn)

	logger.Debug("init", dsn, dbconf["alias"], dbconf["driver"])
	orm.Debug, _ = beego.AppConfig.Bool("debug")
	orm.DefaultTimeLoc = time.UTC
	// 注册Model
	//	orm.RegisterModel(new(GroupTaskModel))
	models := []interface{}{
		NewAppModel(),
		NewMsgInfoModel(),
		NewMsgCidStatusModel(),
	}

	orm.RegisterModel(models...)
}

type BaseModel struct {
	//验证器
	validation.Validation `orm:"-"`
	//当前的Model
	ins interface{}
}

//返回新的orm对象
func (this *BaseModel ) GetOrm() orm.Ormer {
	return orm.NewOrm()
}

/**
 返回字典数组 [{"id":1,"name":"n1"},{"id":2,"name":"n2"}]
 _,err := this.GetOrm().Raw(q.String(),q.GetArgs()).Values(&orm.Params)

 返回单行struct记录 &{Id:1}
 err := this.GetOrm().Raw(q.String(), q.GetArgs()).QueryRow(this)


 */
func (this *BaseModel ) GetQuery() *Query {
	return NewQuery()
	//	qb, _  := orm.NewQueryBuilder(beego.AppConfig.String("db::driver"))
	//	return qb
}

//返回yyyy-mm-dd hh:ii:ss格式的字串
func (this *BaseModel ) GetNow() string {
	return time.Now().Format(TIME_FORMAT)
}

/**
 * 当前时间戳
 */
func (this *BaseModel ) NowTime() int {
	return int(time.Now().Unix())
}

func (this *BaseModel ) GenErrstr() (bool, libs.Errstr) {
	if this.HasErrors() {
		beego.Error(this.Errors)
		var msg string
		for _, err := range this.Errors {
			msg += err.Key + " " + err.Message + ";"
		}
		msg = strings.TrimRight(msg, ";")
		beego.Error(msg)
		return false, libs.Errstr(string(libs.ERR_PARAM) + ":" + msg)
	}
	return true, ""
}

//添加
//func (this *BaseModel ) Insert() error {
//    ins := this.ins
//	_, err := this.GetOrm().Insert(this.ins)
//    this.ins = ins
//	return err
//}

//修改
//func (this *BaseModel ) Update() error {
//    ins := this.ins
//	_, err := this.GetOrm().Update(this.ins)
//    this.ins = ins
//	return err
//}

//删除
//func (this *BaseModel ) Delete() error {
//    ins := this.ins
//	_, err := this.GetOrm().Delete(this.ins)
//    this.ins = ins
//	return err
//}

//查询一行记录
//func (this *BaseModel ) Read(cols ...string) error {
//
//	ins := this.ins
//	err := this.GetOrm().Read(this.ins, cols...)
//	this.ins = ins
//	return err
//
//}

//func (this *BaseModel ) Query() orm.QueryBuilder {
//	return nil
//}

//抛出业务异常
func BizPanic(errstr libs.Errstr) {
	panic("biz:" + string(errstr))
}
