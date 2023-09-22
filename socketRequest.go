package binancedataconnector

type MarketWSRequest struct {
	Method string   `json:"method"`           //
	Params []string `json:"params,omitempty"` //
	ID     int      `json:"id"`     //
}

type MarketWSResponse struct {
	ID     int      `json:"id"`               //
	Result []string `json:"result,omitempty"` //
}
