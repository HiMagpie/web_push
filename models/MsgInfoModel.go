package models

/**
 * 消息内容
 */
type MsgInfoModel  struct {
	Id        int `json:"id"`
	MsgId     int64 `json:"msg_id"`
	Cids      string `json:"cids"`
	MsgTime   int `json:"msg_time"`
	Ring      int `json:"ring"`
	Vibrate   int `json:"vibrate"`
	Cleanable int `json:"cleanable"`
	Trans     int `json:"trans"`
	Begin     string `json:"begin"`
	End       string `json:"end"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Logo      string `json:"logo"`
	Url       string `json:"url"`
	Status    string `json:"status"`
	Enabled   int `json:"enabled"`
	Ctime     string `json:"ctime"`
	Utime     string `json:"utime"`
	BaseModel `orm:"-"`
}

func NewMsgInfoModel() *MsgInfoModel {
	return new(MsgInfoModel)
}

func (this *MsgInfoModel) TableName() string {
	return "hi_msg_info"
}
