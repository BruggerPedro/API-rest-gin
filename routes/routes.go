package routes

import (
	"github.com/BruggerPedro/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	// r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("alunos/:id", controllers.EditaAluno)
	r.GET("alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run() //Caso queira rodar em outra por que não seja a padrão(:8080), é só colocar como primeiro argumento ":5000"
}
