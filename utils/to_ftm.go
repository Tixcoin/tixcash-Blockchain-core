package utils

import "math/big"

// ToTxh number of TXH to Wei
func ToTxh(txh uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(txh), big.NewInt(1e18))
}
