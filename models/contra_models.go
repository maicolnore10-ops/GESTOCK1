package models

import "time"

// Password representa la estructura de la tabla Autenticacion.passwords
type ContraHash struct {
	IDContra       int       `json:"id_contra"`
	IDUsuario         int       `json:"id_usuario"`
	ContraHash      string    `json:"contra_hash"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}