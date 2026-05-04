package controllers

import (
	"encoding/json"
	"net/http"
	"Autenticacion/config"
	"Autenticacion/models"
)

// RegistrarAuditoria es una función interna que puedes llamar desde otros controladores
// para guardar logs automáticamente cuando alguien cree, edite o elimine algo.
func RegistrarAuditoria(a models.Auditoria) error {
	query := `INSERT INTO autenticacion.auditoria 
    (id_usuario, accion, ip, nombre_tabla, id_registro, valores_anteriores, valores_nuevos, agente_usuario) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	valAnt, _ := json.Marshal(a.ValoresAnteriores)
	valNue, _ := json.Marshal(a.ValoresNuevos)

	_, err := config.DB.Exec(query, a.IDUsuario, a.Accion, a.IP, a.NombreTabla, a.IDRegistro, valAnt, valNue, a.AgenteUsuario)
	return err
}

// GetAuditorias obtiene el historial completo para el administrador
func GetAuditorias(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_auditoria, id_usuario, accion, fecha, ip, nombre_tabla, id_registro, agente_usuario 
              FROM autenticacion.auditoria WHERE activo = true ORDER BY fecha DESC`

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": "Error al consultar auditorías"})
		return
	}
	defer rows.Close()

	var lista []models.Auditoria
	for rows.Next() {
		var au models.Auditoria
		rows.Scan(&au.IDAuditoria, &au.IDUsuario, &au.Accion, &au.Fecha, &au.IP, &au.NombreTabla, &au.IDRegistro, &au.AgenteUsuario)
		lista = append(lista, au)
	}

	respondJSON(w, 200, lista)
}