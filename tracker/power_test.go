package tracker

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"metauce-dashboard/consts"
	db2 "metauce-dashboard/db"
	"testing"
)

func TestTrackService_GetSingleUser(t *testing.T) {
	db, err := db2.Init(consts.Dbpassword, consts.Dburl)
	if err != nil {
		t.Fatal(err)
	}
	playContract := common.HexToAddress("0x7Ab359a33F14C23571f79b62f1B8a5dc510c8070")
	codeContract := common.HexToAddress("0xE5547D0378944B786BB38fB1bea4027052FC883E")
	rpcurl := "https://stardust.metis.io/?owner=588"
	service := Init(db, rpcurl, playContract, codeContract, 10000)
	address := common.HexToAddress("0xdfeD40239677b88D6cE4E016273A8FeDd71F4e03")
	userInfo,err := service.GetSingleUser(address)
	fmt.Println(userInfo)


}
