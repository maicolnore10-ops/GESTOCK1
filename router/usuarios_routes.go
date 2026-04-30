package router

import (
	"Autenticacion/controllers" 
	"github.com/gorilla/mux"
)

// RegisterUsuariosRoutes define las rutas para el recurso de usuarios
func RegisterUsuariosRoutes(r *mux.Router) {
	r.HandleFunc("/usuarios", controllers.GetAllUsuarios).Methods("GET")
	r.HandleFunc("/usuarios/{id}", controllers.GetUsuarioByID).Methods("GET")
	r.HandleFunc("/usuarios", controllers.CreateUsuario).Methods("POST")
	r.HandleFunc("/usuarios/{id}", controllers.UpdateUsuario).Methods("PUT")
	r.HandleFunc("/usuarios/{id}", controllers.DeleteUsuario).Methods("DELETE")
}