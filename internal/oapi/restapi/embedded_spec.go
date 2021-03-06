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
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "# Sudoku generation and verification server\n## List of all custom errors\nFirst number is HTTP Status code, second is value of \"code\" field in returned JSON object, text description may or may not match \"message\" field in returned JSON object.\n- 400.700: Provided solution does not match to the original sudoku.\n- 400.701: Original field could not be decrypted \n- 422.702: Solution is invalid\n",
    "title": "Sudoku server",
    "version": "4.0.0"
  },
  "basePath": "/",
  "paths": {
    "/check": {
      "post": {
        "tags": [
          "check"
        ],
        "operationId": "checkSudoku",
        "parameters": [
          {
            "name": "toCheck",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/toCheck"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "validation result",
            "schema": {
              "$ref": "#/definitions/result"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/get": {
      "get": {
        "tags": [
          "get"
        ],
        "operationId": "getSudoku",
        "parameters": [
          {
            "enum": [
              "common",
              "big",
              "huge"
            ],
            "type": "string",
            "default": "common",
            "name": "size",
            "in": "query"
          },
          {
            "enum": [
              "easy",
              "heavy",
              "nightmare"
            ],
            "type": "string",
            "default": "easy",
            "name": "difficulty",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "sudoku to be solved",
            "schema": {
              "$ref": "#/definitions/sudoku"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "field": {
      "type": "array",
      "items": {
        "maxItems": 25,
        "minItems": 9,
        "$ref": "#/definitions/row"
      }
    },
    "original": {
      "description": "encrypted sudoku original returned by the ` + "`" + `/get` + "`" + ` and to be used later with ` + "`" + `/check` + "`" + `",
      "type": "string"
    },
    "result": {
      "type": "object",
      "properties": {
        "sameAsOriginal": {
          "type": "boolean"
        }
      }
    },
    "row": {
      "type": "array",
      "items": {
        "type": "integer",
        "format": "int32",
        "maxItems": 25,
        "minItems": 9
      }
    },
    "sudoku": {
      "type": "object",
      "required": [
        "field",
        "original"
      ],
      "properties": {
        "field": {
          "$ref": "#/definitions/field"
        },
        "original": {
          "$ref": "#/definitions/original"
        }
      }
    },
    "toCheck": {
      "type": "object",
      "required": [
        "field",
        "original"
      ],
      "properties": {
        "field": {
          "$ref": "#/definitions/field"
        },
        "original": {
          "$ref": "#/definitions/original"
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
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "# Sudoku generation and verification server\n## List of all custom errors\nFirst number is HTTP Status code, second is value of \"code\" field in returned JSON object, text description may or may not match \"message\" field in returned JSON object.\n- 400.700: Provided solution does not match to the original sudoku.\n- 400.701: Original field could not be decrypted \n- 422.702: Solution is invalid\n",
    "title": "Sudoku server",
    "version": "4.0.0"
  },
  "basePath": "/",
  "paths": {
    "/check": {
      "post": {
        "tags": [
          "check"
        ],
        "operationId": "checkSudoku",
        "parameters": [
          {
            "name": "toCheck",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/toCheck"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "validation result",
            "schema": {
              "$ref": "#/definitions/result"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/get": {
      "get": {
        "tags": [
          "get"
        ],
        "operationId": "getSudoku",
        "parameters": [
          {
            "enum": [
              "common",
              "big",
              "huge"
            ],
            "type": "string",
            "default": "common",
            "name": "size",
            "in": "query"
          },
          {
            "enum": [
              "easy",
              "heavy",
              "nightmare"
            ],
            "type": "string",
            "default": "easy",
            "name": "difficulty",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "sudoku to be solved",
            "schema": {
              "$ref": "#/definitions/sudoku"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "field": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/row"
      }
    },
    "original": {
      "description": "encrypted sudoku original returned by the ` + "`" + `/get` + "`" + ` and to be used later with ` + "`" + `/check` + "`" + `",
      "type": "string"
    },
    "result": {
      "type": "object",
      "properties": {
        "sameAsOriginal": {
          "type": "boolean"
        }
      }
    },
    "row": {
      "type": "array",
      "items": {
        "type": "integer",
        "format": "int32"
      }
    },
    "sudoku": {
      "type": "object",
      "required": [
        "field",
        "original"
      ],
      "properties": {
        "field": {
          "$ref": "#/definitions/field"
        },
        "original": {
          "$ref": "#/definitions/original"
        }
      }
    },
    "toCheck": {
      "type": "object",
      "required": [
        "field",
        "original"
      ],
      "properties": {
        "field": {
          "$ref": "#/definitions/field"
        },
        "original": {
          "$ref": "#/definitions/original"
        }
      }
    }
  }
}`))
}
