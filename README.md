# About the Project

This project is a tool to benchmark a TimescaleDB instance, reading from a CSV file a list of values to be used to parallel run queries on a TimescaleDB and print metrics of this execution.

The following metrics are calculated:
- Number of Executed Queries - counts how many queries are executed during all the process
- Total Processing Time - calculate the time spent executing all the queries
- Minimum Query Time Execution - calculate the time of the query that spends less time to execute
- Median Query Time Execution - calculate the median time of all executed queries
- Average Query Time Execution - calculate the average time of all executed queries
- Maximum Query Time Execution - calculate the time of the query that spends more time to execute

## Used Libraries
-	[zap](https://pkg.go.dev/go.uber.org/zap) - Improved logging library
- 	[pgx](github.com/jackc/pgx/v5) - PostgreSQL driver
-	[godotenv](github.com/joho/godotenv) - Lbrary to load env vars from a file
-	[testify](github.com/stretchr/testify) - Improved asserts and mocks libray
-	[go-sqlmock](github.com/DATA-DOG/go-sqlmock) - Library to mock sql package interfaces and functions

## Requirements
- [golang](https://go.dev) 1.19
- [golangci-lint](https://golangci-lint.run/usage/install/)
- [docker-compose](https://docs.docker.com/compose/install/)
- [make](https://www.gnu.org/software/make/) tool

# Running the Project

The project has make targets to easly execute some of the project operations.
On the project root folder run the following commands:

## 1. Build
Build the project and create a docker image.

```make build```

## 2. Unit Tests
Run unit tests, using mocks to simulate the database.

```make tests```

## 3. Code Coverage
Generate the code coverage report and open it on your browser.

```make coverage```

## 4. Linter
Run a linter in the project code, using golangci-lint.

```make linter```

## 5. Docker Run

Run the project and all dependencites in docker containers.

```make docker-run```

## 6. Project Dependencies

Start a TimescaleDB in a container and run a flyway container to create the database objets. This is required to run some of the commands bellow.

```make docker-dependencies```

## 7. Standalone App

Run the project as a standalone process, , using the TimescaleDB started in the command 5. 
> Before executing it, you need to copy the file ``.env.sample`` to a file ``.env`` and change its values to your environment.

```make app-run```

## 8. Docker App

Run the project in a docker container.
> Before executing it you need to run the ``Project Dependencies`` command, to start a TimescaleDB instance,

```make docker-app```

## 9. Integration Tests

Run integration tests.
> Before executing it you need to run the ``Project Dependencies`` command, to start a TimescaleDB instance.

```make integration```

## 10. Bencharmk Tests

Run benchmarks test.
> Before executing it you need to run the ``Project Dependencies`` command, to start a TimescaleDB instance,

```make benchmark```
