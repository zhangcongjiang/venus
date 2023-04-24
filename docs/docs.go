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
        "/user": {
            "post": {
                "description": "创建用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "请示参数data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功",
                        "schema": {
                            "$ref": "#/definitions/models.Result"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/models.Result"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/models.Result"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "这是一个查询用户列表信息接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "分页查询所有用户",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "第几页",
                        "name": "page_index",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "每页展示条数",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "结果",
                        "schema": {
                            "$ref": "#/definitions/models.Result"
                        }
                    }
                }
            }
        },
        "/wines": {
            "get": {
                "description": "这是一个查询白酒列表信息接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "白酒相关接口"
                ],
                "summary": "分页查询所有白酒信息",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "第几页",
                        "name": "page_index",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "每页展示条数",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "最低价格",
                        "name": "min_price",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "最高价格",
                        "name": "max_price",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "结果",
                        "schema": {
                            "$ref": "#/definitions/models.Result"
                        }
                    }
                }
            }
        },
        "/wines/top": {
            "get": {
                "description": "这是一个查询白酒各类排行的接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "白酒相关接口"
                ],
                "summary": "查询白酒top排行相关接口",
                "parameters": [
                    {
                        "enum": [
                            "market_price",
                            "reference_price"
                        ],
                        "type": "string",
                        "description": "第几页",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 5,
                        "description": "每页展示条数",
                        "name": "top_num",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "结果",
                        "schema": {
                            "$ref": "#/definitions/models.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8888",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "我的swagger",
	Description:      "这里写描述信息",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
