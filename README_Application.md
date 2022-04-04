# REPL-GO
Simple Golang REPL

# Description
This application is a command line REPL (read-eval-print loop) that drives a simple in-memory key/value storage system. This application also allows nested transactions where a transaction can be committed or aborted.

# Setup Local
1. Clone the repository under GOPATH
2. Install dependencies using ```go mod download```
3) Run the application using ```go run cmd/repl/main.go```

# Run Linter
```golangci-lint run -v -c golangci.yml```

# Run Tests
```go test -v -cover ./...```

# Example
![](https://github.com/karthikpothineni/staticfiles/blob/main/gif/repl-go.gif)