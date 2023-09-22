package binancedataconnector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/Hellizer/lightlogger"
)

const (
	binanceFutures = "https://fapi.binance.com"
)

var httpClient *http.Client = &http.Client{Timeout: 10 * time.Second}

func GetServerTime() int64 {
	endpoint := binanceFutures + "/fapi/v1/time"
	resp, err := http.Get(endpoint)
	if err != nil {
		return 0
	}
	if resp.StatusCode != 200 {
		log.Print(5, log.LogError, "time request", fmt.Sprintf("ошибка: статус запроса времени : %v", resp.StatusCode))
		return 0
	}
	defer resp.Body.Close()
	mt := ExchangeServerTime{}
	err = json.NewDecoder(resp.Body).Decode(&mt)
	if err != nil {
		return 0
	}
	return mt.ServerTime
}

func send(method string, endpoint string, params *url.Values) *http.Response {
	querry := url.Values{}

	if params != nil {
		for k, v := range *params {
			querry.Add(k, v[0])
		}
	}

	qstr := querry.Encode()
	ep := endpoint + qstr
	req, err := http.NewRequest(method, ep, nil)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Print(2, log.LogError, "connector send", fmt.Sprintf("ошибка http клиента  : %v", err))
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		log.Print(2, log.LogError, "connector send", fmt.Sprintf("ошибка: статус ответа : %v", resp.StatusCode))
		return nil
	}
	for k, v := range resp.Header {
		if strings.Contains(k, "X-Mbx-Used-Weight") {
			log.Print(5, log.LogInfo, "client send", fmt.Sprintf("Лимиты:  %v: %v", k, v))
		}
	}
	return resp
}

func GetKlines(symbol string, interval string, limit int32) string {
	endpoint := binanceFutures + "/fapi/v1/klines?"
	querry := url.Values{}
	querry.Add("symbol", symbol)
	querry.Add("interval", interval)
	querry.Add("limit", fmt.Sprintf("%d", limit))
	resp := send("GET", endpoint, &querry)
	if resp == nil {
		log.Print(2, log.LogError, "klines request", "ошибка: пустой ответ")
		return ""
	}
	defer resp.Body.Close()
	//us := [][]Kline{}
	rstr, _ := io.ReadAll(resp.Body)
	return string(rstr[2 : len(rstr)-2])

	//dec := json.NewDecoder(resp.Body)
	//dec.UseNumber()
	//err := dec.Decode(&us)
	//err := json.Unmarshal(bstr, us)
	// if err != nil {
	// 	log.Print(2, log.LogError, "klines request", fmt.Sprintf("ошибка распаковки ответа: %v", err))
	// 	return ""
	// }
	// log.Print(1, log.LogInfo, "klines request", fmt.Sprintf("klines: %v ", us))

	// return ""
	//return us
}

func GetExchangeInfo() *ExchangeInfo {
	endpoint := binanceFutures + "/fapi/v1/exchangeInfo"
	resp := send("GET", endpoint, nil)
	if resp == nil {
		log.Print(2, log.LogError, "Exchange info request", "ошибка: пустой ответ")
		return nil
	}
	defer resp.Body.Close()
	ei := new(ExchangeInfo)
	err := json.NewDecoder(resp.Body).Decode(ei)
	if err != nil {
		log.Print(2, log.LogError, "Exchange info request", fmt.Sprintf("ошибка распаковки ответа: %v", err))
		return nil
	}
	return ei
}
