package xxx

type XBet struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type XBetRequest struct {
	Issue    string `json:"issue"`
	GoldEggs int    `json:"totalgoldeggs"`
	Numbers  int    `json:"numbers"`
	Unix     string `json:"unix"`
	Keycode  string `json:"keycode"`
	PType    string `json:"ptype"`
	DeviceId string `json:"deviceid"`
	Userid   string `json:"userid"`
	Token    string `json:"token"`
}
