package models

import "time"

type Auditoria struct {
	IDAuditoria       int         `json:"id_auditoria"`
	IDUsuario         int         `json:"id_usuario"`
	Accion            string      `json:"accion"`
	Fecha             time.Time   `json:"fecha"`
	IP                string      `json:"ip"`
	NombreTabla       string      `json:"nombre_tabla"`
	IDRegistro        int64       `json:"id_registro"`
	ValoresAnteriores interface{} `json:"valores_anteriores"` // jsonb en Postgres
	ValoresNuevos     interface{} `json:"valores_nuevos"`     // jsonb en Postgres
	AgenteUsuario     string      `json:"agente_usuario"`
	Activo            bool        `json:"activo"`
}