package entity

import "github.com/ethereum/go-ethereum/common"

const (
	ChainIDBase = 8453
	ChainIDArb  = 42161
)

var (
	TopicOrderCreated  = common.HexToHash("0x918554b6bd6e2895ce6553de5de0e1a69db5289aa0e4fe193a0dcd1f14347477")
	BaseUsdc           = common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913")
	TargetJITPool      = common.HexToAddress("0xc62d083faf83f783E01d3f89d913A02cAf4b3120")
	TargetStateManager = common.HexToAddress("0x2BB5Ab403FC9b71bdcA497141F7efe68eaC4876a")
	SrcStateManager    = common.HexToAddress("0x2B4999dBeAeBC622C8ACD9e43a7948aEEB0c7253")
)
