package relayer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type keymanager interface {
	TransactionSignerHook(address common.Address, txn *types.Transaction) (*types.Transaction, error)
	Signer() common.Address
}
