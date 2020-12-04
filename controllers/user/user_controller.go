package user

import (
	"PROYECTOS/go-http-api/config"
	"fmt"
	"go-http-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CrearPersona(c *gin.Context) {

	config.DBConnection()

	var reqBody models.Persona

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request Body",
		})
		return
	}

	res, err := config.DBClient.Exec("INSERT INTO personas(nombre,apellido,celular) VALUES (?,?,?);",
		reqBody.Nombre,
		reqBody.Apellido,
		reqBody.Celular,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error": false,
		"id":    id,
	})
}

func ModificarPersona(c *gin.Context) {

	config.DBConnection()

	var reqBody models.Persona

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request Body",
		})
		return
	}

	res, err := config.DBClient.Exec("UPDATE personas SET nombre = ?, apellido = ?, celular = ? WHERE id_persona = ?;",
		reqBody.Nombre,
		reqBody.Apellido,
		reqBody.Celular,
		id,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	row, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, row)
}

func EliminarPersona(c *gin.Context) {

	config.DBConnection()

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	res, err := config.DBClient.Exec("DELETE FROM personas WHERE id_persona = ?;", id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	row, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, row)
}

func ListarPersonas(c *gin.Context) {

	config.DBConnection()

	var personas []models.Persona

	err := config.DBClient.Select(&personas, "SELECT id_persona,nombre,apellido,celular FROM personas;")

	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, personas)
}

func ObtenerPersona(c *gin.Context) {

	config.DBConnection()

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var persona models.Persona

	err := config.DBClient.Get(&persona, "SELECT id_persona,nombre,apellido,celular FROM personas WHERE id_persona = ?;", id)

	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, persona)
}
