// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/signUp": {
            "post": {
                "description": "sign in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignUp",
                "parameters": [
                    {
                        "description": "new jwt data",
                        "name": "UserForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/forms.Result"
                        }
                    },
                    "460": {
                        "description": "Data is invalid",
                        "schema": {
                            "$ref": "#/definitions/forms.Result"
                        }
                    },
                    "521": {
                        "description": "Cannot create session",
                        "schema": {
                            "$ref": "#/definitions/forms.Result"
                        }
                    }
                }
            }
        },
        "/api/signin": {
            "post": {
                "description": "sign in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignIn",
                "parameters": [
                    {
                        "description": "new jwt data",
                        "name": "UserForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/forms.Result"
                        }
                    },
                    "460": {
                        "description": "Data is invalid",
                        "schema": {
                            "$ref": "#/definitions/forms.Result"
                        }
                    },
                    "521": {
                        "description": "Cannot create session",
                        "schema": {
                            "$ref": "#/definitions/forms.Result"
                        }
                    }
                }
            }
        },
        "/api/signout": {
            "post": {
                "description": "sign in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignOut",
                "parameters": [
                    {
                        "description": "new jwt data",
                        "name": "UserForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/forms.Result"
                        }
                    },
                    "460": {
                        "description": "Data is invalid",
                        "schema": {
                            "$ref": "#/definitions/forms.Result"
                        }
                    },
                    "521": {
                        "description": "Cannot create session",
                        "schema": {
                            "$ref": "#/definitions/forms.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "forms.Result": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "some status"
                }
            }
        },
        "forms.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "hello@example.com"
                },
                "id": {
                    "type": "string",
                    "example": "5125112"
                },
                "name": {
                    "type": "string",
                    "example": "Martin"
                },
                "password": {
                    "type": "string",
                    "example": "1fsgh2rfafas"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
