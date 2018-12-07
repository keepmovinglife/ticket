package response

type Getiomarket struct {
	Result string
	Data   []Marketmodle
}
type Marketmodle struct {
	No           int64    `json:no`
	Symbol       string `json:"symbol"`
	Name         string
	Name_en      string
	Name_cn      string
	Pair         string
	Rate         string
	Vol_a        string
	Vol_b        string
	Curr_a       string
	Curr_b       string
	Curr_suffix  string
	Rate_percent string
	Trend        string
	Supply       float64
	marketcap    string
	Lq           string

}

type GetioTickers struct {
	Tickmap  map[string]Tickersmodel
}
type Tickersmodel struct {
	Pair TickPair
}
type TickPair struct {
	Result        string
	Last          string
	LowestAsk     string
	HighestBid    string
	PercentChange string
	BaseVolume    string
	QuoteVolume   string
	High24hr      string
	Low24hr       string
}
