package loop

import (
	"github.com/galacticship/terra"
	"github.com/shopspring/decimal"
)

type Pair struct {
	*terra.BasePair
}

func NewPair(querier *terra.Querier, contractAddress string, token1 terra.Token, token2 terra.Token, lpToken terra.Cw20Token) (*Pair, error) {
	bp, err := terra.NewBasePair(querier, contractAddress, token1, token2, lpToken, decimal.NewFromFloat(0.003), terra.NewAssetInfoFactory())
	if err != nil {
		return nil, err
	}
	return &Pair{
		bp,
	}, nil
}
