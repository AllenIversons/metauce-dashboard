package types

import "github.com/ethereum/go-ethereum/common"

type AddressRank struct {
	Address common.Address
	Power   float64
	Rank    int
}

type AddressInfo struct {
	Address common.Address
	Power   float64
}
