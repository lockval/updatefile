#!/bin/bash


GOOS=windows    GOARCH=arm64 go build -trimpath -o updatefile.arm64.exe .
GOOS=windows    GOARCH=amd64 go build -trimpath -o updatefile.amd64.exe .
GOOS=linux      GOARCH=arm64 go build -trimpath -o updatefile.arm64.lin .
GOOS=linux      GOARCH=amd64 go build -trimpath -o updatefile.amd64.lin .
GOOS=darwin     GOARCH=arm64 go build -trimpath -o updatefile.arm64.app .
GOOS=darwin     GOARCH=amd64 go build -trimpath -o updatefile.amd64.app .
