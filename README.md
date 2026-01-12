# Hexagonal Minimal

Hexagonal Minimal is a Go-based project that demonstrates the use of the **Hexagonal (Ports and Adapters)** architecture. The core logic is decoupled from its surrounding frameworks, delivering cleanly separated layers and adaptable components.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
    - [Requirements](#requirements)
    - [Setup](#setup)
    - [Running the Application](#running-the-application)
- [Available Commands](#available-commands)
- [Contributing](#contributing)
- [License](#license)

---

## Overview

This project demonstrates a simple User Management API, utilizing `MySQL` as the database and following **DDD (Domain-Driven Design)** principles.

Key components:
- `Port`: Interface to the domain layer.
- `Adapter`: Implements interfaces and interacts with external systems (e.g., HTTP and MySQL).

The app supports:
- Creating Users
- Retrieving User Data
- Deleting Users
- Database migrations using `goose`.

---

## Features

- **Hexagonal Architecture**: Clean separation of concerns between Adapters and Core Business Logic.
- **MySQL Database Integration**: Data persistence using the MySQL driver for Go.
- **Goose Migrations**: Management of database schema changes.
- **Environment Configurations**: Easily configurable `.env` setup, with an example file provided.

---

## Project Structure

```
hexagonal-minimal/
├── cmd/
│   ├── api/
│   │   └── main.go          # Entry point API server
│   └── migrate/
│       └── main.go          # Entry point database migration
│
├── internal/
│   ├── adapter/
│   │   ├── http/
│   │   │   └── handler.go   # HTTP handlers (User API)
│   │   └── mysql/
│   │       └── user_repo.go # MySQL repository (implements domain interface)
│   │
│   └── domain/
│       └── user/
│           ├── entity.go     # User entity
│           ├── repository.go # Repository port/interface
│           └── service.go    # Business logic
│
├── migrations/
│   └── 001_create_users.sql # Database migration
│
├── .env.example             # Environment template
├── Makefile                 # Helper commands
├── go.mod                   # Go module
├── go.sum                   # Dependency checksum
└── .gitignore               # Git ignored files
```

---

## Getting Started

### Requirements

- Go 1.25 or later
- MySQL database
- `make` installed on your system

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/nurmanhabib/go-hexagonal-minimal.git
   cd go-hexagonal-minimal
   ```

2. Copy the example `.env` file:
   ```bash
   cp .env.example .env
   ```

3. Update `.env` with your configuration:
   ```plaintext
   APP_PORT=8080

   DB_HOST=127.0.0.1
   DB_PORT=3306
   DB_USER=root
   DB_PASS=your_password
   DB_NAME=your_database
   ```

4. Install dependencies:
   ```bash
   go mod tidy
   ```

5. Run migrations to set up the database:
   ```bash
   make migrate-up
   ```

---

## Running the Application

### Start the API

To run the HTTP server, execute:

```bash
make run
```

The server will be accessible at `http://localhost:8080`. The available endpoints are:
- `POST /users` - Create a user
- `GET /users/get?id=<user_id>` - Retrieve a user
- `DELETE /users/delete?id=<user_id>` - Delete a user

---

## Available Commands

The `Makefile` includes these commands for convenience:

- `make run`: Start the API server.
- `make migrate-up`: Apply database migrations.
- `make migrate-down`: Roll back the last migration.
- `make migrate-status`: Show migration status.

---

## Contributing

We welcome contributions! Feel free to submit pull requests, create issues, or suggest improvements. Please ensure your code aligns with the repository's architectural goals and coding style.

---

## License

This project is licensed under the **MIT License**.
