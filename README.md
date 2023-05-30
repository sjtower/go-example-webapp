# go-example-webapp
A simple example webapp written in Go

## Summary
A bare-bones webapp that supports entering product information and searching on those fields

Uses Postgres as a data-store that will persist between sessions

## Requirements

1. Golang
2. PostgreSQL

## Run Locally

:information_source: Note that the following instructions assume that you are using MacOS

The program requires a database named `product` and expects the SQL server to run on `localhost:5432`.

One way to do this is to connect a CLI to your local PostgreSQL instance. Open a terminal and run the following commands:

```
psql
CREATE DATABASE product;
```

To run the webserver, open another terminal and run the following commands:

```
$ go build .
$ ./go-example-webapp
```

Navigate to 'http://localhost:8080/add/' to access the web application

## Run Tests

Ensure the webserver is running and run the following command:

```
$ go test
```

## Stretch Goals

1. simple error handling - DONE
2. search page - DONE
3. home bar links - DONE
4. integration test - DONE
5. use docker-compose to add a DB to persist the data 
6. use bootstrap to make the UI look better
7. implement Protobuf / Twirp for the API
8. implement react and typescript for the UI
