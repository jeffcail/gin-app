package _leveldb

import (
	leveldb1 "github.com/mazezen/leveldb"
)

func InitLevelDb(path string) *leveldb1.LevelDB {
	db, err := leveldb1.CreateLevelDB(path)
	if err != nil {
		panic(err)
	}
	return db
}
