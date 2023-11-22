// main.go
package main

import (
	"zocket-api/api"
)

func main() {
	// Initialize the router
	r := api.SetupRouter()

	// Run the server
	r.Run(":8080")
}
