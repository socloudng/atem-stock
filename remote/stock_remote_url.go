package remote

import (
	"strconv"
)

func getUrlOfSohuUrl(stockCode string, startDate string, endDate string) string {
	return "http://q.stock.sohu.com/hisHq?" +
		"code=" + "zs_" + stockCode +
		"&start=" + startDate + "&end=" + endDate +
		"&stat=1&order=D&period=d&callback=historySearchHandler&rt=jsonp&r=0.09105574639477387&0.021587371893673213"
}

//
//fields=f2,f3,f12,f13,f14,f62,f184,f225,f165,f263,f109,f175,f264,f160,f100,f124,f265,f1
func getStockListPageUrl(pageSize int, pageNo int) string {
	return "https://push2.eastmoney.com/api/qt/clist/get?" +
		"fields=f12,f14,f100" +
		"&fs=m:0+t:6+f:!2,m:0+t:13+f:!2,m:0+t:80+f:!2,m:1+t:2+f:!2,m:1+t:23+f:!2,m:0+t:7+f:!2,m:1+t:3+f:!2" +
		"&pz=" + strconv.Itoa(pageSize) +
		"&pn=" + strconv.Itoa(pageNo)
}

func getStockDailyDataListPageUrl(code, fromDate, endDate string) string {

	return "http://app.finance.ifeng.com/hq/stock_daily.php?" +
		"code=" + code +
		"&begin_day" + fromDate +
		"&end_day=" + endDate
}

func getStockListHtmlUrl() (url string, search string) {
	//return "http://quote.eastmoney.com/stocklist.html#sh", "#quotesearch ul li"
	// return "https://data.eastmoney.com/zjlx/list.html", "#dataview tbody tr"
	return "https://quote.eastmoney.com/center/gridlist.html#hs_a_board", "#table_wrapper-table tbody"
}

func getStockHistoryAigaogaoUrl(code string) string {
	url := "https://www.aigaogao.com/tools/history.html?s=" + code
	return url
}
