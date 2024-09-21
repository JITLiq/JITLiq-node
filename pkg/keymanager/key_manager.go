package keymanager

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type KeyManager struct {
	address    common.Address
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
}

func (m *KeyManager) Sign(_ context.Context, hash string) ([]byte, error) {
	signature, err := crypto.Sign(common.HexToHash(hash).Bytes(), m.privateKey)
	if err != nil {
		return nil, err
	}

	if signature[crypto.RecoveryIDOffset] == 0 || signature[crypto.RecoveryIDOffset] == 1 {
		signature[crypto.RecoveryIDOffset] += 27 // Transform yellow paper V from 27/28 to 0/1
	}

	return signature, nil
}

func (m *KeyManager) SignTransaction(txn *types.Transaction) (*types.Transaction, error) {
	return types.SignTx(txn, types.LatestSignerForChainID(m.chainID), m.privateKey)
}

func (m *KeyManager) TransactionSignerHook(address common.Address, txn *types.Transaction) (*types.Transaction, error) {
	fmt.Println(address)
	return types.SignTx(txn, types.LatestSignerForChainID(m.chainID), m.privateKey)
}

func (m *KeyManager) Signer() common.Address {
	return m.address
}

func NewKeyManager(key string, chainID *big.Int) (*KeyManager, error) {
	ecdsaKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}

	return &KeyManager{
		address:    crypto.PubkeyToAddress(ecdsaKey.PublicKey),
		privateKey: ecdsaKey,
		chainID:    chainID,
	}, nil
}
