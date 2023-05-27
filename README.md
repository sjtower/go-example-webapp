# go-example-webapp
A simple example webapp written in Go

## Summary
A bare-bones webapp that supports entering product information and searching on those fields

Provides a simple, in-memory data-store that will not persist between startups.

## Requirements

1. Golang

## Run Locally

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
4. unit tests for search and save - DONE
5. integration test
6. use docker-compose to add a DB to persist the data 
7. use bootstrap to make the UI look better
8. implement Protobuf / Twirp for the API
9. implement react and typescript for the UI