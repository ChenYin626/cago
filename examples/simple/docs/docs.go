// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2022-11-24 00:38:33.618424 +0800 CST m=+0.012050501
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/login": {
            "post": {
                "description": "登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "object",
                        "name": "body",
                        "in": "body"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/BadRequest"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "delete": {
                "description": "登出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "登出",
                "parameters": [
                    {
                        "type": "object",
                        "name": "body",
                        "in": "body"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.LogoutResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/BadRequest"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "type": "object",
                        "name": "body",
                        "in": "body"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/BadRequest"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "BadRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "错误码",
                    "type": "integer",
                    "format": "int32"
                },
                "msg": {
                    "description": "错误信息",
                    "type": "string"
                }
            }
        },
        "user.LoginResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "user.LogoutResponse": {
            "type": "object"
        },
        "user.RegisterResponse": {
            "type": "object"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "api文档",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
