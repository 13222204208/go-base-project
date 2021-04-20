package main

import (
	"firstProject/database"
	"firstProject/router"
)

func main() {

	defer database.DB.Close()
	router.InitRouter()
}
