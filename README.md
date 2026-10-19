# TaskFlow API

A simple RESTful API for managing daily tasks, built with Go.

## Features

- Create a task
- Retrieve all tasks
- Retrieve a task by ID
- Update a task
- Delete a task
- JSON request and response handling
- Input validation and HTTP error handling
- In-memory storage

## Technologies

- Go
- `net/http`
- `encoding/json`
- Postman for API testing

## Requirements

- Go 1.22 or later
- Git
- Postman (optional, for graphical API testing)

## Installation

### 1. Clone the repository

Replace `YOUR_GITHUB_USERNAME` with your GitHub username.

```bash
git clone https://github.com/YOUR_GITHUB_USERNAME/taskflow.git
```

### 2. Enter the project directory

```bash
cd taskflow
```

### 3. Download and verify dependencies

```bash
go mod tidy
```

This project uses Go's standard library and does not require external Go packages.

### 4. Run the server

```bash
go run .
```

The API will be available at:

`http://localhost:8080`

Health check:

`http://localhost:8080/health`

## API Documentation

See [API Documentation](docs/API.md) for available endpoints, request examples, and HTTP status codes.

## Storage

Tasks are stored in memory. All task data is lost when the server stops or restarts.

## Project Structure

```text
taskflow/
├── main.go
├── models/
│   └── task.go
├── handlers/
│   └── task_handler.go
├── docs/
│   └── API.md
├── go.mod
└── README.md
```

## License

This project was created for educational and demonstration purposes.
