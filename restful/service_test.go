package restful

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"metauce-dashboard/types"
	"testing"
)

func TestMarshal(t *testing.T) {
	rank1 := types.AddressRank{
		Address: common.Address{},
		Power:   10.1,
		Rank:    10,
	}
	rank2 := types.AddressRank{
		Address: common.Address{},
		Power:   11.1,
		Rank:    12,
	}
	b, err := json.Marshal([]types.AddressRank{rank1, rank2})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
