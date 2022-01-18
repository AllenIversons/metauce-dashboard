package restful

import (
	"github.com/gorilla/mux"
	"metauce-dashboard/tracker"
	"net/http"
)

type Tracker interface {
	GetRanking() []tracker.RankingInfo
}

type RPCService struct {
	port    string
	router  *mux.Router
	tracker Tracker
}

func Init(port string, tracker Tracker) *RPCService {
	r := mux.NewRouter()
	r.HandleFunc("/ranklist", func(writer http.ResponseWriter, request *http.Request) {

	})
	r.HandleFunc("/addrrank", func(writer http.ResponseWriter, request *http.Request) {

	})
}

func (rpc *RPCService) Start() error {
	address := "0.0.0.0:" + rpc.port
	err := http.ListenAndServe(address, rpc.router)
	if err != nil {
		return err
	}
	return nil
}
