package main

import (
	"fmt"

	"github.com/VictorMilhomem/GoRestApi/src/backend/database"
	"github.com/VictorMilhomem/GoRestApi/src/backend/routes"
)

func main() {

	fmt.Println("Connecting to DB")
	database.ConnectDB()

	fmt.Println("Initiating the server ")
	routes.HandleRequest()
}
