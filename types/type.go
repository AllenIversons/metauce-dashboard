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

type AddressDetailInfo struct {
	Address common.Address `json:"address"`
	Power   float64        `json:"power"`
	CarNum  int            `json:"car_num"`
	MapNum  int            `json:"map_num"`
	Rank    int            `json:"rank"`
}
