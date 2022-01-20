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

type UserDetailInfo struct {
	Address    common.Address `json:"address"`
	TotalPower float64        `json:"total_power"`
	CarPower   float64        `json:"car_power"`
	CarNum     int            `json:"car_num"`
	MapNum     int            `json:"map_num"`
	Rank       int            `json:"rank"`
}
