package controllers

import (
	"Autenticacion/config"
	"encoding/json"
	"net/http"
	"Autenticacion/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type RequestContra struct {
	IDUsuario int    `json:"id_usuario"`
	Password  string `json:"password"`
}

func EstablecerContrasena(w http.ResponseWriter, r *http.Request) {
	var req RequestContra
	
	// 1. Decodificar el JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	// El costo '10' es un equilibrio estándar entre seguridad y velocidad
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": "Error al cifrar contraseña"})
		return
	}

	// 3. Guardar en la base de datos
	query := `
		INSERT INTO Autenticacion.passwords (id_usuario, password_hash)
		VALUES ($1, $2)
		ON CONFLICT (id_usuario) 
		DO UPDATE SET password_hash = $2, fecha_modificacion = NOW()`

	_, err = config.DB.Exec(query, req.IDUsuario, string(hashedPassword))
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": "Error en DB: " + err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Contraseña establecida correctamente"})
}

// GetAllPasswords obtiene todas las contraseñas de un usuario
func GetAllPasswords(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var passwords []models.ContraHash

	query := `SELECT id_password, id_usuario, contrasena_hash, fecha_creacion, fecha_modificacion 
	          FROM Autenticacion.passwords WHERE id_usuario = $1`

	rows, err := config.DB.Query(query, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": "Error al consultar contraseñas: " + err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p models.ContraHash
		err := rows.Scan(&p.IDContra, &p.IDUsuario, &p.ContraHash, &p.FechaCreacion, &p.FechaModificacion)
		if err != nil {
			continue
		}
		passwords = append(passwords, p)
	}

	respondJSON(w, 200, passwords)
}

