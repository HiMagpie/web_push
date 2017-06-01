package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.Index()
}

func (this *HomeController) Index() {
	this.Data["title"] = "Android 消息推送平台"
	this.TplNames = "home/index.html"
}

func (this *HomeController) Doc() {
	this.Data["title"] = "参考文档"
	this.TplNames = "home/doc.html"
}
