package config

import (
	_ "github.com/jihnzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

// returns a hook to the DB
var (
	db *gorm.DB
)

func Conntect() {
	// open db connection. don't want to add usernames and passwords to DBs in here /\.,/\
	d, err := gorm.Open("mysql", "devilears: ")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
