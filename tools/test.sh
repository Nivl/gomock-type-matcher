#!/bin/bash

go test -v -race -coverprofile=coverage.txt -covermode=atomic .
