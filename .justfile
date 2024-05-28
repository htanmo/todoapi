# Show a list of all possible commands
@default:
    just --list

# Run Web Server
@run: build
    ./bin/server

# Build Web Server
@build:
    go build -o ./bin/server ./cmd/app/main.go

# Clean the project
@clean:
    rm -f ./bin/server
