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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetUsersByFilter",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Lấy danh sách Tài khoản",
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
                "description": "CreateUser",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Tạo tài khoản",
                "parameters": [
                    {
                        "description": "request information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateUserPayload"
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
        "/user/add-favourite-motel": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "AddMotelFavourites",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Thêm bài đăng yêu thích",
                "parameters": [
                    {
                        "description": "request information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddFavouritePayload"
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
        "/user/change-pass": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "ChangePass",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Đổi mật khẩu tài khoản",
                "parameters": [
                    {
                        "description": "request information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ChangePassPayload"
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
        "/user/info": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "UpdateUser",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Cập nhật Thông tin tài khoản",
                "parameters": [
                    {
                        "description": "request information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUserInfoPayload"
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
        "/user/login": {
            "post": {
                "description": "LoginUser",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Đăng nhập tài khoản",
                "parameters": [
                    {
                        "description": "request information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginPayload"
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
        "/user/remove-favourite-motel": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "RemoveMotelFavourites",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Bỏ bài đăng yêu thích",
                "parameters": [
                    {
                        "description": "request information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddFavouritePayload"
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
        "/user/{code}": {
            "get": {
                "description": "GetUserByCode",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Lấy Thông tin User theo code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User code",
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
            }
        }
    },
    "definitions": {
        "model.AddFavouritePayload": {
            "type": "object",
            "properties": {
                "MotelCode": {
                    "type": "string"
                }
            }
        },
        "model.ChangePassPayload": {
            "type": "object",
            "properties": {
                "NewPass": {
                    "type": "string",
                    "example": "abc@124"
                },
                "OldPass": {
                    "type": "string",
                    "example": "abc@123"
                }
            }
        },
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
        "model.CreateUserPayload": {
            "type": "object",
            "properties": {
                "FullName": {
                    "type": "string",
                    "example": "Kiều Chí Công"
                },
                "PassWord": {
                    "type": "string",
                    "example": "abc123"
                },
                "RoleCode": {
                    "type": "string",
                    "example": "ADMIN"
                },
                "UserName": {
                    "type": "string",
                    "example": "kieuchicong"
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
        "model.LoginPayload": {
            "type": "object",
            "properties": {
                "PassWord": {
                    "type": "string",
                    "example": "abc123"
                },
                "UserName": {
                    "type": "string",
                    "example": "kieuchicong"
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
        },
        "model.UpdateUserInfoPayload": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string",
                    "example": "example@gmail.com"
                },
                "FullName": {
                    "type": "string",
                    "example": "Kiều Chí Công"
                },
                "Phone": {
                    "type": "string",
                    "example": "0795048768"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
