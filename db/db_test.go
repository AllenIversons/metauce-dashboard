package db

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
	"time"
)

var password = ""
var url = ""

func TestDB(t *testing.T) {
	dbService, err := Init(password, url)
	defer dbService.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = dbService.PutPower(
		common.HexToAddress("0x1ef58bb670b70b75d7c98A2B113EcF50B5aCE323"),
		3,
		int(time.Now().Unix()))
	if err != nil {
		t.Fatal(err)
	}

	err = dbService.PutPower(
		common.HexToAddress("0xdfeD40239677b88D6cE4E016273A8FeDd71F4e03"),
		5,
		int(time.Now().Unix()))
	if err != nil {
		t.Fatal(err)
	}

	err = dbService.PutPower(
		common.HexToAddress("0xbF2175bE2221c431414d0a43447f5636AD1F8fF7"),
		2,
		int(time.Now().Unix()))
	if err != nil {
		t.Fatal(err)
	}

	addrList, err := dbService.GetRanking()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", addrList)
}
