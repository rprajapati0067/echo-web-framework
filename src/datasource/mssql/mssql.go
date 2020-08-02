package mssql

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func GetDbConnection() {

	//	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=candidates_db password=system sslmode=disable")
	db, err := gorm.Open("mssql", "sqlserver://Aditi/ravisp:Jul@4362@LT212-RAVISP:1433?database=dbname")
	if err != nil {
		panic(err)
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	//	logger.Info("connection successful")
	//	db.LogMode(true)

	// /	return db
}
