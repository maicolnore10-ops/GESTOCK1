package routes

import (
	"Autenticacion/controllers"
	"github.com/gorilla/mux"
)

func SetupAuditoriaRoutes(router *mux.Router) {
    // Creamos el prefijo /api/auditoria
	auditoriaRouter := router.PathPrefix("/api/auditoria").Subrouter()

	auditoriaRouter.HandleFunc("/registro", controllers.CrearRegistroAuditoria).Methods("POST")
	
    auditoriaRouter.HandleFunc("/historial", controllers.GetAuditorias).Methods("GET")
}