package binancedataconnector

type Event struct {
	EventType string `json:"e"` //Event Type
	EventTime uint64 `json:"E"` //Event Time
}

type TradeEvent struct {
	Event
	TradeTime    uint64 `json:"T"` //
	Symbol       string `json:"s"` //
	TradeID      uint64 `json:"t"` //
	Price        string `json:"p"` //
	Quantity     string `json:"q"` //
	X            string `json:"X"` //
	IsBuyerMaker bool   `json:"m"` //

}

type StreamEvent struct {
	Stream string `json:"stream"` //
}

type StreamKlineEvent struct {
	Stream string     `json:"stream"` //
	Data   KlineEvent `json:"data"`   //
}

type KlineEvent struct {
	Event
	Symbol string      `json:"s"` //Symbol
	Kline  KlineUpdate `json:"k"` //Kline / Candlestick
}

type KlineUpdate struct {
	OpenTime         int64  `json:"t"` //Kline start time
	CloseTime        int64  `json:"T"` //Kline close time
	Symbol           string `json:"s"` //Symbol
	Interval         string `json:"i"` //Interval
	FirstTradeID     int64  `json:"f"` //First trade ID
	LastTradeID      int64  `json:"L"` //Last trade ID
	OpenPrice        string `json:"o"` //Open price
	ClosePrice       string `json:"c"` //Close price
	HighPrice        string `json:"h"` //High price
	LowPrice         string `json:"l"` //Low price
	BaseAssetVolume  string `json:"v"` //Base asset volume
	TradesCount      int64  `json:"n"` //Number of trades
	IsClosed         bool   `json:"x"` //Is this kline closed?
	QuoteAssetVolume string `json:"q"` //Quote asset volume
	TakerBuyBase     string `json:"V"` //Taker buy base asset volume
	TakerBuyQutoe    string `json:"Q"` //Taker buy quote asset volume
	Ignore           string `json:"B"` //Ignore
}
