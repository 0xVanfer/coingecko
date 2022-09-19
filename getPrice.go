package coingecko

import (
	"errors"

	"github.com/0xVanfer/abigen/erc20"
	"github.com/0xVanfer/coingecko/currencys"
	"github.com/0xVanfer/coingecko/geckoapis"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Return token price.
// Currency can be "", "usd" will be used by default.
func (g *Gecko) GetPriceById(id string, currency string) (float64, error) {
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

// Return token price.
// Some tokens like usdc.e on avalanche cannot be found by coingecko list, need special process.
func (g *Gecko) GetPriceBySymbol(symbol string, network string, currency string) (float64, error) {
	if symbol == "" {
		return 0, errors.New("symbol must not be empty")
	}
	id, err := g.GetId(symbol, network)
	if err != nil {
		return 0, err
	}
	return g.GetPriceById(id, currency)
}

// Return token price.
// Some tokens like usdc.e on avalanche cannot be found by coingecko list, need special process.
func (g *Gecko) GetPriceByAddress(address string, network string, currency string, client bind.ContractBackend) (float64, error) {
	if address == "" {
		return 0, errors.New("address must not be empty")
	}
	if address == "0x0000000000000000000000000000000000000000" {
		return 0, errors.New("address must not be zero")
	}
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
	return g.GetPriceById(id, currency)
}
