package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rprajapati0067/echo-web-framework/src/datasource/mssql"
	"github.com/rprajapati0067/echo-web-framework/src/domain"
)

func GetUser(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	return c.String(http.StatusOK, fmt.Sprintf("your cat name is : %s\nand cat type is : %s\n", catName, catType))
}

func AddUser(c echo.Context) error {
	type User struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Email string `json:"email"`
	}
	user := User{}
	defer c.Request().Body.Close()
	fmt.Println(c.Request().Body)
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is your user %#v", user)
	return c.String(http.StatusOK, "We got your user!!!")
}

func CreateEmployee(c echo.Context) error {

	employee := domain.Employee{}

	if err := c.Bind(&employee); err != nil {
		return err
	}

	dbCon := mssql.GetDbConnection.GetDbConnection()

	db := dbCon.Create(&employee)

	if db.RowsAffected > 0 {
		fmt.Printf("%s", db.Error)
	}

	if db.Error != nil {
		fmt.Printf("%s", db.Error)
	}

	// defer c.Request().Body.Close()
	// fmt.Println(c.Request().Body)
	// err := json.NewDecoder(c.Request().Body).Decode(&employee)
	// if err != nil {
	// 	log.Fatalf("Failed reading the request body %s", err)
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	// }
	// log.Printf("this is your user %#v", employee)
	return c.JSON(http.StatusOK, employee)
}
