package tracker

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"metauce-dashboard/api"
	"metauce-dashboard/types"
	"strconv"
	"sync/atomic"
	"time"
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
	MapMined  float64
	MapRemain float64
	Rank      int
}

type TrackService struct {
	db           DBService
	rpcurl       string
	playContract common.Address
	codeContract common.Address
	userCache    map[common.Address]CacheInfo
	rankList     []types.AddressRank
	duration     int
	running      int32
	exitCh       chan struct{}
}

func (s *TrackService) Start() error {
	go s.loop()
	return nil
}

func (s *TrackService) Close() {
	defer func() {
		if recover() != nil {
		}
	}()
	close(s.exitCh)
	atomic.StoreInt32(&s.running, int32(0))
}

func (s *TrackService) loop() {
	defer func() {
		atomic.StoreInt32(&s.running, int32(0))
	}()
	dur := time.Minute * time.Duration(s.duration)
	atomic.StoreInt32(&s.running, int32(1))
	timer := time.NewTimer(dur)
	log.Info("Start One Cycle")
	err := s.OneCycle()
	if err != nil {
		log.Error("one cycle failed", "err", err)
	}
	for {
		select {
		case <-s.exitCh:
			timer.Stop()
			return
		case <-timer.C:
			err := s.OneCycle()
			if err != nil {
				log.Error("one cycle failed", "err", err)
			}
			timer.Reset(dur)
		}
	}
}

func (s *TrackService) OneCycle() error {
	//get data from contract, update cache and put db
	cacheSuccess := false
	for i := 0; i < 10; i++ {
		err := s.updateCache()
		if err == nil {
			cacheSuccess = true
			break
		} else {
			log.Error("update cache error", "err", err)
		}
		time.Sleep(time.Second * 10)
	}
	if !cacheSuccess {
		return errors.New("update cache retry all failed")
	}

	//query db, update rank in cache and db
	rankSuccess := false
	for i := 0; i < 10; i++ {
		err := s.updateRank()
		if err == nil {
			rankSuccess = true
			break
		} else {
			log.Error("update rank error", "err", err)
		}
		time.Sleep(time.Second * 10)
	}
	if !rankSuccess {
		return errors.New("update rank retry all failed")
	}
	return nil
}

func (s *TrackService) GetRankingList() ([]types.AddressRank, error) {
	return s.rankList, nil
}

func (s *TrackService) GetAddressRank(address common.Address) (types.AddressRank, error) {
	if info, ok := s.userCache[address]; ok {
		return types.AddressRank{
			Address: address,
			Power:   info.Power,
			Rank:    info.Rank,
		}, nil

	} else {
		return types.AddressRank{}, errors.New("no such address in cache")
	}
}

type DBService interface {
	PutPower(address common.Address, total_power float64, timestamp int, carcnt int, mapcnt int, carpower float64, mapremain float64, mapmined float64) error
	GetRanking() ([]types.AddressInfo, error) //return list of addressinfo sorted by power
	PutRank(address common.Address, rank int) error
}

func Init(db DBService, rpcurl string, playContract common.Address, codeContract common.Address, duration int) *TrackService {
	return &TrackService{
		db:           db,
		rpcurl:       rpcurl,
		playContract: playContract,
		codeContract: codeContract,
		userCache:    make(map[common.Address]CacheInfo),
		rankList:     make([]types.AddressRank, 0),
		duration:     duration,
		running:      int32(0),
		exitCh:       make(chan struct{}),
	}
}

func (s *TrackService) updateCache() error {
	log.Info("start update cache")
	client, err := ethclient.Dial(s.rpcurl)
	defer client.Close()
	if err != nil {
		log.Error("Dial metis rpc error", "err", err, "url", s.rpcurl)
		return err
	}

	fromBlock := new(big.Int).SetUint64(128975)
	userAddress, err := GetUsers(client, fromBlock, []common.Address{s.codeContract, s.playContract})
	if err != nil {
		log.Error("GetUsers logs error", "err", err)
		return err
	}

	inveiteCodeInstance, err := abi.NewInvitecode(s.codeContract, client)
	if err != nil {
		log.Error("Failed to call inveiteCodeInstance contract", err)
		return err
	}

	playInstance, err := abi.NewPlay(s.playContract, client)
	if err != nil {
		log.Error("Failed to call playInstance contract", err)
	}
	for _, v := range userAddress {
		carPowerSum := big.NewInt(0)
		codePower, err := inveiteCodeInstance.GetTotalPower(&bind.CallOpts{Pending: false}, v)
		if err != nil {
			log.Error("inveiteCodeInstance GetTotalPower error", "err", err)
			return err
		}
		codePower.Div(codePower, big.NewInt(1e+18))

		carsNFT, err := playInstance.GetUserCarNFT(&bind.CallOpts{Pending: false}, v)
		if err != nil {
			log.Error("playInstance GetUserCarNFT error", "err", err)
			return err
		}
		carsNum := len(carsNFT)

		for _, carNFT := range carsNFT {
			carPowerSum.Add(carPowerSum, carNFT.Power)
		}
		carPowerSum.Div(carPowerSum, big.NewInt(1e+18))

		mapsNFT, err := playInstance.GetUserMapNFT(&bind.CallOpts{Pending: false}, v)
		if err != nil {
			log.Error("playInstance GetUserMapNFT error", "err", err)
			return err
		}
		mapsNum := len(mapsNFT)
		userInfo := CacheInfo{
			Address:   v,
			Power:     bigToFloat64(codePower) + bigToFloat64(carPowerSum),
			CodePower: bigToFloat64(codePower),
			CarPower:  bigToFloat64(carPowerSum),
			CarNums:   carsNum,
			MapNums:   mapsNum,
			MapMined:  0,
			MapRemain: 0,
			Rank:      0,
		}
		s.userCache[v] = userInfo

		err = s.db.PutPower(v, userInfo.Power, int(time.Now().Unix()), userInfo.CarNums, userInfo.MapNums, userInfo.CarPower, userInfo.MapRemain, userInfo.MapMined)
		if err != nil {
			log.Error("db put power error", "err", err)
			return err
		}
	}
	return nil
}

func (s *TrackService) updateRank() error {
	log.Info("start update rank")
	s.rankList = make([]types.AddressRank, 0)
	rankList, err := s.db.GetRanking()
	if err != nil {
		return err
	}
	for i, info := range rankList {
		rank := i + 1
		if cache, ok := s.userCache[info.Address]; ok {
			cache.Rank = rank
			s.userCache[info.Address] = cache
		}
		err := s.db.PutRank(info.Address, rank)
		if err != nil {
			log.Error("db put rank error", "err", err)
			continue
		}
		s.rankList = append(s.rankList, types.AddressRank{
			Address: info.Address,
			Power:   info.Power,
			Rank:    rank,
		})
	}
	return nil
}

func GetUsers(client *ethclient.Client, fromBlock *big.Int, contractAddresses []common.Address) ([]common.Address, error) {
	log.Info("getting users")
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Error("client block number error", "err", err)
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
		log.Error("filter logs error", "err", err)
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
