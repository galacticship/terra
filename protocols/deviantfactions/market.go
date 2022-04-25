package deviantfactions

import (
	"context"

	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Market struct {
	*terra.Contract
}

func NewMarket(q *terra.Querier) (*Market, error) {
	contract, err := terra.NewContract(q, "terra1t37vatdcsmg4qlyy8n4fgapecce5g4gw5mlzlt")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}

	return &Market{
		Contract: contract,
	}, nil
}

type Listing struct {
	Owner        string
	TokenId      string
	AskToken     terra.Token
	Price        decimal.Decimal
	BlockCreated int64
	BlockExpires int64
}

func (m *Market) AllCollections(ctx context.Context) ([]string, error) {
	var res []string
	startAfter := ""
	for {
		tmp, err := m.allCollectionsPage(ctx, 30, startAfter)
		if err != nil {
			return nil, errors.Wrapf(err, "getting page with startafter %s", startAfter)
		}
		res = append(res, tmp...)
		if len(tmp) < 30 {
			break
		}
		startAfter = tmp[len(tmp)-1]
	}
	return res, nil
}

func (m *Market) allCollectionsPage(ctx context.Context, limit int, startAfter string) ([]string, error) {
	var q struct {
		AllCollections struct {
			Limit      int    `json:"limit"`
			StartAfter string `json:"start_after"`
		} `json:"all_collections"`
	}
	q.AllCollections.Limit = limit
	q.AllCollections.StartAfter = startAfter
	var r struct {
		Collections []string `json:"collections"`
	}
	err := m.QueryStore(ctx, q, &r)
	if err != nil {
		return nil, errors.Wrap(err, "querying contract store")
	}

	return r.Collections, nil
}

func (m *Market) AllListedTokens(ctx context.Context, collectionIds []string) ([]Listing, error) {
	var res []Listing
	startAfter := ""
	for {
		tmp, err := m.allListedTokensPage(ctx, collectionIds, 30, startAfter)
		if err != nil {
			return nil, errors.Wrapf(err, "getting page with startafter %s", startAfter)
		}
		res = append(res, tmp...)
		if len(tmp) < 30 {
			break
		}
		startAfter = tmp[len(tmp)-1].TokenId
	}
	return res, nil
}

func (m *Market) allListedTokensPage(ctx context.Context, collectionIds []string, limit int, startAfter string) ([]Listing, error) {
	var q struct {
		AllListedTokens struct {
			CollectionIds []string `json:"collection_ids"`
			Limit         int      `json:"limit"`
			StartAfter    string   `json:"start_after"`
		} `json:"all_listed_tokens"`
	}
	q.AllListedTokens.CollectionIds = collectionIds
	q.AllListedTokens.Limit = limit
	q.AllListedTokens.StartAfter = startAfter
	var r struct {
		Listings []struct {
			Owner   string `json:"owner"`
			TokenId string `json:"token_id"`
			Ask     struct {
				Denom  string          `json:"denom"`
				Amount decimal.Decimal `json:"amount"`
			} `json:"ask"`
			BlockCreated int64 `json:"block_created"`
			BlockExpires int64 `json:"block_expires"`
		} `json:"listings"`
	}

	err := m.QueryStore(ctx, q, &r)
	if err != nil {
		return nil, errors.Wrap(err, "querying store")
	}
	var res []Listing
	for _, listing := range r.Listings {
		token := terra.NativeTokenFromDenom(listing.Ask.Denom)
		res = append(res, Listing{
			Owner:        listing.Owner,
			TokenId:      listing.TokenId,
			AskToken:     token,
			Price:        token.ValueFromTerra(listing.Ask.Amount),
			BlockCreated: listing.BlockCreated,
			BlockExpires: listing.BlockExpires,
		})
	}
	return res, nil
}

func (m *Market) ListedTokensByUser(ctx context.Context, address cosmos.AccAddress) ([]string, error) {
	var res []string
	startAfter := ""
	for {
		tmp, err := m.listedTokensByUserPage(ctx, address, startAfter)
		if err != nil {
			return nil, errors.Wrapf(err, "getting page with startafter %s", startAfter)
		}
		if len(tmp) == 0 {
			break
		}
		res = append(res, tmp...)
		startAfter = tmp[len(tmp)-1]
	}
	return res, nil
}

func (m *Market) listedTokensByUserPage(ctx context.Context, address cosmos.AccAddress, startAfter string) ([]string, error) {
	var q struct {
		ListedTokens struct {
			Owner      string `json:"owner"`
			StartAfter string `json:"start_after"`
		} `json:"listed_tokens"`
	}
	q.ListedTokens.Owner = address.String()
	q.ListedTokens.StartAfter = startAfter
	var r struct {
		Tokens []string `json:"tokens"`
	}
	err := m.QueryStore(ctx, q, &r)
	if err != nil {
		return nil, errors.Wrap(err, "querying store")
	}
	return r.Tokens, nil
}

func (m *Market) TokensByUser(ctx context.Context, address cosmos.AccAddress) ([]string, error) {
	var res []string
	startAfter := ""
	for {
		tmp, err := m.tokensByUserPage(ctx, address, 30, startAfter)
		if err != nil {
			return nil, errors.Wrapf(err, "getting page with startafter %s", startAfter)
		}
		if len(tmp) == 0 {
			break
		}
		res = append(res, tmp...)
		startAfter = tmp[len(tmp)-1]
	}
	return res, nil
}

func (m *Market) tokensByUserPage(ctx context.Context, address cosmos.AccAddress, limit int, startAfter string) ([]string, error) {
	var q struct {
		Tokens struct {
			Owner      string `json:"owner"`
			Limit      int    `json:"limit"`
			StartAfter string `json:"start_after"`
		} `json:"tokens"`
	}
	q.Tokens.Owner = address.String()
	q.Tokens.Limit = limit
	q.Tokens.StartAfter = startAfter
	var r struct {
		Tokens []string `json:"tokens"`
	}
	err := m.QueryStore(ctx, q, &r)
	if err != nil {
		return nil, errors.Wrap(err, "querying store")
	}
	return r.Tokens, nil
}

func (m *Market) NewCancelTokenListingMessage(sender cosmos.AccAddress, tokenId string) (cosmos.Msg, error) {
	var q struct {
		CancelTokenListing struct {
			TokenId string `json:"token_id"`
		} `json:"cancel_token_listing"`
	}
	q.CancelTokenListing.TokenId = tokenId
	return m.NewMsgExecuteContract(sender, q)
}

func (m *Market) NewListTokenMessage(sender cosmos.AccAddress, tokenId string, AskToken terra.NativeToken, AskPrice decimal.Decimal, BlocksToListFor int64) (cosmos.Msg, error) {
	var q struct {
		ListToken struct {
			TokenId string `json:"token_id"`
			Ask     struct {
				Denom  string          `json:"denom"`
				Amount decimal.Decimal `json:"amount"`
			} `json:"ask"`
			BlocksToListFor int64 `json:"blocks_to_list_for"`
		} `json:"list_token"`
	}
	q.ListToken.TokenId = tokenId
	q.ListToken.Ask.Denom = AskToken.Denom()
	q.ListToken.Ask.Amount = AskToken.ValueToTerra(AskPrice)
	q.ListToken.BlocksToListFor = BlocksToListFor
	return m.NewMsgExecuteContract(sender, q)
}

func (m *Market) NewBuyTokenMessage(sender cosmos.AccAddress, tokenId string, token terra.NativeToken, price decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		BuyToken struct {
			TokenId string `json:"token_id"`
		} `json:"buy_token"`
	}
	q.BuyToken.TokenId = tokenId
	return token.NewMsgSendExecute(sender, m.Contract, price, q)
}
