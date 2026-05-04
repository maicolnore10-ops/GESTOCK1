package main

import (
    "Autenticacion/config"
    "Autenticacion/router" 
    "log"
    "net/http"
    "github.com/gorilla/mux"
)
func main() {
	config.ConnectDB() // Conexión a la base de datos

	r := mux.NewRouter()

	// Registro de los dos módulos
	routes.RegisterUsuariosRoutes(r) 
	routes.RegisterRolesRoutes(r)
	routes.RegisterContrasenaRoutes(r)
	routes.SetupAuditoriaRoutes(r)
	routes.SetupRecuperacionRoutes(r)
	

	log.Println("Servidor iniciado en http://localhost:8082")
	http.ListenAndServe(":8082", r)
	}
