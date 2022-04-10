package repo

import (
	"atem-stock/models"
	"log"
)

func InsertStockToMg(mg *Mg, stocks []models.Stock) *Mg {
	if mg == nil {
		mg = NewMongo()
	}
	inface := make([]interface{}, len(stocks))
	for n, s := range stocks {
		inface[n] = s
	}
	if len(inface) > 0 {
		err := mg.Insert("admin", "stock", inface...)
		if err != nil {
			log.Println("插入【"+stocks[0].Code+"】失败", err)
		}

		//log.Println("插入【"+stocks[0].Code+"】成功")
	}
	return mg
}

func GetStockFromMg(mg *Mg, code string) (*Mg, *models.Stock) {
	if mg == nil {
		mg = NewMongo()
	}
	result := &models.Stock{}
	err := mg.FindSortLimit("admin", "stock", map[string]interface{}{"code": code}, "-date", 1, result)
	if err != nil {
		log.Println(" mg.FindSortLimit【"+code+"】未查到数据", err)
		return mg, nil
	}
	return mg, result
}

//使用cache
func GetStockCodesFromMg() []models.StockCode {
	mg := NewMongo()
	stockCodes := make([]models.StockCode, 1024)
	err := mg.FindAll("admin", "stockCode", nil, &stockCodes)
	if err != nil {
		log.Println("查询stockCodes失败")
	}
	return stockCodes
}

func UpdateStockCodeToMg(datas []interface{}) {
	mg := NewMongo()
	err := mg.RemoveAll("admin", "stockCode")
	if err != nil {
		log.Println("删除数据失败", err)
	}
	err = mg.Insert("admin", "stockCode", datas...)
	if err != nil {
		log.Println("插入数据失败", err)
	}
}
