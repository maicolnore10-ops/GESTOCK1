package controllers

import (
	"Autenticacion/config"
	"Autenticacion/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// RecuperacionContra actualiza la contraseña de un usuario validando su correo
func RecuperacionContra(w http.ResponseWriter, r *http.Request) {
	var u models.Usuario
	respondJSON(w, 200, map[string]string{"mensaje": "Recuperación de contraseña iniciada"})

	// 1. Decodificar el JSON
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	// 2. Primero buscamos el ID del usuario por su correo
	var idUsuario int
	err := config.DB.QueryRow(
		"SELECT id_usuario FROM Autenticacion.usuarios WHERE correo = $1",
		u.Correo,
	).Scan(&idUsuario)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "El correo no existe en el sistema"})
		return
	}

	

	// 3. Ciframos la nueva contraseña que viene en el JSON
	// Nota: Asumimos que el JSON trae un campo "password" que mapeas en el modelo
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": "Error al procesar la seguridad"})
		return
	}

	// 3. Update: Actualizamos la tabla de passwords usando el ID encontrado
	_, err = config.DB.Exec(
		"UPDATE Autenticacion.passwords SET password_hash=$1, fecha_modificacion=NOW() WHERE id_usuario=$2",
		string(hash), idUsuario,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"mensaje": "Contraseña restablecida", "correo": u.Correo})
}

// ObtenerEstadoPassword verifica si un usuario ya tiene una contraseña asignada
func ObtenerEstadoPassword(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var existe bool
	err := config.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM Autenticacion.passwords WHERE id_usuario=$1)",
		id,
	).Scan(&existe)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"id_usuario": id, "tiene_password": "true"})
}