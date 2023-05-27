# go-example-webapp
A simple example webapp written in Go

## Summary
A bare-bones webapp that supports entering product information and searching on those fields

Provides a simple, in-memory data-store that will not persist between startups.

## Requirements

1. Golang
2. PostgreSQL

## Run Locally

Connect a CLI to your local PostgreSQL instance and create a database called `products`:

```
psql
createdb -p 5432 -h localhost -e products
```

Open a terminal and run the following commands:

```
$ go build .
$ ./go-example-webapp
```

## Run Tests

Open a terminal and run the following commands:

```
$ go test ./...
```

## Stretch Goals

1. simple error handling - DONE
2. search page - DONE
3. home bar links - DONE
4. Postgres DB - DONE
5. fix unit tests
6. integration test
7. use Gin for the web framework
8. use docker-compose for DB 
9. use bootstrap to make the UI look better
10. implement Protobuf / Twirp for the API
11. implement react and typescript for the UI