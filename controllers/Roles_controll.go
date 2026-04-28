package controllers

import (
	"Autenticacion/config"
	"Autenticacion/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllRoles obtiene la lista de todos los roles activos
func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	// Inicializamos como slice vacío para evitar el "null" en el JSON
	roles := []models.Roles{}

	query := `SELECT id_rol, nombre_rol, activo, fecha_creacion, fecha_modificacion 
	          FROM Autenticacion.roles WHERE activo = true`

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": "Error al consultar roles: " + err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rol models.Roles
		err := rows.Scan(&rol.IDRol, &rol.NombreRol, &rol.Activo, &rol.FechaCreacion, &rol.FechaModificacion)
		if err != nil {
			continue
		}
		roles = append(roles, rol)
	}

	respondJSON(w, 200, roles)
}

// CreateRol inserta un nuevo rol (ej: 'Administrador', 'Vendedor')
func CreateRol(w http.ResponseWriter, r *http.Request) {
	var rol models.Roles
	if err := json.NewDecoder(r.Body).Decode(&rol); err != nil {
		respondJSON(w, 400, map[string]string{"error": "Cuerpo de petición inválido"})
		return
	}

	query := `INSERT INTO Autenticacion.roles (nombre_rol) 
	          VALUES ($1) RETURNING id_rol, fecha_creacion`

	err := config.DB.QueryRow(query, rol.NombreRol).Scan(&rol.IDRol, &rol.FechaCreacion)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": "Error al crear rol: " + err.Error()})
		return
	}

	respondJSON(w, 201, rol)
}

// UpdateRol permite cambiar el nombre de un rol
func UpdateRol(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var rol models.Roles
	json.NewDecoder(r.Body).Decode(&rol)

	query := `UPDATE Autenticacion.roles SET nombre_rol = $1, fecha_modificacion = NOW() 
	          WHERE id_rol = $2`

	_, err := config.DB.Exec(query, rol.NombreRol, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Rol actualizado con éxito"})
}