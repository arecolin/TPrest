#!/bin/bash

cd cmd/restserver/
go clean -cache ./...
go build -o . -v ./...
go run .
