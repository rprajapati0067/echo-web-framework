package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/rprajapati0067/echo-web-framework/src/controllers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// db, err := gorm.Open("mssql", "sqlserver://ravi:System123@LT212-RAVISP:1433?database=employee_db")
	// if err != nil {
	// 	panic(err)
	// }

	// // Ping function checks the database connectivity
	// err = db.DB().Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("connection successful")

	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello")
	})

	e.GET("/getUser", controllers.GetUser)
	e.PUT("/addUser", controllers.AddUser)
	e.PUT("/employees", controllers.CreateEmployee)

	e.Start(":8080")
}
