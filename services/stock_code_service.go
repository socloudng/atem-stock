package services

import (
	"atem-stock/cache"
	"atem-stock/remote"
	"atem-stock/repo"
	"log"
	"sync"
	"time"
)

var size = 200

func CatchStockCodeToCache(wg *sync.WaitGroup) {
	stockCodes, total, _ := remote.GeStockCodesPage(size, 1)
	cache.SetStockCodeToCache(stockCodes)
	pageCount := total/size + 1 //4800+票票,200一次需要循环25+次

	for i := 2; i <= pageCount; i++ {
		stockCodes, _, _ = remote.GeStockCodesPage(size, i)
		cache.SetStockCodeToCache(stockCodes)
		time.Sleep(time.Second * 3) //等待一下,免得被禁IP
	}
	if wg != nil {
		wg.Done()
	}
}

//请使用方法CatchStockCodeToCache
func CatchStockCodeToMg(wg *sync.WaitGroup) {
	log.Println("catch StockCode data begin ....")
	stockCodes := remote.GetStockCodes()
	codeMap := make(map[string]string, len(stockCodes))
	inface := make([]interface{}, len(stockCodes))
	i := 0
	for _, s := range stockCodes {
		if codeMap[s.Code] == "" {
			codeMap[s.Code] = s.Name
			inface[i] = s
			i++
		}
	}
	repo.UpdateStockCodeToMg(inface[:i])
	// time.Sleep(time.Second * 3) //等待一下,免得被禁IP
	log.Println("catch StockCode data end ....")
	wg.Done()
}
