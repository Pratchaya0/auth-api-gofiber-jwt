// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Pratchaya Saenphoi",
            "email": "pratchayaworks@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/current-user": {
            "get": {
                "security": [
                    {
                        "bearerToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "แสดงข้อมูลของผู้ใช้",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "ลงชื่อเข้าใช้งานระบบ",
                "parameters": [
                    {
                        "description": "LoginRequest",
                        "name": "loginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "ออกจากการใช้งานระบบ",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "สมัครเข้าใช้งานระบบ",
                "parameters": [
                    {
                        "description": "RegisterRequest",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/users/list": {
            "post": {
                "security": [
                    {
                        "bearerToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "แสดงข้อมูลผู้ใช้งานทั้งหมดในระบบ",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "requests.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "requests.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "responses.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Fiber Example API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
