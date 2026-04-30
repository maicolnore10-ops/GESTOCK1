package router

import (
	"Autenticacion/controllers" 
	"github.com/gorilla/mux"
)

func RegisterRolesRoutes(r *mux.Router) {

	r.HandleFunc("/roles", controllers.GetAllRoles).Methods("GET")
	r.HandleFunc("/roles", controllers.CreateRol).Methods("POST")
	r.HandleFunc("/roles/{id}", controllers.UpdateRol).Methods("PUT")
	r.HandleFunc("/roles/{id}", controllers.DeleteRol).Methods("DELETE")
}