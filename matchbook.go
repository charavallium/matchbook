package matchbook

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
	//"github.com/valyala/fasthttp"
)

var connectionEndpoints = map[string][]string{
	"login":                    {"https://api.matchbook.com/bpapi/rest/security/session", "POST"},
	"logout":                   {"https://api.matchbook.com/bpapi/rest/security/session", "DELETE"},
	"getSession":               {"https://api.matchbook.com/bpapi/rest/security/session", "GET"},
	"getAccount":               {"https://api.matchbook.com/edge/rest/account", "GET"},
	"getCasinoWalletBalance":   {"https://api.matchbook.com/bpapi/rest/account/balance", "GET"},
	"getSportsWalletBalance":   {"https://api.matchbook.com/edge/rest/account/balance", "GET"},
	"walletTransfer":           {"https://api.matchbook.com/bpapi/rest/account/transfer", "POST"},
	"getEvents":                {"https://api.matchbook.com/edge/rest/events", "GET"},
	"getEvent":                 {"https://api.matchbook.com/edge/rest/events/{event_id}", "GET"},
	"getMarkets":               {"https://api.matchbook.com/edge/rest/events/{event_id}/markets", "GET"},
	"getMarket":                {"https://api.matchbook.com/edge/rest/events/{event_id}/markets/{market_id}", "GET"},
	"getSports":                {"https://api.matchbook.com/edge/rest/lookups/sports", "GET"},
	"getBalance":               {"https://api.matchbook.com/edge/rest/account/balance", "GET"},
	"getNavigation":            {"https://api.matchbook.com/edge/rest/navigation", "GET"},
	"getRunners":               {"https://api.matchbook.com/edge/rest/events/{event_id}/markets/{market_id}/runners", "GET"},
	"getRunner":                {"https://api.matchbook.com/edge/rest/events/{event_id}/markets/{market_id}/runners/{runner_id}", "GET"},
	"getPrices":                {"https://api.matchbook.com/edge/rest/events/{event_id}/markets/{market_id}/runners/{runner_id}/prices", "GET"},
	"getPopularMarkets":        {"https://api.matchbook.com/edge/rest/popular-markets", "GET"},
	"submitOffers":             {"https://api.matchbook.com/edge/rest/offers", "POST"},
	"editOffers":               {"https://api.matchbook.com/edge/rest/offers", "PUT"},
	"editOffer":                {"https://api.matchbook.com/edge/rest/offer/{offer_id}", "PUT"},
	"cancelOffers":             {"https://api.matchbook.com/edge/rest/offers", "DELETE"},
	"deleteOffer":              {"https://api.matchbook.com/edge/rest/offer/{offer_id}", "DELETE"},
	"getOffers":                {"https://api.matchbook.com/edge/rest/offers", "GET"},
	"getOffer":                 {"https://api.matchbook.com/edge/rest/offer/{offer_id}", "GET"},
	"getCancelledMatchedBets":  {"https://api.matchbook.com/bets?status=cancelled", "GET"},
	"getAggregatedMatchedBets": {"https://api.matchbook.com/edge/rest/bets/matched/aggregated", "GET"},
	"getPositions":             {"https://api.matchbook.com/edge/rest/account/positions", "GET"},
	"getOldWalletTransactions": {"https://api.matchbook.com/bpapi/rest/reports/transactions", "GET"},
	"getNewWalletTransactions": {"https://api.matchbook.com/edge/rest/reports/v1/transactions", "GET"},
	"getCurrentOffers":         {"https://api.matchbook.com/edge/rest/reports/v1/offers/current", "GET"},
	"getCurrentBets":           {"https://api.matchbook.com/edge/rest/reports/v1/bets/current", "GET"},
	"getSettledBets":           {"https://api.matchbook.com/edge/rest/reports/v1/bets/settled", "GET"},
	"getCountries":             {"https://api.matchbook.com/bpapi/rest/lookups/countries", "GET"},
	"getRegions":               {"https://api.matchbook.com/bpapi/rest/lookups/regions/{country_id}", "GET"},
	"getCurrencies":            {"https://api.matchbook.com/bpapi/rest/lookups/currencies", "GET"},
	"cancelOffer":              {"https://api.matchbook.com/edge/rest/offers/{offer_id}", "DELETE"},
}

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"session-token"`
}

type Session struct {
	config     *Config
	httpClient *http.Client
	token      string
}

type RequestSpecification struct {
	Url  string
	Type string
}
type Abstract interface{}

// Create a new session.
func (c *Config) NewSession() (*Session, error) {

	s := new(Session)

	cookieJar, _ := cookiejar.New(nil)
	// Configuration
	if c.Username == "" {
		return s, errors.New("Config.Username is empty.")
	}
	if c.Password == "" {
		return s, errors.New("Config.Password is empty.")
	}

	s.config = c
	fmt.Println(c)
	ssl := &tls.Config{
		InsecureSkipVerify: true,
	}
	ssl.Rand = rand.Reader
	s.httpClient = &http.Client{
		Jar: cookieJar,
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Duration(time.Second*3))
			},
			TLSClientConfig: ssl,
		},
	}

	return s, nil
}
func get_uri_params(sp interface{}) (m map[string]string) {
	m = make(map[string]string)
	post := url.Values{}
	v := reflect.ValueOf(sp)
	for i := 0; i < v.NumField(); i++ {
		s := ""
		switch v.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if i := v.Field(i).Int(); i > 0 {
				s = strconv.FormatInt(i, 10)
			}
			break
		case reflect.String:
			s = v.Field(i).String()
			break
		case reflect.Bool:
			s = strconv.FormatBool(v.Field(i).Bool())
			break
		case reflect.Float32, reflect.Float64:
			if i := v.Field(i).Float(); i > 0 {
				dc := decimal.NewFromFloat(i)
				s = dc.String()
			}
			break

		}
		if s != "" {
			nm := v.Type().Field(i).Name
			nm = strings.ToLower(strings.Replace(nm, "_", "-", strings.Count(nm, "_")))
			post.Add(nm, s)
		}
	}

	if pe := post.Encode(); pe != "" {
		m["GET"] = pe

	}
	return m
}

// Builds URLs for API methods.
func (s *Session) getRequestSpec(key, method string, id map[string]string) (RequestSpecification, error) {
	if _, exists := connectionEndpoints[key]; exists == false {
		return RequestSpecification{}, errors.New("Invalid endpoint key: " + key)
	}
	url := connectionEndpoints[key][0] + method
	requestType := connectionEndpoints[key][1]
	for k, v := range id {
		if k != "POST" {
			url = strings.Replace(url, "{"+k+"}", v, 1)
		}
	}

	if requestType == "GET" {
		url += "/"
	}
	if id != nil && id["GET"] != "" {
		url += "?" + id["GET"]
		fmt.Println("URI", url)
	}
	return RequestSpecification{url, requestType}, nil
}
func va(a, b []byte) {
	fmt.Println(a, b)
}
func (s *Session) doRequest(key, method string, body *strings.Reader, id map[string]string) ([]byte, error) {

	reqSpec, err := s.getRequestSpec(key, method, id)
	if err != nil {
		return nil, err
	}
	if v, ok := id["POST"]; ok {
		body = strings.NewReader(v)
	}
	req, err := http.NewRequest(reqSpec.Type, reqSpec.Url, body)

	//	req := fasthttp.AcquireRequest()

	//req.URI().Update(reqSpec.Url)
	//req.Header.SetMethod(reqSpec.Type)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")

	if s.token != "" {
		req.Header.Set("X-Authentication", s.token)
	}
	fmt.Println("----------", req, "--------")
	//res := fasthttp.AcquireResponse()

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(data))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Session) Getter(t string, res interface{}, id map[string]string) (interface{}, error) {
	body := strings.NewReader("")
	if resp, err := s.doRequest(t, "", body, id); err == nil {
		//fmt.Println(string(resp))
		if err := json.Unmarshal(resp, &res); err == nil {
			return res, nil
		} else {
			return res, errors.New("Invalid json")
		}
	}
	return res, errors.New("Unknow error")
}
func (s *Session) GetSpotrsParams() Sports_params {
	return Sports_params{}
}
func (s *Session) GetEventParams() Event_params {
	return Event_params{}
}
func (s *Session) GetEventsParams() Events_params {
	return Events_params{}
}
func (s *Session) GetMarketsParams() Markets_params {
	return Markets_params{}
}
func (s *Session) GetMarketParams() Market_params {
	return Market_params{}
}
func (s *Session) GetSports(qp Sports_params) (*Sports, error) {
	val, err := s.Getter("getSports", reflect.New(reflect.TypeOf(Sports{})).Interface(), get_uri_params(qp))
	return val.(*Sports), err
}
func (s *Session) SubmitOffers(qp Offers_Post) (*Offers, error) {
	m := make(map[string]string)
	if j, e := json.Marshal(qp); e == nil {
		m["POST"] = string(j)
		m["POST"] = strings.Replace(m["POST"], "_", "-", strings.Count(m["POST"], "_"))
		m["POST"] = strings.ToLower(m["POST"])
		m["POST"] = strings.Replace(m["POST"], "decimal", "DECIMAL", 1)
	}
	fmt.Println("=======>", m)
	val, err := s.Getter("submitOffers", reflect.New(reflect.TypeOf(Offers{})).Interface(), m)
	return val.(*Offers), err
}

func (s *Session) GetMarket(qp Market_params, event_id, market_id string) (*Market, error) {
	m := get_uri_params(qp)
	m["event_id"] = event_id
	m["market_id"] = market_id
	val, err := s.Getter("getMarket", reflect.New(reflect.TypeOf(Market{})).Interface(), m)
	return val.(*Market), err
}
func (s *Session) GetMarkets(qp Market_params, event_id string) (*Markets, error) {
	m := get_uri_params(qp)
	m["event_id"] = event_id
	val, err := s.Getter("getMarkets", reflect.New(reflect.TypeOf(Markets{})).Interface(), m)
	return val.(*Markets), err
}
func (s *Session) CancelOffer(offer_id int64) (*Offer, error) {
	m := make(map[string]string)
	m["offer_id"] = strconv.FormatInt(offer_id, 10)
	val, err := s.Getter("cancelOffer", reflect.New(reflect.TypeOf(Offer{})).Interface(), m)
	return val.(*Offer), err
}
func (s *Session) GetEvents(qp Events_params) (*Events, error) {
	val, err := s.Getter("getEvents", reflect.New(reflect.TypeOf(Events{})).Interface(), get_uri_params(qp))
	return val.(*Events), err
}
func (s *Session) GetAcc() (*AccountDetails, error) {
	val, err := s.Getter("getAccount", reflect.New(reflect.TypeOf(AccountDetails{})).Interface(), nil)
	return val.(*AccountDetails), err
}
func (s *Session) GetBalance() (*Balance, error) {
	val, err := s.Getter("getBalance", reflect.New(reflect.TypeOf(Balance{})).Interface(), nil)
	return val.(*Balance), err
}
func (s *Session) GetCanceled() (*MBets, error) {
	val, err := s.Getter("getCancelledMatchedBets", reflect.New(reflect.TypeOf(MBets{})).Interface(), nil)
	return val.(*MBets), err
}

/*
func (s *Session) GetNavigation(offset, per_page int64) (*Navigation, error) {
	val, err := s.Getter("getNavigation", reflect.New(reflect.TypeOf(Navigation{})).Interface(), "")
	return val.(*Navigation), err
}
func (s *Session) GetEvents(after, before, price_depth int64, param map[string]string, minimum_liquidity float64, include_event_participants bool) (*Events, error) {
	val, err := s.Getter("getEvents", reflect.New(reflect.TypeOf(Events{})).Interface(), "")
	return val.(*Events), err
}


func (s *Session) GetMarkets(id string) (*Markets, error) {
	val, err := s.Getter("getMarkets", reflect.New(reflect.TypeOf(Markets{})).Interface(), id)
	return val.(*Markets), err
}

*/
func (s *Session) Login() ([]byte, error) {
	res := LoginResponse{}

	if s.config.Token != "" {
		s.token = s.config.Token
		if resp, err := s.doRequest("getSession", "", strings.NewReader(""), nil); err == nil {
			json.Unmarshal(resp, &res)
			s.token = s.config.Token

			return nil, nil
		}

	}

	body := strings.NewReader(`{ "username": "` + s.config.Username + `", "password": "` + s.config.Password + `"}`)
	if resp, err := s.doRequest("login", "", body, nil); err == nil {
		json.Unmarshal(resp, &res)
		s.token = res.SessionToken

		b, _ := json.Marshal(Config{s.config.Username, s.config.Password, s.token})
		err1 := ioutil.WriteFile("matchbook.json", b, 0644)
		if err1 != nil {
			panic(err1)
		}

		return resp, err
	}
	return nil, nil

}
