package mars

import (
	"context"

	"github.com/galacticship/terra"
	"github.com/pkg/errors"
)

type Mars struct {
	Governance  *Governance
	Bootstrap   *Bootstrap
	ANCUSTField *Field
}

func NewMars(ctx context.Context, querier *terra.Querier) (*Mars, error) {
	governance, err := NewGovernance(querier)
	if err != nil {
		return nil, errors.Wrap(err, "creating governance")
	}
	bootstrap, err := NewBootstrap(ctx, querier)
	if err != nil {
		return nil, errors.Wrap(err, "creating bootstrap")
	}
	ancustfield, err := NewANCUSTField(querier)
	if err != nil {
		return nil, errors.Wrap(err, "creating ancust field")
	}

	return &Mars{
		Governance:  governance,
		Bootstrap:   bootstrap,
		ANCUSTField: ancustfield,
	}, nil
}
