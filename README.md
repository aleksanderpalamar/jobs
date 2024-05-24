# Jobs API

This project is a modern API for registering job vacancies using Golang,currently one of the highest paying programming languages. The API is powered by Go-Gin as a router, GoORM for database communication, SQLite as the database, and Swagger for API documentation and testing. The project follows a modern package structure to keep the codebase organized and maintainable.

## Feature

- Using Go-Gin as a router for route management
- Implementing SQLite as the database for the API
- Using GoORM for communication with the database
- Integrating Swagger for API documentation and testing
- Creating a modern package structure for organizing the project
- Implementing a complete job vacancies API with endpoints for searching, creating, editing, and deleting opening.
- Implementing automated tests to ensure API quality

## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github/aleksanderpalamar/jobs.git`
2. Install the dependencies: `go mod downloads`
3. Build the application: `go build`
4. Run the applications: `./main`

## Makefile Commands

The project includes a makefile to help you manage common task more easily. Here's a list of the available commands and a brief description of what they do:

- `make run`: Runs the application without generating API documentation using Swag.
- `make build`: Build the application and create an executable file named `jobs-api`.
- `make test`: Runs tests for all packages on the project.
- `make docs`: Generate the API documentation using Swag.
- `make clean`: Removes the `jobs-api` executable and delete the `./docs` directory.

To use these commands, simply type `make` followed by the desired command. For example:

```bash
make run
```

## Docker and Docker Compose

This project includes a `Dockerfile` and `docker-compose.yml` file to help you manage common task more easily. Here's a list of the available commands and a brief description of what they do:

- `docker build -t jobs-api .`: Builds the application and creates an executable file named `jobs-api`.
- `docker run -p 8080:8080 -e PORT=8080 jobs-api`: Run a container based on the build image.
Replace `jobs-api` with the name you used when building the image. You can change the port number is necessary.

If you want to use Docker Compose, follow these commands:

- `docker compose build`: Builds the application and creates an executable file named `jobs-api`.
- `docker compose up`: Run the services defined in `docker-compose.yml` file.

To stop and remove containers, network, and volumes defined in `docker-compose.yml` file, use the following commands:

```bash
docker compose down
```

## Contributing

This project is open source and you can contribute to it. You can follow the [Contributor Guide](CONTRIBUTING.md) to learn how to contribute.

## License

This project is licensed under the [MIT License](LICENSE).

## Credits

- [Aleksander Palamar](https://www.aleksanderpalamar.dev/)