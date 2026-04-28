package config

import (
    "database/sql"
    "fmt"
	"log"
	"github.com/lib/pq"
)


var DB *sql.DB


func ConnectDB(){
    // AJUSTA ESTOS VALORES AL NUEVO COMPUTADOR
    host     := "localhost"      
    port     := 5432
    user     := "postgres"   
    password := "postgres"    
	dbname   := "GESTOCK_API"
    schema   := "Autenticacion"          

    psqlInfo := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
        host, port, user, password, dbname, schema,
    )

    db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("error al conectar a la base de datos: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("no se pudo conectar a la base de datos: ", err)
	}

	fmt.Println("Conexión a la base de datos establecida")
	fmt.Println("Base de datos: ", dbname, "y el esquema: ", schema)

	DB=db
}