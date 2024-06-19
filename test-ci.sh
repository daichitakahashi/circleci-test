#!/bin/bash

go test ./... -coverprofile=coverage.out -coverpkg=./... || exit 1;
go tool cover -html=coverage.out -o=coverage.html
echo ""
go tool cover -func=coverage.out
