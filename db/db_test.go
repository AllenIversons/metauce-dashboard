package db

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"metauce-dashboard/consts"
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	dbService, err := Init(consts.Dbpassword, consts.Dburl)
	defer dbService.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = dbService.PutPower(
		common.HexToAddress("0x1ef58bb670b70b75d7c98A2B113EcF50B5aCE323"),
		21,
		int(time.Now().Unix()), 0, 0, 0, 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = dbService.PutPower(
		common.HexToAddress("0xdfeD40239677b88D6cE4E016273A8FeDd71F4e03"),
		22,
		int(time.Now().Unix()), 0, 0, 0, 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = dbService.PutPower(
		common.HexToAddress("0xbF2175bE2221c431414d0a43447f5636AD1F8fF7"),
		23,
		int(time.Now().Unix()), 0, 0, 0, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	err = dbService.PutRank(common.HexToAddress("0x1ef58bb670b70b75d7c98A2B113EcF50B5aCE323"), 100)
	if err != nil {
		t.Fatal(err)
	}
	addrList, err := dbService.GetRanking()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", addrList)
}

func TestDBService_SetLastUserUpdateHeight(t *testing.T) {
	db, err := Init(consts.Dbpassword, consts.Dburl)
	if err != nil {
		t.Fatal(err)
	}

	err = db.SetLastUserUpdateHeight(uint64(128975))
	if err != nil {
		t.Fatal(err)
	}
}

func TestDBService_GetLastUserUpdateHeight(t *testing.T) {
	db, err := Init(consts.Dbpassword, consts.Dburl)
	if err != nil {
		t.Fatal(err)
	}
	height, err := db.GetLastUserUpdateHeight()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(height)
}

func TestDBService_PutUsers(t *testing.T) {
	db, err := Init(consts.Dbpassword, consts.Dburl)
	if err != nil {
		t.Fatal(err)
	}
	users := []string{"0xb9c98869713e83d191b64290f2670CAA4eABE22D", "0x449819f6F560253d8b24f9A38fF14063746429E3", "0xE0bEaC874AC56f4a2cbAc27d73B5ef7a9FB278d3"}
	userAdddress := make([]common.Address, 0)
	for _, u := range users {
		userAdddress = append(userAdddress, common.HexToAddress(u))
	}
	err = db.PutUsers(userAdddress)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDBService_GetUsers(t *testing.T) {
	db, err := Init(consts.Dbpassword, consts.Dburl)
	if err != nil {
		t.Fatal(err)
	}
	userList, err := db.GetUsers()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(userList))
}

func TestDBService_Getcar(t *testing.T) {
	db, err := Init(consts.Dbpassword, consts.Dburl)
	if err != nil {
		t.Fatal(err)
	}
	sqlStr := "SELECT * From carnft where level=4"
	rows, err := db.db.Query(sqlStr)
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	for rows.Next() {
		i++
		var car carnft
		e := rows.Scan(&car.id, &car.level, &car.power, &car.owner, &car.tokenid)
		if e != nil {
			t.Fatal(e)
		}
	}
	fmt.Println(i)
}

func TestDBService_Getmap(t *testing.T) {
	db, err := Init(consts.Dbpassword, consts.Dburl)
	if err != nil {
		t.Fatal(err)
	}
	sqlStr := "SELECT * From mapnft where level=5"
	rows, err := db.db.Query(sqlStr)
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	for rows.Next() {
		i++
		var mapn mapnft
		e := rows.Scan(&mapn.id, &mapn.tokenid, &mapn.level, &mapn.total, &mapn.remain, &mapn.owner)
		if e != nil {
			t.Fatal(e)
		}
	}

	fmt.Println(i)
}
