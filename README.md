# About the Project

This project is a tool to benchmark a TimescaleDB instance. reading from a CSV file a list of values to be used to parallel run queries on a TimescaleDB and print metrics of this execution.

## Metrics

The following metrics are calculated:
- Number of Executed Queries - counts how many queries are executed during all the process
- Total Processing Time - calculate the time spent executing all the queries
- Minimum Query Time Execution - calculate the time of the query that spends less time to execute
- Median Query Time Execution - calculate the median time of all executed queries
- Average Query Time Execution - calculate the average time of all executed queries
- Maximum Query Time Execution - calculate the time of the query that spends more time to execute

## CSV File

The CSV file processed by the tool must have the following layout:
- The first row must have the column names: ``hostname,start_time,end_time``
- From the second line onward:
    - The first column is the hostname
    - The second column is the starte date in the following format: ``YYYY-MM-DD HH:mm:SS``
    - The third column is the end date in the following format: ``YYYY-MM-DD HH:mm:SS``

Example:
```
hostname,start_time,end_time
host_000008,2017-01-01 08:59:22,2017-01-01 09:59:22
host_000001,2017-01-02 13:02:02,2017-01-02 14:02:02
host_000008,2017-01-02 18:50:28,2017-01-02 19:50:28
host_000002,2017-01-02 15:16:29,2017-01-02 16:16:29
host_000003,2017-01-01 08:52:14,2017-01-01 09:52:14
host_000002,2017-01-02 00:25:56,2017-01-02 01:25:56
```

## Used Libraries
-	[zap](https://pkg.go.dev/go.uber.org/zap) - Improved logging library
- 	[pgx](github.com/jackc/pgx/v5) - PostgreSQL driver
-	[godotenv](github.com/joho/godotenv) - Lbrary to load env vars from a file
-	[testify](github.com/stretchr/testify) - Improved asserts and mocks libray
-	[go-sqlmock](github.com/DATA-DOG/go-sqlmock) - Library to mock sql package interfaces and functions

## Requirements
- Linux or MacOS environment
    - WSL Linuix for Windows
- [golang](https://go.dev) 1.19
- [golangci-lint](https://golangci-lint.run/usage/install/)
- [docker engine](https://docs.docker.com/engine/install/) or [docker-desktop](https://www.docker.com/products/docker-desktop/)
- [docker-compose](https://docs.docker.com/compose/install/)
- [make tool](https://www.gnu.org/software/make/)

# Configuration

Before running the project you need to generate the configuration file. Copy the file ``.env.sample`` to a file with the name ``.env``. Use the default parameter values or change its values to your environment.

These parameters below are used when you are running the tool as a standalone app:
- FILE_PATH - path to the csv file.
- DB_URL - connection url to a TimescaleDB instance.

These parameters below are used when you are running the tool in a docker container:
- DOCKER_FILE_PATH - path to the csv file.
- DOCKER_DB_URL - connection url to a TimescaleDB instance.

# Project Commands

The project has make targets to easly execute some of the project operations. On the project root folder you can run the following commands:

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

## 5. TimescaleDB

Starts a TimescaleDB instance in a docker container.

```make timescaledb```

## 6. Migrations

Starts a flyway container and run the migrations to create the database objets in TimescaleDB.
> Before executing it you need to run the ``TimescaleDB`` command.

```make migrations```

## 7. Run Docker App

Run the project in a docker container.
> Before executing it you need to run the ``TimescaleDB`` and ``Migrations`` commands.

```make docker-app```

## 8. Integration Tests

Run integration tests.
> Before executing it you need to run the ``TimescaleDB`` and ``Migrations`` commands.

```make integration```

## 9. Benchmark Tests

Run benchmarks test.
> Before executing it you need to run the ``TimescaleDB`` and ``Migrations`` commands.

```make benchmark```

# Running the Project

To run the project you need to do the following steps:
- Create the environment file as described in the section ``Configuration``;

- Open a terminal in the project root folder;

- Execute the commands bellow to start the timescaledb instance and run the database migrations:

    ```make timescaledb```

    ```make migrations```

- Execute the command bellow to run the tool:

    ```make docker-app```

- The results will be displayed in the console, like the example bellow:
```timescale-assignment-app | 2023-02-06T16:08:30.285Z     INFO    ##### Processing Results #####
Number of Queries:     [200]
Total Processing Time: [220] milliseconds
Minimum Query Time:    [2082952] nanoseconds
Median Query Time:     [5533961] nanoseconds
Average Query Time:    [9359771] nanoseconds
Maximum Query Time:    [65949007] nanoseconds
```
