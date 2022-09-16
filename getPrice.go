package coingecko

import (
	"github.com/0xVanfer/abigen/erc20"
	"github.com/0xVanfer/coingecko/currencys"
	"github.com/0xVanfer/coingecko/geckoapis"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Get price by coingecko id.
// Currency can be "", "usd" will be used by default.
func (g *Gecko) GetPrice(id string, currency string) (float64, error) {
	if currency == "" {
		currency = currencys.CurrencyUSDollar
	}
	mapp, err := geckoapis.GetGeckoPrice(id, currency, g.ApiKey)
	if err != nil {
		return 0, err
	}
	for _, token := range mapp {
		for _, price := range token {
			return price, nil
		}
	}
	return 0, nil
}

// Get token price by its symbol.
// Some tokens like usdc.e on avalanche cannot be found by coingecko list, need special process.
func (g *Gecko) GetPriceBySymbol(symbol string, network string, currency string) (float64, error) {
	id, err := g.GetId(symbol, network)
	if err != nil {
		return 0, err
	}
	return g.GetPrice(id, currency)
}

// Get token price by its address.
// Some tokens like usdc.e on avalanche cannot be found by coingecko list, need special process.
func (g *Gecko) GetPriceByAddress(address string, network string, currency string, client bind.ContractBackend) (float64, error) {
	token, err := erc20.NewErc20(common.HexToAddress(address), client)
	if err != nil {
		return 0, err
	}
	symbol, err := token.Symbol(nil)
	if err != nil {
		return 0, err
	}
	id, err := g.GetId(symbol, network)
	if err != nil {
		return 0, err
	}
	return g.GetPrice(id, currency)
}
