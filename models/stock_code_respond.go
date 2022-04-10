package models

type StockCodeResponse struct {
	Rt     int           `json:"rt"`
	Rc     int           `json:"rc"`
	Svr    int64         `json:"svr"`
	It     int           `json:"it"`
	Result StockCodePage `json:"data"`
}

type StockCodePage struct {
	Total    int                   `json:"total"`
	DataList map[int]StockCodeData `json:"diff"`
}

type StockCodeData struct {
	Code     string `json:"f12"`
	Name     string `json:"f14"`
	Industry string `json:"f100"`
}

type StockDapanResult struct {
	Status int
	Hq     [][]string
	Code   string
}

type StockData struct {
	Code          string `json:"f12"`
	Date          string `json:"date"`
	Close         string `json:"f2"`
	Open          string `json:"f17"`
	Low           string `json:"f16"`
	High          string `json:"f15"`
	Change        string `json:"f4"`
	ChangePercent string `json:"f3"`
	DealCount     string `json:"f5"`
	DealAmount    string `json:"f6"`
}
