package routes

import (
	"log"
	"os"

	"github.com/BruggerPedro/gin-api-rest/controllers"
	docs "github.com/BruggerPedro/gin-api-rest/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	error := godotenv.Load("local.env")
	if error != nil {
		log.Fatal("Error loading .env file")
	}
	port_url := os.Getenv("PORT_URL")
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.LoadHTMLGlob("templates/*") // Carrega as páginas HTML
	r.Static("/assets", "./assets")
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/alunos/", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + port_url) //Caso queira rodar em outra por que não seja a padrão(:8080), é só colocar como primeiro argumento ":5000"
}
