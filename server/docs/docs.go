// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/motel": {
            "get": {
                "description": "GetMotelsByFilter",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Motel"
                ],
                "summary": "Lấy danh sách bài đăng",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetManyResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "CreateMotel",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Motel"
                ],
                "summary": "Tạo bài đăng",
                "parameters": [
                    {
                        "description": "request information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateMotelPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetOneResponse"
                        }
                    }
                }
            }
        },
        "/motel/{code}": {
            "get": {
                "description": "GetMotelByCode",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Motel"
                ],
                "summary": "Lấy bài đăng theo code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Motel code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetOneResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "CreateMotel",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Motel"
                ],
                "summary": "Cập nhật bài đăng",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Motel code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateMotelPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetOneResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateMotelPayload": {
            "type": "object",
            "properties": {
                "Address": {
                    "type": "string",
                    "example": "Ngõ 2, Phạm Văn Đồng, Cầu Giấy, Hà Nội"
                },
                "Cost": {
                    "type": "number",
                    "example": 3.3
                },
                "Description": {
                    "type": "string",
                    "example": "Nhà rộng rãi, thoáng mát"
                },
                "Image": {
                    "type": "string",
                    "example": "https://file4.batdongsan.com.vn/2020/08/28/20200828202318-987e_wm.jpg"
                },
                "Images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "https://file4.batdongsan.com.vn/2020/08/28/20200828202318-987e_wm.jpg"
                    ]
                },
                "Latitude": {
                    "type": "string",
                    "example": "21.0286755"
                },
                "Longitude": {
                    "type": "string",
                    "example": "105.7558943,13z"
                },
                "Tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "hanoi"
                    ]
                },
                "Title": {
                    "type": "string",
                    "example": "CCMN SỐ 1A NGÕ 49 TRIỀU KHÚC, FULL HẾT ĐỒ THANG MÁY BAN CÔNG, BẾP TÁCH RIÊNG, 26M2 GIÁ TỪ 3,3 TR/TH"
                }
            }
        },
        "model.GetManyResponse": {
            "type": "object",
            "properties": {
                "Data": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                },
                "Error": {
                    "type": "string"
                },
                "Message": {
                    "type": "string"
                }
            }
        },
        "model.GetOneResponse": {
            "type": "object",
            "properties": {
                "Data": {
                    "type": "object"
                },
                "Error": {
                    "type": "string"
                },
                "Message": {
                    "type": "string"
                }
            }
        },
        "model.UpdateMotelPayload": {
            "type": "object",
            "properties": {
                "Address": {
                    "type": "string",
                    "example": "Ngõ 2, Phạm Văn Đồng, Cầu Giấy, Hà Nội"
                },
                "Cost": {
                    "type": "number",
                    "example": 3.3
                },
                "Description": {
                    "type": "string",
                    "example": "Nhà rộng rãi, thoáng mát"
                },
                "FinishedAt": {
                    "type": "string",
                    "example": "2021-01-15T22:20:28.227+00:00"
                },
                "Image": {
                    "type": "string",
                    "example": "https://file4.batdongsan.com.vn/2020/08/28/20200828202318-987e_wm.jpg"
                },
                "Images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "https://file4.batdongsan.com.vn/2020/08/28/20200828202318-987e_wm.jpg"
                    ]
                },
                "Latitude": {
                    "type": "string",
                    "example": "21.0286755"
                },
                "Longitude": {
                    "type": "string",
                    "example": "105.7558943,13z"
                },
                "Status": {
                    "type": "boolean",
                    "example": true
                },
                "Tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "hanoi"
                    ]
                },
                "Title": {
                    "type": "string",
                    "example": "CCMN SỐ 1A NGÕ 49 TRIỀU KHÚC, FULL HẾT ĐỒ THANG MÁY BAN CÔNG, BẾP TÁCH RIÊNG, 26M2 GIÁ TỪ 3,3 TR/TH"
                }
            }
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
	Host:        "petstore.swagger.io",
	BasePath:    "/v2",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server Petstore server.",
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
