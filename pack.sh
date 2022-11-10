#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o updatefile.amd64.lin .
GOOS=linux GOARCH=arm64 go build -o updatefile.arm64.lin .
GOOS=windows GOARCH=amd64 go build -o updatefile.amd64.exe .
GOOS=windows GOARCH=arm64 go build -o updatefile.arm64.exe .
GOOS=darwin GOARCH=amd64 go build -o updatefile.amd64.mac .
GOOS=darwin GOARCH=arm64 go build -o updatefile.arm64.mac .
