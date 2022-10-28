package geckoapis

import (
	"github.com/imroc/req"
)

type CoingeckoToken struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// Get token list to pair token symbol with token id.
func GetGeckoTokenList(apiKey string) (list []CoingeckoToken, err error) {
	// if user has api key, use the pro version
	url := "https://api.coingecko.com/api/v3/coins/list"
	if apiKey != "" {
		url = "https://api.coingecko.com/api/v3/coins/list?x_cg_pro_api_key=" + apiKey
	}
	r, err := req.Get(url)
	if err != nil {
		return
	}
	err = r.ToJSON(&list)
	if err != nil {
		return
	}
	return
}
