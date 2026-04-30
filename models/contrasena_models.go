package models

import "time"

// ContrasenaHash representa la estructura de la tabla Autenticacion.contrasena_hash
type ContrasenaHash struct {
	IDUsuario         int       `json:"id_usuario"`
	ContrasenaHash    string    `json:"contrasena_hash"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}

// LoginRequest se usa para capturar las credenciales cuando un usuario intenta ingresar
type LoginRequest struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}