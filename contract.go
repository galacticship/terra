package terra

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/galacticship/terra/cosmos"
	"github.com/galacticship/terra/terra"
	"github.com/pkg/errors"
)

type Contract struct {
	q               *Querier
	contractAddress cosmos.AccAddress
}

func NewContract(querier *Querier, contractAddress string) (*Contract, error) {
	accAddress, err := cosmos.AccAddressFromBech32(contractAddress)
	if err != nil {
		return nil, errors.Wrap(err, "validating address")
	}
	return &Contract{
		q:               querier,
		contractAddress: accAddress,
	}, nil
}

func (c *Contract) Querier() *Querier {
	return c.q
}

func (c *Contract) Address() cosmos.AccAddress {
	return c.contractAddress
}

func (c *Contract) QueryStore(ctx context.Context, query interface{}, result interface{}) error {
	var envelope struct {
		QueryResult json.RawMessage `json:"query_result"`
	}
	queryBytes, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "marshalling query execmsg")
	}
	params := url.Values{}
	params.Set("query_msg", base64.StdEncoding.EncodeToString(queryBytes))
	err = c.q.GET(ctx, fmt.Sprintf("terra/wasm/v1beta1/contracts/%s/store", c.contractAddress.String()), params, &envelope)
	if err != nil {
		return errors.Wrap(err, "executing get request")
	}
	if result != nil {
		err = json.Unmarshal(envelope.QueryResult, result)
		if err != nil {
			return errors.Wrap(err, "unmarshalling query result")
		}
	}
	return nil
}

func (c Contract) NewMsgExecuteContract(sender cosmos.AccAddress, execMsg interface{}) (*terra.MsgExecuteContract, error) {
	return terra.NewMsgExecuteContract(sender, c.Address(), execMsg, nil)
}
