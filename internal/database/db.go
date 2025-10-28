package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB es una variable global que contendrá la conexión activa a la base de datos
var DB *gorm.DB

// ConnectDB inicializa la conexión a PostgreSQL usando GORM
func ConnectDB() {
	// Configuración de conexión — normalmente se obtienen de variables de entorno
	// host := "dpg-d404rmjuibrs73b2bn0g-a"
	// port := "5432"
	// user := "keinermesa"
	// password := "pq_uqwb"
	// dbname := "books_db"

	// Cadena de conexión
	dsn := "postgresql://keinermesa:V5EazhXBUcg2No0MGz0paUHYXWYPh4Fu@dpg-d404rmjuibrs73b2bn0g-a/pq_uqwb"

	//  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	host, user, password, dbname, port)

	// Intentamos abrir la conexión
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Error al conectar con la base de datos: %v", err)
		os.Exit(1)
	}

	// Guardamos la conexión global
	DB = db
	fmt.Println("✅ Conexión a la base de datos exitosa")
}
