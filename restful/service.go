package restful

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/mux"
	"metauce-dashboard/db"
	"metauce-dashboard/types"
	"net/http"
)

type Tracker interface {
	GetRankingList() ([]types.AddressRank, error)
	GetAddressRank(address common.Address) (types.AddressRank, error)
}

type RPCService struct {
	port    string
	tracker Tracker
	db      *db.DBService
}

func Init(port string, tracker Tracker, db *db.DBService) *RPCService {
	return &RPCService{
		port:    port,
		tracker: tracker,
		db:      db,
	}
}

func (rpc *RPCService) Start() error {
	address := "0.0.0.0:" + rpc.port
	r := mux.NewRouter()
	r.HandleFunc("/ranklist", func(writer http.ResponseWriter, request *http.Request) {
		rankList, err := rpc.db.GetDetailRanking()
		//rankList, err := rpc.tracker.GetRankingList()
		if err != nil {
			log.Error("db get rank list error", "err", err)
			_, e := writer.Write([]byte{})
			if e != nil {
				log.Error("write response err", "err", err)
			}
			return
		}
		b, err := json.Marshal(rankList)
		if err != nil {
			log.Error("json marshal ranklist error", "err", err)
			_, e := writer.Write([]byte{})
			if e != nil {
				log.Error("write response err", "err", err)
			}
			return
		}
		_, e := writer.Write(b)
		if e != nil {
			log.Error("write response err", "err", err)
		}
	})
	//r.HandleFunc("/addrrank/{address}", func(writer http.ResponseWriter, request *http.Request) {
	//	vars := mux.Vars(request)
	//	addrStr := vars["address"]
	//	addr := common.HexToAddress(addrStr)
	//	addrRank, err := rpc.tracker.GetAddressRank(addr)
	//	if err != nil {
	//		log.Error("tracker get address rank error", "err", err)
	//		_, e := writer.Write([]byte{})
	//		if e != nil {
	//			log.Error("write response err", "err", err)
	//		}
	//		return
	//	}
	//	b, err := json.Marshal(addrRank)
	//	if err != nil {
	//		log.Error("json marshal address rank error", "err", err)
	//		_, e := writer.Write([]byte{})
	//		if e != nil {
	//			log.Error("write response err", "err", err)
	//		}
	//		return
	//	}
	//	_, e := writer.Write(b)
	//	if e != nil {
	//		log.Error("write response err", "err", err)
	//	}
	//})
	go func() {
		err := http.ListenAndServe(address, r)
		if err != nil {
			utils.Fatalf("http listen error", err)
		}
	}()
	return nil
}
