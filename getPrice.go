package coingecko

import (
	"errors"

	"github.com/0xVanfer/abigen/erc20"
	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/coingecko/geckoapis"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// Return token price.
//
// Currency can be "", "usd" will be used by default.
func (g *Gecko) GetPriceById(id string, currency string) (decimal.Decimal, error) {
	if currency == "" {
		currency = "usd"
	}
	mapp, err := geckoapis.GetGeckoPrice(id, currency, g.ApiKey)
	if err != nil {
		return decimal.Zero, err
	}
	for _, token := range mapp {
		for _, price := range token {
			return price, nil
		}
	}
	return decimal.Zero, nil
}

// Return token price.
//
// Some tokens like usdc.e on avalanche cannot be found by coingecko list, need special process.
func (g *Gecko) GetPriceBySymbol(symbol string, network string, currency string) (decimal.Decimal, error) {
	if symbol == "" {
		return decimal.Zero, errors.New("symbol must not be empty")
	}
	id, err := g.GetId(symbol, network)
	if err != nil {
		return decimal.Zero, err
	}
	return g.GetPriceById(id, currency)
}

// Return token price.
//
// Some tokens like usdc.e on avalanche cannot be found by coingecko list, need special process.
func (g *Gecko) GetPriceByAddress(address string, network string, currency string, client bind.ContractBackend) (decimal.Decimal, error) {
	if address == "" {
		return decimal.Zero, errors.New("address must not be empty")
	}
	if address == "0x0000000000000000000000000000000000000000" {
		return decimal.Zero, errors.New("address must not be zero")
	}
	token, err := erc20.NewErc20(common.HexToAddress(address), client)
	if err != nil {
		return decimal.Zero, err
	}
	symbol, err := token.Symbol(nil)
	if err != nil {
		return decimal.Zero, err
	}
	id, err := g.GetId(symbol, network)
	if err != nil {
		return decimal.Zero, err
	}
	return g.GetPriceById(id, currency)
}

// Return chain token price.
func (g *Gecko) GetChainTokenPrice(network string, currency string) (decimal.Decimal, error) {
	return g.GetPriceBySymbol(chainId.ChainTokenSymbolList[network], network, currency)
}
