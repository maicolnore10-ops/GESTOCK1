package models

import "time"

// Role representa la estructura de la tabla Autenticacion.roles
type Roles struct {
	IDRol             int       `json:"id_rol"`
	NombreRol         string    `json:"nombre_rol"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}