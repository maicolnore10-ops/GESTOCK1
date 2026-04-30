package models

import "time"

type RecuperacionContra struct {
	IDUsuario         int       `json:"id_usuario"`
	ContraHash      string    `json:"contra_hash"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}