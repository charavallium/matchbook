package matchbook

import (
	"time"

	"github.com/shopspring/decimal"
)

type LoginResponse struct {
	SessionToken string         `json:"session-token"`
	UserId       string         `json:"user-id"`
	Role         string         `json:"role"`
	Account      AccountDetails `json:"account"`
	Email        string         `json:"email"`
	PhoneNumber  string         `json:"phone-number"`
	Address      AddressDetails `json:"address"`
}

type AccountDetails struct {
	Id       int         `json:"id"`
	Username string      `json:"username"`
	Name     NameDetails `json:"name"`
}

type NameDetails struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

type BaseResponse struct {
	Total   int `json:"total"`
	PerPage int `json:"per-page"`
	Offset  int `json:"offset"`
}
type Cancelled struct {
	BaseResponse
	Bets []MBets
}
type MBets struct {
	Id        int64   `json:"id"`
	Offer_id  int64   `json:"offer-id"`
	Odds      float64 `json:"odds"`
	Odds_type string  `json:"odds-type"`
	Decimal   float64 `json:"decimal-odds"`
	Stake     float64 `json:"stake"`
	Pl        float64 `json:"potential-liability"`
	Comission float64 `json:"comission"`
	Currency  string  `json:"currency"`
	Date      string  `json:"created-at"`
	Status    string  `json:"status"`
}
type AddressDetails struct {
}
type Balance struct {
	Id                 int     `json:"id"`
	Balance            float64 `json:"balance"`
	Exposure           float64 `json:"exposure"`
	Commission_reserve float64 `json:"commission-reserve"`
	Free_funds         float64 `json:"free-funds"`
}
type Sport struct {
	BaseResponse
	Name string `json:"name"`
	Type string `json:"type"`
	Id   int    `json:"id"`
}

type Events struct {
	BaseResponse
	Events []Event `json:"events"`
}

type Event struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	SportId          int       `json:"sport-id"`
	Start            time.Time `json:"start"`
	InRunningFlag    bool      `json:"in-running-flag"`
	AllowLiveBetting bool      `json:"allow-live-betting"`
	CategoryId       []int     `json:"category-id"`
	Status           string    `json:"status"`
	Volume           float32   `json:"volume"`
	Markets          []Market  `json:"markets"`
	MetaTags         []MetaTag `json:"meta-tags"`
}
type Event_params struct {
	Exchange_type string
	/*Prices are returned in accordance with
	the specified exchange-type. Allowed values:
	back-lay, binary.*/
	Odds_type string
	/*Prices are returned in accordance with the
	specified odds-type. Allowed values:
	DECIMAL, US, HK, MALAY, INDO, %.*/
	Price_depth string
	/*Value indicating, for each runner,
	the maximum number of prices that will
	be returned in the response for each side.*/
	Side string
	/*Only prices on the specified side are included
	 in the response. Allowed values are:
	 back, lay or win, lose.
	Both sides are returned if no side is specified.*/
	Currency string
	/*Prices are returned in the specified currency.
	Allowed values: USD, EUR, GBP, AUD, CAD. HKD*/
	Minimum_liquidity float64
	/*Only prices with available-amount greater than
	or equal to this value are included in the response.*/
	Include_event_participants bool
	/*A boolean indicating whether to return the event
	participants information.*/

}
type Events_params struct {
	Offset string
	/*Used for paging. The offset of the first entry in the response.*/
	Per_page string
	/*Used for paging. The number of entries to return in a single response.*/
	After int64
	/*A unix epoch timestamp. Only events starting after the
	  timestamp will be returned.*/
	Before int64
	/*A unix epoch timestamp. Only events starting before the
	  timestamp will be returned.*/
	Category_ids string
	/*A comma separated list of category ids.
	   Only events on the provided categories are
	  included in the response.
	*/
	Ids string
	/*A comma separated list of event ids. Only events with id
	  in the provided list are included in the response.*/
	Sport_ids string
	/*A comma separated list of sport ids. Only events with
	  sport in the provided list are included in the response.*/
	States string
	/*A comma separated list of event states. Only events with
	  status in the provided list are included in the response.*/
	Tag_url_names string
	/*A comma separated list of url-names. Only events with tags
	  having url-name in the provided list are included in the response.*/
	Exchange_type string
	/*Prices are returned in accordance with the specified exchange-type.
	  Allowed values: back-lay, binary.*/
	Odds_type string
	/*Prices are returned in accordance with the specified odds-type.
	  Allowed values: DECIMAL, US, HK, MALAY, INDO, %.*/
	Price_depth int32
	/*Value indicating, for each runner, the maximum number of
	  prices that will be returned in the response for each side.*/
	Side string
	/*Only prices on the specified side are included in the response.
	  Allowed values are: back, lay or win, lose. Both sides are returned
	  if no side is specified.*/
	Currency string
	/*Prices are returned in the specified currency.
	  Allowed values: USD, EUR, GBP, AUD, CAD. HKD*/
	Minimum_liquidity float64
	/*Only prices with available-amount greater than or equal to this
	  value are included in the response.*/
	Include_event_participants bool
	/*A boolean indicating whether to return the
	  event participants information.*/

}
type Market struct {
	//Live             bool      `json:"live"`
	EventId          int       `json:"event-id"`
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Runners          []Runner  `json:"runners"`
	Start            time.Time `json:"start"`
	InRunningFlag    bool      `json:"in-running-flag"`
	AllowLiveBetting bool      `json:"allow-live-betting"`
	Status           string    `json:"status"`
	MarketType       string    `json:"market-type"`
	Type             string    `json:"type"`
	Volume           float32   `json:"volume"`
	BackOverround    float32   `json:"back-overround"`
	LayOverround     float32   `json:"lay-overround"`
	BaseResponse
}
type Runners struct {
	Runners []Runner `json:"runners"`
	Total   int64    `json:"total"`
}
type Market_params struct {
	Exchange_type string
	/*Prices are returned in accordance with
	 the specified exchange-type.
	Allowed values: back-lay, binary.*/
	Odds_type string
	/*Prices are returned in accordance with the
	  specified odds-type. Allowed values:
	   DECIMAL, US, HK, MALAY, INDO, %.*/
	Price_depth int32
	/*Value indicating, for each runner, the maximum
	  number of prices that will be returned in the
	  response for each side.*/
	Side string
	/*Only prices on the specified side are included
	  in the response. Allowed values are:
	   back, lay or win, lose.
	  Both sides are returned if no side is specified.*/
	Currency string
	/*Prices are returned in the specified currency.
	  Allowed values: USD, EUR, GBP, AUD, CAD. HKD*/
	Minimum_liquidity float64
}
type Markets struct {
	Markets []Market `json:markets`
	BaseResponse
}
type Markets_params struct {
	Offset string
	/*Used for paging. The offset of the first entry
	  in the response.*/
	Per_page string
	/*Used for paging. The number of entries to return
	  in a single response.*/
	Names string
	/*A comma separated list of market names. Only markets
	  with name in the provided list are included in the response.*/
	States string
	/*A comma separated list of market states. Only markets with status
	  in the provided list are included in the response.*/
	Types string
	/*A comma separated list of market types. Only markets with type in
	  the provided list are included in the response.*/
	Exchange_type string
	/*Prices are returned in accordance with the specified exchange-type.
	  Allowed values: back-lay, binary.*/
	Odds_type string
	/*Prices are returned in accordance with the specified odds-type.
	  Allowed values: DECIMAL, US, HK, MALAY, INDO, %.*/
	Price_depth string
	/*Value indicating, for each runner, the maximum number of prices
	  that will be returned in the response for each side.*/
	Side string
	/*Only prices on the specified side are included in the response.
	  Allowed values are: back, lay or win, lose. Both sides are returned
	  if no side is specified.*/
	Currency string
	/*Prices are returned in the specified currency.
	  Allowed values: USD, EUR, GBP, AUD, CAD. HKD*/
	Minimum_liquidity float64
	/*Only prices with available-amount greater
	  than or equal to this value are included in the response.*/

}
type Sports struct {
	BaseResponse
	Sports []Sport `json:sports`
}
type Sports_params struct {
	Offset int64
	/*Used for paging. The offset of the first entry in the response.*/
	Per_page int64
	/*Used for paging. The number of entries to return in a single response.*/
	Order string
	/*Used for sorting results by name or id in ascending or descending order. Allowed values: name asc, name desc, id asc, id desc.*/
	Status string
	/*A comma separated list of statuses. Only sports with status in the provided list are included in the response.*/
}
type Navigation_Params struct {
	Offset int64
	/*Used for paging. The offset of the first entry in the response.*/
	Per_page int64
	/*Used for paging. The number of entries to return in a single response.*/
}
type Runner struct {
	Prices             []PriceDetail `json:"prices"`
	EventId            int           `json:"event-id"`
	Id                 int           `json:"id"`
	MarketId           int           `json:"market-id"`
	Name               string        `json:"name"`
	Status             string        `json:"status"`
	Volume             float32       `json:"volume"`
	EventParticipantId int           `json:"event-participant-id"`
}
type Runner_params struct {
	States string
	/*A comma separated list of runner states.
	  Only runners with status in the provided list
	   are included in the response.*/
	Include_withdrawn bool
	/*A boolean for returning or not the withdrawn
	  runners in the response.*/
	Include_prices bool
	/*A boolean indicating whether to return the prices for the runners.*/
	Exchange_type string
	/*Prices are returned in accordance with the specified exchange-type.
	  Allowed values: back-lay, binary.*/
	Odds_type string
	/*Prices are returned in accordance with the specified odds-type.
	  Allowed values: DECIMAL, US, HK, MALAY, INDO, %.*/
	Price_depth int32
	/*
	   Value indicating, for each runner, the maximum number of prices
	   that will be returned in the response for each side.
	*/
	Side string
	/*Only prices on the specified side are included in the response.
	  Allowed values are: back, lay or win, lose. Both sides are returned if no side is specified.
	*/
	Currency string
	/*Prices are returned in the specified currency. Allowed values: USD, EUR, GBP, AUD, CAD. HKD*/
	Minimum_liquidity float64
	/*Only prices with available-amount greater than or equal to this value are included in the response.*/

}
type Runners_params struct {
	Include_prices bool
	/*A boolean indicating whether to return the prices for this runner.*/
	Exchange_type string
	/*Prices are returned in accordance with the specified exchange-type.
	  Allowed values: back-lay, binary.*/
	Odds_type string
	/*Prices are returned in accordance with the specified odds-type.
	  Allowed values: DECIMAL, US, HK, MALAY, INDO, %.*/
	Price_depth int32
	/*Value indicating, for each runner, the maximum number of prices that
	  will be returned in the response for each side.*/
	Side string
	/*Only prices on the specified side are included in the response.
	  Allowed values are: back, lay or win, lose. Both sides are
	 returned if no side is specified.*/
	Currency string
	/*Prices are returned in the specified currency.
	  Allowed values: USD, EUR, GBP, AUD, CAD. HKD*/
	Minimum_liquidity float64
	/*Only prices with available-amount greater than
	  or equal to this value are included in the response.*/

}
type PopularMarkets struct {
	Events  []Events  `json:"events"`
	Markets []Markets `json:"markets"`
}
type PopularMarkets_params struct {
	Exchange_type string
	/*Prices are returned in accordance with the specified
	  exchange-type. Allowed values: back-lay, binary.*/
	Odds_type string
	/*Prices are returned in accordance with the specified
	  odds-type. Allowed values: DECIMAL, US, HK, MALAY, INDO, %.*/
	Price_depth int32
	/*Value indicating, for each runner, the maximum number
	  of prices that will be returned in the response for each side.*/
	Side string
	/*Only prices on the specified side are included in the response.
	  Allowed values are: back, lay or win, lose. Both sides are returned if no side is specified.*/
	Currency string
	/*Prices are returned in the specified currency.
	  Allowed values: USD, EUR, GBP, AUD, CAD. HKD*/
	Minimum_liquidity float64
	/*Only prices with available-amount greater than or
	  equal to this value are included in the response.*/
	Old_format bool
	/*A boolean whether to return the events & markets hierarchically.*/

}
type PriceDetail struct {
	AvailableAmount float32 `json:"available-amount"`
	Currency        string  `json:"currency"`
	OddsType        string  `json:"odds-type"`
	Odds            float32 `json:"odds"`
	DecimalOdds     float32 `json:"decimal-odds"`
	Side            string  `json:"side"`
	ExchangeType    string  `json:"exchange-type"`
}
type Prices struct {
	Prices []PriceDetail `json:"prices"`
}
type Prices_params struct {
	Exchange_type string
	/*Prices are returned in accordance with the specified exchange-type.
	  Allowed values: back-lay, binary.*/
	Odds_type string
	/*Prices are returned in accordance with the specified odds-type.
	  Allowed values: DECIMAL, US, HK, MALAY, INDO, %.*/
	Depth int32
	/*Value indicating, for each runner, the maximum number of prices that
	  will be returned in the response for each side.*/
	Side string
	/*Only prices on the specified side are included in the response.
	  Allowed values are: back, lay or win, lose.
	 Both sides are returned if no side is specified.*/
	Currency string
	/*Prices are returned in the specified currency.
	  Allowed values: USD, EUR, GBP, AUD, CAD. HKD*/
	Minimum_liquidity float64
	/*Only prices with available-amount greater than or equal
	  to this value are included in the response.*/

}
type MetaTag struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	UrlName string `json:"url-names"`
}
type Navigation struct {
	Navigation [][][][]MetaTag
}
type Offer struct {
	Id                         int64   `json:"id"`
	Event_id                   int64   `json:"event-id"`
	Event_name                 string  `json:"event-name"`
	Market_id                  int64   `json:"market-id"`
	Market_name                string  `json:"market-name"`
	Market_type                string  `json:"market-type"`
	Runner_id                  int64   `json:"runner-id"`
	Runner_name                string  `json:"runner-name"`
	Exchange_type              string  `json:"exchange-type"`
	Side                       string  `json:"side"`
	Odds                       float64 `json:"odds"`
	Odds_type                  string  `json:"odds-type"`
	Decimal_odds               float64 `json:"decimal-odds"`
	Stake                      float64 `json:"stake"`
	Remaining                  float64 `json:"remaining"`
	Potential_profit           float64 `json:"potential-profit"`
	Remaining_potential_profit float64 `json:"remaining-potential-profit"`
	Commission_type            string  `json:"comission-type"`
	Originator_commission_rate float64 `json:"originator-commission-rate"`
	Acceptor_commission_rate   float64 `json:"acceptor-commission-rate"`
	Commission_reserve         float64 `json:"commission-reserve"`
	Currency                   string  `json:"currency"`
	Created_at                 string  `json:"created-at"`
	Status                     string  `json:"status"`
	In_play                    bool    `json:"in-play"`
}
type OfferE struct {
	Offer       Offer
	Offer_edits []Offer_Edits
}
type Offer_Edits struct {
	Id                  int64   `json:"id"`
	Offer_id            int64   `json:"offer-id"`
	Odds_type           string  `json:"odds-type"`
	Odds_before         float64 `json:"odds-before"`
	Decimal_odds_before float64 `json:"decimal-odds-before"`
	Odds_after          float64 `json:"odds-after"`
	Decimal_odds_after  float64 `json:"decimal-odds-after"`
	Stake_before        float64 `json:"stake-before"`
	Stake_after         float64 `json:"stake-after"`
	Edit_time           string  `json:"edit-time"`
}
type Offer_head struct {
	Language      string `json:"in-play"`
	Currency      string `json:"currency"`
	Exchange_type string `json:"exchange-type"`
	Odds_type     string `json:"odds-type"`
}
type Offers struct {
	Offer_head
	Offers []Offer `json:"offers"`
}
type Offer_Post struct {
	Runner_id int64
	Side      string
	Odds      float32
	Stake     float32
}
type Offers_Post struct {
	Odds_type     string
	Exchange_type string
	Offers        []Offer_Post
}
type Offer_Edit struct {
	Id            string
	Current_odds  decimal.Decimal
	New_odds      decimal.Decimal
	Current_stake decimal.Decimal
	New_stake     decimal.Decimal
}
type Offers_Edit struct {
	Offers []Offer_Edit
}
type Offers_Cancel struct {
	Event_ids string
	/*A comma separated list of event ids. Only offers on the provided events are cancelled.*/
	Market_ids string
	/*A comma separated list of market ids. Only offers on the provided markets are cancelled.*/
	Runner_ids string
	/*A comma separated list of runner ids. Only offers on the provided runners are cancelled.*/
	Offer_ids string
	/*A comma separated list of offer ids. Only the specified offers are cancelled.*/

}
type Offers_EditResult struct {
	Offer_head
	Offers []OfferE `json:"offers"`
}
