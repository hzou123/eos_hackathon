package main

import (
	"fmt"
	"../../../router"
	"../../../service"

)
func main() {
	fmt.Printf("Starting %v\n", "order service")

	service.StartService("6769",router.OrderRoutes)

}




