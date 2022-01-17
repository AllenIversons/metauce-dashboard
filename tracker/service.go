package tracker

import (
	"github.com/ethereum/go-ethereum/common"
	"metauce-dashboard/db"
)

type RankingInfo struct {
	Address common.Address
	Power   int
	Rank    int
}

type DBService interface {
	PutPower(address common.Address, power int) error
	Flush() error
	GetRanking() []db.AddressInfo //return list of addressinfo sorted by power
}
