# Go Back-End Project

## Description
A Go-based back-end project for managing recipes and ingredients. This project uses a clean architecture to keep the codebase modular and maintainable.

## Directory Structure

- **`cmd`**: Entry points for your application. Each executable application should have its own subdirectory with a `main.go` file containing the `main()` function.

- **`internal`**: Core of your application. The `internal` folder is for your application code that should not be imported by other projects.
  - **`api`**: Handlers for your HTTP endpoints and middleware.
  - **`dto`** (Data Transfer Objects): Structures used for data exchange between your API and clients.
  - **`model`**: Data structures representing your business logic and database tables if using an ORM.
  - **`repo`** (Repositories): Abstraction layer over the database. Here you define functions to save and retrieve data.
  - **`service`**: Contains the business logic of your application. Services call repositories to manipulate data and implement business rules.
  - **`validator`**: Validation layer for incoming data.

- **`deployments`**: Docker Compose files for running your application in a containerized environment.

### Installation Prerequisites

- Go 1.16+
- Docker (for using Docker Compose)
- Git

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
