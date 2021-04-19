package main

import (
	"firstProject/database"
	"firstProject/router"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	dictionary := make(map[int32]*string, 10000000)
	s := "hello"
	for i := int32(0); i < int32(10000000); i++ {
		dictionary[i] = &s
	}
	duration := time.Since(start)
	fmt.Println("time used:", duration)

	defer database.DB.Close()
	router.InitRouter()
}
