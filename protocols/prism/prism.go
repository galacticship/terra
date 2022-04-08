package prism

import (
	"github.com/galacticship/terra"
	"github.com/pkg/errors"
)

type Prism struct {
	Amps         *Amps
	YLUNAStaking *YLUNAStaking
	Farm         *Farm
	Governance   *Governance
}

func NewPrism(querier *terra.Querier) (*Prism, error) {
	amps, err := NewAmps(querier)
	if err != nil {
		return nil, errors.Wrap(err, "creating amps")
	}
	ylunastaking, err := NewYLUNAStaking(querier)
	if err != nil {
		return nil, errors.Wrap(err, "creating yluna staking")
	}
	farm, err := NewFarm(querier)
	if err != nil {
		return nil, errors.Wrap(err, "creating farm")
	}
	gov, err := NewGovernance(querier)
	if err != nil {
		return nil, errors.Wrap(err, "creating governance")
	}
	return &Prism{
		Amps:         amps,
		YLUNAStaking: ylunastaking,
		Farm:         farm,
		Governance:   gov,
	}, nil
}
