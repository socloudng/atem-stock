package models

type Stock struct {
	Id            int    `json:"id" 	gorm:"column:id;primaryKey;autoIncrement;comment:主键编码"`
	Code          string `json:"code" 	gorm:"column:stock_code"`
	Date          string `json:"date" 	gorm:"column:stock_date;	comment:日期"`
	Start         string `json:"open" 	gorm:"column:stock_open;	comment:开盘"`
	End           string `json:"close" 	gorm:"column:stock_close;	comment:收盘"`
	Low           string `json:"low" 	gorm:"column:stock_low;		comment:最低"`
	High          string `json:"high" 	gorm:"column:stock_high;	comment:最高"`
	Change        string `json:"change" gorm:"column:stock_change;	comment:涨幅"`
	ChangePercent string `json:"change_percent" gorm:"column:stock_change_percent;comment:涨幅百分比"`
	DealCount     string `json:"vol" 	gorm:"column:stock_vol;		comment:成交量"`
	DealAmount    string `json:"amount" gorm:"column:stock_amount;	comment:成交额"`
}
