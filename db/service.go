package db

import (
	"database/sql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	_ "github.com/go-sql-driver/mysql"
)

type AddressInfo struct {
	Address common.Address
	Power   int
}

type powerRow struct {
	id        int
	address   string
	power     int
	timestamp int
}

type DBService struct {
	db *sql.DB
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

func (db *DBService) PutPower(address common.Address, power int, timestamp int) error {
	sqlStr := "INSERT INTO power(address, power, timestamp) values (?,?,?) ON DUPLICATE KEY UPDATE address=?, power=?, timestamp=?"
	addrStr := address.String()
	_, err := db.db.Exec(sqlStr, addrStr, power, timestamp, addrStr, power, timestamp)
	if err != nil {
		log.Error("db put power err", "err", err)
		return err
	}
	return nil
}

func (db *DBService) GetRanking() ([]AddressInfo, error) {
	sqlStr := "SELECT * from power ORDER BY power DESC"
	rows, err := db.db.Query(sqlStr)
	if err != nil {
		log.Error("db get ranking error", "err", err)
		return nil, err
	}
	defer rows.Close()

	addrList := make([]AddressInfo, 0)

	for rows.Next() {
		var r powerRow
		e := rows.Scan(&r.id, &r.address, &r.power, &r.timestamp)
		if e != nil {
			log.Error("db row scan error", "err", e)
			return nil, e
		}
		addrList = append(addrList, AddressInfo{
			Address: common.HexToAddress(r.address),
			Power:   r.power,
		})
	}

	return addrList, nil
}