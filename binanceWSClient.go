package binancedataconnector

//TODO:сделать поддержку открытия новых потоков для поддержки более 200 событий
import (
	"encoding/json"
	"fmt"
	"sync"

	ws "github.com/Hellizer/wsclient"
	log "github.com/Hellizer/lightlogger"
)

const (
	binanceMarketStream = "wss://fstream.binance.com/stream?streams="
)

type MsgHandler func(msg []byte)
type MarketWsClient struct {
	ws *ws.WsClient //
	//	reqID   int          //
	handler MsgHandler //
}

var instance *MarketWsClient = nil
var lock sync.Mutex

func GetMarketWsClient() *MarketWsClient {
	if instance != nil {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()
	if instance != nil {
		return instance
	}
	wsClient := new(MarketWsClient)
	wsClient.handler = defaultHandler
	wsClient.ws = ws.NewClient(wsClient.baseEventHandler)
	//wsClient.Open(startStream)
	//TODO: сделать переподключение
	instance = wsClient
	return instance
}

func (mc *MarketWsClient) SetHandler(handler MsgHandler) {
	mc.handler = handler
}

func (mc *MarketWsClient) Open(stream string) error {
	//str := strings.ToLower(stream)
	err := mc.ws.Open(binanceMarketStream + stream)
	if err != nil {
		log.Print(1, log.LogError, "binance market ws", fmt.Sprintf("error: %v", err))
		return err
	}
	mc.ws.Serve()
	return nil
}

func (mc *MarketWsClient) Subscribe(stream []string) error {
	err := mc.sendRequest(&MarketWSRequest{Method: "SUBSCRIBE", Params: stream})
	if err != nil {
		return err
	}
	//log.Print(1, log.LogInfo, "binance market ws", fmt.Sprint(resp))
	return nil
}

func (mc *MarketWsClient) UnSubscribe(stream []string) error {
	err := mc.sendRequest(&MarketWSRequest{Method: "UNSUBSCRIBE", Params: stream})
	if err != nil {
		return err
	}
	//log.Print(1, log.LogInfo, "binance market ws", fmt.Sprint(resp))
	return nil
}

func (mc *MarketWsClient) GetSubscription() error {
	err := mc.sendRequest(&MarketWSRequest{Method: "LIST_SUBSCRIPTIONS"})
	if err != nil {
		return err
	}
	//log.Print(1, log.LogInfo, "binance market ws", fmt.Sprint(resp))
	return nil
}

func (mc *MarketWsClient) sendRequest(req *MarketWSRequest) error {
	//lock.Lock()
	//req.ID = mc.reqID
	//mc.reqID++
	//TODO: добавить ожидание результата
	//	lock.Unlock()
	bstr, _ := json.Marshal(req)
	err := mc.ws.SendText(bstr)
	if err != nil {
		return err
	}

	return nil
}

func defaultHandler(msg []byte) {
	// if msg.Err != nil {
	// 	log.Print(1, log.LogError, "msg handler", string(msg.Err.Error()))
	// }
	log.Print(1, log.LogInfo, "msg default handler", string(msg))
}

func (mc *MarketWsClient) baseEventHandler(msg ws.WSMessage) {
	//TODO: добавить отлов событий
	//fmt.Println(string(msg.RawMsg))
	if mc.handler != nil {
		mc.handler(msg.RawMsg)
	} else {
		defaultHandler(msg.RawMsg)
	}
}

func (mc *MarketWsClient) Close() error {
	log.Print(1, log.LogInfo, "binance ws close", "закрываем вебсокет")
	mc.handler = nil
	err := mc.ws.Close()
	if err != nil {
		log.Print(1, log.LogError, "binance ws close", fmt.Sprintf("closing error +%v ", err))
		return err //
	}
	log.Print(1, log.LogInfo, "binace ws close", "вебсокет закрыт")
	return nil
}
