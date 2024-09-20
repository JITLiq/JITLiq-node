package entity

import (
	"github.com/ethereum/go-ethereum/common"
)

type Fillers []common.Address

func ToFillers(fillers []string) Fillers {
	addr := make([]common.Address, 0)
	for _, owner := range fillers {
		addr = append(addr, common.HexToAddress(owner))
	}

	return addr
}
