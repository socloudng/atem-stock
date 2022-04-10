package main

import (
	"atem-stock/configs"
	"atem-stock/services"
	"log"
	
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	appConfig := &configs.AppConfig{}
	appConfig.LoadConfigByViper("application.yml")
	configs.AppConfigInstance = appConfig

	if appConfig.StockConfig.InitStockCode {
		wg.Add(1)
		go services.CatchStockCodeToCache(wg)
	}

	if appConfig.StockConfig.InitStock {
		wg.Add(1)
		go services.CatchStock(wg)
	}
	if appConfig.StockConfig.InitDapan {
		wg.Add(1)
		go services.CatchDapanData(wg) //如果需要大盘数据可以取消此行注释获取
	}
	wg.Wait()

	go func() {
		startAfterRun := appConfig.StockConfig.RunWait
		if startAfterRun > 0 {
			log.Println("需要等待"+ strconv.Itoa(startAfterRun) +"小时")
			
			time.Sleep(time.Hour * time.Duration(startAfterRun))
		}
		for {
			services.CatchStock(nil)
			time.Sleep(time.Hour * 12)
		}
	}()

	for {
		time.Sleep(time.Hour * 24 * 365 * 10)
	}
}
