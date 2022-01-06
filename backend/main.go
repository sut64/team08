package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sut64/team08/entity"
)

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	// Run the server
	r.Run()
}
