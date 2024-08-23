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
        "/attachments/{ticketId}": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Attachments"
                ],
                "summary": "upload an attachment for a trouble ticket",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Attachment ID",
                        "name": "ticketId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Attachment file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AttachmentDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/auth/signIn": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "summary": "Sign In",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Auth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AuthJwtPayload"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/auth/signUp": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Done by System Management Personnel's",
                "tags": [
                    "Auth"
                ],
                "summary": "Sign Up",
                "parameters": [
                    {
                        "description": "Signup credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/troubleTickets": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Trouble Tickets"
                ],
                "summary": "fetch all trouble tickets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.TroubleTicketDTO"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/troubleTickets/filters": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Trouble Tickets"
                ],
                "summary": "fetch all related trouble tickets filters / dropdown",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.FiltersDTO"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AttachmentDTO": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "mime_type": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "original_name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "ref": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.Auth": {
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
        "models.AuthJwtPayload": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "id_token": {
                    "type": "string"
                },
                "not-before-policy": {
                    "type": "integer"
                },
                "refresh_expires_in": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "scope": {
                    "type": "string"
                },
                "session_state": {
                    "type": "string"
                },
                "token_type": {
                    "type": "string"
                }
            }
        },
        "models.ChannelDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.ExternalIdentifierDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                },
                "ref": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/models.TypeDTO"
                }
            }
        },
        "models.FiltersDTO": {
            "type": "object",
            "properties": {
                "channels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ChannelDTO"
                    }
                },
                "priorities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.PriorityDTO"
                    }
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RoleDTO"
                    }
                },
                "severities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SeverityDTO"
                    }
                },
                "statuses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.StatusDTO"
                    }
                },
                "types": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TypeDTO"
                    }
                }
            }
        },
        "models.NoteDTO": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.PartyDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/models.RoleDTO"
                }
            }
        },
        "models.PriorityDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "sequence": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.RelatedEntityDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "ref": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.RelatedPartyDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "party": {
                    "$ref": "#/definitions/models.PartyDTO"
                }
            }
        },
        "models.RoleDTO": {
            "type": "object",
            "properties": {
                "filter": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                }
            }
        },
        "models.SeverityDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.SignUpDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "realmRoles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.StatusChangeDTO": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "reason": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/models.StatusDTO"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        },
        "models.StatusDTO": {
            "type": "object",
            "properties": {
                "filter": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                }
            }
        },
        "models.TroubleTicketDTO": {
            "type": "object",
            "properties": {
                "attachments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.AttachmentDTO"
                    }
                },
                "channel": {
                    "$ref": "#/definitions/models.ChannelDTO"
                },
                "created_at": {
                    "description": "autoPopulate with the current timestamp on record creation",
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "deleted_by": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "expected_resolution_date": {
                    "type": "string"
                },
                "external_identifiers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ExternalIdentifierDTO"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "notes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.NoteDTO"
                    }
                },
                "priority": {
                    "$ref": "#/definitions/models.PriorityDTO"
                },
                "ref": {
                    "type": "string"
                },
                "related_entities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RelatedEntityDTO"
                    }
                },
                "related_parties": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RelatedPartyDTO"
                    }
                },
                "resolution_date": {
                    "type": "string"
                },
                "severity": {
                    "$ref": "#/definitions/models.SeverityDTO"
                },
                "status": {
                    "$ref": "#/definitions/models.StatusDTO"
                },
                "status_changes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.StatusChangeDTO"
                    }
                },
                "type": {
                    "$ref": "#/definitions/models.TypeDTO"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        },
        "models.TypeDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
