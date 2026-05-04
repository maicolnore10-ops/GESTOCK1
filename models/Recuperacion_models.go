package models

import "time"

type Recuperacion struct {
	IDRecuperacion     int       `json:"id_recuperacion"`
	IDUsuario          int       `json:"id_usuario"`
	CodigoVerificacion string    `json:"codigo_verificacion"`
	FechaExpiracion    time.Time `json:"fecha_expiracion"`
	Usado              bool      `json:"usado"`
	Activo             bool      `json:"activo"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaModificacion  time.Time `json:"fecha_modificacion"`
}