package tracker

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"metauce-dashboard/api"
	"metauce-dashboard/db"
	"metauce-dashboard/types"
	"strconv"
)

type RankingInfo struct {
	Address common.Address
	Power   float64
	Rank    int
}

type CacheInfo struct {
	Address   common.Address
	Power     float64
	CodePower float64
	CarPower  float64
	CarNums   int
	MapNums   int
}

type TrackService struct {
	db     DBService
	client *ethclient.Client
}

type DBService interface {
	PutPower(address common.Address, total_power int, timestamp int, carcnt int, mapcnt int, carpower float64, mapremain float64, mapmined float64) error
	GetRanking() ([]types.AddressInfo, error) //return list of addressinfo sorted by power
}

func Init(db DBService, rpcurl string) *TrackService {
	client, err := ethclient.Dial(rpcurl)
	if err != nil {
		log.Error("Dial metis rpc error", "err", err, "url", rpcurl)
		return nil
	}
	return &TrackService{db: db, client: client}
}

func (s *TrackService) PutPower(address common.Address, power float64, timestamp int) error {
	return s.db.PutPower(address, power, timestamp)
}

func (s *TrackService) GetRanking() ([]types.AddressInfo, error) {
	address, err := s.db.GetRanking()
	if err != nil {
		return nil, err
	}
	return address, nil
}
func (s *TrackService) GetUsersMap() (map[common.Address]CacheInfo, error) {
	userInfoMap := make(map[common.Address]CacheInfo)
	rpcUrl := "https://stardust.metis.io/?owner=588"
	fromBlock := new(big.Int).SetUint64(128975)
	contractAddress := []common.Address{
		common.HexToAddress("0x7Ab359a33F14C23571f79b62f1B8a5dc510c8070"),
		common.HexToAddress("0xE5547D0378944B786BB38fB1bea4027052FC883E")}
	userAddress, err := GetUsers(rpcUrl, fromBlock, contractAddress)
	if err != nil {
		log.Error("GetUsers logs error", "err", err)
		return nil, err
	}
	//timestamp := time.Now().Unix()
	//邀请码合约实例
	inveiteCodeInstance, err := abi.NewInvitecode(contractAddress[1], s.client)
	if err != nil {
		log.Error("Failed to call inveiteCodeInstance contract", err)
		return nil, err
	}
	//play 游戏合约实例
	playInstance, err := abi.NewPlay(contractAddress[0], s.client)
	if err != nil {
		log.Error("Failed to call playInstance contract", err)
	}
	for _, v := range userAddress {
		carPowerSum := big.NewInt(0)
		codePower, _ := inveiteCodeInstance.GetTotalPower(nil, v)
		codePower.Div(codePower, big.NewInt(1e+18))

		carsNFT, _ := playInstance.GetUserCarNFT(nil, v)
		carsNum := len(carsNFT)

		for _, carNFT := range carsNFT {
			carPowerSum.Add(carPowerSum, carNFT.Power)
		}
		carPowerSum.Div(carPowerSum, big.NewInt(1e+18))

		mapsNFT, _ := playInstance.GetUserMapNFT(nil, v)
		mapsNum := len(mapsNFT)
		userInfoMap[v] = CacheInfo{v,
			bigToFloat64(codePower) + bigToFloat64(carPowerSum),
			bigToFloat64(codePower),
			bigToFloat64(carPowerSum),
			carsNum,
			mapsNum}
	}
	return userInfoMap, nil
}
func GetUsers(rpcurl string, fromBlock *big.Int, contractAddresses []common.Address) ([]common.Address, error) {
	client, err := ethclient.Dial(rpcurl)
	if err != nil {
		log.Error("Dial metis rpc error", "err", err, "url", rpcurl)
		return nil, err
	}
	//currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Error("client block number error", "err", err, "url", rpcurl)
		return nil, err
	}
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   new(big.Int).SetUint64(136562),
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

func bigToFloat64(num *big.Int) float64 {
	numStr, _ := strconv.ParseFloat(num.String(), 32)
	return numStr
}