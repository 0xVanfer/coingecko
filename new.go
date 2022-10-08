package coingecko

import (
	"time"

	"github.com/0xVanfer/coingecko/geckoapis"
)

type Gecko struct {
	ApiKey    string
	TokenList []geckoapis.CoingeckoToken
	UpdatedAt time.Time
}

// Create a new coingeco reader.
//
// "apiKey" is the key for pro users, to unlock the request limits.
//
// "" can be used as the key, and you will have a request limit of 10-50 times/min.
func New(apiKey string) (*Gecko, error) {
	list, err := geckoapis.GetGeckoTokenList(apiKey)
	if err != nil {
		return nil, err
	}
	var gecko = Gecko{
		ApiKey:    apiKey,
		TokenList: list,
		UpdatedAt: time.Now(),
	}
	return &gecko, nil
}

// Update the token list in gecko.
func (g *Gecko) Update() {
	newList, err := geckoapis.GetGeckoTokenList(g.ApiKey)
	if err != nil {
		return
	}
	g.TokenList = newList
	g.UpdatedAt = time.Now()
}
