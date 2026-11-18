# Go Hello World Service

A small HTTP API built with Go's standard `net/http` package. It returns a greeting
when `name` starts with a letter in A-M and returns a 400 response for names starting
with N-Z or for missing/empty names.

## Requirements

- Go 1.18 or later

## Run

From the repository root:

```bash
go run ./cmd/server
```

The server listens on port `8889` by default. Set another port with `PORT`:

```bash
PORT=3000 go run ./cmd/server
```

Using Make:

```bash
make run
make run-port PORT=3000
make build
./bin/server
```

## Test

```bash
go test ./...
```

For verbose tests:

```bash
make test
```

## API

| Request | Status | Response |
|---|---:|---|
| `GET /hello-world?name=Alice` | 200 | `{"message":"Hello Alice"}` |
| `GET /hello-world?name=alice` | 200 | `{"message":"Hello Alice"}` |
| `GET /hello-world?name=Zane` | 400 | `{"error":"Invalid Input"}` |
| `GET /hello-world` | 400 | `{"error":"Invalid Input"}` |

Example:

```bash
curl "http://localhost:8889/hello-world?name=Alice"
curl "http://localhost:8889/hello-world?name=Zane"
curl "http://localhost:8889/health"
```

## Assumptions

- Valid names begin with an ASCII letter A-M, case-insensitively.
- Names beginning with N-Z or another character are invalid.
- Missing, empty, or whitespace-only names are invalid.
- Only GET requests are accepted for `/hello-world`.
- Names are cosmetic text: no minimum or maximum length is enforced.
