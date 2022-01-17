package tracker

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"metauce-dashboard/db"
)

type RankingInfo struct {
	Address common.Address
	Power   int
	Rank    int
}

type DBService interface {
	PutPower(address common.Address, power int, timestamp int) error
	GetRanking() ([]db.AddressInfo, error) //return list of addressinfo sorted by power
}

func GetUsers(rpcurl string, fromBlock *big.Int, contractAddresses []common.Address) ([]common.Address, error) {
	client, err := ethclient.Dial(rpcurl)
	if err != nil {
		log.Error("Dial metis rpc error", "err", err, "url", rpcurl)
		return nil, err
	}
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Error("client block number error", "err", err, "url", rpcurl)
		return nil, err
	}
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   new(big.Int).SetUint64(currentBlock),
		Addresses: contractAddresses,
	}

	users := make(map[common.Address]bool)

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Error("filter logs error", "err", err, "url", rpcurl)
		return nil, err
	}
	for _, l := range logs {
		if l.Topics != nil && len(l.Topics) == 2 {
			addr := l.Topics[1]
			users[common.HexToAddress(addr.Hex())] = true
		}
	}

	usersList := make([]common.Address, 0)
	for k, _ := range users {
		usersList = append(usersList, k)
	}
	return usersList, nil
}
