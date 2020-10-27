package routes

import (
	"net/http"

	"github.com/danielwetan/kdigital-backend/controllers"
)

func Auth() {
	http.HandleFunc("/auth/register", controllers.Register)
	http.HandleFunc("/auth/login", controllers.Login)
}
