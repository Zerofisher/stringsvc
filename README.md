# StringService - A Go kit Learning Project

## Introduction

This project is a refactored version of the "stringsvc" example from the Go kit tutorial. The original tutorial can be found at [https://gokit.io/examples/stringsvc.html](https://gokit.io/examples/stringsvc.html).

The primary purpose of this project is to serve as a learning resource for those interested in:

1. Understanding Go kit's architecture and components
2. Implementing microservices using Go kit
3. Structuring a Go project following standard practices

## Getting Started

To run the service locally:

1. Clone the repository
2. Navigate to the project root
3. Run `go mod tidy` to download dependencies
4. Start the server with `go run cmd/server/main.go`

The service will be available at `http://localhost:8080`.

## Testing the Service

You can test the service using curl:

```bash
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count