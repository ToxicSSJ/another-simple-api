package main

import (
	"net/http"
	"os"
	"strconv"
	"text/template"

	_ "eafit.edu.co/asapi/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Operaciones Matemáticas
// @version 1.0
// @description Esta es una API simple para realizar operaciones matemáticas básicas como suma, resta, multiplicación y división.

// @BasePath /
func main() {

	// Lee la variable de entorno "HOST"
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost:80" // Utiliza un valor predeterminado si la variable de entorno no está establecida
	}

	// Carga la plantilla personalizada de Swagger
	tmpl, err := template.ParseFiles("templates/swagger.tmpl")
	if err != nil {
		panic(err)
	}

	// Carga la plantilla de la calculadora
	calculatorTmpl, err := template.ParseFiles("templates/calculator.tmpl")
	if err != nil {
		panic(err)
	}

	// Crea un archivo de documentación personalizado en tiempo de ejecución
	swaggerJSON := struct {
		Host string
	}{
		Host: host,
	}
	err = tmpl.Execute(os.Stdout, swaggerJSON)
	if err != nil {
		panic(err)
	}

	// Creamos un router y lo configuramos
	router := crearRouter(host, calculatorTmpl)

	// Iniciamos el servidor en el puerto 80
	router.Run(":80")
}

// Función para crear un router y configurar las rutas
func crearRouter(host string, calculatorTmpl *template.Template) *gin.Engine {
	// Creamos una nueva instancia de Gin
	router := gin.Default()

	// Definimos los endpoints para cada operación
	router.GET("/sumar", sumar)
	router.GET("/restar", restar)
	router.GET("/multiplicar", multiplicar)
	router.GET("/dividir", dividir)

	// Endpoint para servir la calculadora
	router.GET("/calculadoratest", func(c *gin.Context) {
		calculatorTmpl.Execute(c.Writer, struct {
			Host string
		}{
			Host: host,
		})
	})

	// Con esta línea:
	router.GET("/swagger/*any", ginSwagger.CustomWrapHandler(&ginSwagger.Config{
		URL: "http://" + host + "/swagger/doc.json",
	}, swaggerFiles.Handler))

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
