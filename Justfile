# Show a list of all possible commands
@default:
    just --list

# Run Web Server
@run: build
    bin/todoapi

# Build Web Server
@build:
    go build -o bin/

# Clean the project
@clean:
    rm -rf bin/
