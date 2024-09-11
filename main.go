package main

import (
	"bank-queue-system/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080") 
}
