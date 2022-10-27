package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type LogDeposit struct {
	From  common.Address
	Value *big.Int
}

type LogClose struct {
	From common.Address
}

var (
	LogCreateSigHash = crypto.Keccak256Hash([]byte("CreateDeposit(address,uint256)")).Hex()
	LogCloseSigHash  = crypto.Keccak256Hash([]byte("CloseDeposit(address)")).Hex()
)
