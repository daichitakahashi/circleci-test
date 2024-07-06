#!/bin/bash

gotestsum --format=testname --junitfile=test-results.xml -- ./... -coverprofile=coverage.out -coverpkg=./... || exit 1;
go tool cover -html=coverage.out -o=coverage.html
echo ""
go tool cover -func=coverage.out
