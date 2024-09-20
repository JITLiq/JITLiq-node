package entity

import "github.com/ethereum/go-ethereum/common"

const (
	ChainIDBase = 8453
	ChainIDArb  = 42161
)

var (
	TopicOrderCreated = common.HexToHash("0x918554b6bd6e2895ce6553de5de0e1a69db5289aa0e4fe193a0dcd1f14347477")
	BaseUsdc          = common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913")
	TargetJITPool     = common.HexToAddress("")
	SrcStateManager   = common.HexToAddress("0x7c62010e5CC7A839826A50a541bf20767750d345")
)
