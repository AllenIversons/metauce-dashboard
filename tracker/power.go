package tracker

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	abi "metauce-dashboard/api"
	"metauce-dashboard/types"
	"time"
)

func (s *TrackService)GetSingleUser(address common.Address)(types.UserDetailInfo,error){

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
		inveitiCodeInstance, err := abi.NewInvitecode(s.codeContract, client)
		if err != nil {
			log.Error("Failed to call inveiteCodeInstance contract", err)
			continue
		}
		playInstance, err := abi.NewPlay(s.playContract, client)
		if err != nil {
			fmt.Println(err)
			log.Error("Failed to call playInstance contract", err)
			continue
		}
		if err != nil {
			log.Error("db GetUsers error", "err", err)
			continue
		}

		ownerAddress := common.HexToAddress("")
		codeList ,_:=inveitiCodeInstance.GetMyAllInviteCode(&bind.CallOpts{Pending: false,From: ownerAddress},address)
		fmt.Printf("当前用户为:%s,邀请码个数为:%d\n",address.String(),len(codeList))
		carPowerSum := big.NewInt(0)
		codePowerSum := big.NewInt(0)
		for _,code := range codeList {
			fmt.Println("邀请码:",code)
			if code.Owner != address {
				continue
			}
			codePowerSum.Add(codePowerSum,code.Power)
		}

		fmt.Printf("邀请码总算力为:%s\n",codePowerSum.String())
		carsNFT, err := playInstance.GetUserCarNFT(&bind.CallOpts{Pending: false}, address)
		if err != nil {
			fmt.Println(err)
			log.Error("playInstance GetUserCarNFT error", "err", err)
			continue
		}

		for _, carNFT := range carsNFT {
			if carNFT.Level.Cmp(big.NewInt(0)) == 0 {
				continue
			}

			carPowerSum.Add(carPowerSum, carNFT.Power)
		}
		totalPower := new(big.Int).Add(codePowerSum, carPowerSum)
		fmt.Printf("CarNFT总算力为:%s\n",carPowerSum.String())
		fmt.Printf("总算力为:%s\n",totalPower.String())
		userDetailInfo := types.UserDetailInfo{
			Address: address,
			TotalPower: bigToFloat64(totalPower),
			CarPower: bigToFloat64(carPowerSum),
		}

		return userDetailInfo,nil
	}
	return types.UserDetailInfo{}, errors.New("get single user info" + address.String())
}
