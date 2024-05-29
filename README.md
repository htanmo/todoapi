# Todoapi

This repository contains source code for a todo REST api backend written in 
Golang's standard library using [gorm](https://gorm.io/) and [sqlite](https://www.sqlite.org/) database.

## Endpoints

### Base URL

```txt
http://localhost:8080/api/{api_version}
```

### API Endpoints

| Method | Endpoint | Description |
| ------ | -------- |---------- |
| GET | /todo | Return all available todos |
| POST | /todo | Create a new todo |
| GET | /todo/{id} | Return a specific todo |
| PUT | /todo/{id} | Edit a specific todo |
| DELETE | /todo/{id} | Delete a specific todo |

## Building the Project

Default path to store the database is `./todo.db` you can set a custom database name and path to store the database using the `DB_URI` environment variable.

### Justfile (Recommended)

Builds the server executable under `bin/` directory and also runs the server.

```sh
just run
```

### Manually

#### Build Server

```sh
go build -o ./bin/server ./cmd/app/
```

#### Run Server

```sh
./bin/server
```
