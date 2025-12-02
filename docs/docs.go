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
    "consumes": ["application/json"],
    "produces": ["application/json"],
    "paths": {
        "/ping": {
            "get": {
                "tags": ["System"],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {"type": "string"}
                    }
                }
            }
        },
        "/api/v1/organization": {
            "get": {
                "tags": ["Organization"],
                "summary": "List organizations with pagination and filtering",
                "description": "Returns paginated list of organizations. Supports filtering by name (substring search, case-insensitive) and pagination.",
                "parameters": [
                    {
                        "name": "name",
                        "in": "query",
                        "required": false,
                        "type": "string",
                        "description": "Filter organizations by name (partial match, case-insensitive)"
                    },
                    {
                        "name": "page",
                        "in": "query",
                        "required": false,
                        "type": "integer",
                        "format": "int32",
                        "default": 1,
                        "description": "Page number (1-based)"
                    },
                    {
                        "name": "page_size",
                        "in": "query",
                        "required": false,
                        "type": "integer",
                        "format": "int32",
                        "default": 10,
                        "minimum": 1,
                        "maximum": 100,
                        "description": "Number of items per page (max 100)"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Paginated list of organizations",
                        "schema": {"$ref": "#/definitions/PaginatedOrganizationResponse"}
                    },
                    "400": {
                        "description": "Invalid pagination parameters",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            },
            "post": {
                "tags": ["Organization"],
                "summary": "Create organization",
                "parameters": [
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {"$ref": "#/definitions/CreateOrganizationRequest"}
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created organization",
                        "schema": {"$ref": "#/definitions/Organization"}
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/organization/{id}": {
            "get": {
                "tags": ["Organization"],
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
                        "schema": {"$ref": "#/definitions/Organization"}
                    },
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "404": {
                        "description": "Organization not found",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            },
            "put": {
                "tags": ["Organization"],
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
                        "schema": {"$ref": "#/definitions/UpdateOrganizationRequest"}
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated organization",
                        "schema": {"$ref": "#/definitions/Organization"}
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            },
            "delete": {
                "tags": ["Organization"],
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
                    "204": {"description": "Deleted"},
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/organization/search/by-inn": {
            "get": {
                "tags": ["Organization"],
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
                        "schema": {"$ref": "#/definitions/Organization"}
                    },
                    "400": {
                        "description": "Invalid query parameter",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "404": {
                        "description": "Organization not found",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/museum": {
            "get": {
                "tags": ["Museum"],
                "summary": "List museums with pagination and filtering",
                "description": "Returns paginated list of museums. Supports filtering by name (substring search) and museum type.",
                "parameters": [
                    {
                        "name": "name",
                        "in": "query",
                        "required": false,
                        "type": "string",
                        "description": "Filter museums by name (partial match, case-insensitive)"
                    },
                    {
                        "name": "museum_type",
                        "in": "query",
                        "required": false,
                        "type": "string",
                        "enum": ["art", "memorial_reserve", "historical_memorial_reserve", "museum_reserve", "estate", "palace_park_ensemble", "historical_architectural_reserve"],
                        "description": "Filter museums by type. Valid values: art, memorial_reserve, historical_memorial_reserve, museum_reserve, estate, palace_park_ensemble, historical_architectural_reserve"
                    },
                    {
                        "name": "page",
                        "in": "query",
                        "required": false,
                        "type": "integer",
                        "format": "int32",
                        "default": 1,
                        "description": "Page number (1-based)"
                    },
                    {
                        "name": "page_size",
                        "in": "query",
                        "required": false,
                        "type": "integer",
                        "format": "int32",
                        "default": 10,
                        "minimum": 1,
                        "maximum": 100,
                        "description": "Number of items per page (max 100)"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Paginated list of museums",
                        "schema": {"$ref": "#/definitions/PaginatedMuseumResponse"}
                    },
                    "400": {
                        "description": "Invalid pagination parameters",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            },
            "post": {
                "tags": ["Museum"],
                "summary": "Create museum",
                "parameters": [
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {"$ref": "#/definitions/CreateMuseumRequest"}
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created museum",
                        "schema": {"$ref": "#/definitions/Museum"}
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/museum/{id}": {
            "get": {
                "tags": ["Museum"],
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
                        "schema": {"$ref": "#/definitions/Museum"}
                    },
                    "404": {
                        "description": "Museum not found",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            },
            "put": {
                "tags": ["Museum"],
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
                        "schema": {"$ref": "#/definitions/UpdateMuseumRequest"}
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated museum",
                        "schema": {"$ref": "#/definitions/Museum"}
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            },
            "delete": {
                "tags": ["Museum"],
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
                    "204": {"description": "Deleted"},
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/museum/search/by-inn": {
            "get": {
                "tags": ["Museum"],
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
                        "schema": {"$ref": "#/definitions/Museum"}
                    },
                    "404": {
                        "description": "Museum not found",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/museum/owner/{owner_id}": {
            "get": {
                "tags": ["Museum"],
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
                            "items": {"$ref": "#/definitions/Museum"}
                        }
                    }
                }
            }
        },
        "/api/v1/activity": {
            "get": {
                "tags": ["Activity"],
                "summary": "List activities with pagination",
                "description": "Returns paginated list of activities.",
                "parameters": [
                    {
                        "name": "page",
                        "in": "query",
                        "required": false,
                        "type": "integer",
                        "format": "int32",
                        "default": 1,
                        "description": "Page number (1-based)"
                    },
                    {
                        "name": "page_size",
                        "in": "query",
                        "required": false,
                        "type": "integer",
                        "format": "int32",
                        "default": 10,
                        "minimum": 1,
                        "maximum": 100,
                        "description": "Number of items per page (max 100)"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Paginated list of activities",
                        "schema": {"$ref": "#/definitions/PaginatedActivityResponse"}
                    },
                    "400": {
                        "description": "Invalid pagination parameters",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            },
            "post": {
                "tags": ["Activity"],
                "summary": "Create activity record",
                "parameters": [
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {"$ref": "#/definitions/CreateActivityRequest"}
                    }
                ],
                "responses": {
                    "201": {"description": "Created"},
                    "400": {
                        "description": "Invalid payload",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            },
            "put": {
                "tags": ["Activity"],
                "summary": "Update activity record",
                "parameters": [
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {"$ref": "#/definitions/UpdateActivityRequest"}
                    }
                ],
                "responses": {
                    "200": {"description": "Updated"},
                    "400": {
                        "description": "Invalid payload",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            },
            "delete": {
                "tags": ["Activity"],
                "summary": "Delete activity record",
                "parameters": [
                    {"name": "id", "in": "query", "required": true, "type": "integer", "format": "int64"}
                ],
                "responses": {
                    "204": {"description": "Deleted"},
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/activity/search/by-inn": {
            "get": {
                "tags": ["Activity"],
                "summary": "List activities by INN",
                "parameters": [
                    {"name": "inn", "in": "query", "required": true, "type": "string"}
                ],
                "responses": {
                    "200": {
                        "description": "List of activities",
                        "schema": {
                            "type": "array",
                            "items": {"$ref": "#/definitions/Activity"}
                        }
                    }
                }
            }
        },
        "/api/v1/activity/museum/{museum_id}": {
            "get": {
                "tags": ["Activity"],
                "summary": "List activities by museum ID",
                "parameters": [
                    {
                        "name": "museum_id",
                        "in": "path",
                        "required": true,
                        "type": "integer",
                        "format": "int64"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of activities",
                        "schema": {
                            "type": "array",
                            "items": {"$ref": "#/definitions/Activity"}
                        }
                    },
                    "400": {
                        "description": "Invalid museum id",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/labor/organization/{org_id}": {
            "get": {
                "tags": ["Labor"],
                "summary": "Get labor data by organization ID",
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
                        "description": "Labor data",
                        "schema": {"$ref": "#/definitions/LaborResponse"}
                    },
                    "400": {
                        "description": "Invalid organization id",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "404": {
                        "description": "Labor data not found",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/labor/search/by-inn": {
            "get": {
                "tags": ["Labor"],
                "summary": "Get labor data by organization INN",
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
                        "description": "Labor data",
                        "schema": {"$ref": "#/definitions/LaborResponse"}
                    },
                    "400": {
                        "description": "Invalid query parameter",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "404": {
                        "description": "Labor data not found",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/form/labor/organization/{org_id}": {
            "get": {
                "tags": ["Reporting Form"],
                "summary": "Get labor form data by organization ID",
                "description": "Returns data for labor resources form including staff counts and payroll (FOT) information.",
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
                        "description": "Labor form data",
                        "schema": {"$ref": "#/definitions/LaborFormData"}
                    },
                    "400": {
                        "description": "Invalid organization id",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/form/labor/search/by-inn": {
            "get": {
                "tags": ["Reporting Form"],
                "summary": "Get labor form data by organization INN",
                "description": "Returns data for labor resources form by organization INN.",
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
                        "description": "Labor form data",
                        "schema": {"$ref": "#/definitions/LaborFormData"}
                    },
                    "400": {
                        "description": "Invalid query parameter",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/form/activities/organization/{org_id}": {
            "get": {
                "tags": ["Reporting Form"],
                "summary": "Get activities form data by organization ID and year",
                "description": "Returns data for activities (volumes) form including all activity types with their volume indicators, costs, revenues, and counts.",
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
                        "in": "query",
                        "required": true,
                        "type": "integer",
                        "format": "int32",
                        "description": "Reporting year"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Activities form data",
                        "schema": {"$ref": "#/definitions/ActivitiesFormData"}
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/form/activities/search/by-inn": {
            "get": {
                "tags": ["Reporting Form"],
                "summary": "Get activities form data by organization INN and year",
                "description": "Returns data for activities (volumes) form by organization INN and year.",
                "parameters": [
                    {
                        "name": "inn",
                        "in": "query",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "year",
                        "in": "query",
                        "required": true,
                        "type": "integer",
                        "format": "int32",
                        "description": "Reporting year"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Activities form data",
                        "schema": {"$ref": "#/definitions/ActivitiesFormData"}
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/form/financial/organization/{org_id}": {
            "get": {
                "tags": ["Reporting Form"],
                "summary": "Get financial form data by organization ID and year",
                "description": "Returns data for revenues and expenses form including all income sources and expense categories.",
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
                        "in": "query",
                        "required": true,
                        "type": "integer",
                        "format": "int32",
                        "description": "Reporting year"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Financial form data",
                        "schema": {"$ref": "#/definitions/FinancialFormData"}
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "404": {
                        "description": "Financial data not found",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        },
        "/api/v1/form/financial/search/by-inn": {
            "get": {
                "tags": ["Reporting Form"],
                "summary": "Get financial form data by organization INN and year",
                "description": "Returns data for revenues and expenses form by organization INN and year.",
                "parameters": [
                    {
                        "name": "inn",
                        "in": "query",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "year",
                        "in": "query",
                        "required": true,
                        "type": "integer",
                        "format": "int32",
                        "description": "Reporting year"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Financial form data",
                        "schema": {"$ref": "#/definitions/FinancialFormData"}
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "404": {
                        "description": "Financial data not found",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {"$ref": "#/definitions/ErrorResponse"}
                    }
                }
            }
        }
    },
    "definitions": {
        "Activity": {
            "type": "object",
            "properties": {
                "ID": {"type": "integer", "format": "int64"},
                "IDOwner": {"type": "integer", "format": "int64"},
                "INN": {"type": "string"},
                "ActivityTypeID": {"type": "integer", "format": "int64"},
                "ActivityTypeName": {"type": "string"},
                "VolumeIndicator": {"type": "string", "description": "Volume indicator for the activity type"},
                "CustomActivityID": {"type": "integer", "format": "int64"},
                "CustomActivityName": {"type": "string"},
                "VisitorCategory": {"type": "string", "enum": ["internal","external"]},
                "CostSharePercent": {"type": "number", "format": "float"},
                "RevenueAmount": {"type": "number", "format": "float"},
                "TotalCount": {"type": "integer", "format": "int64"},
                "StateTaskCount": {"type": "integer", "format": "int64"},
                "RevenueActivityCount": {"type": "integer", "format": "int64"},
                "Year": {"type": "integer", "format": "int32"}
            }
        },
        "CreateActivityRequest": {
            "type": "object",
            "required": ["inn","visitor_category","year"],
            "properties": {
                "inn": {"type": "string"},
                "activity_type_id": {"type": "integer", "format": "int64"},
                "custom_activity_id": {"type": "integer", "format": "int64"},
                "visitor_category": {"type": "string", "enum": ["internal","external"]},
                "cost_share_percent": {"type": "number", "format": "float"},
                "revenue_amount": {"type": "number", "format": "float"},
                "total_count": {"type": "integer", "format": "int64"},
                "state_task_count": {"type": "integer", "format": "int64"},
                "revenue_activity_count": {"type": "integer", "format": "int64"},
                "year": {"type": "integer", "format": "int32"}
            }
        },
        "UpdateActivityRequest": {
            "type": "object",
            "required": ["id"],
            "properties": {
                "id": {"type": "integer", "format": "int64"},
                "cost_share_percent": {"type": "number", "format": "float"},
                "revenue_amount": {"type": "number", "format": "float"},
                "total_count": {"type": "integer", "format": "int64"},
                "state_task_count": {"type": "integer", "format": "int64"},
                "revenue_activity_count": {"type": "integer", "format": "int64"},
                "year": {"type": "integer", "format": "int32"}
            }
        },
        "CreateMuseumRequest": {
            "type": "object",
            "required": ["id_owner","inn","museum_activity_in_charter","name"],
            "properties": {
                "id_owner": {"type": "integer", "format": "int64"},
                "inn": {"type": "string"},
                "kpp": {"type": "string"},
                "founder": {"type": "string"},
                "museum_activity_in_charter": {"type": "boolean"},
                "name": {"type": "string"},
                "museum_legal_status": {"type": "string"},
                "is_memorial_reserve_museum": {"type": "boolean"},
                "is_historical_memorial_reserve": {"type": "boolean"},
                "is_art_museum": {"type": "boolean"},
                "is_museum_reserve": {"type": "boolean"},
                "is_estate_museum": {"type": "boolean"},
                "is_palace_park_ensemble": {"type": "boolean"},
                "is_historical_architectural_reserve": {"type": "boolean"},
                "annual_visitor_capacity": {"type": "integer", "format": "int64"},
                "internal_visitors_count": {"type": "integer", "format": "int64"},
                "external_visitors_count": {"type": "integer", "format": "int64"},
                "is_valuable_cultural_heritage": {"type": "boolean"},
                "valuable_museum_items_count": {"type": "integer", "format": "int64"}
            }
        },
        "UpdateMuseumRequest": {
            "type": "object",
            "required": ["id","id_owner","inn","museum_activity_in_charter","name"],
            "properties": {
                "id": {"type": "integer", "format": "int64"},
                "id_owner": {"type": "integer", "format": "int64"},
                "inn": {"type": "string"},
                "kpp": {"type": "string"},
                "founder": {"type": "string"},
                "museum_activity_in_charter": {"type": "boolean"},
                "name": {"type": "string"},
                "museum_legal_status": {"type": "string"},
                "is_memorial_reserve_museum": {"type": "boolean"},
                "is_historical_memorial_reserve": {"type": "boolean"},
                "is_art_museum": {"type": "boolean"},
                "is_museum_reserve": {"type": "boolean"},
                "is_estate_museum": {"type": "boolean"},
                "is_palace_park_ensemble": {"type": "boolean"},
                "is_historical_architectural_reserve": {"type": "boolean"},
                "annual_visitor_capacity": {"type": "integer", "format": "int64"},
                "internal_visitors_count": {"type": "integer", "format": "int64"},
                "external_visitors_count": {"type": "integer", "format": "int64"},
                "is_valuable_cultural_heritage": {"type": "boolean"},
                "valuable_museum_items_count": {"type": "integer", "format": "int64"}
            }
        },
        "CreateOrganizationRequest": {
            "type": "object",
            "required": ["inn","name","exist_museum"],
            "properties": {
                "inn": {"type": "string"},
                "name": {"type": "string"},
                "exist_museum": {"type": "boolean"}
            }
        },
        "UpdateOrganizationRequest": {
            "type": "object",
            "required": ["inn","name","exist_museum"],
            "properties": {
                "inn": {"type": "string"},
                "name": {"type": "string"},
                "exist_museum": {"type": "boolean"}
            }
        },
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {"type": "string"}
            }
        },
        "Museum": {
            "type": "object",
            "properties": {
                "Id": {"type": "integer", "format": "int64"},
                "IdOwner": {"type": "integer", "format": "int64"},
                "INN": {"type": "string"},
                "KPP": {"type": "string"},
                "Founder": {"type": "string"},
                "MuseumActivityInCharter": {"type": "boolean"},
                "Name": {"type": "string"},
                "MuseumLegalStatus": {"type": "string"},
                "IsMemorialReserveMuseum": {"type": "boolean"},
                "IsHistoricalMemorialReserve": {"type": "boolean"},
                "IsArtMuseum": {"type": "boolean"},
                "IsMuseumReserve": {"type": "boolean"},
                "IsEstateMuseum": {"type": "boolean"},
                "IsPalaceParkEnsemble": {"type": "boolean"},
                "IsHistoricalArchitecturalReserve": {"type": "boolean"},
                "AnnualVisitorCapacity": {"type": "integer", "format": "int64"},
                "InternalVisitorsCount": {"type": "integer", "format": "int64"},
                "ExternalVisitorsCount": {"type": "integer", "format": "int64"},
                "IsValuableCulturalHeritage": {"type": "boolean"},
                "ValuableMuseumItemsCount": {"type": "integer", "format": "int64"}
            }
        },
        "Organization": {
            "type": "object",
            "properties": {
                "ID": {"type": "integer", "format": "int64"},
                "INN": {"type": "string"},
                "Name": {"type": "string"},
                "ExistMuseum": {"type": "boolean"}
            }
        },
        "LaborResponse": {
            "type": "object",
            "properties": {
                "total_staff_annual": {"type": "number", "format": "float"},
                "research_staff_internal": {"type": "number", "format": "float"},
                "core_operational_staff_internal": {"type": "number", "format": "float"},
                "admin_support_staff_internal": {"type": "number", "format": "float"},
                "research_staff_external": {"type": "number", "format": "float"},
                "core_operational_staff_external": {"type": "number", "format": "float"},
                "admin_support_staff_external": {"type": "number", "format": "float"},
                "fot": {"$ref": "#/definitions/LaborFotResponse"}
            }
        },
        "LaborFotResponse": {
            "type": "object",
            "properties": {
                "total_staff_annual": {"type": "number", "format": "float"},
                "research_staff_internal": {"type": "number", "format": "float"},
                "core_operational_staff_internal": {"type": "number", "format": "float"},
                "admin_support_staff_internal": {"type": "number", "format": "float"},
                "research_staff_external": {"type": "number", "format": "float"},
                "core_operational_staff_external": {"type": "number", "format": "float"},
                "admin_support_staff_external": {"type": "number", "format": "float"}
            }
        },
        "PaginatedOrganizationResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {"$ref": "#/definitions/Organization"}
                },
                "page": {"type": "integer", "format": "int32", "description": "Current page number"},
                "page_size": {"type": "integer", "format": "int32", "description": "Number of items per page"},
                "total_count": {"type": "integer", "format": "int64", "description": "Total number of items"},
                "total_pages": {"type": "integer", "format": "int32", "description": "Total number of pages"}
            }
        },
        "PaginatedMuseumResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {"$ref": "#/definitions/Museum"}
                },
                "page": {"type": "integer", "format": "int32", "description": "Current page number"},
                "page_size": {"type": "integer", "format": "int32", "description": "Number of items per page"},
                "total_count": {"type": "integer", "format": "int64", "description": "Total number of items"},
                "total_pages": {"type": "integer", "format": "int32", "description": "Total number of pages"}
            }
        },
        "PaginatedActivityResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {"$ref": "#/definitions/Activity"}
                },
                "page": {"type": "integer", "format": "int32", "description": "Current page number"},
                "page_size": {"type": "integer", "format": "int32", "description": "Number of items per page"},
                "total_count": {"type": "integer", "format": "int64", "description": "Total number of items"},
                "total_pages": {"type": "integer", "format": "int32", "description": "Total number of pages"}
            }
        },
        "LaborFormData": {
            "type": "object",
            "properties": {
                "year": {"type": "integer", "format": "int32", "description": "Reporting year"},
                "labor_data": {"$ref": "#/definitions/LaborResponse"}
            }
        },
        "ActivitiesFormData": {
            "type": "object",
            "properties": {
                "year": {"type": "integer", "format": "int32", "description": "Reporting year"},
                "activities": {
                    "type": "array",
                    "items": {"$ref": "#/definitions/ActivityFormItem"}
                }
            }
        },
        "ActivityFormItem": {
            "type": "object",
            "properties": {
                "id": {"type": "integer", "format": "int64"},
                "activity_type_id": {"type": "integer", "format": "int64"},
                "activity_type_name": {"type": "string"},
                "volume_indicator": {"type": "string", "description": "Volume indicator for the activity type"},
                "custom_activity_id": {"type": "integer", "format": "int64"},
                "custom_activity_name": {"type": "string"},
                "cost_share_percent": {"type": "number", "format": "float", "description": "Cost share percentage"},
                "revenue_amount": {"type": "number", "format": "float", "description": "Revenue amount in rubles"},
                "visitor_category": {"type": "string", "enum": ["internal", "external"]},
                "total_count": {"type": "integer", "format": "int64", "description": "Total count"},
                "state_task_count": {"type": "integer", "format": "int64", "description": "State task count"},
                "revenue_activity_count": {"type": "integer", "format": "int64", "description": "Revenue activity count"}
            }
        },
        "FinancialFormData": {
            "type": "object",
            "properties": {
                "year": {"type": "integer", "format": "int32", "description": "Reporting year"},
                "financial": {"$ref": "#/definitions/FinancialData"}
            }
        },
        "FinancialData": {
            "type": "object",
            "properties": {
                "total_revenue": {"type": "number", "format": "float", "description": "Total revenue in rubles"},
                "state_assignment_subsidy": {"type": "number", "format": "float", "description": "State assignment subsidy in rubles"},
                "earned_revenue": {"type": "number", "format": "float", "description": "Earned revenue in rubles"},
                "other_funding_sources": {"type": "number", "format": "float", "description": "Other funding sources in rubles"},
                "total_expenses": {"type": "number", "format": "float", "description": "Total expenses in rubles"},
                "inventory_assets_acquisition": {"type": "number", "format": "float", "description": "Inventory assets acquisition in rubles"},
                "utility_services": {"type": "number", "format": "float", "description": "Utility services in rubles"},
                "communication_services": {"type": "number", "format": "float", "description": "Communication services in rubles"},
                "transportation_services": {"type": "number", "format": "float", "description": "Transportation services in rubles"},
                "valuable_assets_acquisition": {"type": "number", "format": "float", "description": "Valuable assets acquisition in rubles"},
                "real_estate_maintenance": {"type": "number", "format": "float", "description": "Real estate maintenance in rubles"},
                "valuable_assets_maintenance": {"type": "number", "format": "float", "description": "Valuable assets maintenance in rubles"},
                "general_administrative_costs": {"type": "number", "format": "float", "description": "General administrative costs in rubles"},
                "tax_payments": {"type": "number", "format": "float", "description": "Tax payments in rubles"},
                "other_expenses": {"type": "number", "format": "float", "description": "Other expenses in rubles"}
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
