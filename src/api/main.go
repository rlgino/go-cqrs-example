package main

import (
	"log"
	"net/http"

	productController "github.com/rlgino/go-cqrs-example/src/api/modules/product/infraestructure/controller"
	userController "github.com/rlgino/go-cqrs-example/src/api/modules/user/infraestructure/controller"
)

func main() {
	userController.Run()
	productController.Run()

	log.Println("Running in port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
