package dashboard

import "time"

const UpdateInterval = 5 * time.Second

type OperatorsStake struct {
	Holding float64
	Staked  float64
	Address string
}

type Order struct {
	TxHash   string
	OrderID  string
	BlockNum uint64
}

var orders = map[string]Order{}
