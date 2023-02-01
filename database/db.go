package database

import (
	"fmt"
	"log"
	"os"

	"github.com/BruggerPedro/gin-api-rest/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	error := godotenv.Load("local.env")
	if error != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("HOST")
	user_dbpost := os.Getenv("USER_DBPOST")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")

	stringDeConexao := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user_dbpost, password, dbname, port)
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados.")
	}
	DB.AutoMigrate(&models.Aluno{}) // Cria uma tabela baseada nos dados que estão no models
}
