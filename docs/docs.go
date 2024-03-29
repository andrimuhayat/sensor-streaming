// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/streaming/sensor-generate": {
            "post": {
                "description": "stream data sensor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Generate stream data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "example: 2",
                        "name": "frequency",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Stream data request",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SensorDataGenerateRequest"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "dto.SensorDataGenerateRequest": {
            "type": "object",
            "properties": {
                "ID1": {
                    "type": "string"
                },
                "ID2": {
                    "type": "integer"
                },
                "sensor_type": {
                    "type": "string"
                },
                "sensor_value": {
                    "type": "number"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
