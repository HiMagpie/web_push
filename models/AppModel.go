package models

import (
	"web_push/libs/logger"
	"web_push/libs"
	"encoding/base64"
)

/**
 * 第三方应用的信息
 */
type AppModel  struct {
	Id           int `json:"id"`
	Name         string `json:"name"`
	AppId        string  `json:"app_id"`
	AppSecret    string `json:"app_secret"`
	MasterSecret string `json:"master_secret"`
	Enabled      int `json:"enabled"`
	Ctime        string `json:"ctime"`
	Utime        string `json:"utime"`
	BaseModel `orm:"-"`
}

func NewAppModel() *AppModel {
	return new(AppModel)
}

func (this *AppModel) TableName() string {
	return "hi_app_1"
}

/**
 * 获取分页列表
 */
func (this *AppModel) GetList(keyword string, page int, pageSize int) ([]AppModel, error) {
	var apps []AppModel

	offset := (page - 1) / pageSize
	q := this.GetQuery().From(this.TableName()).
	Where("enabled", ENABLED).OrderBy("id DESC").
	Offset(offset).Limit(pageSize)
	_, err := this.GetOrm().Raw(q.String(), q.GetArgs()).QueryRows(&apps)
	if err != nil {
		logger.Error("app.model", err.Error())
		return nil, err
	}

	return apps, nil
}

func (this *AppModel) GetByAppIdMaster(appId, masterSecret string) (*AppModel, error) {
	return this.GetByParams(map[string]interface{}{
		"app_id": appId,
		"master_secret": masterSecret,
	})
}

func (this *AppModel) GetByAppIdSecret(appId, appSecret string) (*AppModel, error) {
	return this.GetByParams(map[string]interface{}{
		"app_id": appId,
		"app_secret": appSecret,
	})
}

func (this *AppModel) GetByParams(params map[string]interface{}) (*AppModel, error) {
	q := this.GetQuery().From(this.TableName())
	for k, v := range params {
		q.Where(k, v)
	}
	q.Limit(1)

	err := this.GetOrm().Raw(q.String(), q.GetArgs()).QueryRow(this)
	if err != nil {
		logger.Error("app.model", logger.Format("err", err.Error()))
		return nil, err
	}

	return this, nil
}

func (this *AppModel) UpdateByParams(params map[string]interface{}, cond map[string]interface{}) error {
	q := this.GetQuery().Update(this.TableName(), params)
	for k, v := range cond {
		q.Where(k, v)
	}

	_, err := this.GetOrm().Raw(q.String(), q.GetArgs()).Exec()
	if err != nil {
		logger.Error("app.model.update.by.params", logger.Format("err", err.Error(), "sql", q.String(), "args", q.GetArgs()))
		return err
	}

	return nil
}

func (this *AppModel) DeleteById(id int) error {
	return this.UpdateByParams(map[string]interface{}{
		"enabled": DELETED,
	}, map[string]interface{}{"id": id})
}

func (this *AppModel) Add(name string) (int64, error) {
	if e, _ := this.IsExist(name); e {
		return int64(this.Id), nil
	}

	this.Name = name
	this.AppId = libs.Md5Str(name + this.GetNow())
	this.AppSecret = base64.StdEncoding.EncodeToString([]byte("as" + this.AppId))
	this.MasterSecret = base64.StdEncoding.EncodeToString([]byte("ms" + this.AppId))
	this.Enabled = ENABLED
	this.Ctime = this.GetNow()
	this.Utime = this.GetNow()
	return this.GetOrm().Insert(this)
}

func (this *AppModel) IsExist(name string) (bool, error) {
	q := this.GetQuery().From(this.TableName()).Where("name", name).Where("enabled", ENABLED)
	err := this.GetOrm().Raw(q.String(), q.GetArgs()).QueryRow(this)
	if err != nil {
		return false, err
	}

	return true, nil
}
