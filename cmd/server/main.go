package main

import "go_ecommerce_backend/internal/initialize"

func main() {
	// router := routers.NewRouters()
	// router.Run(":8002") // listen and serve on 0.0.0.0:8080
	initialize.Run()
}
