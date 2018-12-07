package servier

import (
	"github.com/robfig/cron"
	"log"
	"github.com/ticket/model/response"
	"net/http"
	"strings"
	"io/ioutil"
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"encoding/json"
	"github.com/ticket/dao"
	"github.com/ticket/model"
	"time"
	"strconv"
)
const KEY  = "your api key" // gate.io api key
const SECRET = "your api secret"  // gate.io api secret

func TimeGetioTicketFunc()  {
	c:=cron.New()
	spec:="3/5 * * * * ?"
	c.AddFunc(spec, func() {
		var dat map[string]interface{}
		ticklist:=tickers()
		err := json.Unmarshal(ticklist,&dat)
		if err !=nil{
			log.Println("json.Unmarsha1 出现错误",err)
		}
			UpdateTicker("btc_usdt","getio",dat)
			UpdateTicker("eth_usdt","getio",dat)
			UpdateTicker("eth_btc","getio",dat)
	})
	c.Start()
}
func TimeGetiomarketFunc()  {
	c:=cron.New()
	spec:="3/5 * * * * ?"
	c.AddFunc(spec, func() {
		var market response.Getiomarket
		marklist:=marketlist()
		err := json.Unmarshal(marklist,&market)
		if err !=nil{
			log.Println("json.Unmarsha1 出现错误",err)
		}
		for _,v:=range market.Data{
			Updatemarket("getio",v)
		}
	})
	c.Start()
}

func Updatemarket(channelname string ,market response.Marketmodle)  {
	if dao.MarketIsExit(market.Pair,channelname){
		// update
		mar:=model.Market{
			No:market.No,
			Symbol:market.Symbol,
			ChanelName:"getio",
			Name:market.Name,
			NameEn:market.Name_en,
			NameCn:market.Name_cn,
			Pair:market.Pair,
			Rate:market.Rate,
			VolA:market.Vol_a,
			VolB:market.Vol_b,
			CurrA:market.Curr_a,
			CurrB:market.Curr_b,
			CurrSuffix:market.Curr_suffix,
			RatePercent:market.Rate_percent,
			Trend:market.Trend,
			Supply:strconv.FormatFloat(float64(market.Supply), 'f', -1, 64),
			Plot:market.Lq,
			Updatetime:time.Now(),
		}
		result,err:=dao.EngineXorm.Where("pair=?",market.Pair).Update(&mar)
		if err !=nil{
			log.Fatal("btc_balance添加失败:",result, err)
		}
	}else{
		//insert

		mar:=model.Market{
			No:market.No,
			Symbol:market.Symbol,
			ChanelName:"getio",
			Name:market.Name,
			NameEn:market.Name_en,
			NameCn:market.Name_cn,
			Pair:market.Pair,
			Rate:market.Rate,
			VolA:market.Vol_a,
			VolB:market.Vol_b,
			CurrA:market.Curr_a,
			CurrB:market.Curr_b,
			CurrSuffix:market.Curr_suffix,
			RatePercent:market.Rate_percent,
			Trend:market.Trend,
			Supply:strconv.FormatFloat(float64(market.Supply), 'f', -1, 64),
			Plot:market.Lq,
			CreateTime:time.Now(),
			Updatetime:time.Now(),
		}
		result,err:=dao.EngineXorm.Insert(&mar)
		if err !=nil{
			log.Fatal("btc_balance添加失败:",result, err)
		}
	}

}
func UpdateTicker(pair string ,channelname string , tickermap map[string]interface{}) {
	mappair := tickermap[pair].(map[string]interface {})
	if dao.TickerIsExit(pair,channelname){
		// update
		ticket:=model.Ticket{
		ChanelName:"getio",
		Pair:pair,
		Result:mappair["result"].(string),
		Last:mappair["last"].(string),
		Lowestask:mappair["lowestAsk"].(string),
		Highestbid:mappair["highestBid"].(string),
		Percentchange:mappair["percentChange"].(string),
		Basevolume:mappair["baseVolume"].(string),
		Quotevolume:mappair["quoteVolume"].(string),
		High24hr:mappair["high24hr"].(string),
		Low24hr:mappair["low24hr"].(string),
		Updatetime:time.Now(),
		}
		result,err:=dao.EngineXorm.Where("pair=?",pair).Update(&ticket)
		if err !=nil{
			log.Fatal("btc_balance添加失败:",result, err)
		}
	}else{
		//inset
		var ticket model.Ticket
		ticket.Result=mappair["result"].(string)
		ticket.ChanelName="getio"
		ticket.Pair=pair
		ticket.Last=mappair["last"].(string)
		ticket.Lowestask=mappair["lowestAsk"].(string)
		ticket.Highestbid=mappair["highestBid"].(string)
		ticket.Percentchange=mappair["percentChange"].(string)
		ticket.Basevolume=mappair["baseVolume"].(string)
		ticket.Quotevolume=mappair["quoteVolume"].(string)
		ticket.High24hr=mappair["high24hr"].(string)
		ticket.Low24hr=mappair["low24hr"].(string)
		ticket.CreateTime=time.Now()
		ticket.Updatetime=time.Now()
		result,err:=dao.EngineXorm.Insert(&ticket)
		if err !=nil{
			log.Fatal("ticket insert failed:",result, err)
		}
	}
}

func httpDo(method string,url string, param string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(param))
	if err != nil {
	}
	var sign = getSign(param)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("key", KEY)
	req.Header.Set("sign", sign)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	return body
}
func getSign( params string) string {
	key := []byte(SECRET)
	mac := hmac.New(sha512.New, key)
	mac.Write([]byte(params))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

//// all support pairs
//func getPairs() string {
//	method := "GET"
//	url := "http://data.gateio.io/api2/1/pairs"
//	param := ""
//	ret := string(httpDo(method,url,param))
//	return ret
//}
//
//// Market Info
//func marketinfo() string {
//	method := "GET"
//	url := "http://data.gateio.io/api2/1/marketinfo"
//	param := ""
//	ret := string(httpDo(method,url,param))
//	return ret
//}


// Market Details
func marketlist() []byte {
	method :=  "GET"
	url := "http://data.gateio.io/api2/1/marketlist"
	param := ""
	ret := httpDo(method,url,param)
	return ret
}


// tickers
func tickers() []byte {
	method := "GET"
	url := "http://data.gateio.io/api2/1/tickers"
	param := ""
	 ret:=httpDo(method,url,param)
	return ret
}

//// ticker
//func ticker(ticker string) string {
//	var method string = "GET"
//	var url string = "http://data.gateio.io/api2/1/ticker" + "/" + ticker
//	var param string = ""
//	var ret string = string(httpDo(method,url,param))
//	return ret
//}
//
//
//// Depth
//func orderBooks() string {
//	var method string = "GET"
//	var url string = "http://data.gateio.io/api2/1/orderBooks"
//	var param string = ""
//	var ret string =string(httpDo(method,url,param))
//	return ret
//}
//
//
//// Depth of pair
//func orderBook(params string) string {
//	var method string = "GET"
//	var url string = "http://data.gateio.io/api2/1/orderBook/" + params
//	var param string = ""
//	var ret string = string(httpDo(method,url,param))
//	return ret
//}
//
//// Trade History
//func tradeHistory(params string) string {
//	var method string = "GET"
//	var url string = "http://data.gateio.io/api2/1/tradeHistory/" + params
//	var param string = ""
//	var ret string = string(httpDo(method,url,param))
//	return ret
//}
//
//
//// Get account fund balances
//func balances() string {
//	var method string = "POST"
//	var url string = "https://api.gateio.io/api2/1/private/balances"
//	var param string = ""
//	var ret string = string(httpDo(method,url,param))
//	return ret
//}
//
//
//
//// get deposit address
//func depositAddress(currency string) string {
//	var method string = "POST"
//	var url string = "https://api.gateio.io/api2/1/private/depositAddress"
//	var param string = "currency=" + currency
//	var ret string = string(httpDo(method,url,param))
//	return ret
//}
//
//
//// get deposit withdrawal history
//func depositsWithdrawals(start string, end string) string {
//	var method string = "POST"
//	var url string = "https://api.gateio.io/api2/1/private/depositsWithdrawals"
//	var param string = "start=" + start + "&end=" + end
//	var ret string =string(httpDo(method,url,param))
//	return ret
//}
//
//
//// Place order buy
//func buy(currencyPair string, rate string, amount string) string {
//	var method string = "POST"
//	var url string = "https://api.gateio.io/api2/1/private/buy"
//	var param string = "currencyPair=" + currencyPair + "&rate=" + rate + "&amount=" + amount
//	var ret string =string(httpDo(method,url,param))
//	return ret
//}
//
//// Place order sell
//func sell(currencyPair string, rate string, amount string) string {
//	var method string = "POST"
//	var url string = "https://api.gateio.io/api2/1/private/sell"
//	var param string = "currencyPair=" + currencyPair + "&rate=" + rate + "&amount=" + amount
//	var ret string = string(httpDo(method,url,param))
//	return ret
//}
//
//
//// Cancel order
//func cancelOrder(orderNumber string, currencyPair string ) string {
//	var method string = "POST"
//	var url string = "https://api.gateio.io/api2/1/private/cancelOrder"
//	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
//	var ret string = string(httpDo(method,url,param))
//	return ret
//}
//
//// Cancel all orders
//func cancelAllOrders( types string, currencyPair string ) string {
//	var method string = "POST"
//	var url string = "https://api.gateio.io/api2/1/private/cancelAllOrders"
//	var param string = "type=" + types + "&currencyPair=" + currencyPair
//	var ret string = string(httpDo(method,url,param))
//	return ret
//}
//
//
//// Get order status
//func getOrder( orderNumber string, currencyPair string ) string {
//	var method string = "POST"
//	var url string = "https://api.gateio.io/api2/1/private/getOrder"
//	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
//	var ret string = string(httpDo(method,url,param))
//	return ret
//}
//
//
//// Get my open order list
//func openOrders() string {
//	var method string = "POST"
//	var url string = "https://api.gateio.io/api2/1/private/openOrders"
//	var param string = ""
//	var ret string = string(httpDo(method,url,param))
//	return ret
//}
//
//
//// 获取我的24小时内成交记录
//func myTradeHistory( currencyPair string, orderNumber string) string {
//	var method string = "POST"
//	var url string = "https://api.gateio.io/api2/1/private/tradeHistory"
//	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
//	var ret string = string(httpDo(method,url,param))
//	return ret
//}
//// Get my last 24h trades
//func withdraw( currency string, amount string, address string) string {
//	var method string = "POST"
//	var url string = "https://api.gateio.io/api2/1/private/withdraw"
//	var param string = "currency=" + currency + "&amount=" + amount + "address=" + address
//	var ret string = string(httpDo(method,url,param))
//	return ret
//}