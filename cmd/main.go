package main

import (
	"gotest/bookTest/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	handlers.BookRouter(router)
	router.Run(":8080")
}
