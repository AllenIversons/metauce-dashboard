package tracker

import (
	"context"
	"errors"
	"fmt"
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

const MaxRequestHeight = 5000

type DBService interface {
	PutPower(address common.Address, total_power float64, timestamp int, carcnt int, mapcnt int, carpower float64, mapremain float64, mapmined float64) error
	GetRanking() ([]types.AddressInfo, error) //return list of addressinfo sorted by power
	PutRank(address common.Address, rank int) error
	GetUsers() ([]common.Address, error)
	PutUsers(addressList []common.Address) error
	GetLastUserUpdateHeight() (uint64, error)
	SetLastUserUpdateHeight(height uint64) error
	PutCarInfo(tokenid int, power float64, level int, owner common.Address) error
	PutMapInfo(tokenid int, level int, total float64, remain float64, owner common.Address) error
}

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
	db                 DBService
	rpcurl             string
	playContract       common.Address
	codeContract       common.Address
	userCache          map[common.Address]CacheInfo
	rankList           []types.AddressRank
	infoLoopDur        int
	userLoopDur        int
	infoLoopRunning    int32
	getUserLoopRunning int32
	exitInfoLoopCh     chan struct{}
	exitGetUserLoopCh  chan struct{}
}

func (s *TrackService) Start() error {
	go s.userInfoLoop()
	go s.getUserLoop()
	return nil
}

func (s *TrackService) Close() {
	defer func() {
		if recover() != nil {
		}
	}()
	close(s.exitInfoLoopCh)
	close(s.exitGetUserLoopCh)
	atomic.StoreInt32(&s.infoLoopRunning, int32(0))
	atomic.StoreInt32(&s.getUserLoopRunning, int32(0))
}

func (s *TrackService) userInfoLoop() {
	defer func() {
		atomic.StoreInt32(&s.infoLoopRunning, int32(0))
	}()
	dur := time.Minute * time.Duration(s.infoLoopDur)
	atomic.StoreInt32(&s.infoLoopRunning, int32(1))
	timer := time.NewTimer(dur)
	log.Info("Start One Cycle")
	err := s.OneInfoCheck()
	if err != nil {
		log.Error("one cycle failed", "err", err)
	}
	for {
		select {
		case <-s.exitInfoLoopCh:
			timer.Stop()
			return
		case <-timer.C:
			err := s.OneInfoCheck()
			if err != nil {
				log.Error("one cycle failed", "err", err)
			}
			timer.Reset(dur)
		}
	}
}

func (s *TrackService) OneInfoCheck() error {
	//get data from contract, update cache and put db
	cacheSuccess := false
	for i := 0; i < 10; i++ {
		fmt.Println("start update cache")
		err := s.updateCache()
		if err == nil {
			cacheSuccess = true
			break
		} else {
			fmt.Println(err)
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
		fmt.Println("start update rank")
		err := s.updateRank()
		if err == nil {
			rankSuccess = true
			break
		} else {
			fmt.Println(err)
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

func Init(db DBService, rpcurl string, playContract common.Address, codeContract common.Address, duration int) *TrackService {
	return &TrackService{
		db:                 db,
		rpcurl:             rpcurl,
		playContract:       playContract,
		codeContract:       codeContract,
		userCache:          make(map[common.Address]CacheInfo),
		rankList:           make([]types.AddressRank, 0),
		infoLoopDur:        duration,
		userLoopDur:        10,
		infoLoopRunning:    int32(0),
		getUserLoopRunning: int32(0),
		exitInfoLoopCh:     make(chan struct{}),
		exitGetUserLoopCh:  make(chan struct{}),
	}
}

func (s *TrackService) checkOneAddress(addr common.Address) (CacheInfo, error) {
	for i := 0; i < 10; i++ {
		if i != 0 {
			time.Sleep(time.Second)
		}
		client, err := ethclient.Dial(s.rpcurl)
		if err != nil {
			fmt.Println(err)
			log.Error("Dial metis rpc error", "err", err, "url", s.rpcurl)
			continue
		}
		inveiteCodeInstance, err := abi.NewInvitecode(s.codeContract, client)
		if err != nil {
			fmt.Println(err)
			log.Error("Failed to call inveiteCodeInstance contract", err)
			continue
		}

		playInstance, err := abi.NewPlay(s.playContract, client)
		if err != nil {
			fmt.Println(err)
			log.Error("Failed to call playInstance contract", err)
			continue
		}

		fmt.Printf("checking %v\n", addr.String())
		carPowerSum := big.NewInt(0)
		codePower, err := inveiteCodeInstance.GetTotalPower(&bind.CallOpts{Pending: false}, addr)
		if err != nil {
			fmt.Println(err)
			log.Error("inveiteCodeInstance GetTotalPower error", "err", err)
			continue
		}

		carsNFT, err := playInstance.GetUserCarNFT(&bind.CallOpts{Pending: false}, addr)
		if err != nil {
			fmt.Println(err)
			log.Error("playInstance GetUserCarNFT error", "err", err)
			continue
		}
		carsNum := len(carsNFT)

		for _, carNFT := range carsNFT {
			if carNFT.Level.Cmp(big.NewInt(0)) == 0 {
				continue
			}
			carPowerSum.Add(carPowerSum, carNFT.Power)
			e := s.db.PutCarInfo(int(carNFT.TokenId.Uint64()), bigToFloat64(carNFT.Power), int(carNFT.Level.Uint64()), carNFT.Owner)
			if e != nil {
				fmt.Println(e)
				log.Error("db PutCarInfo error", "err", e)
			}
		}

		mapsNFT, err := playInstance.GetUserMapNFT(&bind.CallOpts{Pending: false}, addr)
		if err != nil {
			fmt.Println(err)
			log.Error("playInstance GetUserMapNFT error", "err", err)
			continue
		}
		mapsNum := len(mapsNFT)
		remainSum := big.NewInt(0)
		minedSum := big.NewInt(0)
		for _, mapNFT := range mapsNFT {
			remain := new(big.Int).Sub(mapNFT.TotalUsdt, mapNFT.ConsumeUsdt)
			remainSum.Add(remainSum, remain)
			minedSum.Add(minedSum, mapNFT.ConsumeUsdt)
			e := s.db.PutMapInfo(int(mapNFT.TokenId.Uint64()), int(mapNFT.Level.Int64()), bigToFloat64(mapNFT.TotalUsdt), bigToFloat64(remain), mapNFT.Owner)
			if e != nil {
				fmt.Println(e)
				log.Error("db PutCarInfo error", "err", e)
			}
		}

		totalPower := new(big.Int).Add(codePower, carPowerSum)
		userInfo := CacheInfo{
			Address:   addr,
			Power:     bigToFloat64(totalPower),
			CodePower: bigToFloat64(codePower),
			CarPower:  bigToFloat64(carPowerSum),
			CarNums:   carsNum,
			MapNums:   mapsNum,
			MapMined:  bigToFloat64(minedSum),
			MapRemain: bigToFloat64(remainSum),
			Rank:      0,
		}
		return userInfo, nil
	}
	return CacheInfo{}, errors.New("check one info" + addr.String())
}

func (s *TrackService) updateCache() error {
	log.Info("start update cache")
	userCache := make(map[common.Address]CacheInfo)
	userAddress, err := s.db.GetUsers()
	if err != nil {
		log.Error("db GetUsers error", "err", err)
		return err
	}
	fmt.Println(len(userAddress))
	for i, addr := range userAddress {
		fmt.Println(i)
		cache, e := s.checkOneAddress(addr)
		if e != nil {
			return e
		}
		userCache[addr] = cache
	}

	//put into db
	for _, userInfo := range userCache {
		fmt.Println("put db userinfo" + userInfo.Address.String())
		for i := 0; i < 10; i++ {
			if i != 0 {
				time.Sleep(time.Second)
			}
			err = s.db.PutPower(userInfo.Address, userInfo.Power, int(time.Now().Unix()), userInfo.CarNums, userInfo.MapNums, userInfo.CarPower, userInfo.MapRemain, userInfo.MapMined)
			if err != nil {
				fmt.Println(err)
				log.Error("db put power error", "err", err)
				continue
			} else {
				break
			}
		}
	}

	s.userCache = userCache
	return nil
}

func (s *TrackService) updateRank() error {
	log.Info("start update rank")
	rankListCache := make([]types.AddressRank, 0)
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
		fmt.Println(info.Address.String(), rank)
		e := s.db.PutRank(info.Address, rank)
		if e != nil {
			fmt.Println(e)
			log.Error("db put rank error", "err", e)
			continue
		}
		rankListCache = append(rankListCache, types.AddressRank{
			Address: info.Address,
			Power:   info.Power,
			Rank:    rank,
		})
	}
	s.rankList = rankListCache
	return nil
}

func getRecentUsers(client *ethclient.Client, fromBlock *big.Int, toBlock *big.Int, contractAddresses []common.Address) ([]common.Address, error) {
	log.Info("getting users")
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
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

func (s *TrackService) getUserLoop() {
	defer func() {
		atomic.StoreInt32(&s.getUserLoopRunning, int32(0))
	}()
	dur := time.Minute * time.Duration(s.userLoopDur)
	atomic.StoreInt32(&s.getUserLoopRunning, int32(1))
	timer := time.NewTimer(dur)

	for {
		select {
		case <-s.exitGetUserLoopCh:
			timer.Stop()
			return
		case <-timer.C:
			timer.Reset(dur)
			err := s.OneUserCheck()
			if err != nil {
				log.Error("OneUserCheck error", "err", err)
				continue
			}
		}
	}
}

func (s *TrackService) OneUserCheck() error {
	//check users
	users := make([]common.Address, 0)
	toHeight := uint64(0)
	checkSuccess := false
	for i := 0; i < 10; i++ {
		fmt.Println("check users")
		if i != 0 {
			time.Sleep(time.Second * 1)
		}
		client, err := ethclient.Dial(s.rpcurl)
		if err != nil {
			fmt.Println(err)
			log.Error("ethclient dial error", "err", err)
			continue
		}
		currentHeight, err := client.BlockNumber(context.Background())
		if err != nil {
			fmt.Println(err)
			log.Error("client blocknumber error", "err", err)
			continue
		}
		fromHeight, err := s.db.GetLastUserUpdateHeight()
		if err != nil {
			fmt.Println(err)
			log.Error("db GetLastUserUpdateHeight error", "err", err)
			continue
		}

		if (currentHeight - fromHeight) > MaxRequestHeight {
			toHeight = fromHeight + MaxRequestHeight
		} else {
			toHeight = currentHeight
		}

		fmt.Printf("from %v to %v\n", fromHeight, toHeight)
		users, err = getRecentUsers(client, big.NewInt(int64(fromHeight)), big.NewInt(int64(toHeight)), []common.Address{s.codeContract, s.playContract})
		if err != nil {
			fmt.Println(err)
			log.Error("getRecentUsers error", "err", err)
			continue
		}
		checkSuccess = true
		break
	}
	if !checkSuccess {
		return errors.New("OneUserCheck check users failed")
	}

	//write db
	writeSuccess := false
	for i := 0; i < 10; i++ {
		fmt.Println("write db")
		if i != 0 {
			time.Sleep(time.Second)
		}
		err := s.db.PutUsers(users)
		if err != nil {
			fmt.Println(err)
			log.Error("db put users error", "err", err)
			continue
		}
		err = s.db.SetLastUserUpdateHeight(toHeight)
		if err != nil {
			fmt.Println(err)
			log.Error("db SetLastUserUpdateHeight error", "err", err)
			continue
		}
		writeSuccess = true
		break
	}

	if writeSuccess {
		return nil
	} else {
		return errors.New("OneUserCheck write db failed")
	}
}

func bigToFloat64(num *big.Int) float64 {
	n := new(big.Int).Div(num, big.NewInt(1e+18))
	mod := new(big.Int).Mod(num, big.NewInt(1e+18))
	modInt := new(big.Int).Div(mod, big.NewInt(1e+16))
	nStr := n.String() + "." + modInt.String()
	numStr, _ := strconv.ParseFloat(nStr, 64)
	return numStr
}
