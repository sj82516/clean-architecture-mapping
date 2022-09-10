package main

import (
	"fmt"
	"mapping/internal/adapter/in"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	route := gin.Default()

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/local",
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	in.NewOrderController(db, route)

	route.Run(":8080")
}
