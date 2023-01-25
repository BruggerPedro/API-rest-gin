package database

import (
	"log"

	"github.com/BruggerPedro/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	stringDeConexao := "host=localhost user=root password=root dbname=postgres port=5436 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados.")
	}
	DB.AutoMigrate(&models.Aluno{}) // Cria uma tabela baseada nos dados que estão no models
}
