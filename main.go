package main

import (
    "Autenticacion/config"
    "Autenticacion/router" // Verifica que coincida con el nombre en go.mod
    "log"
    "net/http"
    "github.com/gorilla/mux"
)
func main() {
	config.ConnectDB() // Conexión a la base de datos

	r := mux.NewRouter()

	// Registro de los dos módulos
	router.RegisterUsuariosRoutes(r) // Verifica que coincida con el nombre en go.mod
	{
		router.RegisterUsuariosRoutes(r)
		router.RegisterRolesRoutes(r)
	

	log.Println("Servidor iniciado en http://localhost:8082")
	http.ListenAndServe(":8082", r)
	}
}