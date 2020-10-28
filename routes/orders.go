package routes

import (
	"net/http"

	"github.com/danielwetan/kdigital-backend/controllers"
)

func Orders() {
	http.HandleFunc("/orders", controllers.Orders)
}
