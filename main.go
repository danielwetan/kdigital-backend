package main

import (
	"fmt"
	"net/http"

	"github.com/danielwetan/kdigital-backend/routes"
)

func main() {

	routes.Auth()
	routes.Orders()

	PORT := ":3000"
	fmt.Println("App running on PORT", PORT[1:])
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
