package entity

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ORDER_PUBSUB_TOPIC = "jit-liq-order"
)

type AttestOrderPayload struct {
	OrderID         string         `json:"orderId"`
	Operator        common.Address `json:"operator"`
	Signature       string         `json:"signature"`
	ChainID         int64          `json:"chainId"`
	TransactionHash common.Hash    `json:"transactionHash"`
}
