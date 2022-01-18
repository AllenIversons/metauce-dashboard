package tracker

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

func TestGetUsers(t *testing.T) {
	url := "https://stardust.metis.io/?owner=588"
	contractAddresses := []common.Address{common.HexToAddress("0x7Ab359a33F14C23571f79b62f1B8a5dc510c8070")}
	client, err := ethclient.Dial(url)
	if err != nil {
		t.Fatal(err)
	}
	users, err := GetUsers(client, new(big.Int).SetInt64(int64(128975)), contractAddresses)
	if err != nil {
		t.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user.String())
	}
	fmt.Println(len(users))
}
