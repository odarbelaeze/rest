# rest

A minimal Go library for handling JSON in REST APIs.

## Installation

```bash
go get github.com/odarbelaeze/rest
```

## Usage

### Decoding Requests

Use `Load` to decode a JSON request body into a struct.

```go
data, err := rest.Load[MyStruct](r)
```

### Writing Responses

Use `JSON` to write a structured JSON response.

```go
rest.JSON(ctx, w, http.StatusOK, data)
```

### Handling Errors

Use `Err` to write standardized JSON error responses. The message is
automatically derived from the HTTP status code.

```go
rest.Err(ctx, w, http.StatusNotFound, rest.WithDetails("ID: 123"))
```
