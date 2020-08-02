package main

import (
	"gorm"
	"gorm/logger"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {

		db, err := gorm.Open("mssql", "sqlserver://Aditi/ravisp:Jul@4362@LT212-RAVISP:1433?database=dbname")
		if err != nil {
			panic(err)
		}

		// Ping function checks the database connectivity
		err = db.DB().Ping()
		if err != nil {
			panic(err)
		}
		logger.Info("connection successful")
		db.LogMode(true)
		return c.String(http.StatusOK, "Hello")
	})

	e.GET("/getUser", controllers.GetUser)
	e.PUT("/addUser", controllers.AddUser)

	e.Start(":8080")
}
