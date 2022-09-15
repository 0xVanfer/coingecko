package geckoapis

import "github.com/imroc/req"

func GetGeckoPrice(id string, currency string, apiKey string) (map[string]map[string]float64, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + id + "&vs_currencies=" + currency
	if apiKey != "" {
		url = "https://api.coingecko.com/api/v3/simple/price?x_cg_pro_api_key=" + apiKey + "&ids=" + id + "&vs_currencies=usd"
	}
	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	var v map[string]map[string]float64
	err = r.ToJSON(&v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
