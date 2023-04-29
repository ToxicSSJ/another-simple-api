// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/dividir": {
            "get": {
                "description": "Divide dos números",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Dividir",
                "parameters": [
                    {
                        "type": "number",
                        "description": "Número a",
                        "name": "a",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Número b",
                        "name": "b",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "resultado",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/multiplicar": {
            "get": {
                "description": "Multiplica dos números",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Multiplicar",
                "parameters": [
                    {
                        "type": "number",
                        "description": "Número a",
                        "name": "a",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Número b",
                        "name": "b",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "resultado",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/restar": {
            "get": {
                "description": "Resta dos números",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Restar",
                "parameters": [
                    {
                        "type": "number",
                        "description": "Número a",
                        "name": "a",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Número b",
                        "name": "b",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "resultado",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/sumar": {
            "get": {
                "description": "Suma dos números",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Sumar",
                "parameters": [
                    {
                        "type": "number",
                        "description": "Número a",
                        "name": "a",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Número b",
                        "name": "b",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "resultado",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:80",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API de Operaciones Matemáticas",
	Description:      "Esta es una API simple para realizar operaciones matemáticas básicas como suma, resta, multiplicación y división.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
