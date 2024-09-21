package entity

import (
	"math/big"

	"github.com/JITLiq/node/abis/srcstatemanager"
	"github.com/ethereum/go-ethereum/common"
)

type OrderData struct {
	Fulfilled   bool
	Expiry      uint32
	OrderAmount *big.Int
	DestAddress common.Address
	Operator    common.Address
	Fees        srcstatemanager.IEntityFeesData
}
