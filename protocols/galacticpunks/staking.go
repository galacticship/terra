package galacticpunks

import (
	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
)

type Staking struct {
	*terra.Contract
}

func NewStaking(querier *terra.Querier) (*Staking, error) {
	contract, err := terra.NewContract(querier, "terra10t4pgfs6s3qeykqgfq9r74s89jmu7zx5gfkga5")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}

	return &Staking{
		Contract: contract,
	}, nil
}

func (s *Staking) NewWithdrawRewardsMessage(sender cosmos.AccAddress, tokenId string) (cosmos.Msg, error) {
	var q struct {
		WithdrawRewards struct {
			TokenId string `json:"token_id"`
		} `json:"withdraw_rewards"`
	}
	q.WithdrawRewards.TokenId = tokenId
	return s.NewMsgExecuteContract(sender, q)
}
