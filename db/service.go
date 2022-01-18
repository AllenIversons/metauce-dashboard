package db

import (
	"database/sql"
	"fmt"
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

type DBService struct {
	db *sql.DB
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

func (db *DBService) GetRanking() ([]types.AddressInfo, error) {
	sqlStr := "SELECT * from userinfo ORDER BY total_power DESC"
	rows, err := db.db.Query(sqlStr)
	if err != nil {
		log.Error("db get ranking error", "err", err)
		return nil, err
	}
	defer rows.Close()

	addrList := make([]types.AddressInfo, 0)

	for rows.Next() {
		var r powerRow
		e := rows.Scan(&r.id, &r.address, &r.total_power, &r.timestamp, &r.carcnt, &r.mapcnt, &r.carpower, &r.mapremain, &r.mapmined, &r.ranking)
		fmt.Println(r.mapremain)
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
