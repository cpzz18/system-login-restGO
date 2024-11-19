package main

import (
	"myapi/config"
	"myapi/routes"

)

func main() {
	config.ConnectDB()

	routes.Route()
	
	
	
}
