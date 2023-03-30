package xxx

var standard = map[int]int{
	0:  1000,
	1:  333,
	2:  166,
	3:  100,
	4:  66,
	5:  48,
	6:  36,
	7:  28,
	8:  22,
	9:  18,
	10: 16,
	11: 15,
	12: 14,
	13: 13,
	14: 13,
	15: 14,
	16: 15,
	17: 16,
	18: 18,
	19: 22,
	20: 28,
	21: 36,
	22: 48,
	23: 66,
	24: 100,
	25: 166,
	26: 333,
	27: 1000,
}

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
