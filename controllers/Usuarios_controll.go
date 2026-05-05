package controllers

import (
	"Autenticacion/config"
	"Autenticacion/models"
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

// Función helper para estandarizar las respuestas JSON
func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// GetAllUsuarios obtiene todos los usuarios activos
func GetAllUsuarios(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_usuario, id_rol, correo, nombres, apellidos, telefono, 
	          fecha_nacimiento, documento, estado, activo, two_factor_activo, 
	          fecha_creacion, fecha_modificacion FROM Autenticacion.usuarios WHERE activo = true`

	rows, err := config.DB.Query(query)
	fmt.Println(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()


	var usuarios []models.Usuario
	for rows.Next() {
		var u models.Usuario
		err := rows.Scan(&u.IDUsuario, &u.IDRol, &u.Correo, &u.Nombres, &u.Apellidos, 
			&u.Telefono, &u.FechaNacimiento, &u.Documento,&u.Estado, &u.Activo, 
			&u.TwoFactorActivo, &u.FechaCreacion, &u.FechaModificacion)
		
		if err != nil {
			continue
		}
		usuarios = append(usuarios, u)
	}

	respondJSON(w, 200, usuarios)
}

// GetUsuarioByID busca un usuario por su ID primario
func GetUsuarioByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var u models.Usuario

	query := `SELECT id_usuario, id_rol, correo, nombres, apellidos, telefono, 
	          fecha_nacimiento, documento, estado, activo, two_factor_activo 
	          FROM Autenticacion.usuarios WHERE id_usuario = $1`

	err := config.DB.QueryRow(query, id).Scan(&u.IDUsuario, &u.IDRol, &u.Correo, 
		&u.Nombres, &u.Apellidos, &u.Telefono, &u.FechaNacimiento, &u.Estado,  &u.Documento, 
		 &u.Activo, &u.TwoFactorActivo)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Usuario no encontrado"})
		return
	}

	respondJSON(w, 200, u)
}

// CreateUsuario inserta un nuevo usuario en la base de datos
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var u models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondJSON(w, 400, map[string]string{"error": "Datos inválidos"})
		return
	}

	query := `INSERT INTO Autenticacion.usuarios (id_rol, correo, nombres, apellidos, 
	          telefono, fecha_nacimiento, documento) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_usuario`

	err := config.DB.QueryRow(query, u.IDRol, u.Correo, u.Nombres, u.Apellidos, 
		u.Telefono, u.FechaNacimiento, u.Documento).Scan(&u.IDUsuario)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, u)
}

// UpdateUsuario modifica los datos de un usuario existente
func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var u models.Usuario
	json.NewDecoder(r.Body).Decode(&u)

	query := `UPDATE Autenticacion.usuarios SET id_rol=$1, correo=$2, nombres=$3, 
	          apellidos=$4, telefono=$5, documento=$6 WHERE id_usuario=$7`

	_, err := config.DB.Exec(query, u.IDRol, u.Correo, u.Nombres, u.Apellidos, 
		u.Telefono, u.Documento, id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Usuario actualizado correctamente"})
}

// DeleteUsuario realiza un borrado lógico (cambia activo a false)
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	query := `UPDATE Autenticacion.usuarios SET activo = false WHERE id_usuario = $1`
	_, err := config.DB.Exec(query, id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Usuario desactivado correctamente"})
}