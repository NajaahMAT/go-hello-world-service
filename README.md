# Go Hello World Service

A simple RESTful HTTP API written in Go using only the standard `net/http`
package. The service validates a `name` query parameter and returns an appropriate
greeting based on the first letter.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Running Tests](#running-tests)
- [Postman Collection](#postman-collection)
- [Project Structure](#project-structure)
- [Make Commands](#make-commands)
- [Assumptions](#assumptions)

## Features

- A `GET /hello-world` endpoint
- A health-check endpoint at `/health`
- Alphabet-range validation for A-M and N-Z
- JSON success and error responses
- Request logging and panic recovery middleware
- CORS headers
- Graceful server shutdown
- Unit and handler test coverage

## Prerequisites

- Go 1.18 or later
- Make (optional)
- curl or Postman (optional)

Verify Go:

```bash
go version
```

## Getting Started

Clone the repository and enter the project directory:

```bash
git clone https://github.com/NajaahMAT/go-hello-world-service.git
cd go-hello-world-service
```

Download module dependencies:

```bash
go mod download
```

## Running the Application

Run in development mode:

```bash
go run ./cmd/server
```

The server listens on port `8889` by default.

Run on a custom port:

```bash
PORT=3000 go run ./cmd/server
```

Build and run the executable:

```bash
make build
./bin/server
```

Stop a running server with `Ctrl+C`.

## API Documentation

### Get Greeting

```http
GET /hello-world?name=<name>
```

#### Valid Request

**Request**

```bash
curl "http://localhost:8889/hello-world?name=Alice"
```

**Response**

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "message": "Hello Alice"
}
```

Names beginning with A-M (case-insensitive) are valid.

#### Invalid Name

**Request**

```bash
curl "http://localhost:8889/hello-world?name=Zane"
```

**Response**

```http
HTTP/1.1 400 Bad Request
Content-Type: application/json
```

```json
{
  "error": "Invalid Input"
}
```

Names beginning with N-Z, symbols, digits, or Unicode letters outside the
English alphabet are invalid.

#### Missing or Empty Name

```bash
curl "http://localhost:8889/hello-world"
curl "http://localhost:8889/hello-world?name="
```

Both requests return `400 Bad Request`:

```json
{
  "error": "Invalid Input"
}
```

### Health Check

**Request**

```bash
curl "http://localhost:8889/health"
```

**Response**

```json
{
  "status": "ok"
}
```

## Running Tests

Run the complete test suite:

```bash
go test ./...
```

Run verbose tests:

```bash
make test
```

Generate a coverage report:

```bash
make test-coverage
```

Open the HTML coverage report:

```bash
go tool cover -html=coverage.out
```

## Postman Collection

Import `postman_collection.json` into Postman:

1. Open Postman.
2. Select **Import**.
3. Choose `postman_collection.json`.
4. Set the `baseUrl` collection variable if the server is not running on
   `http://localhost:8889`.

The collection covers:

- Valid A-M names
- Invalid N-Z names
- Missing and empty names
- Whitespace-only and non-letter names
- Uppercase and lowercase boundaries
- POST and PUT method rejection
- Health-check behavior

## Project Structure

```text
.
├── cmd/server/            Application entry point
├── internal/
│   ├── domain/            Domain models and validation
│   ├── handler/           HTTP handlers
│   ├── middleware/        Logging, recovery, and CORS middleware
│   └── service/           Greeting business logic
├── pkg/response/          JSON response helpers
├── Makefile
├── go.mod
├── postman_collection.json
└── README.md
```

## Make Commands

```text
make help             Show available commands
make build            Build bin/server
make run              Run the application
make run-port PORT=3000
                      Run on a custom port
make test             Run tests verbosely
make test-coverage    Generate a coverage report
make format           Format Go source files
make clean            Remove build and coverage artifacts
make lint             Run golangci-lint if installed
make install-tools    Install golangci-lint
```

## Assumptions

- Valid names begin with an ASCII letter A-M, case-insensitively.
- Names beginning with N-Z or starting with another character are invalid.
- Missing, empty, and whitespace-only names are invalid.
- Name casing is normalized in successful greetings.
- Only `GET` requests are accepted by `/hello-world`.
- No external dependencies or database are required.
