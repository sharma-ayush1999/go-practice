# Useful Go Commands

## Running Code
```bash
go run .                     # compile + run the current package (use in project root)
go run main.go               # run a specific file
go run . -flag value         # run with flags
```

## Building
```bash
go build .                   # compile but don't run — checks for errors
go build -o server .         # compile and name the output binary "server"
./server                     # run the compiled binary
```

## Testing
```bash
go test ./...                # run all tests in all packages
go test ./jobs/...           # run tests only in the jobs package
go test -v ./...             # verbose — shows each test name and result
go test -race ./...          # run with race detector — catches data races
go test -cover ./...         # show test coverage percentage
go test -run TestJobStore    # run only tests matching this name
go test -bench=. ./...       # run benchmarks
go test -bench=. -benchmem   # benchmarks + memory allocation stats
```

> Always run `go test -race ./...` before pushing code. Data races are silent bugs.

## Dependencies
```bash
go mod init github.com/you/project  # create go.mod (like npm init)
go get github.com/google/uuid       # add a dependency (like npm install)
go mod tidy                         # add missing + remove unused dependencies
go mod download                     # download all dependencies to local cache
go list -m all                      # list all dependencies
```

## Code Quality
```bash
go fmt ./...                 # format all code — run this constantly
go vet ./...                 # catch common bugs (wrong format strings, unreachable code, etc.)
go doc fmt.Println           # read docs for any function in the terminal
go doc sync.RWMutex          # read docs for any type
```

## Exploring Code
```bash
go env                       # print all Go environment variables
go env GOPATH                # print a specific variable
go version                   # print installed Go version
go list ./...                # list all packages in the project
```

## Installing Tools
```bash
go install tool@latest       # install a CLI tool written in Go
                             # e.g. go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## Cleaning
```bash
go clean                     # remove compiled binaries
go clean -cache              # clear build cache (fixes weird build issues)
```

---

## Quick Reference — Most Used

| Command | When to use |
|---|---|
| `go run .` | During development — just run it |
| `go build ./...` | Check everything compiles |
| `go test -race ./...` | Before every commit |
| `go fmt ./...` | After writing code |
| `go vet ./...` | After writing code |
| `go mod tidy` | After adding/removing imports |
| `go get package` | Adding a new dependency |
