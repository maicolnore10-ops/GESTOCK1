package controllers

import (
	"Autenticacion/config"
	"Autenticacion/models"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Estructura para recibir datos de contraseña
type PasswordUpdate struct {
	IDUsuario int    `json:"id_usuario"`
	Password  string `json:"password"`
}

// SetPassword guarda o actualiza el hash de la contraseña de un usuario
func SetPassword(w http.ResponseWriter, r *http.Request) {
	var p PasswordUpdate
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondJSON(w, 400, map[string]string{"error": "Datos inválidos"})
		return
	}

	// 1. Encriptar la contraseña (Bcrypt)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": "Error al procesar la contraseña"})
		return
	}

	// 2. Guardar en la tabla contrasena_hash (usando ON CONFLICT para actualizar si ya existe)
	query := `
		INSERT INTO Autenticacion.contrasena_hash (id_usuario, contrasena_hash)
		VALUES ($1, $2)
		ON CONFLICT (id_usuario) 
		DO UPDATE SET contrasena_hash = EXCLUDED.contrasena_hash, fecha_modificacion = CURRENT_TIMESTAMP`

	_, err = config.DB.Exec(query, p.IDUsuario, string(hashedPassword))
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": "Error al guardar en la base de datos: " + err.Error()})
		return
	}

	respondJSON(w, 201, map[string]string{"message": "Contraseña establecida correctamente"})
}

// VerifyPassword es una función de ejemplo para el Login
func VerifyContrasena(w http.ResponseWriter, r *http.Request) {
	var loginData PasswordUpdate
	json.NewDecoder(r.Body).Decode(&loginData)

	var storedHash string
	query := `SELECT contrasena_hash FROM Autenticacion.contrasena_hash WHERE id_usuario = $1`

	err := config.DB.QueryRow(query, loginData.IDUsuario).Scan(&storedHash)
	if err != nil {
		respondJSON(w, 401, map[string]string{"error": "Usuario no tiene contraseña configurada"})
		return
	}

	// Comparar hash con la contraseña recibida
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(loginData.Password))
	if err != nil {
		respondJSON(w, 401, map[string]string{"error": "Contraseña incorrecta"})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Autenticación exitosa"})
}
func VerifyPassword(w http.ResponseWriter, r *http.Request) {
	var creds models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		respondJSON(w, 400, map[string]string{"error": "Datos inválidos"})
		return
	}

	var storedHash string
	query := `SELECT contrasena_hash FROM Autenticacion.contrasena_hash WHERE id_usuario = $1`
	err := config.DB.QueryRow(query, creds.IDUsuario).Scan(&storedHash)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Usuario no encontrado"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(creds.Password))
	if err != nil {
		respondJSON(w, 401, map[string]string{"valido": "false", "mensaje": "Contraseña incorrecta"})
		return
	}

	respondJSON(w, 200, map[string]string{"valido": "true", "mensaje": "¡Es la misma contraseña!"})
}
