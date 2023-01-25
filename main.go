package main

import (
	"github.com/BruggerPedro/gin-api-rest/database"
	"github.com/BruggerPedro/gin-api-rest/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
