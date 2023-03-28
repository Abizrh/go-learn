package main

import "gin-framework/routes"

func main() {
	var PORT = ":3000"

	routes.StartServer().Run(PORT)
}