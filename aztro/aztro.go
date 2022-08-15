package aztro

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const AztroBaseURL = "https://aztro.sameerkumar.website"

type Sign int

const (
	Aries Sign = iota
	Taurus
	Gemini
	Cancer
	Leo
	Virgo
	Libra
	Scorpio
	Sagittarius
	Capricorn
	Aquarius
	Pisces
)

// String for sign is to convert into string
func (s Sign) String() string {
	return [...]string{
		"aries", "taurus", "gemini", "cancer", "leo",
		"virgo", "libra", "scorpio", "sagittarius", "capricorn",
		"aquarius", "pisces",
	}[s]
}

type Day int

const (
	Yesterday Day = iota
	Today
	Tomorrow
)

// String for Day is to convert into string
func (d Day) String() string {
	return [...]string{"yesterday", "today", "tomorrow"}[d]
}

// Horoscope struct represent a horoscope
type Horoscope struct {
	DateRange     string `json:"date_range"`
	CurrentDate   string `json:"current_date"`
	Description   string `json:"description"`
	Compatibility string `json:"compatibility"`
	Mood          string `json:"mood"`
	Color         string `json:"color"`
	LuckyNumber   string `json:"lucky_number"`
	LuckyTime     string `json:"lucky_time"`
}

// AztroErr struct represent error object from Aztro API
type AztroErr struct {
	Message        string `json:"message"`
	HTTPStatusCode int    `json:"status_code,omitempty"`
}

// AztroRequestParam struct represent a request parameter to query from Aztro API
type AztroRequestParam struct {
	Sign Sign
	Day  Day
}

// NewAztroRequestParam function to create a new AztroRequestParam, Default Day=Today
func NewAztroRequestParam(sign Sign, opts ...func(*AztroRequestParam)) *AztroRequestParam {
	param := &AztroRequestParam{
		Sign: sign,
		Day:  Today,
	}

	for _, opt := range opts {
		opt(param)
	}

	return param
}

// WithDay function to add Day into AztroRequestParam
func WithDay(day Day) func(*AztroRequestParam) {
	return func(param *AztroRequestParam) {
		param.Day = day
	}
}

// IAztroClient represent an interface to query to Aztro API
type IAztroClient interface {
	GetHoroscope(param *AztroRequestParam) (Horoscope, *AztroErr)
}

// AztroClient struct represent a client to query to Aztro API
type AztroClient struct {
	HTTPReq    *http.Request
	HTTPClient *http.Client
}

// NewAztroClient is function to create AztroClient
func NewAztroClient() (*AztroClient, error) {
	httpReq, err := http.NewRequest(http.MethodPost, AztroBaseURL, nil)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	return &AztroClient{HTTPReq: httpReq, HTTPClient: httpClient}, nil
}

// GetHoroscope function to get a horoscope based on requested param
func (ac *AztroClient) GetHoroscope(param *AztroRequestParam) (Horoscope, *AztroErr) {
	ac.setQueryParameter(param)

	httpResp, err := ac.HTTPClient.Do(ac.HTTPReq)
	if err != nil {
		return Horoscope{}, &AztroErr{HTTPStatusCode: httpResp.StatusCode, Message: "Error sending request to Aztro API"}
	}

	return parseHoroscopeResp(httpResp)
}

func (ac *AztroClient) setQueryParameter(param *AztroRequestParam) {
	q := url.Values{}
	q.Add("sign", param.Sign.String())
	q.Add("day", param.Day.String())
	ac.HTTPReq.URL.RawQuery = q.Encode()
}

func parseHoroscopeResp(resp *http.Response) (Horoscope, *AztroErr) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Horoscope{}, &AztroErr{Message: "Error Parsing Response Body", HTTPStatusCode: resp.StatusCode}
	}

	if resp.StatusCode != http.StatusOK {
		azErr := AztroErr{HTTPStatusCode: resp.StatusCode}
		if err := json.Unmarshal(body, &azErr); err != nil {
			return Horoscope{}, &AztroErr{Message: "Error Unmarshal Response Body"}
		}

		return Horoscope{}, &azErr
	}

	horoscope := Horoscope{}
	if err := json.Unmarshal(body, &horoscope); err != nil {
		return Horoscope{}, &AztroErr{Message: "Error Unmarshal Response Body", HTTPStatusCode: resp.StatusCode}
	}

	return horoscope, nil
}
