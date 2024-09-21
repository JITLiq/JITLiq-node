package entity

import "github.com/ethereum/go-ethereum/common"

const (
	ChainIDBase = 8453
	ChainIDArb  = 42161
)

var (
	TopicOrderCreated  = common.HexToHash("0x918554b6bd6e2895ce6553de5de0e1a69db5289aa0e4fe193a0dcd1f14347477")
	BaseUsdc           = common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913")
	TargetJITPool      = common.HexToAddress("0x5eB42e519408ec6E2e76928Aac347352d2375904")
	TargetStateManager = common.HexToAddress("0xFE4A47a399A76db4dd98bbfb6913C2cC3d585BAf")
	SrcStateManager    = common.HexToAddress("0x6dA38F0A1EB71cD8279F073604c859eA67a50D6F")
)
