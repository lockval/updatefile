#!/bin/bash


GOOS=windows    GOARCH=amd64 go build -trimpath -o updatefile.exe   .
GOOS=linux      GOARCH=amd64 go build -trimpath -o updatefile       .
