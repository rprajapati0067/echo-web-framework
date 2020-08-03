package mssql

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var (
	GetDbConnection databaseConnectionInterface = &databaseCon{}
)

type databaseConnectionInterface interface {
	GetDbConnection() *gorm.DB
}

type databaseCon struct {
}

func (dbCon *databaseCon) GetDbConnection() *gorm.DB {

	db, err := gorm.Open("mssql", "sqlserver://ravi:System123@LT212-RAVISP:1433?database=employee_db")
	if err != nil {
		panic(err)
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connection successful")

	return db
}
