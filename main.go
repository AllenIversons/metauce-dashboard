package main

import (
	"github.com/ethereum/go-ethereum/common"
	"metauce-dashboard/tracker"
	"time"
)

type TrackService interface {
	Start() error
	Close()
	PutPower(address common.Address, power int) error
	UpdateCycleTime(duration time.Duration)
	GetRanking() []tracker.RankingInfo
}

func main() {

}
