package db

import "github.com/ethereum/go-ethereum/common"

type mapnft struct {
	id      int
	tokenid int
	level   int
	total   float64
	remain  float64
	owner   string
}

func (db *DBService) PutMapInfo(tokenid int, level int, total float64, remain float64, owner common.Address) error {
	sqlStr := "INSERT INTO mapnft(tokenid, level, total, remain, owner) values (?,?,?,?,?) ON DUPLICATE KEY UPDATE tokenid=?, level=?, total=?, remain=?, owner=?"
	_, err := db.db.Exec(sqlStr, tokenid, level, total, remain, owner.String(), tokenid, level, total, remain, owner.String())
	return err
}
