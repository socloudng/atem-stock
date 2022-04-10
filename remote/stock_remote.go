package remote

import (
	"atem-stock/models"
	"atem-stock/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
)

//分页获取StockCodes
func GeStockCodesPage(pageSize, pageNo int) ([]models.StockCode, int, error) {
	url := getStockListPageUrl(pageSize, pageNo)
	getUrlResp, err := http.Get(url)
	if err != nil {
		log.Println("查询股票代码失败 :("+url+")", err)
		return nil, 0, err
	}
	bodyData, err := ioutil.ReadAll(getUrlResp.Body)
	if err != nil {
		log.Println("股票代码读取失败 :("+url+")", bodyData)
		return nil, 0, err
	}
	response := &models.StockCodeResponse{}
	err = json.Unmarshal(bodyData, response)
	if err != nil {
		log.Println("股票代码解析失败 :("+url+")", err)
		return nil, 0, err
	}
	stockCodes := make([]models.StockCode, len(response.Result.DataList))
	for i, value := range response.Result.DataList {
		stockCodes[i] = models.StockCode{
			Name:     value.Name,
			Code:     value.Code,
			Industry: value.Industry,
		}
	}
	return stockCodes, response.Result.Total, nil
}

func GetDpData(startDate string, endDate string, stockCode string) ([]models.Stock, error) {
	url := getUrlOfSohuUrl(stockCode, startDate, endDate)
	req, err := http.NewRequest("GET", url, nil)
	fmt.Println("url :" + url)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http request error")
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("result : " + string(result))
	if len(result) < 100 {
		return nil, nil
	}
	result = []byte(string(result)[strings.Index(string(result), "{"):strings.LastIndex(string(result), "]")])
	stockResult := &models.StockDapanResult{}
	err = json.Unmarshal(result, stockResult)
	if err != nil {
		return nil, errors.New("解析报文异常")
	}
	if stockResult.Status != 0 {
		return nil, errors.New("get data error")
	}
	if len(stockResult.Hq) <= 0 {
		return nil, errors.New("not get any data")
	}

	stocks := make([]models.Stock, len(stockResult.Hq))
	for n, hq := range stockResult.Hq {
		s := models.Stock{
			Date:          string(hq[0]),
			Start:         string(hq[1]),
			End:           string(hq[2]),
			Change:        string(hq[3]),
			ChangePercent: string(hq[4]),
			Low:           string(hq[5]),
			High:          string(hq[6]),
			DealCount:     string(hq[7]),
			DealAmount:    string(hq[8]),
			Code:          stockCode,
		}
		stocks[n] = s
	}
	return stocks, nil
}

func GetStocks(code string) ([]models.Stock, error) {
	url := getStockHistoryAigaogaoUrl(code)
	//decoder := simplifiedchinese.GBK.NewDecoder()
	docs, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	selections := docs.Find("#ctl16_contentdiv table tr")

	if selections == nil || len(selections.Nodes) == 0 {
		return nil, errors.New("not find data " + code)
	}

	arrys := make([]string, len(selections.Nodes))
	var index = 0
	selections.Each(func(i int, selection *goquery.Selection) {
		tds := selection.Find(".altertd")
		if tds != nil && len(tds.Nodes) > 0 {
			var str string
			tds.Each(func(i int, selection *goquery.Selection) {
				str += (selection.Text() + ";")

			})
			if len(strings.TrimSpace(str)) != 0 {
				arrys[index] = str
				index++
			}
		}
	})
	stocks := make([]models.Stock, index)
	for i := 0; i < index; i++ {
		tds := strings.Split(arrys[i], ";")
		t, _ := time.Parse(utils.Layout_4, tds[0])
		stocks[i] = models.Stock{
			Code:          code,
			Date:          t.Format(utils.Layout_2),
			Start:         tds[1],
			High:          tds[2],
			Low:           tds[3],
			End:           tds[4],
			DealCount:     tds[5],
			DealAmount:    tds[6],
			Change:        tds[7],
			ChangePercent: tds[8],
		}

	}
	return stocks, nil
}

//因东方财富接口变更,已失效
func GetStockCodes() []models.StockCode {
	url, search := getStockListHtmlUrl()
	decoder := simplifiedchinese.GBK.NewDecoder()
	docs, err := goquery.NewDocument(url)
	if err != nil {
		log.Println("查询股票代码失败 :("+url+")", err)
	}
	size := len(docs.Find(search).Nodes)
	stockCodes := make([]models.StockCode, size)
	docs.Find(search).Each(
		func(i int, contentSelection *goquery.Selection) {
			dst := make([]byte, 2*len(contentSelection.Text()))
			decoder.Transform(dst, []byte(contentSelection.Text()), true)
			codeStr := utils.ByteString(dst)
			arry := strings.Split(codeStr, "(")
			stockCodes[i] = models.StockCode{
				Name: arry[0],
				Code: arry[1][:len(arry[1])-1],
			}
		})

	return stockCodes
}
