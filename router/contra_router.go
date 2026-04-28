package router

import (
	"Autenticacion/controllers"
	"github.com/gorilla/mux"
)

func RegisterContraRoutes(r *mux.Router) {
	r.HandleFunc("/contraseña", controllers.EstablecerContrasena).Methods("GET")
	r.HandleFunc("/passwords/{id}", controllers.EstablecerContrasena).Methods("GET")
	r.HandleFunc("/passwords", controllers.EstablecerContrasena).Methods("POST")
	r.HandleFunc("/passwords/{id}", controllers.EstablecerContrasena).Methods("PUT")
	r.HandleFunc("/passwords/{id}", controllers.EstablecerContrasena).Methods("DELETE")

	r.HandleFunc("/set-password", controllers.EstablecerContrasena).Methods("POST")
}