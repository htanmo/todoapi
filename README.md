# Todoapi

A simple REST API for todo applications.

## Endpoints

### Base URL

```txt
http://localhost:8080/api
```

By default the `todoapi` uses port 8080 you can set a different port using
`PORT` Environment variable. 

### API Endpoints

| Method | Endpoint | Description |
| ------ | -------- |---------- |
| GET | /todo | Return all available todos |
| POST | /todo | Create a new todo |
| GET | /todo/{id} | Return a specific todo |
| PUT | /todo/{id} | Edit a specific todo |
| DELETE | /todo/{id} | Delete a specific todo |

## Building the Project

Default path to store the database is `todo.db` you can set a custom database
name and path to store the database using the `DB_URI` environment variable.

### Justfile (Recommended)

Builds the server executable under `bin/` directory and also runs the server.

```sh
just run
```

### Manually

#### Build Server

```sh
go build -o bin/
```

#### Run Server

```sh
./bin/todoapi
```
