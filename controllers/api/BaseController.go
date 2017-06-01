package api

import (
	"github.com/astaxie/beego"
	"strings"
	"os"
	"web_push/libs/logger"
	"web_push/libs"
)

type BaseController struct {
	// 返回数据
	Msg map[string]interface{}

	beego.Controller
}

// API 返回数据
type Data map[string]interface{}

/**
 * Action 的前置钩子
 */
func (this *BaseController) Prepare() {

	// @TODO 签名验证

	this.initMsg()
}

/**
 * Action 的后置钩子
 */
func (this *BaseController) Finish() {

}

func (this *BaseController ) initMsg() {
	this.Msg = make(map[string]interface{})
	this.Msg["errno"] = 0
	this.Msg["errmsg"] = ""
	this.Msg["data"] = []string{}
	this.Msg["node"], _ = os.Hostname()
	this.Data["json"] = this.Msg
}

/**
 * 操作成功, 设置返回数据
 */
func (this *BaseController) setData(data interface{}) {
	this.Msg["data"] = data
	this.Data["json"] = this.Msg

	logger.Debug("biz." + this.GetActionId(), map[string]interface{}{
		"request": this.Input(),
		"data": data,
	})

	this.ServeJson()
}

/**
 * 操作失败, 设置错误信息
 */
func (this *BaseController ) setError(errstr libs.Errstr, mixed ... interface{}) {
	this.Msg["errno"] = errstr.GetErrno()
	this.Msg["errmsg"] = errstr.GetErrmsg()
	if len(mixed) > 0 {
		this.Msg["data"] = mixed[0]
	}
	this.Data["json"] = this.Msg

	logger.Error("biz." + this.GetActionId(), map[string]interface{} {
		"err": errstr,
		"request": this.Input(),
		"data": mixed,
	})
	this.ServeJson()
}

func (this *BaseController ) GetActionId() string {
	_, actionId := this.GetControllerAndAction()
	return strings.ToLower(actionId)
}