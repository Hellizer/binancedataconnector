package binancedataconnector

type ExchangeServerTime struct {
	ServerTime int64 `json:"serverTime"` //Test connectivity to the Rest API and get the current server time.
}

type UntypedStruct struct {
	Data interface{}
}

type Klines struct {
	Klines []Kline
}

type Kline struct {
	OpenTime                 int64
	Open                     string
	High                     string
	Low                      string
	Close                    string
	Volume                   string
	CloseTime                int64
	QuoteAssetVolume         string
	NumberOfTrades           int64
	TakerBuyBaseAssetVolume  string
	TakerBuyQuoteAssetVolume string
	Ignore                   string
}

type ExchangeInfo struct {
	TimeZone        string        `json:"timezone"`
	ServerTime      int64         `json:"serverTime"`
	FuturesType     string        `json:"futuresType"`
	RateLimits      []RateLimit   `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Assets          []Asset       `json:"assets"`
	Symbols         []Symbol      `json:"symbols"`
}
type RateLimit struct {
	Ratelimittype string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	Intervalnum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
}
type Asset struct {
	Asset             string `json:"asset"`
	MarginAvailable   bool   `json:"marginAvailable"`
	AutoAssetExchange string `json:"autoAssetExchange"`
}
type Filter struct {
	MinPrice          string `json:"minPrice,omitempty"`
	MaxPrice          string `json:"maxPrice,omitempty"`
	FilterType        string `json:"filterType"`
	TickSize          string `json:"tickSize,omitempty"`
	StepSize          string `json:"stepSize,omitempty"`
	MaxQty            string `json:"maxQty,omitempty"`
	MinQty            string `json:"minQty,omitempty"`
	Limit             int    `json:"limit,omitempty"`
	Notional          string `json:"notional,omitempty"`
	MultiplierDown    string `json:"multiplierDown,omitempty"`
	MultiplierUp      string `json:"multiplierUp,omitempty"`
	MultiplierDecimal string `json:"multiplierDecimal,omitempty"`
}
type Symbol struct {
	Symbol                string   `json:"symbol"`
	Pair                  string   `json:"pair"`
	ContractType          string   `json:"contractType"`
	DeliveryDate          int64    `json:"deliveryDate"`
	OnboardDate           int64    `json:"onboardDate"`
	Status                string   `json:"status"`
	MaintMarginPercent    string   `json:"maintMarginPercent"`
	RequiredMarginPercent string   `json:"requiredMarginPercent"`
	BaseAsset             string   `json:"baseAsset"`
	QuoteAsset            string   `json:"quoteAsset"`
	MarginAsset           string   `json:"marginAsset"`
	PricePrecision        int      `json:"pricePrecision"`
	QuantityPrecision     int      `json:"quantityPrecision"`
	BasAassetPrecision    int      `json:"baseAssetPrecision"`
	QuotePrecision        int      `json:"quotePrecision"`
	UnderlyingType        string   `json:"underlyingType"`
	UnderlyingSubType     []string `json:"underlyingSubType"`
	SettlePlan            int      `json:"settlePlan"`
	TriggerProtect        string   `json:"triggerProtect"`
	LiquidationFee        string   `json:"liquidationFee"`
	MarketTakeBound       string   `json:"marketTakeBound"`
	Filters               []Filter `json:"filters"`
	OrderTypes            []string `json:"orderTypes"`
	TimeInForce           []string `json:"timeInForce"`
}
