package cache

import (
	"atem-stock/models"
)

var stock_code_key = "stock_code"

func SetStockCodeToCache(stockCodes []models.StockCode) {
	codeMap := make(map[string]interface{}, len(stockCodes))
	for _, sc := range stockCodes {
		codeMap[sc.Code] = &sc
	}
	redisHMSet(stock_code_key, codeMap)
}

func GetStockCodesFromCache() []string {
	return redisHGetAllFields(stock_code_key)
}
func GetStockCodesFromCacheMap() []string {
	return redisHGetAllFields(stock_code_key)
}
