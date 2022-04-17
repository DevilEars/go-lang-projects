package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// returns a hook to the DB
var (
	db *gorm.DB
)

func Connect() {
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
