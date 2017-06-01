package api

import (
	"web_push/libs"
	"web_push/models"
)

type AppController struct {
	BaseController
}

/**
 * 验证第三方应用是否有权限发送推送消息
 */
func (this *AppController) ValidMaster() {
	appId := this.GetString("app_id", "")
	masterSecret := this.GetString("master_secret", "")

	app, err := models.NewAppModel().GetByAppIdMaster(appId, masterSecret)
	if err != nil {
		this.setError(libs.ERR_OPERATE, Data{
			"app_id": appId,
			"master_secret": masterSecret,
		})
		return
	}

	this.setData(Data{
		"name": app.Name,
	})
	return
}

/**
 * 判断第三方应用是否合法
 */
func (this *AppController) ValidApp()  {
	appId := this.GetString("app_id", "")
	appSecret := this.GetString("app_secret", "")

	app, err := models.NewAppModel().GetByAppIdSecret(appId, appSecret)
	if err != nil {
		this.setError(libs.ERR_OPERATE, Data{
			"app_id": appId,
			"app_secret": appSecret,
		})
		return
	}

	this.setData(Data{
		"name": app.Name,
	})
	return
}