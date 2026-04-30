package models

import "time"

// Usuario representa la estructura de la tabla autenticacion.usuarios
type Usuario struct {
	IDUsuario         int       `json:"id_usuario"`
	IDRol             int       `json:"id_rol"`
	Correo            string    `json:"correo"`
	Nombres           string    `json:"nombres"`
	Apellidos         string    `json:"apellidos"`
	Telefono          string    `json:"telefono"`
	FechaNacimiento   string    `json:"fecha_nacimiento"` 
	Documento         int       `json:"documento"`
	Password          string    `json:"password"`
	Estado            string    `json:"estado"` 
	Activo            bool      `json:"activo"`
	TwoFactorActivo   bool      `json:"two_factor_activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}