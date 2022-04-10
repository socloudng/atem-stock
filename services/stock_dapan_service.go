package services

import (
	"atem-stock/remote"
	"atem-stock/repo"
	"atem-stock/utils"
	"log"
	"sync"
	"time"
)

func CatchDapanData(wg *sync.WaitGroup) {
	var mg *repo.Mg
	log.Println("catch dapan data begin ....")
	code := "000001"                           //上证指数
	initDate := "20010101"                     //初始时间
	mg, stock := repo.GetStockFromMg(mg, code) //获取已经存取的大盘数据
	if stock != nil {
		d, err := time.Parse(utils.Layout_2, stock.Date)
		if err != nil {
			log.Fatal("时间解析错误", err)
		} else {
			initDate = d.Format(utils.Layout)
		}
	}
	t, err := time.Parse(utils.Layout, initDate)
	if err != nil {
		log.Fatal("时间解析错误", err)
	}
	beginDate, err := time.Parse(utils.Layout, initDate) //确定开始下载的时间
	endDate := t.Add(1 * 24 * time.Hour)
	for n := 1; endDate.Before(time.Now()); n++ { //按日期遍历下载数据
		stocks, err := remote.GetDpData(beginDate.Format(utils.Layout), endDate.Format(utils.Layout), code)
		if err != nil {
			log.Println(err)
		}
		if stocks != nil {
			mg = repo.InsertStockToMg(mg, stocks)
		}
		beginDate = endDate
		endDate = endDate.Add(1 * 24 * time.Hour)
	}
	log.Println("catch dapan data end ....")
	if wg != nil {
		wg.Done()
	}
}
