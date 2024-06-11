# Go Server Template

Welcome to the Go Server Template repository! 🎉

*This repository serves as a template for creating a new Golang server. It is designed to save you time and effort by providing a solid starting point so that you don't have to write your Golang server from scratch every time.*

## Features

- **Structured Project Layout**: A well-organized project structure that follows best practices.
- **Router Setup**: Pre-configured routing using popular routing libraries.
- **Middleware**: Basic middleware included for logging, authentication, etc.
- **Configuration Management**: Easy configuration management using environment variables.
- **Error Handling**: Standardized error handling across the application.
- **Database Integration**: Boilerplate code for PostgreSQL integration.
- **Testing**: Example unit and integration tests to get you started.
- **Docker Support**: Dockerfile included for containerized deployments.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.22+)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerization)

### Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/isaka-james/go-server.git
    cd go-server
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Set up environment variables:**

    Edit the `.env` file in the root directory with your configurations:

    ```env
    # THESE ARE POSTGRES CONFIG
    SERVERNAME_DB=localhost
    USERNAME_DB=admin
    PASSWORD_DB=group7
    DATABASE=mydb
    PORT_DB=5432

    PORT_SERVER=80
    ```

4. **Run the server:**

    ```sh
    go run main.go
    ```

    The server should now be running on `http://localhost:80`.


### Using Docker

1. **Build the Docker image:**

    ```sh
    docker build -t go-server-template .
    ```

2. **Run the Docker container:**

    ```sh
    docker run -p 80:80 go-server-template
    ```

    The server should now be running on `http://localhost:80`.

## Project Structure

```plaintext
go-server/
├── README.md
├── api
│   ├── config
│   │   └── config.go
│   ├── handlers
│   │   ├── login_handler.go
│   │   └── notification_handler.go
│   ├── models
│   │   ├── credentials.go
│   │   ├── post.go
│   │   ├── responses.go
│   │   └── user.go
│   ├── routes.go
│   ├── scripts
│   │   └── mydb.sql
│   └── utils
│       └── database.go
├── go.mod
├── go.sum
└── main.go
```

- `api/`: Main directory for API-related code.
  - `config/`: Configuration management.
  - `handlers/`: Request handlers.
  - `models/`: Database models.
  - `scripts/`: SQL scripts.
  - `utils/`: Utility functions.
- `go.mod` and `go.sum`: Go modules files.
- `main.go`: Entry point of the application.



## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


