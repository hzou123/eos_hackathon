package main

import (
	"fmt"
	"../../../router"
	"../../../service"
)

func main() {
	fmt.Printf("Starting %v\n", "product service")
	service.StartService("6768",router.ProductRoutes)
}




