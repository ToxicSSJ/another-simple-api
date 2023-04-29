package main

import (
	"net/http"
	"strconv"

	_ "eafit.edu.co/asapi/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Operaciones Matemáticas
// @version 1.0
// @description Esta es una API simple para realizar operaciones matemáticas básicas como suma, resta, multiplicación y división.

// @host localhost:80
// @BasePath /
func main() {
	// Creamos un router y lo configuramos
	router := crearRouter()

	// Iniciamos el servidor en el puerto 80
	router.Run(":80")
}

// Función para crear un router y configurar las rutas
func crearRouter() *gin.Engine {
	// Creamos una nueva instancia de Gin
	router := gin.Default()

	// Definimos los endpoints para cada operación
	router.GET("/sumar", sumar)
	router.GET("/restar", restar)
	router.GET("/multiplicar", multiplicar)
	router.GET("/dividir", dividir)

	// Establecemos la ruta para Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

// @Summary Sumar
// @Description Suma dos números
// @Accept  json
// @Produce  json
// @Param a query float64 true "Número a"
// @Param b query float64 true "Número b"
// @Success 200 {object} map[string]interface{} "resultado"
// @Failure 400 {object} map[string]interface{} "error"
// @Router /sumar [get]
func sumar(c *gin.Context) {
	// Obtenemos los parámetros 'a' y 'b' de la URL
	a, errA := strconv.ParseFloat(c.Query("a"), 64)
	b, errB := strconv.ParseFloat(c.Query("b"), 64)

	if errA != nil || errB != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	// Realizamos la suma
	resultado := a + b

	// Retornamos el resultado en formato JSON
	c.JSON(http.StatusOK, gin.H{"resultado": resultado})
}

// @Summary Restar
// @Description Resta dos números
// @Accept  json
// @Produce  json
// @Param a query float64 true "Número a"
// @Param b query float64 true "Número b"
// @Success 200 {object} map[string]interface{} "resultado"
// @Failure 400 {object} map[string]interface{} "error"
// @Router /restar [get]
func restar(c *gin.Context) {
	a, errA := strconv.ParseFloat(c.Query("a"), 64)
	b, errB := strconv.ParseFloat(c.Query("b"), 64)

	if errA != nil || errB != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	resultado := a - b
	c.JSON(http.StatusOK, gin.H{"resultado": resultado})
}

// @Summary Multiplicar
// @Description Multiplica dos números
// @Accept  json
// @Produce  json
// @Param a query float64 true "Número a"
// @Param b query float64 true "Número b"
// @Success 200 {object} map[string]interface{} "resultado"
// @Failure 400 {object} map[string]interface{} "error"
// @Router /multiplicar [get]
func multiplicar(c *gin.Context) {
	a, errA := strconv.ParseFloat(c.Query("a"), 64)
	b, errB := strconv.ParseFloat(c.Query("b"), 64)

	if errA != nil || errB != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	resultado := a * b
	c.JSON(http.StatusOK, gin.H{"resultado": resultado})
}

// @Summary Dividir
// @Description Divide dos números
// @Accept  json
// @Produce  json
// @Param a query float64 true "Número a"
// @Param b query float64 true "Número b"
// @Success 200 {object} map[string]interface{} "resultado"
// @Failure 400 {object} map[string]interface{} "error"
// @Router /dividir [get]
func dividir(c *gin.Context) {
	a, errA := strconv.ParseFloat(c.Query("a"), 64)
	b, errB := strconv.ParseFloat(c.Query("b"), 64)

	if errA != nil || errB != nil || b == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	resultado := a / b
	c.JSON(http.StatusOK, gin.H{"resultado": resultado})
}
