package xxx

type HistoryItem struct {
	Issue  string `json:"issue"`
	Result string `json:"lresult"`
}

type HistoryData struct {
	Items []HistoryItem `json:"items"`
}

type QHistory struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`

	Data HistoryData `json:"data"`
}

type QHistoryRequest struct {
	PageSize int    `json:"pagesize"`
	Unix     string `json:"unix"`
	KeyCode  string `json:"keycode"`
	PType    string `json:"ptype"`
	DeviceId string `json:"deviceid"`
	UserId   string `json:"userid"`
	Token    string `json:"token"`
}
