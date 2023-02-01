package main

import (
	"github.com/BruggerPedro/gin-api-rest/database"
	"github.com/BruggerPedro/gin-api-rest/routes"
)

// @title APP test Gin
// @version ¿
// @description Uma API que fornece cadastro e dados de alunoss
func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
