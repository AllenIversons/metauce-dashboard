package db

import "github.com/ethereum/go-ethereum/common"

type carnft struct {
	id      int
	level   int
	power   float64
	owner   string
	tokenid int
}

func (db *DBService) PutCarInfo(tokenid int, power float64, level int, owner common.Address) error {
	sqlStr := "INSERT INTO carnft(level, power, owner, tokenid) values (?,?,?,?) ON DUPLICATE KEY UPDATE level=?, power=?, owner=?, tokenid=?"
	_, err := db.db.Exec(sqlStr, level, power, owner.String(), tokenid, level, power, owner.String(), tokenid)
	return err
}