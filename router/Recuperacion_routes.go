package routes

import (
	"Autenticacion/controllers"
	"github.com/gorilla/mux"
)

func SetupRecuperacionRoutes(router *mux.Router) {
    recRouter := router.PathPrefix("/api/recuperacion").Subrouter()

    recRouter.HandleFunc("/solicitar", controllers.GenerarCodigoRecuperacion).Methods("POST")
    recRouter.HandleFunc("/validar", controllers.ValidarCodigo).Methods("POST")
}