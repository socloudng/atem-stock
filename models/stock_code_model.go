package models

import "encoding/json"

type StockCode struct {
	Id       int    `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:主键编码"`
	Code     string `json:"code" gorm:"column:stock_code"`
	Name     string `json:"name" gorm:"column:stock_name"`
	Industry string `json:"industry" gorm:"column:stock_industry"`
}

func (sc *StockCode) MarshalBinary() (data []byte, err error) { return json.Marshal(sc) }
