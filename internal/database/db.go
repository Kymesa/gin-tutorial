package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := ""

	if os.Getenv("DEV") == "TRUE" {
		dsn = os.Getenv("DB_EXTERNAL")
	} else {
		dsn = os.Getenv("DB_INTERNAL")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Error al conectar con la base de datos: %v", err)
		os.Exit(1)
	}

	DB = db
	fmt.Println("✅ Conexión a la base de datos exitosa")
}
