package services

import (
	"atem-stock/cache"
	"atem-stock/models"
	"atem-stock/remote"
	"atem-stock/repo"
	"log"
	"sync"
	"time"

	"atem-stock/utils"
)

func CatchStock(wg *sync.WaitGroup) {
	var mg *repo.Mg
	log.Println("catch cacthStock data begin ....")
	stockCodes := cache.GetStockCodesFromCache()
	for _, code := range stockCodes {
		stocks, err := remote.GetStocks(code)
		if err != nil {
			log.Println("查询 CacthStock【"+code+"】失败", err)
			continue
		}
		mg, savedStock := repo.GetStockFromMg(mg, code)
		if savedStock != nil {
			stocks = filterStock(stocks, savedStock) //过滤出未保存的
		}
		if len(stocks) < 1 {
			continue
		}
		mg = repo.InsertStockToMg(mg, stocks)
	}
	log.Println("catch cacthStock data end ....")
	if wg != nil {
		wg.Done()
	}
}

func filterStock(stocks []models.Stock, st *models.Stock) []models.Stock {
	result := make([]models.Stock, len(stocks))
	stTime, _ := time.Parse(utils.Layout_2, st.Date)
	i := 0
	for _, stock := range stocks {
		stockTime, _ := time.Parse(utils.Layout_2, stock.Date)
		if stockTime.After(stTime) {
			result[i] = stock
			i++
		}
	}
	return result[:i]
}
