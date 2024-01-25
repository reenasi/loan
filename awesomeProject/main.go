package main

import (
	"awesomeProject/dao"
	"awesomeProject/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := dao.NewDatabase("mysql", "root:password@tcp(localhost:3306)/awesome")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	r := gin.Default()

	router.Router.SetupRoutes(r)

	port := ":8080"
	r.Run(port)
}
