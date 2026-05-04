package routes

import (
	"Autenticacion/controllers"
	"github.com/gorilla/mux"
)

func RegisterContrasenaRoutes(router *mux.Router) {
	router.HandleFunc("/contrasena_hash", controllers.SetPassword).Methods("POST")
	router.HandleFunc("/verifyContrasena", controllers.VerifyContrasena).Methods("POST")
}