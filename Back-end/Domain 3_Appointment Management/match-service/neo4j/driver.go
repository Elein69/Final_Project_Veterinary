package neo4j

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var Driver neo4j.DriverWithContext

func InitDriver() {
	// 📌 Obtener ruta absoluta del .env
	wd, _ := os.Getwd()
	envPath := filepath.Join(wd, "..", ".env")

	// 📥 Cargar archivo .env desde la ruta absoluta
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("❌ Error loading .env file from %s", envPath)
	} else {
		log.Println("✅ .env file loaded")
	}

	// 🔑 Obtener valores de entorno
	uri := os.Getenv("NEO4J_URI")
	user := os.Getenv("NEO4J_USER")
	password := os.Getenv("NEO4J_PASSWORD")

	// 🔌 Crear el driver de conexión
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(user, password, ""))
	if err != nil {
		log.Fatalf("❌ Could not create Neo4j driver: %v", err)
	}

	// ✅ Verificar conexión
	ctx := context.Background()
	if err = driver.VerifyConnectivity(ctx); err != nil {
		log.Fatalf("❌ Neo4j connection failed: %v", err)
	}

	Driver = driver
	fmt.Println("✅ Connected to Neo4j")
}
