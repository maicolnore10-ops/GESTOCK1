package controllers

import (
	"encoding/json"
	"net/http"
	"time"
	"math/rand"
	"fmt"
	"Autenticacion/config"
	"Autenticacion/models"
)

// GenerarCodigoRecuperacion crea un código de 6 dígitos válido por 15 minutos
func GenerarCodigoRecuperacion(w http.ResponseWriter, r *http.Request) {
	var input struct {
		IDUsuario int `json:"id_usuario"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID de usuario requerido"})
		return
	}

	// Generar un código aleatorio de 6 cifras
	rand.Seed(time.Now().UnixNano())
	codigo := fmt.Sprintf("%06d", rand.Intn(1000000))
	expiracion := time.Now().Add(15 * time.Minute)

	query := `INSERT INTO autenticacion.recuperacion_contrasena 
	(id_usuario, codigo_verificacion, fecha_expiracion) 
	VALUES ($1, $2, $3)`

	_, err := config.DB.Exec(query, input.IDUsuario, codigo, expiracion)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": "No se pudo generar el código"})
		return
	}

	respondJSON(w, 201, map[string]string{
		"mensaje": "Código generado con éxito",
		"codigo":  codigo,
	})
}

// ValidarCodigo permite verificar si el código ingresado es correcto y no ha expirado
func ValidarCodigo(w http.ResponseWriter, r *http.Request) {
	var input struct {
		IDUsuario int    `json:"id_usuario"`
		Codigo    string `json:"codigo"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondJSON(w, 400, map[string]string{"error": "Datos inválidos: usuario no encontrado"})
		return
	}

	var rec models.Recuperacion
	query := `SELECT id_recuperacion, fecha_expiracion, usado 
	          FROM autenticacion.recuperacion_contrasena 
	          WHERE id_usuario = $1 AND codigo_verificacion = $2 AND activo = true`

	err := config.DB.QueryRow(query, input.IDUsuario, input.Codigo).Scan(&rec.IDRecuperacion, &rec.FechaExpiracion, &rec.Usado)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Código incorrecto o no encontrado"})
		return
	}

	if rec.Usado {
		respondJSON(w, 400, map[string]string{"error": "Este código ya fue utilizado"})
		return
	}

	if time.Now().After(rec.FechaExpiracion) {
		respondJSON(w, 400, map[string]string{"error": "El código ha expirado"})
		return
	}

	// Marcar como usado si todo es correcto
	updateQuery := `UPDATE autenticacion.recuperacion_contrasena SET usado = true WHERE id_recuperacion = $1`
	config.DB.Exec(updateQuery, rec.IDRecuperacion)

	respondJSON(w, 200, map[string]string{"mensaje": "Código validado. Procede a cambiar tu contraseña."})
}