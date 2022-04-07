package terra

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

type Factory struct {
	*Contract
}

func NewFactory(querier *Querier, contractAddress string) (*Factory, error) {
	contract, err := NewContract(querier, contractAddress)
	if err != nil {
		return nil, errors.Wrap(err, "creating base contract")
	}

	return &Factory{
		contract,
	}, nil
}

func (f *Factory) Pairs(ctx context.Context) ([]string, error) {
	const limit = 30
	var res []string
	var startAfter []StandardAssetInfo
	cpt := 1
	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("context canceled")
		default:
			pagePairs, lastAssets, err := f.pairsPage(ctx, startAfter, limit)
			if err != nil {
				return nil, errors.Wrapf(err, "querying page %d", cpt)
			}
			if pagePairs == nil {
				return res, nil
			}
			res = append(res, pagePairs...)
			startAfter = lastAssets
			cpt++
		}
	}
}

func (f *Factory) pairsPage(ctx context.Context, startAfter []StandardAssetInfo, limit int) ([]string, []StandardAssetInfo, error) {
	var q struct {
		Pairs struct {
			StartAfter []StandardAssetInfo `json:"start_after,omitempty"`
			Limit      *int                `json:"limit,omitempty"`
		} `json:"pairs"`
	}
	q.Pairs.StartAfter = startAfter
	q.Pairs.Limit = &limit

	type response struct {
		Pairs []struct {
			AssetInfos   []StandardAssetInfo `json:"asset_infos"`
			ContractAddr string              `json:"contract_addr"`
		} `json:"pairs"`
	}
	var resp response
	err := f.QueryStore(ctx, q, &resp)
	if err != nil {
		return nil, nil, errors.Wrap(err, "querying contract store")
	}

	var res []string
	var lastAssets []StandardAssetInfo
	for _, pair := range resp.Pairs {
		res = append(res, pair.ContractAddr)
		fmt.Println(pair.ContractAddr)
		lastAssets = pair.AssetInfos
	}
	fmt.Println("")

	return res, lastAssets, nil
}
