// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "photo API",
    "title": "BA API",
    "version": "0.0.1"
  },
  "basePath": "/api/v1",
  "paths": {
    "/photos/{id}": {
      "get": {
        "security": [
          {
            "ui-api": []
          }
        ],
        "tags": [
          "ui"
        ],
        "summary": "Returns list of photos",
        "operationId": "getPhotos",
        "parameters": [
          {
            "type": "string",
            "description": "for now, flickr user id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "list of photos",
            "schema": {
              "$ref": "#/definitions/Photos"
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          },
          "500": {
            "$ref": "#/responses/ServerError"
          }
        }
      }
    },
    "/users": {
      "post": {
        "security": [
          {
            "intern-api": []
          }
        ],
        "tags": [
          "intern"
        ],
        "summary": "Create user from internal ba webserver",
        "operationId": "saveUser",
        "parameters": [
          {
            "description": "user to create.",
            "name": "user",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/users/{id}": {
      "get": {
        "security": [
          {
            "intern-api": []
          }
        ],
        "tags": [
          "intern"
        ],
        "summary": "Get user from internal ba webserver",
        "operationId": "getUsersId",
        "parameters": [
          {
            "type": "string",
            "description": "flickr user id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Photos": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "description": {
            "type": "string"
          },
          "id": {
            "type": "string"
          },
          "title": {
            "type": "string"
          },
          "url": {
            "type": "string"
          }
        }
      }
    },
    "User": {
      "type": "object",
      "required": [
        "id",
        "flickrToken",
        "flickrSecretToken"
      ],
      "properties": {
        "flickrSecretToken": {
          "type": "string"
        },
        "flickrToken": {
          "type": "string"
        },
        "id": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Bad request",
      "schema": {
        "type": "object",
        "properties": {
          "code": {
            "type": "integer",
            "format": "int64"
          },
          "message": {
            "type": "string"
          }
        }
      }
    },
    "NotFound": {
      "description": "Not Found"
    },
    "ServerError": {
      "description": "Server Error"
    }
  },
  "securityDefinitions": {
    "intern-api": {
      "type": "apiKey",
      "name": "x-intern-ba-token",
      "in": "header"
    },
    "ui-api": {
      "type": "apiKey",
      "name": "x-ui-ba-token",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "photo API",
    "title": "BA API",
    "version": "0.0.1"
  },
  "basePath": "/api/v1",
  "paths": {
    "/photos/{id}": {
      "get": {
        "security": [
          {
            "ui-api": []
          }
        ],
        "tags": [
          "ui"
        ],
        "summary": "Returns list of photos",
        "operationId": "getPhotos",
        "parameters": [
          {
            "type": "string",
            "description": "for now, flickr user id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "list of photos",
            "schema": {
              "$ref": "#/definitions/Photos"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "type": "integer",
                  "format": "int64"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "404": {
            "description": "Not Found"
          },
          "500": {
            "description": "Server Error"
          }
        }
      }
    },
    "/users": {
      "post": {
        "security": [
          {
            "intern-api": []
          }
        ],
        "tags": [
          "intern"
        ],
        "summary": "Create user from internal ba webserver",
        "operationId": "saveUser",
        "parameters": [
          {
            "description": "user to create.",
            "name": "user",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/users/{id}": {
      "get": {
        "security": [
          {
            "intern-api": []
          }
        ],
        "tags": [
          "intern"
        ],
        "summary": "Get user from internal ba webserver",
        "operationId": "getUsersId",
        "parameters": [
          {
            "type": "string",
            "description": "flickr user id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Photos": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "description": {
            "type": "string"
          },
          "id": {
            "type": "string"
          },
          "title": {
            "type": "string"
          },
          "url": {
            "type": "string"
          }
        }
      }
    },
    "User": {
      "type": "object",
      "required": [
        "id",
        "flickrToken",
        "flickrSecretToken"
      ],
      "properties": {
        "flickrSecretToken": {
          "type": "string"
        },
        "flickrToken": {
          "type": "string"
        },
        "id": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Bad request",
      "schema": {
        "type": "object",
        "properties": {
          "code": {
            "type": "integer",
            "format": "int64"
          },
          "message": {
            "type": "string"
          }
        }
      }
    },
    "NotFound": {
      "description": "Not Found"
    },
    "ServerError": {
      "description": "Server Error"
    }
  },
  "securityDefinitions": {
    "intern-api": {
      "type": "apiKey",
      "name": "x-intern-ba-token",
      "in": "header"
    },
    "ui-api": {
      "type": "apiKey",
      "name": "x-ui-ba-token",
      "in": "header"
    }
  }
}`))
}
