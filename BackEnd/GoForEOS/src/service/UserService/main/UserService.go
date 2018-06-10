package main

import (
	"fmt"
	"../../../router"
	"../../../service"

)

func main() {
	fmt.Printf("Starting %v\n", "user service")
	service.StartService("6767",router.UserRoutes)
}




