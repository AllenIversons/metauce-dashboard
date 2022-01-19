package tracker

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"metauce-dashboard/consts"
	db2 "metauce-dashboard/db"
	"testing"
	"time"
)

func TestGetRecentUsers(t *testing.T) {
	url := "https://stardust.metis.io/?owner=588"
	contractAddresses := []common.Address{common.HexToAddress("0x7Ab359a33F14C23571f79b62f1B8a5dc510c8070")}
	client, err := ethclient.Dial(url)
	if err != nil {
		t.Fatal(err)
	}
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	users, err := getRecentUsers(client, new(big.Int).SetInt64(int64(currentBlock-1000)), new(big.Int).SetInt64(int64(currentBlock)), contractAddresses)
	if err != nil {
		t.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user.String())
	}
	fmt.Println(len(users))
}

func TestOneUserCheck(t *testing.T) {
	db, err := db2.Init(consts.Dbpassword, consts.Dburl)
	if err != nil {
		t.Fatal(err)
	}
	rpcurl := "https://stardust.metis.io/?owner=588"
	playContract := common.HexToAddress("0x7Ab359a33F14C23571f79b62f1B8a5dc510c8070")
	codeContract := common.HexToAddress("0xE5547D0378944B786BB38fB1bea4027052FC883E")
	service := Init(db, rpcurl, playContract, codeContract, 10000)
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		err = service.OneUserCheck()
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second * 3)
	}
}

func TestTrackService_OneInfoCheck(t *testing.T) {
	db, err := db2.Init(consts.Dbpassword, consts.Dburl)
	if err != nil {
		t.Fatal(err)
	}
	rpcurl := "https://stardust.metis.io/?owner=588"
	playContract := common.HexToAddress("0x7Ab359a33F14C23571f79b62f1B8a5dc510c8070")
	codeContract := common.HexToAddress("0xE5547D0378944B786BB38fB1bea4027052FC883E")
	service := Init(db, rpcurl, playContract, codeContract, 10000)
	start := time.Now()
	err = service.OneInfoCheck()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(service.userCache))
	fmt.Println(len(service.rankList))
	fmt.Println(time.Now().Sub(start).Seconds())
}
