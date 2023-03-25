package main

import (
	"github.com/labstack/echo"
)

// Main method to start the server in port: 8000
func main() {
	server := echo.New()
	server.POST("/receipts/process/", getId)
	server.GET("/receipts/:id/points/", getPoints)
	server.Start(":8000")
}
