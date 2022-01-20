package db

import (
	"database/sql"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	_ "github.com/go-sql-driver/mysql"
	"metauce-dashboard/types"
)

type powerRow struct {
	id          int
	address     string
	total_power float64
	timestamp   int
	carcnt      int
	mapcnt      int
	carpower    float64
	mapremain   float64
	mapmined    float64
	ranking     int
}

type LastUpdateHeight struct {
	id       int
	getusers int
}

type DBService struct {
	db *sql.DB
}

func (db *DBService) GetUsers() ([]common.Address, error) {
	userList := make([]common.Address, 0)
	sqlStr := "SELECT address From userinfo"
	rows, err := db.db.Query(sqlStr)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var r powerRow
		e := rows.Scan(&r.address)
		if e != nil {
			return nil, e
		}
		addr := common.HexToAddress(r.address)
		userList = append(userList, addr)
	}
	return userList, nil
}

func (db *DBService) PutUsers(addressList []common.Address) error {
	sqlStr := "INSERT IGNORE INTO userinfo(address) values (?)"
	for _, address := range addressList {
		addrStr := address.String()
		_, err := db.db.Exec(sqlStr, addrStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DBService) GetLastUserUpdateHeight() (uint64, error) {
	sqlStr := "SELECT * from updateheight WHERE id=1"
	rows, err := db.db.Query(sqlStr)
	if err != nil {
		return 0, err
	}

	var r LastUpdateHeight
	for rows.Next() {
		e := rows.Scan(&r.id, &r.getusers)
		if e != nil {
			return 0, e
		}
		if r.id == 1 {
			return uint64(r.getusers), nil
		}
	}

	return 0, errors.New("no id 1")
}

func (db *DBService) SetLastUserUpdateHeight(height uint64) error {
	sqlStr := "update updateheight set getusers=? where id=1"
	_, err := db.db.Exec(sqlStr, height)
	return err
}

func (db *DBService) PutRank(address common.Address, rank int) error {
	sqlStr := "update userinfo set ranking=? where address=?"
	_, err := db.db.Exec(sqlStr, rank, address.String())
	if err != nil {
		log.Error("db put power err", "err", err)
		return err
	}
	return nil
}

func Init(password string, url string) (*DBService, error) {
	dsn := "root:" + password + "@tcp(" + url + ")/metauce"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DBService{db: db}, nil
}

func (db *DBService) Close() error {
	return db.db.Close()
}

func (db *DBService) PutPower(address common.Address, total_power float64, timestamp int, carcnt int, mapcnt int, carpower float64, mapremain float64, mapmined float64) error {
	sqlStr := "INSERT INTO userinfo(address, total_power, timestamp, carcnt, mapcnt, carpower, mapremain, mapmined) values (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE address=?, total_power=?, timestamp=?, carcnt=?, mapcnt=?, carpower=?, mapremain=?, mapmined=?"
	addrStr := address.String()
	_, err := db.db.Exec(sqlStr, addrStr, total_power, timestamp, carcnt, mapcnt, carpower, mapremain, mapmined,
		addrStr, total_power, timestamp, carcnt, mapcnt, carpower, mapremain, mapmined)
	if err != nil {
		log.Error("db put power err", "err", err)
		return err
	}
	return nil
}

func (db *DBService) GetDetailRanking() ([]types.UserDetailInfo, error) {
	sqlStr := "SELECT * from userinfo ORDER BY total_power DESC"
	rows, err := db.db.Query(sqlStr)
	if err != nil {
		log.Error("db get ranking error", "err", err)
		return nil, err
	}
	defer rows.Close()

	addrList := make([]types.UserDetailInfo, 0)

	for rows.Next() {
		var r powerRow
		e := rows.Scan(&r.id, &r.address, &r.total_power, &r.timestamp, &r.carcnt, &r.mapcnt, &r.carpower, &r.mapremain, &r.mapmined, &r.ranking)
		if e != nil {
			log.Error("db row scan error", "err", e)
			return nil, e
		}
		addrList = append(addrList, types.UserDetailInfo{
			Address:    common.HexToAddress(r.address),
			TotalPower: r.total_power,
			CarPower:   r.carpower,
			CarNum:     r.carcnt,
			MapNum:     r.mapcnt,
			Rank:       r.ranking,
		})
	}

	return addrList, nil
}

func (db *DBService) GetRanking() ([]types.AddressInfo, error) {
	sqlStr := "SELECT (address, total_power) from userinfo ORDER BY total_power DESC"
	rows, err := db.db.Query(sqlStr)
	if err != nil {
		log.Error("db get ranking error", "err", err)
		return nil, err
	}
	defer rows.Close()

	addrList := make([]types.AddressInfo, 0)

	for rows.Next() {
		var r powerRow
		e := rows.Scan(&r.address, &r.total_power)
		if e != nil {
			log.Error("db row scan error", "err", e)
			return nil, e
		}
		addrList = append(addrList, types.AddressInfo{
			Address: common.HexToAddress(r.address),
			Power:   r.total_power,
		})
	}
	return addrList, nil
}

func (db *DBService) GetSingleUser(addr common.Address) (types.UserDetailInfo, error) {
	sqlStr := "SELECT (total_power, carcnt, mapcnt,carpower, rank) from userinfo WHERE address=?"
	rows, err := db.db.Query(sqlStr, addr.String())
	if err != nil {
		return types.UserDetailInfo{}, err
	}
	for rows.Next() {
		var userDetail types.UserDetailInfo
		userDetail.Address = addr
		err = rows.Scan(&userDetail.TotalPower, &userDetail.CarNum, &userDetail.MapNum, &userDetail.CarPower, &userDetail.Rank)
		if err != nil {
			return userDetail, nil
		}
	}
	return types.UserDetailInfo{}, err
}
