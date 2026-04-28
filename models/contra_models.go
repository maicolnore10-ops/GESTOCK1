package models

import "time"

// Password representa la estructura de la tabla Autenticacion.passwords
type Password struct {
	IDContraseña        int       `json:"id_password"`
	IDUsuario         int       `json:"id_usuario"`
	ContraseñaHash      string    `json:"contrasena_hash"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}