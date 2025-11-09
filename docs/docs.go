package docs

import (
	"strings"

	"github.com/swaggo/swag"
)

const docTemplate = `{
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "schemes": {{.Schemes}},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "paths": {
        "/ping": {
            "get": {
                "tags": [
                    "System"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/organization": {
            "get": {
                "tags": [
                    "Organization"
                ],
                "summary": "List organizations",
                "responses": {
                    "200": {
                        "description": "List of organizations",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Organization"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Organization"
                ],
                "summary": "Create organization",
                "parameters": [
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateOrganizationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created organization",
                        "schema": {
                            "$ref": "#/definitions/Organization"
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/organization/{id}": {
            "get": {
                "tags": [
                    "Organization"
                ],
                "summary": "Get organization by ID",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Organization",
                        "schema": {
                            "$ref": "#/definitions/Organization"
                        }
                    },
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Organization not found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "Organization"
                ],
                "summary": "Update organization",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpdateOrganizationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated organization",
                        "schema": {
                            "$ref": "#/definitions/Organization"
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Organization"
                ],
                "summary": "Delete organization",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Deleted"
                    },
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/organization/search/by-inn": {
            "get": {
                "tags": [
                    "Organization"
                ],
                "summary": "Find organization by INN",
                "parameters": [
                    {
                        "name": "inn",
                        "in": "query",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Organization",
                        "schema": {
                            "$ref": "#/definitions/Organization"
                        }
                    },
                    "400": {
                        "description": "Invalid query parameter",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Organization not found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/museum": {
            "get": {
                "tags": [
                    "Museum"
                ],
                "summary": "List museums",
                "responses": {
                    "200": {
                        "description": "List of museums",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Museum"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Museum"
                ],
                "summary": "Create museum",
                "parameters": [
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateMuseumRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created museum",
                        "schema": {
                            "$ref": "#/definitions/Museum"
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/museum/{id}": {
            "get": {
                "tags": [
                    "Museum"
                ],
                "summary": "Get museum by ID",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Museum",
                        "schema": {
                            "$ref": "#/definitions/Museum"
                        }
                    },
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Museum not found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "Museum"
                ],
                "summary": "Update museum",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpdateMuseumRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated museum",
                        "schema": {
                            "$ref": "#/definitions/Museum"
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Museum"
                ],
                "summary": "Delete museum",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Deleted"
                    },
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/museum/search/by-inn": {
            "get": {
                "tags": [
                    "Museum"
                ],
                "summary": "Find museum by INN",
                "parameters": [
                    {
                        "name": "inn",
                        "in": "query",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Museum",
                        "schema": {
                            "$ref": "#/definitions/Museum"
                        }
                    },
                    "400": {
                        "description": "Invalid query parameter",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Museum not found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/museum/owner/{owner_id}": {
            "get": {
                "tags": [
                    "Museum"
                ],
                "summary": "List museums by owner",
                "parameters": [
                    {
                        "name": "owner_id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of museums",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Museum"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid owner ID",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/activity": {
            "get": {
                "tags": [
                    "Activity"
                ],
                "summary": "List activities",
                "responses": {
                    "200": {
                        "description": "List of activities",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Activity"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Activity"
                ],
                "summary": "Create activity record",
                "parameters": [
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateActivityRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "Activity"
                ],
                "summary": "Update activity record",
                "parameters": [
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpdateActivityRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated"
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Activity"
                ],
                "summary": "Delete activity record",
                "parameters": [
                    {
                        "name": "inn",
                        "in": "query",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "activity_type_id",
                        "in": "query",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    },
                    {
                        "name": "visitor_category",
                        "in": "query",
                        "required": true,
                        "type": "string",
                        "enum": [
                            "internal",
                            "external",
                            "total"
                        ]
                    },
                    {
                        "name": "year",
                        "in": "query",
                        "required": true,
                        "type": "integer",
                        "format": "int32"
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Deleted"
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/activity/search/by-inn": {
            "get": {
                "tags": [
                    "Activity"
                ],
                "summary": "List activities by INN",
                "parameters": [
                    {
                        "name": "inn",
                        "in": "query",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of activities",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Activity"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid query parameter",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/reporting-form": {
            "get": {
                "tags": [
                    "ReportingForm"
                ],
                "summary": "List reporting forms",
                "responses": {
                    "200": {
                        "description": "List of reporting forms",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ReportingForm"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "ReportingForm"
                ],
                "summary": "Create reporting form",
                "parameters": [
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ReportingFormRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created reporting form",
                        "schema": {
                            "$ref": "#/definitions/ReportingForm"
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/reporting-form/{id}": {
            "get": {
                "tags": [
                    "ReportingForm"
                ],
                "summary": "Get reporting form by ID",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Reporting form",
                        "schema": {
                            "$ref": "#/definitions/ReportingForm"
                        }
                    },
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Reporting form not found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "ReportingForm"
                ],
                "summary": "Update reporting form",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ReportingFormUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated reporting form",
                        "schema": {
                            "$ref": "#/definitions/ReportingForm"
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "ReportingForm"
                ],
                "summary": "Delete reporting form",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Deleted"
                    },
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/reporting-form/organization/{org_id}/year/{year}": {
            "get": {
                "tags": [
                    "ReportingForm"
                ],
                "summary": "Get reporting form by organization and year",
                "parameters": [
                    {
                        "name": "org_id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    },
                    {
                        "name": "year",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int32"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Reporting form",
                        "schema": {
                            "$ref": "#/definitions/ReportingForm"
                        }
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Reporting form not found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/reporting-form/organization/{org_id}": {
            "get": {
                "tags": [
                    "ReportingForm"
                ],
                "summary": "List reporting forms by organization",
                "parameters": [
                    {
                        "name": "org_id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of reporting forms",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ReportingForm"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid organization ID",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/reporting-form/year/{year}": {
            "get": {
                "tags": [
                    "ReportingForm"
                ],
                "summary": "List reporting forms by year",
                "parameters": [
                    {
                        "name": "year",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int32"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of reporting forms",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ReportingForm"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid year",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Activity": {
            "type": "object",
            "properties": {
                "INN": {
                    "type": "string"
                },
                "ActivityTypeID": {
                    "type": "integer",
                    "format": "int64"
                },
                "VisitorCategory": {
                    "type": "string",
                    "enum": [
                        "internal",
                        "external",
                        "total"
                    ]
                },
                "CostSharePercent": {
                    "type": "number",
                    "format": "float"
                },
                "RevenueAmount": {
                    "type": "number",
                    "format": "float"
                },
                "TotalCount": {
                    "type": "integer",
                    "format": "int64"
                },
                "StateTaskCount": {
                    "type": "integer",
                    "format": "int64"
                },
                "RevenueActivityCount": {
                    "type": "integer",
                    "format": "int64"
                },
                "Year": {
                    "type": "integer",
                    "format": "int32"
                }
            }
        },
        "CreateActivityRequest": {
            "type": "object",
            "required": [
                "inn",
                "activity_type_id",
                "visitor_category",
                "year"
            ],
            "properties": {
                "inn": {
                    "type": "string"
                },
                "activity_type_id": {
                    "type": "integer",
                    "format": "int64"
                },
                "visitor_category": {
                    "type": "string",
                    "enum": [
                        "internal",
                        "external",
                        "total"
                    ]
                },
                "cost_share_percent": {
                    "type": "number",
                    "format": "float"
                },
                "revenue_amount": {
                    "type": "number",
                    "format": "float"
                },
                "total_count": {
                    "type": "integer",
                    "format": "int64"
                },
                "state_task_count": {
                    "type": "integer",
                    "format": "int64"
                },
                "revenue_activity_count": {
                    "type": "integer",
                    "format": "int64"
                },
                "year": {
                    "type": "integer",
                    "format": "int32"
                }
            },
            "description": "Fields mirror Activity entity; optional counters may be omitted."
        },
        "UpdateActivityRequest": {
            "type": "object",
            "required": [
                "inn",
                "activity_type_id",
                "visitor_category",
                "year"
            ],
            "properties": {
                "inn": {
                    "type": "string"
                },
                "activity_type_id": {
                    "type": "integer",
                    "format": "int64"
                },
                "visitor_category": {
                    "type": "string",
                    "enum": [
                        "internal",
                        "external",
                        "total"
                    ]
                },
                "cost_share_percent": {
                    "type": "number",
                    "format": "float"
                },
                "revenue_amount": {
                    "type": "number",
                    "format": "float"
                },
                "total_count": {
                    "type": "integer",
                    "format": "int64"
                },
                "state_task_count": {
                    "type": "integer",
                    "format": "int64"
                },
                "revenue_activity_count": {
                    "type": "integer",
                    "format": "int64"
                },
                "year": {
                    "type": "integer",
                    "format": "int32"
                }
            }
        },
        "CreateMuseumRequest": {
            "type": "object",
            "required": [
                "id_owner",
                "inn",
                "museum_activity_in_charter",
                "name"
            ],
            "properties": {
                "id_owner": {
                    "type": "integer",
                    "format": "int64"
                },
                "inn": {
                    "type": "string"
                },
                "kpp": {
                    "type": "string"
                },
                "founder": {
                    "type": "string"
                },
                "museum_activity_in_charter": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "museum_legal_status": {
                    "type": "string"
                },
                "is_memorial_reserve_museum": {
                    "type": "boolean"
                },
                "is_historical_memorial_reserve": {
                    "type": "boolean"
                },
                "is_art_museum": {
                    "type": "boolean"
                },
                "is_museum_reserve": {
                    "type": "boolean"
                },
                "is_estate_museum": {
                    "type": "boolean"
                },
                "is_palace_park_ensemble": {
                    "type": "boolean"
                },
                "is_historical_architectural_reserve": {
                    "type": "boolean"
                },
                "annual_visitor_capacity": {
                    "type": "integer",
                    "format": "int64"
                },
                "internal_visitors_count": {
                    "type": "integer",
                    "format": "int64"
                },
                "external_visitors_count": {
                    "type": "integer",
                    "format": "int64"
                },
                "is_valuable_cultural_heritage": {
                    "type": "boolean"
                },
                "valuable_museum_items_count": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "UpdateMuseumRequest": {
            "type": "object",
            "required": [
                "id",
                "id_owner",
                "inn",
                "museum_activity_in_charter",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "format": "int64"
                },
                "id_owner": {
                    "type": "integer",
                    "format": "int64"
                },
                "inn": {
                    "type": "string"
                },
                "kpp": {
                    "type": "string"
                },
                "founder": {
                    "type": "string"
                },
                "museum_activity_in_charter": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "museum_legal_status": {
                    "type": "string"
                },
                "is_memorial_reserve_museum": {
                    "type": "boolean"
                },
                "is_historical_memorial_reserve": {
                    "type": "boolean"
                },
                "is_art_museum": {
                    "type": "boolean"
                },
                "is_museum_reserve": {
                    "type": "boolean"
                },
                "is_estate_museum": {
                    "type": "boolean"
                },
                "is_palace_park_ensemble": {
                    "type": "boolean"
                },
                "is_historical_architectural_reserve": {
                    "type": "boolean"
                },
                "annual_visitor_capacity": {
                    "type": "integer",
                    "format": "int64"
                },
                "internal_visitors_count": {
                    "type": "integer",
                    "format": "int64"
                },
                "external_visitors_count": {
                    "type": "integer",
                    "format": "int64"
                },
                "is_valuable_cultural_heritage": {
                    "type": "boolean"
                },
                "valuable_museum_items_count": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "CreateOrganizationRequest": {
            "type": "object",
            "required": [
                "inn",
                "name",
                "exist_museum"
            ],
            "properties": {
                "inn": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "exist_museum": {
                    "type": "boolean"
                }
            }
        },
        "UpdateOrganizationRequest": {
            "type": "object",
            "required": [
                "inn",
                "name",
                "exist_museum"
            ],
            "properties": {
                "inn": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "exist_museum": {
                    "type": "boolean"
                }
            }
        },
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "Museum": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer",
                    "format": "int64"
                },
                "IdOwner": {
                    "type": "integer",
                    "format": "int64"
                },
                "INN": {
                    "type": "string"
                },
                "KPP": {
                    "type": "string"
                },
                "Founder": {
                    "type": "string"
                },
                "MuseumActivityInCharter": {
                    "type": "boolean"
                },
                "Name": {
                    "type": "string"
                },
                "MuseumLegalStatus": {
                    "type": "string"
                },
                "IsMemorialReserveMuseum": {
                    "type": "boolean"
                },
                "IsHistoricalMemorialReserve": {
                    "type": "boolean"
                },
                "IsArtMuseum": {
                    "type": "boolean"
                },
                "IsMuseumReserve": {
                    "type": "boolean"
                },
                "IsEstateMuseum": {
                    "type": "boolean"
                },
                "IsPalaceParkEnsemble": {
                    "type": "boolean"
                },
                "IsHistoricalArchitecturalReserve": {
                    "type": "boolean"
                },
                "AnnualVisitorCapacity": {
                    "type": "integer",
                    "format": "int64"
                },
                "InternalVisitorsCount": {
                    "type": "integer",
                    "format": "int64"
                },
                "ExternalVisitorsCount": {
                    "type": "integer",
                    "format": "int64"
                },
                "IsValuableCulturalHeritage": {
                    "type": "boolean"
                },
                "ValuableMuseumItemsCount": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "Organization": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer",
                    "format": "int64"
                },
                "INN": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "ExistMuseum": {
                    "type": "boolean"
                }
            }
        },
        "ReportingForm": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer",
                    "format": "int64"
                },
                "OrganizationID": {
                    "type": "integer",
                    "format": "int64"
                },
                "Year": {
                    "type": "integer",
                    "format": "int32"
                },
                "Status": {
                    "type": "string",
                    "enum": [
                        "draft",
                        "submitted",
                        "approved"
                    ]
                },
                "GeneralInfo": {
                    "$ref": "#/definitions/ReportingFormGeneralInfo"
                },
                "MuseumData": {
                    "$ref": "#/definitions/ReportingFormMuseumDetails"
                },
                "Labor": {
                    "$ref": "#/definitions/ReportingFormLabor"
                },
                "Finances": {
                    "$ref": "#/definitions/ReportingFormFinancials"
                }
            }
        },
        "ReportingFormRequest": {
            "type": "object",
            "required": [
                "organization_id",
                "year",
                "status",
                "general_info",
                "museum_data",
                "labor",
                "finances"
            ],
            "properties": {
                "organization_id": {
                    "type": "integer",
                    "format": "int64"
                },
                "year": {
                    "type": "integer",
                    "format": "int32"
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "draft",
                        "submitted",
                        "approved"
                    ]
                },
                "general_info": {
                    "$ref": "#/definitions/ReportingFormGeneralInfoRequest"
                },
                "museum_data": {
                    "$ref": "#/definitions/ReportingFormMuseumDetailsRequest"
                },
                "labor": {
                    "$ref": "#/definitions/ReportingFormLaborRequest"
                },
                "finances": {
                    "$ref": "#/definitions/ReportingFormFinancialsRequest"
                }
            }
        },
        "ReportingFormUpdateRequest": {
            "allOf": [
                {
                    "type": "object",
                    "properties": {
                        "id": {
                            "type": "integer",
                            "format": "int64"
                        }
                    },
                    "required": [
                        "id"
                    ]
                },
                {
                    "$ref": "#/definitions/ReportingFormRequest"
                }
            ]
        },
        "ReportingFormFinancials": {
            "type": "object",
            "properties": {
                "TotalRevenue": {
                    "type": "number",
                    "format": "float"
                },
                "StateAssignmentSubsidy": {
                    "type": "number",
                    "format": "float"
                },
                "EarnedRevenue": {
                    "type": "number",
                    "format": "float"
                },
                "OtherFundingSources": {
                    "type": "number",
                    "format": "float"
                },
                "TotalExpenses": {
                    "type": "number",
                    "format": "float"
                },
                "InventoryAssetsAcquisition": {
                    "type": "number",
                    "format": "float"
                },
                "UtilityServices": {
                    "type": "number",
                    "format": "float"
                },
                "CommunicationServices": {
                    "type": "number",
                    "format": "float"
                },
                "TransportationServices": {
                    "type": "number",
                    "format": "float"
                },
                "ValuableAssetsAcquisition": {
                    "type": "number",
                    "format": "float"
                },
                "RealEstateMaintenance": {
                    "type": "number",
                    "format": "float"
                },
                "ValuableAssetsMaintenance": {
                    "type": "number",
                    "format": "float"
                },
                "GeneralAdministrativeCosts": {
                    "type": "number",
                    "format": "float"
                },
                "TaxPayments": {
                    "type": "number",
                    "format": "float"
                },
                "OtherExpenses": {
                    "type": "number",
                    "format": "float"
                }
            }
        },
        "ReportingFormFinancialsRequest": {
            "type": "object",
            "properties": {
                "total_revenue": {
                    "type": "number",
                    "format": "float"
                },
                "state_assignment_subsidy": {
                    "type": "number",
                    "format": "float"
                },
                "earned_revenue": {
                    "type": "number",
                    "format": "float"
                },
                "other_funding_sources": {
                    "type": "number",
                    "format": "float"
                },
                "total_expenses": {
                    "type": "number",
                    "format": "float"
                },
                "inventory_assets_acquisition": {
                    "type": "number",
                    "format": "float"
                },
                "utility_services": {
                    "type": "number",
                    "format": "float"
                },
                "communication_services": {
                    "type": "number",
                    "format": "float"
                },
                "transportation_services": {
                    "type": "number",
                    "format": "float"
                },
                "valuable_assets_acquisition": {
                    "type": "number",
                    "format": "float"
                },
                "real_estate_maintenance": {
                    "type": "number",
                    "format": "float"
                },
                "valuable_assets_maintenance": {
                    "type": "number",
                    "format": "float"
                },
                "general_administrative_costs": {
                    "type": "number",
                    "format": "float"
                },
                "tax_payments": {
                    "type": "number",
                    "format": "float"
                },
                "other_expenses": {
                    "type": "number",
                    "format": "float"
                }
            }
        },
        "ReportingFormGeneralInfo": {
            "type": "object",
            "properties": {
                "INN": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "ExistMuseum": {
                    "type": "boolean"
                }
            }
        },
        "ReportingFormGeneralInfoRequest": {
            "type": "object",
            "properties": {
                "inn": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "exist_museum": {
                    "type": "boolean"
                }
            }
        },
        "ReportingFormLabor": {
            "type": "object",
            "properties": {
                "TotalStaffAnnual": {
                    "type": "number",
                    "format": "float"
                },
                "ResearchStaffInternal": {
                    "type": "number",
                    "format": "float"
                },
                "CoreOperationalStaffInternal": {
                    "type": "number",
                    "format": "float"
                },
                "AdminSupportStaffInternal": {
                    "type": "number",
                    "format": "float"
                },
                "ResearchStaffExternal": {
                    "type": "number",
                    "format": "float"
                },
                "CoreOperationalStaffExternal": {
                    "type": "number",
                    "format": "float"
                },
                "AdminSupportStaffExternal": {
                    "type": "number",
                    "format": "float"
                },
                "FOT": {
                    "$ref": "#/definitions/ReportingFormLaborFOT"
                }
            }
        },
        "ReportingFormLaborRequest": {
            "type": "object",
            "properties": {
                "total_staff_annual": {
                    "type": "number",
                    "format": "float"
                },
                "research_staff_internal": {
                    "type": "number",
                    "format": "float"
                },
                "core_operational_staff_internal": {
                    "type": "number",
                    "format": "float"
                },
                "admin_support_staff_internal": {
                    "type": "number",
                    "format": "float"
                },
                "research_staff_external": {
                    "type": "number",
                    "format": "float"
                },
                "core_operational_staff_external": {
                    "type": "number",
                    "format": "float"
                },
                "admin_support_staff_external": {
                    "type": "number",
                    "format": "float"
                },
                "fot": {
                    "$ref": "#/definitions/ReportingFormLaborFOTRequest"
                }
            }
        },
        "ReportingFormLaborFOT": {
            "type": "object",
            "properties": {
                "TotalStaffAnnual": {
                    "type": "number",
                    "format": "float"
                },
                "ResearchStaffInternal": {
                    "type": "number",
                    "format": "float"
                },
                "CoreOperationalStaffInternal": {
                    "type": "number",
                    "format": "float"
                },
                "AdminSupportStaffInternal": {
                    "type": "number",
                    "format": "float"
                },
                "ResearchStaffExternal": {
                    "type": "number",
                    "format": "float"
                },
                "CoreOperationalStaffExternal": {
                    "type": "number",
                    "format": "float"
                },
                "AdminSupportStaffExternal": {
                    "type": "number",
                    "format": "float"
                }
            }
        },
        "ReportingFormLaborFOTRequest": {
            "type": "object",
            "properties": {
                "total_staff_annual": {
                    "type": "number",
                    "format": "float"
                },
                "research_staff_internal": {
                    "type": "number",
                    "format": "float"
                },
                "core_operational_staff_internal": {
                    "type": "number",
                    "format": "float"
                },
                "admin_support_staff_internal": {
                    "type": "number",
                    "format": "float"
                },
                "research_staff_external": {
                    "type": "number",
                    "format": "float"
                },
                "core_operational_staff_external": {
                    "type": "number",
                    "format": "float"
                },
                "admin_support_staff_external": {
                    "type": "number",
                    "format": "float"
                }
            }
        },
        "ReportingFormMuseumDetails": {
            "type": "object",
            "properties": {
                "INN": {
                    "type": "string"
                },
                "KPP": {
                    "type": "string"
                },
                "Founder": {
                    "type": "string"
                },
                "MuseumActivityInCharter": {
                    "type": "boolean"
                },
                "Name": {
                    "type": "string"
                },
                "MuseumLegalStatus": {
                    "type": "string"
                },
                "IsMemorialReserveMuseum": {
                    "type": "boolean"
                },
                "IsHistoricalMemorialReserve": {
                    "type": "boolean"
                },
                "IsArtMuseum": {
                    "type": "boolean"
                },
                "IsMuseumReserve": {
                    "type": "boolean"
                },
                "IsEstateMuseum": {
                    "type": "boolean"
                },
                "IsPalaceParkEnsemble": {
                    "type": "boolean"
                },
                "IsHistoricalArchitecturalReserve": {
                    "type": "boolean"
                },
                "AnnualVisitorCapacity": {
                    "type": "integer",
                    "format": "int64"
                },
                "InternalVisitorsCount": {
                    "type": "integer",
                    "format": "int64"
                },
                "ExternalVisitorsCount": {
                    "type": "integer",
                    "format": "int64"
                },
                "IsValuableCulturalHeritage": {
                    "type": "boolean"
                },
                "ValuableMuseumItemsCount": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "ReportingFormMuseumDetailsRequest": {
            "type": "object",
            "properties": {
                "inn": {
                    "type": "string"
                },
                "kpp": {
                    "type": "string"
                },
                "founder": {
                    "type": "string"
                },
                "museum_activity_in_charter": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "museum_legal_status": {
                    "type": "string"
                },
                "is_memorial_reserve_museum": {
                    "type": "boolean"
                },
                "is_historical_memorial_reserve": {
                    "type": "boolean"
                },
                "is_art_museum": {
                    "type": "boolean"
                },
                "is_museum_reserve": {
                    "type": "boolean"
                },
                "is_estate_museum": {
                    "type": "boolean"
                },
                "is_palace_park_ensemble": {
                    "type": "boolean"
                },
                "is_historical_architectural_reserve": {
                    "type": "boolean"
                },
                "annual_visitor_capacity": {
                    "type": "integer",
                    "format": "int64"
                },
                "internal_visitors_count": {
                    "type": "integer",
                    "format": "int64"
                },
                "external_visitors_count": {
                    "type": "integer",
                    "format": "int64"
                },
                "is_valuable_cultural_heritage": {
                    "type": "boolean"
                },
                "valuable_museum_items_count": {
                    "type": "integer",
                    "format": "int64"
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

var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8081",
	BasePath:    "",
	Schemes:     []string{"http"},
	Title:       "VKR API",
	Description: "API documentation for museum reporting service.",
}

type s struct{}

func (s *s) ReadDoc() string {
	doc := docTemplate
	doc = strings.Replace(doc, "{{.Version}}", SwaggerInfo.Version, -1)
	doc = strings.Replace(doc, "{{.Host}}", SwaggerInfo.Host, -1)
	doc = strings.Replace(doc, "{{.BasePath}}", SwaggerInfo.BasePath, -1)
	doc = strings.Replace(doc, "{{.Title}}", SwaggerInfo.Title, -1)
	doc = strings.Replace(doc, "{{.Description}}", SwaggerInfo.Description, -1)
	if len(SwaggerInfo.Schemes) == 0 {
		doc = strings.Replace(doc, "\"schemes\": {{.Schemes}}", "\"schemes\": []", 1)
	} else {
		doc = strings.Replace(doc, "\"schemes\": {{.Schemes}}", "\"schemes\": [\""+strings.Join(SwaggerInfo.Schemes, "\",\"")+"\"]", 1)
	}
	return doc
}

func init() {
	swag.Register(swag.Name, &s{})
}
