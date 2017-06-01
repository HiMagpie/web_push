package center

import (
	"web_push/controllers"
	"web_push/models"
	"web_push/libs"
)

type DashboardController struct {
	controllers.BaseController
}

func (this *DashboardController) Get() {
	this.Index()
}

func (this *DashboardController) Index() {
	keyword := this.GetString("keyword", "")
	page, _ := this.GetInt("page", 1)
	pageSize, _ := this.GetInt("page_size", 20)
	apps, _ := models.NewAppModel().GetList(keyword, page, pageSize)

	this.Data["apps"] = apps
	this.Data["title"] = "Android 消息推送平台"
	this.TplNames = "dashboard/index.html"
}

func (this *DashboardController) Edit() {
	name := this.GetString("name", "")
	id, _ := this.GetInt("id", 0)
	if name == "" || id == 0 {
		this.SetError(libs.ERR_PARAM)
		return
	}

	err := models.NewAppModel().UpdateByParams(map[string]interface{}{"name": name}, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		this.SetError(libs.ERR_OPERATE)
		return
	}

	this.SetData(map[string]string{"name": name})
	return
}

func (this *DashboardController) DeleteById() {
	id, _ := this.GetInt("id", 0)
	if id == 0 {
		this.SetError(libs.ERR_PARAM)
		return
	}

	err := models.NewAppModel().DeleteById(id)
	if err != nil {
		this.SetError(libs.ERR_OPERATE)
		return
	}

	this.SetData(map[string]int{"id": id})
	return
}

func (this *DashboardController) Add() {
	name := this.GetString("name", "")
	if name == "" {
		this.SetError(libs.ERR_PARAM)
		return
	}

	id, err := models.NewAppModel().Add(name)
	if err != nil {
		this.SetError(libs.ERR_OPERATE)
		return
	}

	this.SetData(map[string]int64{"id": id})
}
