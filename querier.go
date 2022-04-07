package terra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	terraappparams "github.com/terra-money/core/app/params"
	"golang.org/x/net/context/ctxhttp"
)

type Querier struct {
	url            string
	httpClient     *http.Client
	encodingConfig terraappparams.EncodingConfig
	chainId        string
}

type QuerierOption func(q *Querier) *Querier

func WithChainId(chainId string) QuerierOption {
	return func(q *Querier) *Querier {
		q.chainId = chainId
		return q
	}
}

func NewQuerier(httpClient *http.Client, url string, options ...QuerierOption) *Querier {
	q := &Querier{
		url:            url,
		httpClient:     httpClient,
		encodingConfig: terraappparams.MakeEncodingConfig(),
		chainId:        "columbus-5",
	}
	for _, option := range options {
		q = option(q)
	}
	return q
}

func (q Querier) ChainId() string {
	return q.chainId
}

func (q Querier) POST(ctx context.Context, method string, payload interface{}, result interface{}) error {
	u, err := url.Parse(q.url)
	if err != nil {
		return errors.Wrapf(err, "parsing lcd URL %s", q.url)
	}
	u.Path = path.Join(u.Path, method)

	reqBytes, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "marshalling payload to json")
	}

	resp, err := ctxhttp.Post(ctx, q.httpClient, u.String(), "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return errors.Wrap(err, "executing http post request")
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "reading response")
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("non-200 response code %d: %s", resp.StatusCode, string(out))
	}
	if result != nil {
		err = json.Unmarshal(out, &result)
		if err != nil {
			return errors.Wrap(err, "unmarshalling response from json")
		}
	}
	return nil
}

func (q Querier) POSTProto(ctx context.Context, method string, payload proto.Message, result proto.Message) error {
	u, err := url.Parse(q.url)
	if err != nil {
		return errors.Wrapf(err, "parsing lcd URL %s", q.url)
	}
	u.Path = path.Join(u.Path, method)

	reqBytes, err := q.encodingConfig.Marshaler.MarshalJSON(payload)
	if err != nil {
		return errors.Wrap(err, "marshalling payload to json")
	}

	resp, err := ctxhttp.Post(ctx, q.httpClient, u.String(), "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return errors.Wrap(err, "executing http post request")
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "reading response")
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("non-200 response code %d: %s", resp.StatusCode, string(out))
	}
	if result != nil {
		err = q.encodingConfig.Marshaler.UnmarshalJSON(out, result)
		if err != nil {
			return errors.Wrap(err, "unmarshalling response from json")
		}
	}
	return nil
}

func (q Querier) GET(ctx context.Context, method string, params url.Values, res interface{}) error {
	u, err := url.Parse(q.url)
	if err != nil {
		return errors.Wrapf(err, "parsing lcd URL %s", q.url)
	}
	u.Path = path.Join(u.Path, method)
	if params != nil {
		u.RawQuery = params.Encode()
	}
	resp, err := ctxhttp.Get(ctx, q.httpClient, u.String())
	if err != nil {
		return errors.Wrapf(err, "http requesting %s", u.String())
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "reading response")
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("non-200 response code %d: %s", resp.StatusCode, string(out))
	}
	if res != nil {
		err = json.Unmarshal(out, &res)
		if err != nil {
			return errors.Wrap(err, "unmarshalling response from json")
		}
	}
	return nil
}

func (q *Querier) LatestBlockInfo(ctx context.Context) (int64, time.Time, error) {
	var res struct {
		Block struct {
			Header struct {
				Height int64     `json:"height,string"`
				Time   time.Time `json:"time"`
			} `json:"header"`
		} `json:"block"`
	}
	err := q.GET(ctx, "blocks/latest", nil, &res)
	if err != nil {
		return 0, time.Time{}, errors.Wrap(err, "executing get request")
	}
	return res.Block.Header.Height, res.Block.Header.Time, nil
}
