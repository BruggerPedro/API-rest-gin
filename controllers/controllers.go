package controllers

import (
	"net/http"

	"github.com/BruggerPedro/gin-api-rest/database"
	"github.com/BruggerPedro/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// ExibeTodosOsAlunos godoc
//
//	@Summary		Exibe todos os alunos
//	@Description	Rota para exibir todos os alunos
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Aluno
//	@Failure		400	{object}	httputil.HTTPError
//	@Router			/alunos [get]
func ExibeTodosOsAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos) // 200 = resposta de sucesso
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E aí " + nome + ", tudo beleza?",
	})
}

// CriaNovoAluno godoc
//
//	@Summary		Cadastra novo aluno
//	@Description	Rota para cadastrar aluno
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Param			alunos	body	models.Aluno	true	"Modelo do aluno"
//	@Success		200	{object}	models.Aluno
//	@Failure		400	{object}	httputil.HTTPError
//	@Router			/alunos [post]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorID godoc
//
//	@Summary		Exibe aluno pelo id
//	@Description	Rota para exibir aluno pelo id
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int		true	"ID aluno" Format(int64)
//	@Success		200		{object}	models.Aluno
//	@Failure		404		{object}	httputil.HTTPError
//	@Router			/alunos/{id} [get]
func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado."})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

// DeletaAluno godoc
//
//	@Summary		Deleta aluno
//	@Description	Rota para deletar aluno
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int		true	"ID aluno" Format(int64)
//	@Success		200	{object}	models.Aluno
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/alunos/{id} [delete]
func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado."})
		return
	}
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado."})
}

// EditaAluno godoc
//
//	@Summary		Edita aluno
//	@Description	Rota para editar aluno
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"ID aluno"	Format(int64)
//	@Param			alunos	body	models.Aluno	true	"Modelo do aluno"
//	@Success		200	{object}	models.Aluno
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Router			/alunos/{id} [patch]
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id) // acessa o aluno pelo id

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	} else if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado."})
		return
	}

	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorCPF godoc
//
//	@Summary		Exibe aluno pelo cpf
//	@Description	Rota para exibir aluno pelo cpf
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Param			cpf	path	string	true	"CPF do aluno"
//	@Success		200	{object}	models.Aluno
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Router			/alunos/cpf/{cpf} [get]
func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
	}
	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
