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

// Create a new gecko.
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
func (g *Gecko) Update() (*Gecko, error) {
	newList, err := geckoapis.GetGeckoTokenList(g.ApiKey)
	if err != nil {
		return g, err
	}
	var newGecko = Gecko{
		ApiKey:    g.ApiKey,
		TokenList: newList,
		UpdatedAt: time.Now(),
	}
	return &newGecko, nil
}
