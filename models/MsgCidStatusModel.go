package models

/**
 * 消息对应的Cid的推送状态
 */
type MsgCidStatusModel struct {
	Id      int `json:"id"`
	Cid     string `json:"cid"`
	MsgId   int64 `json:"msg_id"`
	Status  int `json:"status"`
	Enabled int `json:"enabled"`
	Ctime   string `json:"ctime"`
	Utime   string `json:"utime"`
	BaseModel `orm:"-"`
}

func NewMsgCidStatusModel() *MsgCidStatusModel {
	return new(MsgCidStatusModel)
}

func (this *MsgCidStatusModel) TableName() string {
	return "hi_msg_cid_status"
}
