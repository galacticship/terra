package terra

import (
	customauthtx "github.com/terra-money/core/custom/auth/tx"
)

type (
	ComputeTaxRequest struct {
		customauthtx.ComputeTaxRequest
	}
	ComputeTaxResponse struct {
		customauthtx.ComputeTaxResponse
	}
)

func NewComputeTaxRequest(txBytes []byte) *ComputeTaxRequest {
	return &ComputeTaxRequest{
		customauthtx.ComputeTaxRequest{
			TxBytes: txBytes,
		},
	}
}
