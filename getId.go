package coingecko

import (
	"errors"
	"strings"

	"github.com/0xVanfer/chainId"
)

// Return the id accroding to symbol.
//
// Some tokens like usdc.e on avalanche cannot be found by coingecko list, need special process.
func (g *Gecko) GetId(symbol string, network string) (string, error) {
	// avalanche bridge tokens have ".e" after the origin symbol
	if network == chainId.AvalancheChainName {
		symbol = strings.Split(symbol, ".")[0]
	}
	// other networks rules: todo
	// find the id
	for _, token := range g.TokenList {
		// wormhole is a bridge on solana
		// wormholes use the same symbols as normal coins but have different ids
		if strings.Contains(token.ID, "wormhole") {
			continue
		}
		if strings.EqualFold(token.Symbol, symbol) {
			return token.ID, nil
		}
	}
	return "", errors.New("symbol not in the token list")
}
