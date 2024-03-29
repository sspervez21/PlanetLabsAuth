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
    "description": "PlantLabs Auth Service",
    "title": "PlanetAuth",
    "version": "0.1.0"
  },
  "paths": {
    "/groups": {
      "post": {
        "operationId": "CreateGroup",
        "parameters": [
          {
            "name": "CreateGroupInput",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Group"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "409": {
            "description": "Group already exists",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      }
    },
    "/groups/{groupName}": {
      "get": {
        "operationId": "GetGroup",
        "parameters": [
          {
            "type": "string",
            "name": "groupName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "minItems": 1,
              "items": {
                "type": "string"
              }
            }
          },
          "404": {
            "description": "Group not found",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      },
      "put": {
        "operationId": "UpdateGroup",
        "parameters": [
          {
            "type": "string",
            "name": "groupName",
            "in": "path",
            "required": true
          },
          {
            "name": "UpdateGroupInput",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/GroupList"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "404": {
            "description": "Group does not exists",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      },
      "delete": {
        "operationId": "DeleteGroup",
        "parameters": [
          {
            "type": "string",
            "name": "groupName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "404": {
            "description": "Group does not exists",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      }
    },
    "/users": {
      "post": {
        "operationId": "CreateUser",
        "parameters": [
          {
            "name": "CreateUserInput",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserRecord"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "409": {
            "description": "User already exists",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      }
    },
    "/users/{userId}": {
      "get": {
        "operationId": "GetUser",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/UserRecord"
            }
          },
          "404": {
            "description": "User not found",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      },
      "put": {
        "operationId": "UpdateUser",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          },
          {
            "name": "UpdateUserInput",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserRecord"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "404": {
            "description": "User does not exists",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      },
      "delete": {
        "operationId": "DeleteUser",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "404": {
            "description": "User does not exists",
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
      "description": "Bad Request object",
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Group": {
      "description": "A Group of Users",
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "GroupList": {
      "description": "A list of groups",
      "type": "object",
      "required": [
        "list"
      ],
      "properties": {
        "list": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "string"
          }
        }
      }
    },
    "NotFound": {
      "description": "Not Found object",
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "UserRecord": {
      "description": "A User Record",
      "type": "object",
      "required": [
        "firstName",
        "lastName",
        "userId",
        "groups"
      ],
      "properties": {
        "firstName": {
          "type": "string"
        },
        "groups": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "string"
          }
        },
        "lastName": {
          "type": "string"
        },
        "userId": {
          "type": "string"
        }
      }
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
    "description": "PlantLabs Auth Service",
    "title": "PlanetAuth",
    "version": "0.1.0"
  },
  "paths": {
    "/groups": {
      "post": {
        "operationId": "CreateGroup",
        "parameters": [
          {
            "name": "CreateGroupInput",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Group"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "409": {
            "description": "Group already exists",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      }
    },
    "/groups/{groupName}": {
      "get": {
        "operationId": "GetGroup",
        "parameters": [
          {
            "type": "string",
            "name": "groupName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "minItems": 1,
              "items": {
                "type": "string"
              }
            }
          },
          "404": {
            "description": "Group not found",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      },
      "put": {
        "operationId": "UpdateGroup",
        "parameters": [
          {
            "type": "string",
            "name": "groupName",
            "in": "path",
            "required": true
          },
          {
            "name": "UpdateGroupInput",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/GroupList"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "404": {
            "description": "Group does not exists",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      },
      "delete": {
        "operationId": "DeleteGroup",
        "parameters": [
          {
            "type": "string",
            "name": "groupName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "404": {
            "description": "Group does not exists",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      }
    },
    "/users": {
      "post": {
        "operationId": "CreateUser",
        "parameters": [
          {
            "name": "CreateUserInput",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserRecord"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "409": {
            "description": "User already exists",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      }
    },
    "/users/{userId}": {
      "get": {
        "operationId": "GetUser",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/UserRecord"
            }
          },
          "404": {
            "description": "User not found",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      },
      "put": {
        "operationId": "UpdateUser",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          },
          {
            "name": "UpdateUserInput",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserRecord"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "404": {
            "description": "User does not exists",
            "schema": {
              "$ref": "#/definitions/BadRequest"
            }
          }
        }
      },
      "delete": {
        "operationId": "DeleteUser",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "404": {
            "description": "User does not exists",
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
      "description": "Bad Request object",
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Group": {
      "description": "A Group of Users",
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "GroupList": {
      "description": "A list of groups",
      "type": "object",
      "required": [
        "list"
      ],
      "properties": {
        "list": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "string"
          }
        }
      }
    },
    "NotFound": {
      "description": "Not Found object",
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "UserRecord": {
      "description": "A User Record",
      "type": "object",
      "required": [
        "firstName",
        "lastName",
        "userId",
        "groups"
      ],
      "properties": {
        "firstName": {
          "type": "string"
        },
        "groups": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "string"
          }
        },
        "lastName": {
          "type": "string"
        },
        "userId": {
          "type": "string"
        }
      }
    }
  }
}`))
}
