package controllers

import (
	"encoding/json"
	"net/http"
	"fmt" 
	"Autenticacion/config"
	"Autenticacion/models"
)

func RegistrarAuditoria(a models.Auditoria) error {
	query := `INSERT INTO autenticacion.auditoria 
    (id_usuario, accion, ip, nombre_tabla, id_registro, valores_anteriores, valores_nuevos, agente_usuario, activo) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	valAnt, _ := json.Marshal(a.ValoresAnteriores)
	valNue, _ := json.Marshal(a.ValoresNuevos)

	_, err := config.DB.Exec(query, a.IDUsuario, a.Accion, a.IP, a.NombreTabla, a.IDRegistro, valAnt, valNue, a.AgenteUsuario, true)
	return err
}

func GetAuditorias(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_auditoria, id_usuario, accion, fecha, ip, nombre_tabla, id_registro, agente_usuario 
              FROM autenticacion.auditoria 
              WHERE activo = true 
              ORDER BY fecha DESC`

	rows, err := config.DB.Query(query)
	if err != nil {
		fmt.Println("Error en Query:", err) // Esto sale en tu terminal de VS Code
		respondJSON(w, 500, map[string]string{"error": "Error al consultar auditorías"})
		return
	}
	defer rows.Close()

	lista := []models.Auditoria{}

	for rows.Next() {
		var au models.Auditoria
		err := rows.Scan(
			&au.IDAuditoria, 
			&au.IDUsuario, 
			&au.Accion, 
			&au.Fecha, 
			&au.IP, 
			&au.NombreTabla, 
			&au.IDRegistro, 
			&au.AgenteUsuario,
		)
		
		if err != nil {
			fmt.Println("Error en Scan:", err)
			continue
		}
		lista = append(lista, au)
	}

	respondJSON(w, 200, lista)
}

// CrearRegistroAuditoria es una función interna que puedes llamar desde otros controladores


func CrearRegistroAuditoria(w http.ResponseWriter, r *http.Request) {
    var au models.Auditoria
    if err := json.NewDecoder(r.Body).Decode(&au); err != nil {
        respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
        return
    }

    err := RegistrarAuditoria(au)
    if err != nil {
        respondJSON(w, 500, map[string]string{"error": "No se pudo guardar"})
        return
    }

    respondJSON(w, 201, map[string]string{"message": "Registrado con éxito"})
}