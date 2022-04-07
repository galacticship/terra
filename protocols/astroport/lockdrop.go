package astroport

import (
	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
)

type Lockdrop struct {
	*terra.Contract
}

func NewLockdrop(querier *terra.Querier) (*Lockdrop, error) {
	contract, err := terra.NewContract(querier, "terra1627ldjvxatt54ydd3ns6xaxtd68a2vtyu7kakj")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}

	return &Lockdrop{
		Contract: contract,
	}, nil
}

func (l *Lockdrop) NewClaimRewardsMessage(sender cosmos.AccAddress, duration int, withdrawLpStake bool, terraswapLpToken string) (cosmos.Msg, error) {
	type query struct {
		Claim struct {
			Duration         int    `json:"duration"`
			WithdrawLpStake  bool   `json:"withdraw_lp_stake"`
			TerraswapLpToken string `json:"terraswap_lp_token"`
		} `json:"claim_rewards_and_optionally_unlock"`
	}
	var q query
	q.Claim.Duration = duration
	q.Claim.WithdrawLpStake = withdrawLpStake
	q.Claim.TerraswapLpToken = terraswapLpToken
	return l.NewMsgExecuteContract(sender, q)
}
