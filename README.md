# TODOs API Server using Go Fiber

![Go Fiber](https://img.shields.io/badge/Go%20Fiber-1.17-green) ![License](https://img.shields.io/badge/license-MIT-blue)

A simple and efficient TODOs API server built using the Go Fiber framework. This API provides CRUD operations for managing your TODO tasks with support for database interactions and data validation.

## Features

- **Create, Read, Update, Delete (CRUD)** operations for TODO items
- Database integration with GORM
- Data validation using structs
- RESTful API structure
- Environment configuration support
- Hot-reloading with Air for development

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/gofiber-todos.git
   cd gofiber-todos
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Copy the example environment configuration:**
   ```bash
   cp .env.example .env
   ```

4. **Setup Database Schema:**
   ```bash
   go run commands/schema.go
   ```

5. **Start the application:**
   ```bash
   go run .
   ```
   or, for development with hot-reloading:
   ```bash
   air
   ```

## Usage

Once the application is running, you can interact with the TODOs API. The default base URL is `http://localhost:3000`.

## API Endpoints

| Method | Endpoint         | Description                          |
|--------|------------------|--------------------------------------|
| GET    | `/todos`         | Retrieve all TODOs                  |
| POST   | `/todos`         | Create a new TODO                   |
| GET    | `/todos/:id`     | Retrieve a TODO by ID               |
| PUT    | `/todos/:id`     | Update a TODO by ID                 |
| DELETE | `/todos/:id`     | Delete a TODO by ID                 |

## Project Structure

```plaintext
gofiber-todos/
│
├── commands/           # Command files for setup and migration
│   └── schema.go       # Database schema setup
│
├── configs/            # Configuration files
│   └── config.go       # Configuration handling
│
├── models/             # GORM models
│   └── todo.go         # TODO model definition
│
├── routes/             # API routes
│   └── todo.go         # Routes for TODO operations
│
├── services/           # Service layer
│   └── database.go     # Database connection handling
│
├── .env                # Environment variables
├── .env.example        # Example environment file
├── go.mod              # Go module file
└── main.go             # Entry point of the application
```



