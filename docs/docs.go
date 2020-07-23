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
        "contact": {},
        "license": {
            "name": "MIT License"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account/:id/credentials": {
            "put": {
                "description": "Non-authenticated endpoint that adds grafana and confluence-server users to an account. Assumes entries are pre-validated",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Add credentials to an account",
                "parameters": [
                    {
                        "description": "Add credentials",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/credentials.SetCredentialsV1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/credentials.SetCredentialsV1"
                        }
                    }
                }
            }
        },
        "/account/{id}": {
            "get": {
                "description": "Non-authenticated endpoint fetches account at specified key",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Get account record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/record.RecordViewV1"
                        }
                    }
                }
            },
            "put": {
                "description": "Non-authenticated endpoint creates an empty record at the specified key. Overwrites any record that already exists",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Create account record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Create account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/record.AccountViewV1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/credentials/check": {
            "post": {
                "description": "Non-authenticated endpoint Check credentials for validity. Returns an array of user objects with check result",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "credentials"
                ],
                "summary": "Check credentials for validity",
                "parameters": [
                    {
                        "description": "Check credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/credentials.CheckCredentialsV1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/credentials.CheckUsersResultV1"
                        }
                    }
                }
            }
        },
        "/health/hello": {
            "get": {
                "description": "Non-authenticated endpoint that returns 200 with hello message. Used to validate that the service is responsive.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Hello sanity endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/health.Ping"
                        }
                    }
                }
            }
        },
        "/schedule/check": {
            "post": {
                "description": "Non-authenticated endpoint which checks and runs a schedule to validate connectivity and storage behaviour by the end user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "summary": "Check and execute schedule",
                "parameters": [
                    {
                        "description": "Check schedule",
                        "name": "schedule",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schedule.CheckScheduleV1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/report.CheckV1View"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Auth": {
            "type": "object",
            "properties": {
                "basic": {
                    "type": "object",
                    "$ref": "#/definitions/common.Basic"
                },
                "bearerToken": {
                    "type": "object",
                    "$ref": "#/definitions/common.BearerToken"
                }
            }
        },
        "common.Basic": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "common.BearerToken": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "common.ConfluenceServerUserV1": {
            "type": "object",
            "properties": {
                "auth": {
                    "type": "object",
                    "$ref": "#/definitions/common.Auth"
                },
                "description": {
                    "type": "string"
                }
            }
        },
        "common.GrafanaDashBoard": {
            "type": "object",
            "properties": {
                "excludePanelsIDs": {
                    "description": "blank means exclude nothing",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "host": {
                    "type": "string"
                },
                "includePanelsIDs": {
                    "description": "blank means include all panels. Will include newly added panels",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "port": {
                    "type": "integer"
                },
                "uid": {
                    "type": "string"
                },
                "user": {
                    "type": "object",
                    "$ref": "#/definitions/common.GrafanaUserV1"
                }
            }
        },
        "common.GrafanaUserV1": {
            "type": "object",
            "properties": {
                "auth": {
                    "type": "object",
                    "$ref": "#/definitions/common.Auth"
                },
                "description": {
                    "type": "string"
                }
            }
        },
        "credentials.CheckCredentialsV1": {
            "type": "object",
            "properties": {
                "ConfluenceServerUsers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/credentials.CheckUserV1"
                    }
                },
                "GrafanaAPIUsers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/credentials.CheckUserV1"
                    }
                }
            }
        },
        "credentials.CheckUserResultV1": {
            "type": "object",
            "properties": {
                "Cause": {
                    "type": "string"
                },
                "auth": {
                    "type": "object",
                    "$ref": "#/definitions/common.Auth"
                },
                "host": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "result": {
                    "type": "boolean"
                }
            }
        },
        "credentials.CheckUserV1": {
            "type": "object",
            "properties": {
                "auth": {
                    "type": "object",
                    "$ref": "#/definitions/common.Auth"
                },
                "host": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        },
        "credentials.CheckUsersResultV1": {
            "type": "object",
            "properties": {
                "confluenceServerUserCheck": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/credentials.CheckUserResultV1"
                    }
                },
                "grafanaReadUserCheck": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/credentials.CheckUserResultV1"
                    }
                }
            }
        },
        "credentials.SetCredentialsV1": {
            "type": "object",
            "properties": {
                "ConfluenceServerUsers": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/common.ConfluenceServerUserV1"
                    }
                },
                "GrafanaAPIUsers": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/common.GrafanaUserV1"
                    }
                }
            }
        },
        "health.Ping": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "hello"
                }
            }
        },
        "record.AccountViewV1": {
            "type": "object",
            "properties": {
                "Alias": {
                    "description": "Optional arg. Won't be returned if missing.",
                    "type": "string"
                },
                "Email": {
                    "type": "string"
                }
            }
        },
        "record.ConfluenceServerUser": {
            "type": "object",
            "properties": {
                "auth": {
                    "type": "object",
                    "$ref": "#/definitions/common.Auth"
                },
                "description": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        },
        "record.CredentialsView1": {
            "type": "object",
            "properties": {
                "ConfluenceServerUser": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/record.ConfluenceServerUser"
                    }
                },
                "GrafanaAPIUsers": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/record.GrafanaAPIUser"
                    }
                }
            }
        },
        "record.GrafanaAPIUser": {
            "type": "object",
            "properties": {
                "auth": {
                    "type": "object",
                    "$ref": "#/definitions/common.Auth"
                },
                "description": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        },
        "record.MetadataViewV1": {
            "type": "object",
            "properties": {
                "CreateTimeUTC": {
                    "type": "string"
                },
                "LastUpdate": {
                    "type": "string"
                },
                "PrimaryKey": {
                    "type": "string"
                },
                "Version": {
                    "type": "string"
                }
            }
        },
        "record.RecordViewV1": {
            "type": "object",
            "properties": {
                "Account": {
                    "type": "object",
                    "$ref": "#/definitions/record.AccountViewV1"
                },
                "Credentials": {
                    "type": "object",
                    "$ref": "#/definitions/record.CredentialsView1"
                },
                "Metadata": {
                    "type": "object",
                    "$ref": "#/definitions/record.MetadataViewV1"
                }
            }
        },
        "report.CheckV1View": {
            "type": "object"
        },
        "schedule.CheckScheduleV1": {
            "type": "object",
            "properties": {
                "dataStores": {
                    "type": "object",
                    "$ref": "#/definitions/schedule.DataStores"
                },
                "graphDashboards": {
                    "type": "object",
                    "$ref": "#/definitions/schedule.DashBoards"
                }
            }
        },
        "schedule.DashBoards": {
            "type": "object",
            "properties": {
                "grafana": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/common.GrafanaDashBoard"
                    }
                }
            }
        },
        "schedule.DataStores": {
            "type": "object",
            "properties": {
                "confluencePages": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/schedule.ParentConfluencePage"
                    }
                }
            }
        },
        "schedule.ParentConfluencePage": {
            "type": "object",
            "properties": {
                "parentPageID": {
                    "type": "string"
                },
                "spaceKey": {
                    "type": "string"
                },
                "user": {
                    "type": "object",
                    "$ref": "#/definitions/common.ConfluenceServerUserV1"
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
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Graph Snapper API",
	Description: "Takes and updates snapshots from a graph service to a document store",
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
